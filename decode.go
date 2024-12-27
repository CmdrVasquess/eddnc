package eddnc

import (
	"bytes"
	"compress/zlib"
	"io"
)

func Decode(zmqmsg []byte, limit int64, into []byte) ([]byte, error) {
	zrd, err := zlib.NewReader(bytes.NewReader(zmqmsg))
	if err != nil {
		return nil, err
	}
	defer zrd.Close()
	txt := bytes.NewBuffer(into)
	if limit > 0 {
		_, err = io.Copy(txt, &io.LimitedReader{R: zrd, N: limit})
		if err != nil {
			return nil, err
		}
	} else {
		_, err = io.Copy(txt, zrd)
		if err != nil {
			return nil, err
		}
	}
	return txt.Bytes(), nil
}

func SchemaString(event []byte) string {
	return string(Schema(event))
}

func Schema(raw []byte) []byte {
	tmp := fieldStart(raw, schemaKey)
	if tmp < 0 {
		return nil
	}
	q1 := bytes.IndexByte(raw[tmp:], '"')
	if q1 < 0 {
		return nil
	}
	q1 += tmp + 1
	q2 := bytes.IndexByte(raw[q1:], '"')
	if q2 < 0 {
		return nil
	}
	return raw[q1 : q1+q2]
}

func Message(event []byte) []byte {
	s := MessageStart(event)
	if s < 0 {
		return nil
	}
	e := MessageEnd(event[s:])
	if e < 0 {
		return nil
	}
	return event[s : s+e]
}

func MessageStart(event []byte) int { return fieldStart(event, messageKey) }

var (
	schemaKey  = []byte(`"$schemaRef"`)
	messageKey = []byte(`"message"`)
	eventKey   = []byte(`"event"`)
)

func fieldStart(raw, name []byte) int {
	key := bytes.Index(raw, name)
	if key < 0 {
		return -1
	}
	sep := bytes.IndexByte(raw[key:], ':')
	if sep < 0 {
		return -1
	}
	return key + sep + 1
}

func MessageEnd(event []byte) int {
	lvl := 0
	l := len(event)
	for i := 0; i < l; i++ {
		c := event[i]
		switch c {
		case '{':
			lvl++
		case '}':
			lvl--
			if lvl == 0 {
				return i + 1
			}
		case '"':
			i = jsonSkipString(event, i+1)
		}
	}
	return -1
}

// Check inlining with go build -gcflags -m decode.go
func jsonSkipString(b []byte, i int) int {
	for i <= len(b) {
		switch b[i] {
		case '\\':
			i++
		case '"':
			return i
		}
		i++
	}
	return i
}
