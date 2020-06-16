package down

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"io"
	"sync/atomic"
	"time"

	"git.fractalqb.de/fractalqb/c4hgol"
	"git.fractalqb.de/fractalqb/qbsllm"
	"github.com/CmdrVasquess/goEDDNc/schema"
	zmq "github.com/pebbe/zmq4"
)

var (
	log    = qbsllm.New(qbsllm.Lnormal, "down", nil, nil)
	LogCfg = c4hgol.Config(qbsllm.NewConfig(log))
)

type Subscriber struct {
	Blackmarket <-chan *schema.Blackmarket
	Commodity   <-chan *schema.Commodity
	Journal     <-chan *schema.Journal
	Outfitting  <-chan *schema.Outfitting
	Shipyard    <-chan *schema.Shipyard

	chanNo  int
	relay   string
	timeout time.Duration
	closing int32
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
}

func New(cfg Config) *Subscriber {
	var (
		bChan chan *schema.Blackmarket
		cChan chan *schema.Commodity
		jChan chan *schema.Journal
		oChan chan *schema.Outfitting
		sChan chan *schema.Shipyard
	)
	chanNo := 0
	if cfg.QCapBlackmarket >= 0 {
		bChan = make(chan *schema.Blackmarket, cfg.QCapBlackmarket)
		chanNo++
	}
	if cfg.QCapCommodity >= 0 {
		cChan = make(chan *schema.Commodity, cfg.QCapCommodity)
		chanNo++
	}
	if cfg.QCapJournal >= 0 {
		jChan = make(chan *schema.Journal, cfg.QCapJournal)
		chanNo++
	}
	if cfg.QCapOutfitting >= 0 {
		oChan = make(chan *schema.Outfitting, cfg.QCapOutfitting)
		chanNo++
	}
	if cfg.QCapShipyard >= 0 {
		sChan = make(chan *schema.Shipyard, cfg.QCapShipyard)
		chanNo++
	}
	res := &Subscriber{
		Blackmarket: bChan,
		Commodity:   cChan,
		Journal:     jChan,
		Outfitting:  oChan,
		Shipyard:    sChan,
		chanNo:      chanNo,
		relay:       cfg.Relay,
		timeout:     cfg.Timeout,
	}
	if res.relay == "" {
		res.relay = DefaultRelay
	}
	go res.loop(bChan, cChan, jChan, oChan, sChan)
	return res
}

func (s *Subscriber) UsedChannels() int { return s.chanNo }

func (s *Subscriber) Close() bool {
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

func (s *Subscriber) loop(
	bChan chan<- *schema.Blackmarket,
	cChan chan<- *schema.Commodity,
	jChan chan<- *schema.Journal,
	oChan chan<- *schema.Outfitting,
	sChan chan<- *schema.Shipyard,
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
	var txt bytes.Buffer
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
		if err != nil {
			log.Errore(err)
			continue
		}
		zrd, err := zlib.NewReader(bytes.NewReader(msg))
		if err != nil {
			log.Errore(err)
			continue
		}
		txt.Reset()
		io.Copy(&txt, zrd)
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
				msg := new(schema.Blackmarket)
				if err = json.Unmarshal(line, msg); err != nil {
					log.Errore(err)
				} else {
					bChan <- msg
				}
			}
		case bytes.Index(scm, commodityTag) >= 0:
			if cChan != nil {
				msg := new(schema.Commodity)
				if err = json.Unmarshal(line, msg); err != nil {
					log.Errore(err)
				} else {
					cChan <- msg
				}
			}
		case bytes.Index(scm, journalTag) >= 0:
			if jChan != nil {
				msg := new(schema.Journal)
				if err = json.Unmarshal(line, &msg); err != nil {
					log.Errore(err)
				} else {
					jChan <- msg
				}
			}
		case bytes.Index(scm, outfittingTag) >= 0:
			if oChan != nil {
				msg := new(schema.Outfitting)
				if err = json.Unmarshal(line, msg); err != nil {
					log.Errore(err)
				} else {
					oChan <- msg
				}
			}
		case bytes.Index(scm, shipyardTag) >= 0:
			if sChan != nil {
				msg := new(schema.Shipyard)
				if err = json.Unmarshal(line, msg); err != nil {
					log.Errore(err)
				} else {
					sChan <- msg
				}
			}
		default:
			log.Errora("unknown `schema`", string(scm))
		}
	}
}
