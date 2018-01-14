package log

import (
	"os"
	"strconv"
)

type FileWriter interface {
	Write([]byte) (int, error)
	SetMaxBytes(int)
	SetMaxLines(int)
}

type FileWrite struct {
	fd       *os.File
	cname    string //config name
	curPath  string //real path
	curFile  string //real name
	maxBytes int
	curBytes int
	maxLines int
	curLines int
	cnt      int //count
}

func NewFile(cname string) FileWriter {
	w := &FileWrite{cname: cname}
	w.open()
	return w
}

func (w *FileWrite) SetMaxBytes(maxBytes int) {
	w.maxBytes = maxBytes
}

func (w *FileWrite) SetMaxLines(maxLines int) {
	w.maxLines = maxLines
}

func (w *FileWrite) Name() string {
	name := NameWithTime(w.cname)
	if w.maxBytes > 0 || w.maxLines > 0 {
		name = name + strconv.Itoa(w.cnt)
	}
	if w.cnt == 0 {
		name = NameWithTime(w.cname)
	}
	return name
}

func (w *FileWrite) open() (err error) {
	w.curFile = w.Name()
	w.fd, err = os.OpenFile(w.curFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	return
}

func (w *FileWrite) rorate() error {
	if err := w.fd.Close(); err != nil {
		return err
	}
	w.cnt = 0
	return w.open()
}

func (w *FileWrite) Write(p []byte) (int, error) {
	if w.curFile != w.Name() {
		w.rorate()
	}
	if w.maxBytes > 0 && w.maxBytes <= w.curBytes || w.maxLines > 0 && w.maxLines <= w.curBytes {
		if err := w.fd.Close(); err != nil {
			return 0, nil
		}
		w.curBytes = 0
		w.curLines = 0
		w.cnt++
		w.open()
	}
	n, err := w.fd.Write(p)
	w.curBytes += n
	w.curLines += 1
	return n, err
}
