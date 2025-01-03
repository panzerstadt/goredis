package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestSimpleResp(t *testing.T) {
	got := SimpleResp("$5\r\nAhmed\r\n")
	want := "Ahmed"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestNewResp(t *testing.T) {
	t.Run("works with bulk arrays (single bulk string)", func(t *testing.T) {
		input := "*1\r\n$5\r\nhello\r\n"
		reader := bufio.NewReader(strings.NewReader(input))

		resp := NewResp(reader)
		got, err := resp.Read()
		if err != nil {
			fmt.Println(err.Error())
			t.Errorf("test returned with an error")
		}

		arrays := make([]Value, 1)
		arrays[0].bulk = "hello"
		want := Value{typ: "array", array: arrays}

		for i := 0; i < len(got.array); i++ {
			if got.array[i].bulk != want.array[i].bulk {
				t.Errorf("got %q want %q", got, want)
			}
		}
	})
	t.Run("works with bulk arrays (multiple bulk strings)", func(t *testing.T) {
		input := "*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n"
		reader := bufio.NewReader(strings.NewReader(input))

		resp := NewResp(reader)
		got, err := resp.Read()
		if err != nil {
			t.Errorf("test returned with an error")
		}

		arrays := make([]Value, 2)
		arrays[0].bulk = "hello"
		arrays[1].bulk = "world"
		want := Value{typ: "array", array: arrays}

		for i := 0; i < len(got.array); i++ {
			if got.array[i].bulk != want.array[i].bulk {
				t.Errorf("got %q want %q", got, want)
			}
		}
	})
}
