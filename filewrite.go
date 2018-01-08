package log

import (
	"os"
)

type FileWrite struct {
	fd      *os.File
	cname   string //config name
	line    chan []byte
	maxSize int64
}

func NewFileWrite(cname string, maxSize int64) *FileWrite {
	var w FileWrite
	var err error
	w.fd, err = os.OpenFile(cname, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil
	}
	return &w
}

func (w *FileWrite) Flush() error {
	return w.fd.Close()
}

func (w *FileWrite) open(name string) (err error) {
	w.fd, err = os.OpenFile(name, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	return
}
func (w *FileWrite) rorate() error {
	if err := w.Flush(); err != nil {
		return err
	}
	return w.open(NameWithTime(w.cname))
}

func (w *FileWrite) Write(p []byte) (int, error) {
	state, err := w.fd.Stat()
	if err != nil {
		return 0, err
	}
	if state.Name() != NameWithTime(w.cname) {
		w.rorate()
	}
	if w.maxSize < state.Size() {
		if err = w.fd.Close(); err != nil {
			return 0, nil
		}
		//w.open()
	}
	return w.Write(p)
}
