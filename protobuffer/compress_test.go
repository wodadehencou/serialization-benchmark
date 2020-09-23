package protobuffer

import (
	"bytes"
	"compress/gzip"
	"testing"
)

func TestCompress(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		bs := NewSimpleBytes()
		var b bytes.Buffer
		w := gzip.NewWriter(&b)
		_, _ = w.Write(bs)
		w.Close()
		t.Logf("Before Compress Simple bytes length is %d", len(bs))
		t.Logf("After gzip Compress length is %d", b.Len())
	})
	t.Run("simpleArr", func(t *testing.T) {
		bs := NewSimpleArrBytes()
		var b bytes.Buffer
		w := gzip.NewWriter(&b)
		_, _ = w.Write(bs)
		w.Close()
		t.Logf("Before Compress SimpleArr bytes length is %d", len(bs))
		t.Logf("After gzip Compress length is %d", b.Len())
	})
}
