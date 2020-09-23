//go:generate capnp compile -I../../evaluate/go-capnproto2/std -ogo simple.capnp

package capnp

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"testing"
	"time"

	capnp "zombiezen.com/go/capnproto2"
)

func Benchmark_Simple(b *testing.B) {
	total := 32
	simples := make([]Simple, 0, total)
	simplesBytes := make([][]byte, 0, total)
	for i := 0; i < total; i++ {
		s := newSimple()
		bs, _ := s.Segment().Message().Marshal()
		simples = append(simples, s)
		simplesBytes = append(simplesBytes, bs)
	}
	b.Run(fmt.Sprintf("marshal %d simples", total), func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, s := range simples {
				_, _ = s.Segment().Message().Marshal()
			}
		}
	})
	b.Run(fmt.Sprintf("unmarshal %d simples", total), func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, bs := range simplesBytes {
				msg, err := capnp.Unmarshal(bs)
				if err != nil {
					panic(err)
				}
				s, err := ReadRootSimple(msg)
				if err != nil {
					panic(err)
				}
				_ = s
			}
		}
	})
}

func Benchmark_SimpleArr(b *testing.B) {
	total := 32
	simples := make([]SimpleArr, 0, total)
	simplesBytes := make([][]byte, 0, total)
	for i := 0; i < total; i++ {
		s := newSimpleArr()
		bs, _ := s.Segment().Message().Marshal()
		simples = append(simples, s)
		simplesBytes = append(simplesBytes, bs)
	}
	b.Run(fmt.Sprintf("marshal %d simpleArr4", total), func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, s := range simples {
				_, _ = s.Segment().Message().Marshal()
			}
		}
	})
	b.Run(fmt.Sprintf("unmarshal %d simpleArr4", total), func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, bs := range simplesBytes {
				msg, err := capnp.Unmarshal(bs)
				if err != nil {
					panic(err)
				}
				s, err := ReadRootSimpleArr(msg)
				if err != nil {
					panic(err)
				}
				_ = s
			}
		}
	})

}

func newSimple() Simple {
	_, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic(err)
	}

	r, err := NewRootSimple(seg)
	if err != nil {
		panic(err)
	}
	strBytes := make([]byte, 16)
	io.ReadFull(rand.Reader, strBytes)
	r.SetStrField(hex.EncodeToString(strBytes))
	bytes1Field := make([]byte, 1024)
	bytes2Field := make([]byte, 64)
	io.ReadFull(rand.Reader, bytes1Field)
	io.ReadFull(rand.Reader, bytes2Field)
	r.SetBytes1Field(bytes1Field)
	r.SetBytes2Field(bytes2Field)
	r.SetTimestamp(time.Now().Unix())
	return r
}

func newSimpleArr() SimpleArr {
	_, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic(err)
	}

	r, err := NewRootSimpleArr(seg)
	if err != nil {
		panic(err)
	}
	simples, err := r.NewSimples(4)
	if err != nil {
		panic(err)
	}
	simples.Set(0, newSimple())
	simples.Set(1, newSimple())
	simples.Set(2, newSimple())
	simples.Set(3, newSimple())
	return r
}
