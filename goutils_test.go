package goutils

import (
	"testing"
)

func TestStat(t *testing.T) {
	s, _ := Stat("go.mod")
	t.Log(s)
}
