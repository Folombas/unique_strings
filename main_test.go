package main

import (
	"testing"
)

func TestOK(t *testing.T) {

	err := uniq()
	if err != nil {
		t.Errorf("test for OK failed")
	}
}
