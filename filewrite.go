package log

import (
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// FileWriter implements io.Writer interface,
type FileWriter interface {
	Write([]byte) (int, error)
	Close() error
	SetMaxBytes(int)
	SetMaxLines(int)
	SetMaxIndex(int)
}

// FileWrite implements FileWriter interface.
type FileWrite struct {
	fd       *os.File
	cname    string //config name
	curPath  string //real path
	curFile  string //real name
	maxBytes int
	curBytes int
	maxLines int
	curLines int
	maxIndex int
	cnt      int //count
}

// NewFile create a FileWrite with implements FileWriter interface.
func NewFile(cname string) FileWriter {
	w := &FileWrite{cname: cname, curPath: path.Dir(cname)}
	w.open()
	return w
}

// SetMaxBytes set max bytes to rorate.
func (w *FileWrite) SetMaxBytes(maxBytes int) {
	w.maxBytes = maxBytes
}

//SetMaxLines sets max lines to rorate.
func (w *FileWrite) SetMaxLines(maxLines int) {
	w.maxLines = maxLines
}

//SetMaxIndex sets max lines to rorate.
func (w *FileWrite) SetMaxIndex(maxIndex int) {
	w.maxIndex = maxIndex
}

// Name gets current filename log
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
	if w.maxIndex > 0 {
		var files files
		filepath.Walk(w.curPath, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() || !strings.Contains(info.Name(), ".log") {
				return nil
			}
			files = append(files, info)
			return nil
		})
		if files.Len() > w.maxIndex {
			sort.Sort(files)
			for _, info := range files[w.maxIndex-1:] {
				os.Remove(path.Join(w.curPath, info.Name()))
			}
		}
	}
	w.curFile = w.Name()
	w.fd, err = os.OpenFile(w.curFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	return
}

func (w *FileWrite) rorate() error {
	if err := w.Close(); err != nil {
		return err
	}
	w.cnt = 0
	return w.open()
}

// Write
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
	w.curLines++
	return n, err
}

// Close close file
func (w *FileWrite) Close() error {
	return w.fd.Close()
}
