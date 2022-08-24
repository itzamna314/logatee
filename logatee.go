package logatee

import "testing"

func New(t *testing.T) *logrus.Logger {
	log := logrus.New()
	log.Out = &writer{t}

	return log
}
