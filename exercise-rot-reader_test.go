package main

import (
	"bufio"
	"strings"
	"testing"
)

type testpair struct {
	value string
	rot13 string
}

var tests = []testpair{
	{"Lbh penpxrq gur pbqr!", "You cracked the code!"},
	{"You cracked the code!", "Lbh penpxrq gur pbqr!"},
	{"Uryyb, Jbeyq!", "Hello, World!"},
	{"Hello, World!", "Uryyb, Jbeyq!"},
	{"", ""},
	{"a", "n"},
	{"n", "a"},
	{"A", "N"},
	{"N", "A"},
	{"z", "m"},
	{"m", "z"},
	{"Z", "M"},
	{"M", "Z"},
}

func TestRot13Reader_Read(t *testing.T) {
	for _, pair := range tests {
		s := strings.NewReader(pair.value)
		r := Rot13Reader{s}
		reader := bufio.NewReader(r)
		result, _ := reader.ReadString('\n')
		if result != pair.rot13 {
			t.Error("For", pair.value, "expected", pair.rot13, "got", result)
		}

	}
}
