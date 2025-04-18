package utils

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"unsafe"

	"github.com/fatih/color"
)

//go:inline
func String2Bytes(str string) []byte {
	return unsafe.Slice(unsafe.StringData(str), len(str))
}

//go:inline
func Bytes2String(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

//go:inline
func CountByte(s []byte, sep byte) int {
	n := 0
	for {
		i := bytes.IndexByte(s, sep)
		if i == -1 {
			return n
		}
		n++
		s = s[i+1:]
	}
}

//go:inline
func FileExist(path string) bool {
	ok, _ := fileExists(path)
	return ok
}

//go:inline
func Slice(str string, start, end int) string {
	if end < 0 {
		end = len(str) + end
	}
	if end > len(str) {
		end = len(str)
	}
	return str[start:end]
}

//go:inline
func WriteFile(path string, b []byte) error {
	bLength := len(b)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	// 获取当前文件大小
	oldSize, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}
	if int64(bLength) != oldSize {
		if err = file.Truncate(int64(bLength)); err != nil {
			return err
		}
	}
	file.Seek(0, io.SeekStart)
	_, err = file.Write(b)
	if err != nil {
		return err
	}
	return file.Close()
}

//go:inline
func PathJoin(elem ...string) string {
	size := 0
	for _, e := range elem {
		size += len(e)
	}
	if size == 0 {
		return ""
	}
	buf := make([]byte, 0, size+len(elem)-1)
	for _, e := range elem {
		if len(buf) > 0 || e != "" {
			if len(buf) > 0 {
				buf = append(buf, '/')
			}
			buf = append(buf, e...)
		}
	}
	return path.Clean(Bytes2String(buf))
}

//go:inline
func DirWalk(dir string, wdf filepath.WalkFunc) error {
	fileInfo, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, entry := range fileInfo {
		sub := path.Join(dir, entry.Name())
		if entry.IsDir() {
			err = DirWalk(sub, wdf)
		} else {
			err = wdf(sub, nil, err)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

//go:inline
func FileCount(dir string) (n int) {
	fileInfo, err := os.ReadDir(dir)
	if err != nil {
		return 0
	}

	for _, entry := range fileInfo {
		if entry.IsDir() {
			sub := path.Join(dir, entry.Name())
			FileCount(sub)
			n++
		}
	}
	return n
}

//go:inline
func SplitByte(s []byte, sep byte) [][]byte {
	if len(s) == 0 {
		return nil
	}
	n := CountByte(s, sep) + 1
	a := make([][]byte, n)
	n--
	i := 0
	for i < n {
		m := bytes.IndexByte(s, sep)
		if m < 0 {
			break
		}
		a[i] = s[:m:m]
		s = s[m+1:]
		i++
	}
	a[i] = s
	return a[:i+1]
}

func init() {
	log.SetFlags(0)
}

const (
	Warning = iota
	Info
	Error
	Debug
	Success
)

var (
	warnStr    = color.YellowString("Warning") + ":"
	infoStr    = color.BlueString("Info") + ":"
	errorStr   = color.RedString("Error") + ":"
	debugStr   = color.CyanString("Debug") + ":"
	successStr = color.GreenString("Success") + ":"
)

//go:inline
func Log(l int, v ...any) {
	var t string
	switch l {
	case Warning:
		t = warnStr
	case Info:
		t = infoStr
	case Error:
		t = errorStr
	case Debug:
		t = debugStr
	case Success:
		t = successStr
	}
	t2 := []any{t}
	log.Println(append(t2, v...)...)
}

func Exec(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
