package tools

import (
	"regexp"
	"testing"
)

func TestGID(t *testing.T) {
	for i := 1; i <= 100000000; i++ {
		gid := GID(11)

		match, err := regexp.MatchString("[0-9a-zA-Z]+", gid)
		if err != nil {
			t.Error(err.Error())
		}

		if !match {
			t.Errorf("%s doesn't match", gid)
		}
	}
}
