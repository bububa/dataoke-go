package util

import (
	"bytes"
	"net/url"
	"strings"
	"sync"
)

var urlValuesPool = sync.Pool{
	New: func() any {
		return make(url.Values)
	},
}

func GetUrlValues() url.Values {
	return urlValuesPool.Get().(url.Values)
}

func PutUrlValues(values url.Values) {
	for k := range values {
		values.Del(k)
	}
	urlValuesPool.Put(values)
}

var bytesBufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func GetBytesBuffer() *bytes.Buffer {
	return bytesBufferPool.Get().(*bytes.Buffer)
}

func PutBytesBuffer(b *bytes.Buffer) {
	b.Reset()
	bytesBufferPool.Put(b)
}

var stringsBuilderPool = sync.Pool{
	New: func() any {
		return new(strings.Builder)
	},
}

func GetStringsBuilder() *strings.Builder {
	return stringsBuilderPool.Get().(*strings.Builder)
}

func PutStringsBuilder(b *strings.Builder) {
	b.Reset()
	stringsBuilderPool.Put(b)
}

func StringsJoin(elems ...string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return elems[0]
	}
	var n int
	for i := 0; i < len(elems); i++ {
		n += len(elems[i])
	}

	b := GetStringsBuilder()
	defer PutStringsBuilder(b)
	b.Grow(n)
	for _, s := range elems {
		b.WriteString(s)
	}
	return b.String()
}
