package recorderx

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"testing"
)

func TestSlog(t *testing.T) {
	var b = make([]byte, 0)
	buffer := bytes.NewBuffer(b)
	//buffer := &bytes.Buffer{}
	multiWriter := io.MultiWriter(buffer)
	handler := slog.NewJSONHandler(multiWriter, nil)
	logger := slog.New(handler)
	logger.Info("hello...")
	//var nb = make([]byte, buffer.Len())
	//copy(nb, buffer.Bytes())
	//fmt.Println(string(nb))
	//fmt.Println(string(b))
	//fmt.Println(string(buffer.AvailableBuffer()))
	fmt.Println(string(buffer.Bytes()))
}
