//go:generate protoc --gogoslick_out=paths=source_relative:. ./pb.proto

package gogoprotobuf

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
)

func Benchmark_Simple(b *testing.B) {
	total := 32
	simples := make([]*Simple, 0, total)
	simplesBytes := make([][]byte, 0, total)
	for i := 0; i < total; i++ {
		s := newSimple()
		bs, _ := proto.Marshal(s)
		simples = append(simples, s)
		simplesBytes = append(simplesBytes, bs)
	}
	b.Run(fmt.Sprintf("marshal %d simples", total), func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, s := range simples {
				// proto.Marshal(s)
				s.Marshal()
			}
		}
	})
	b.Run(fmt.Sprintf("unmarshal %d simples", total), func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, bs := range simplesBytes {
				s := new(Simple)
				// proto.Unmarshal(bs, s)
				s.Unmarshal(bs)
			}
		}
	})
}

func Benchmark_SimpleArr(b *testing.B) {
	total := 32
	simples := make([]*SimpleArr, 0, total)
	simplesBytes := make([][]byte, 0, total)
	for i := 0; i < total; i++ {
		s := newSimpleArr()
		bs, _ := proto.Marshal(s)
		simples = append(simples, s)
		simplesBytes = append(simplesBytes, bs)
	}
	b.Run(fmt.Sprintf("marshal %d simpleArr4", total), func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, s := range simples {
				// proto.Marshal(s)
				s.Marshal()
			}
		}
	})
	b.Run(fmt.Sprintf("unmarshal %d simpleArr4", total), func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, bs := range simplesBytes {
				s := new(SimpleArr)
				// proto.Unmarshal(bs, s)
				s.Unmarshal(bs)
			}
		}
	})

}

func newSimple() *Simple {
	r := &Simple{}
	strBytes := make([]byte, 16)
	io.ReadFull(rand.Reader, strBytes)
	r.StrField = hex.EncodeToString(strBytes)
	r.Bytes1Field = make([]byte, 1024)
	r.Bytes2Field = make([]byte, 64)
	io.ReadFull(rand.Reader, r.Bytes1Field)
	io.ReadFull(rand.Reader, r.Bytes2Field)
	r.Timestamp = time.Now().Unix()
	return r
}

func newSimpleArr() *SimpleArr {
	return &SimpleArr{
		Simples: []*Simple{newSimple(), newSimple(), newSimple(), newSimple()},
	}
}
