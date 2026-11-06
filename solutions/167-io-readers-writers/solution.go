package piscine

import (
	"bytes"
	"io"
)

func ReadAll(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
}

func WriteString(w io.Writer, s string) (int, error) {
	return io.WriteString(w, s)
}

func Copy(dst io.Writer, src io.Reader) (int64, error) {
	return io.Copy(dst, src)
}

func CountLines(r io.Reader) (int, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	return bytes.Count(data, []byte{'\n'}), nil
}

type upperWriter struct{ w io.Writer }

func (u upperWriter) Write(p []byte) (int, error) {
	return u.w.Write(bytes.ToUpper(p))
}

func UpperWriter(w io.Writer) io.Writer {
	return upperWriter{w: w}
}
