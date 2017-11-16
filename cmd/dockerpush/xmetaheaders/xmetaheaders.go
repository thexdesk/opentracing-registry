package xmetaheaders

import (
	"fmt"
	"strings"

	opentracing "github.com/opentracing/opentracing-go"
)

type XMetaHeadersCarrier struct {
	opentracing.TextMapWriter
	opentracing.TextMapReader
}

func (c XMetaHeadersCarrier) Set(key, val string) {
	key = fmt.Sprintf("X-Meta-%s", key)
	c.TextMapWriter.Set(key, val)
}

func (c XMetaHeadersCarrier) ForeachKey(handler func(key, val string) error) error {
	return c.TextMapReader.ForeachKey(func(key, val string) error {
		key = strings.TrimPrefix(key, "X-Meta-")
		return handler(key, val)
	})
}
