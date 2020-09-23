package flatbuffers

import "testing"

func BenchmarkFlatbuffersSimpleMarshal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		NewSimpleBytes()
	}
}

func BenchmarkFlatbuffersSimpleUnmarshal(b *testing.B) {
	bs := NewSimpleBytes()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ReadSimple(bs)
	}
}

func BenchmarkFlatbuffersSimpleArrMarshal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		NewSimpleArrBytes()
	}
}

func BenchmarkFlatbuffersSimpleArrUnmarshal(b *testing.B) {
	bs := NewSimpleArrBytes()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ReadSimpleArr(bs)
	}
}
