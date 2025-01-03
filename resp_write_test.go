package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestWriter(t *testing.T) {
	t.Run("can write in RESP format", func(t *testing.T) {
		var buffer bytes.Buffer
		writer := bufio.NewWriter(&buffer)

		w := NewWriter(writer)
		w.Write(Value{typ: "string", str: "OK"})
		writer.Flush() // print that output from the buffer and clear buffer

		got := buffer.String()
		want := "+OK\r\n"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
