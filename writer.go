package logatee

import "testing"

type writer struct {
	t *testing.T
}

func (w *writer) Write(p []byte) (int, error) {
	w.t.Log(string(p))
	return len(p), nil
}
