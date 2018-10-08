package main

import (
	"testing"
)

func TestMessage(t *testing.T) {
	cases := []struct {
		subtest  string
		input    string
		expected string
	}{
		{
			"hoge -> Hello, hoge",
			"hoge",
			"Hello, hoge",
		},
		{
			"hogehoge -> Hello, hogehoge",
			"hogehoge",
			"Hello, hogehoge",
		},
		{
			"hogehogehogehoge -> Hello, hogehogehogehoge",
			"hogehogehogehoge",
			"Hello, hogehogehogehoge",
		},
	}

	for _, c := range cases {
		t.Run(c.subtest, func(t *testing.T) {
			msg := NewMessage(c.input)
			actual := msg.String()
			msg.Destroy()

			if actual != c.expected {
				t.Errorf("expected %#v but got %#v", c.expected, actual)
			}
		})
	}
}
