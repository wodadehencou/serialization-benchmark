//go:generate flatc --go --grpc ./simple.fbs

// can use `flatc --proto ../protobuffer/simple.proto` to generate fbs from proto

package flatbuffers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"io"
	"time"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/wodadehencou/serialization-benchmark/flatbuffers/fb"
)

func newSimple(builder *flatbuffers.Builder) flatbuffers.UOffsetT {

	strBytes := make([]byte, 16)
	_, _ = io.ReadFull(rand.Reader, strBytes)
	strField := builder.CreateString(hex.EncodeToString(strBytes))

	bf1 := make([]byte, 1024)
	_, _ = io.ReadFull(rand.Reader, bf1)
	bytes1Field := builder.CreateByteVector(bf1)

	bf2 := make([]byte, 64)
	_, _ = io.ReadFull(rand.Reader, bf2)
	bytes2Field := builder.CreateByteVector(bf2)

	fb.SimpleStart(builder)
	fb.SimpleAddStrField(builder, strField)
	fb.SimpleAddBytes1Field(builder, bytes1Field)
	fb.SimpleAddBytes2Field(builder, bytes2Field)
	fb.SimpleAddTimestamp(builder, time.Now().Unix())
	s := fb.SimpleEnd(builder)

	return s
}

func NewSimpleBytes() []byte {
	builder := flatbuffers.NewBuilder(1_200)
	s := newSimple(builder)

	builder.Finish(s)
	buf := builder.FinishedBytes()
	return buf
}

func newSimpleArr(builder *flatbuffers.Builder) flatbuffers.UOffsetT {

	n := 4

	offsets := make([]flatbuffers.UOffsetT, n)
	for i := 0; i < n; i++ {
		offsets[i] = newSimple(builder)
	}

	fb.SimpleArrStartSimplesVector(builder, n)
	for i := n - 1; i >= 0; i-- {
		builder.PrependUOffsetT(offsets[i])
	}
	simples := builder.EndVector(n)
	fb.SimpleArrStart(builder)
	fb.SimpleArrAddSimples(builder, simples)
	s := fb.SimpleArrEnd(builder)

	return s
}

func NewSimpleArrBytes() []byte {
	builder := flatbuffers.NewBuilder(5_000)
	s := newSimpleArr(builder)

	builder.Finish(s)
	buf := builder.FinishedBytes()
	return buf
}

func UnmarshalSimple(bs []byte) *fb.Simple {
	return fb.GetRootAsSimple(bs, 0)
}

func UnmarshalSimpleArr(bs []byte) *fb.SimpleArr {
	return fb.GetRootAsSimpleArr(bs, 0)
}

func readSimple(strFiled, bytes1Field, bytes2Field []byte, timestamp int64) {
	h := sha256.New()
	_, _ = h.Write(strFiled)
	_, _ = h.Write(bytes1Field)
	_, _ = h.Write(bytes2Field)
	ts := make([]byte, 8)
	binary.LittleEndian.PutUint64(ts, uint64(timestamp))
	_, _ = h.Write(ts)
	h.Sum(nil)
}

func ReadSimple(bs []byte) {
	s := UnmarshalSimple(bs)
	readSimple(
		s.StrField(),
		s.Bytes1FieldBytes(),
		s.Bytes2FieldBytes(),
		s.Timestamp(),
	)
}

func ReadSimpleArr(bs []byte) {
	s := UnmarshalSimpleArr(bs)
	v := new(fb.Simple)
	for i := 0; i < s.SimplesLength(); i++ {
		if s.Simples(v, i) {
			readSimple(
				v.StrField(),
				v.Bytes1FieldBytes(),
				v.Bytes2FieldBytes(),
				v.Timestamp(),
			)
		}

	}
}
