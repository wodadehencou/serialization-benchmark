//go:generate protoc --go_out=. --go_opt=paths=source_relative ./simple.proto

package protobuffer

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"io"
	"time"

	"google.golang.org/protobuf/proto"
)

func NewSimple() *Simple {
	r := &Simple{}
	strBytes := make([]byte, 16)
	_, _ = io.ReadFull(rand.Reader, strBytes)
	r.StrField = hex.EncodeToString(strBytes)
	r.Bytes1Field = make([]byte, 1024)
	r.Bytes2Field = make([]byte, 64)
	_, _ = io.ReadFull(rand.Reader, r.Bytes1Field)
	_, _ = io.ReadFull(rand.Reader, r.Bytes2Field)
	r.Timestamp = time.Now().Unix()
	return r
}

func NewSimpleBytes() []byte {
	s := NewSimple()
	bs, err := proto.Marshal(s)

	if err != nil {
		panic(err)
	}
	return bs
}

func NewSimpleArr() *SimpleArr {
	r := &SimpleArr{}
	r.Simples = make([]*Simple, 4)
	for i := range r.Simples {
		r.Simples[i] = NewSimple()
	}
	return r
}

func NewSimpleArrBytes() []byte {
	s := NewSimpleArr()
	bs, err := proto.Marshal(s)

	if err != nil {
		panic(err)
	}
	return bs
}

func UnmarshalSimple(bs []byte) *Simple {
	s := new(Simple)
	err := proto.Unmarshal(bs, s)
	if err != nil {
		panic(err)
	}
	return s
}

func UnmarshalSimpleArr(bs []byte) *SimpleArr {
	s := new(SimpleArr)
	err := proto.Unmarshal(bs, s)
	if err != nil {
		panic(err)
	}
	return s
}

func readSimple(s *Simple) {
	h := sha256.New()
	_, _ = h.Write([]byte(s.StrField))
	_, _ = h.Write(s.Bytes1Field)
	_, _ = h.Write(s.Bytes2Field)
	ts := make([]byte, 8)
	binary.LittleEndian.PutUint64(ts, uint64(s.Timestamp))
	_, _ = h.Write(ts)
	h.Sum(nil)
}

func ReadSimple(bs []byte) {
	s := UnmarshalSimple(bs)
	readSimple(s)
}

func ReadSimpleArr(bs []byte) {
	s := UnmarshalSimpleArr(bs)
	for _, v := range s.Simples {
		readSimple(v)
	}
}
