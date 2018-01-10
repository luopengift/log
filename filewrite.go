package log

import (
	"os"
)

type FileWrite struct {
	fd      *os.File
	cname   string //config name
	maxSize int64
}

func NewFile(cname string, maxSize int64) Logger {
	w := &FileWrite{cname: cname, maxSize: maxSize}
	w.open()
	return w
}

func (w *FileWrite) open() (err error) {
	name := NameWithTime(w.cname)
	w.fd, err = os.OpenFile(name, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	return
}

func (w *FileWrite) rorate() error {
	if err := w.fd.Close(); err != nil {
		return err
	}
	return w.open()
}

func (w *FileWrite) Write(p []byte) (int, error) {
	state, err := w.fd.Stat()
	if err != nil {
		return 0, err
	}
	if state.Name() != NameWithTime(w.cname) {
		w.rorate()
	}
	if w.maxSize > 0 && w.maxSize < state.Size() {
		if err = w.fd.Close(); err != nil {
			return 0, nil
		}
	}
	return w.fd.Write(p)
}

func (w *FileWrite) Flush() error {
	return nil
}
