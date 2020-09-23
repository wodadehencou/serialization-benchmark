package protobuffer

import "testing"

func BenchmarkProtobufSimpleMarshal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		NewSimpleBytes()
	}
}

func BenchmarkProtobufSimpleUnmarshal(b *testing.B) {
	bs := NewSimpleBytes()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ReadSimple(bs)
	}
}

func BenchmarkProtobufSimpleArrMarshal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		NewSimpleArrBytes()
	}
}

func BenchmarkProtobufSimpleArrUnmarshal(b *testing.B) {
	bs := NewSimpleArrBytes()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ReadSimpleArr(bs)
	}
}

func BenchmarkReadSimple(b *testing.B) {
	s := NewSimple()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		readSimple(s)
	}
}

func BenchmarkReadSimpleArr(b *testing.B) {
	s := NewSimpleArr()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, v := range s.Simples {
			readSimple(v)
		}
	}
}
