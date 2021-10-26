package subscriber

import (
	"bytes"
	"compress/zlib"
	"io"
	"sync"
	"sync/atomic"
	"time"

	"git.fractalqb.de/fractalqb/c4hgol"
	"git.fractalqb.de/fractalqb/qbsllm"
	zmq "github.com/pebbe/zmq4"
)

var (
	log    = qbsllm.New(qbsllm.Lnormal, "fromeddn", nil, nil)
	LogCfg = c4hgol.Config(qbsllm.NewConfig(log))
)

type EnqueueFunc func(c chan<- []byte, data []byte, queue string)

func Blocking(c chan<- []byte, data []byte, _ string) {
	c <- data
}

func Dropping(c chan<- []byte, data []byte, queue string) {
	select {
	case c <- data:
	default:
		log.Warna("dropping message from `queue`", queue)
	}
}

type DropStats struct{ Total, Dropped uint64 }

type DropWithStats map[string]*DropStats

func NewDropWithStats() DropWithStats { return make(DropWithStats) }

func (dws DropWithStats) Enqueue(c chan<- []byte, data []byte, queue string) {
	stats := dws[queue]
	if stats == nil {
		stats = new(DropStats)
		dws[queue] = stats
	}
	stats.Total++
	select {
	case c <- data:
	default:
		stats.Dropped++
		log.Warna("`drop` message of `total` from `queue`",
			stats.Dropped,
			stats.Total,
			queue)
	}
}

type S struct {
	Blackmarket <-chan []byte
	Commodity   <-chan []byte
	Journal     <-chan []byte
	Outfitting  <-chan []byte
	Shipyard    <-chan []byte

	rtuxm   int64
	chanNo  int
	relay   string
	timeout time.Duration
	closing int32
	enqueue EnqueueFunc
}

const (
	DefaultRelay = "tcp://eddn.edcd.io:9500"
	GoodTimeout  = 6 * time.Second
)

type Config struct {
	Relay           string
	Timeout         time.Duration
	QCapBlackmarket int
	QCapCommodity   int
	QCapJournal     int
	QCapOutfitting  int
	QCapShipyard    int
	Enqueue         EnqueueFunc
}

func New(cfg *Config) *S {
	var (
		bChan chan []byte
		cChan chan []byte
		jChan chan []byte
		oChan chan []byte
		sChan chan []byte
	)
	chanNo := 0
	if cfg.QCapBlackmarket >= 0 {
		bChan = make(chan []byte, cfg.QCapBlackmarket)
		chanNo++
	}
	if cfg.QCapCommodity >= 0 {
		cChan = make(chan []byte, cfg.QCapCommodity)
		chanNo++
	}
	if cfg.QCapJournal >= 0 {
		jChan = make(chan []byte, cfg.QCapJournal)
		chanNo++
	}
	if cfg.QCapOutfitting >= 0 {
		oChan = make(chan []byte, cfg.QCapOutfitting)
		chanNo++
	}
	if cfg.QCapShipyard >= 0 {
		sChan = make(chan []byte, cfg.QCapShipyard)
		chanNo++
	}
	res := &S{
		Blackmarket: bChan,
		Commodity:   cChan,
		Journal:     jChan,
		Outfitting:  oChan,
		Shipyard:    sChan,
		chanNo:      chanNo,
		relay:       cfg.Relay,
		timeout:     cfg.Timeout,
		enqueue:     cfg.Enqueue,
	}
	if res.relay == "" {
		res.relay = DefaultRelay
	}
	if res.enqueue == nil {
		res.enqueue = Blocking
	}
	go res.loop(bChan, cChan, jChan, oChan, sChan)
	return res
}

func (s *S) Return(rawEvent []byte) {
	bufPool.Put(rawEvent[:0])
}

func (s *S) UsedChannels() int { return s.chanNo }

func (s *S) TRecvUnixMilli() int64 { return atomic.LoadInt64(&s.rtuxm) }

func (s *S) Close() bool {
	return atomic.CompareAndSwapInt32(&s.closing, 0, 1)
}

func must(err error) {
	if err != nil {
		log.Panice(err)
	}
}

var (
	schemaRefTag   = []byte("$schemaRef")
	blackmarketTag = []byte("blackmarket")
	commodityTag   = []byte("commodity")
	journalTag     = []byte("journal")
	outfittingTag  = []byte("outfitting")
	shipyardTag    = []byte("shipyard")

	bufPool = sync.Pool{
		New: func() interface{} { return []byte{} }, // TODO good default size
	}
)

func pickSchema(text []byte) []byte {
	idx := bytes.Index(text, schemaRefTag)
	if idx < 0 {
		return nil
	}
	text = text[idx+len(schemaRefTag)+1:]
	if idx = bytes.IndexByte(text, '"'); idx < 0 {
		return nil
	}
	text = text[idx+1:]
	if idx = bytes.IndexByte(text, '"'); idx < 0 {
		return nil
	}
	return text[:idx]
}

func (s *S) loop(
	bChan chan<- []byte,
	cChan chan<- []byte,
	jChan chan<- []byte,
	oChan chan<- []byte,
	sChan chan<- []byte,
) {
	zctx, err := zmq.NewContext()
	if err != nil {
		log.Panice(err)
	}
	subs, err := zctx.NewSocket(zmq.SUB)
	if err != nil {
		log.Panice(err)
	}
	defer subs.Close()
	must(subs.SetSubscribe(""))
	must(subs.SetConnectTimeout(s.timeout))
	must(subs.Connect(s.relay))
	for {
		if atomic.CompareAndSwapInt32(&s.closing, 1, -1) {
			if bChan != nil {
				close(bChan)
			}
			if cChan != nil {
				close(cChan)
			}
			if jChan != nil {
				close(jChan)
			}
			if oChan != nil {
				close(oChan)
			}
			if sChan != nil {
				close(sChan)
			}
			return
		}
		msg, err := subs.RecvBytes(0)
		atomic.StoreInt64(&s.rtuxm, time.Now().UnixMilli())
		if err != nil {
			log.Errore(err)
			continue
		}
		zrd, err := zlib.NewReader(bytes.NewReader(msg))
		if err != nil {
			log.Errore(err)
			continue
		}
		txt := bytes.NewBuffer(bufPool.Get().([]byte))
		io.Copy(txt, zrd)
		zrd.Close()
		line := txt.Bytes()
		scm := pickSchema(line)
		if scm == nil {
			log.Errora("no $schemaRef in `message`", string(line))
			continue
		}
		switch {
		case bytes.Index(scm, blackmarketTag) >= 0:
			if bChan != nil {
				s.enqueue(bChan, line, "blackmarket")
			}
		case bytes.Index(scm, commodityTag) >= 0:
			if cChan != nil {
				s.enqueue(cChan, line, "commodity")
			}
		case bytes.Index(scm, journalTag) >= 0:
			if jChan != nil {
				s.enqueue(jChan, line, "journal")
			}
		case bytes.Index(scm, outfittingTag) >= 0:
			if oChan != nil {
				s.enqueue(oChan, line, "outfitting")
			}
		case bytes.Index(scm, shipyardTag) >= 0:
			if sChan != nil {
				s.enqueue(sChan, line, "shipyard")
			}
		default:
			bufPool.Put(line)
			log.Errora("unknown `schema`", string(scm))
		}
	}
}
