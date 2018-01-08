package log

import (
	"fmt"
)

type Logger interface {
	Init(...string) error
	Write([]byte) (int, error)
}

type Formatter interface {
	Format(ctn *Content) string
}

type JSONFormat struct{}

func (f *JSONFormat) Format(ctn *Content) string {
	return ctn.JSON()
}

type PlainFormat struct {
}

func (f *PlainFormat) Format(ctn *Content) string {
	return ctn.String()
}

type KvFormat struct{}

func (f *KvFormat) Format(ctn *Content) string {
	return fmt.Sprintf("TODO:kv format, %s", ctn.body)
}
