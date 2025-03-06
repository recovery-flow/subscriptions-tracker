package cron

import (
	"github.com/sirupsen/logrus"
)

type logger struct {
	in *logrus.Entry
}

func newLogger(in *logrus.Entry) *logger {
	return &logger{in: in.WithField("who", "gocron-scheduler")}
}

func (l *logger) Debug(msg string, args ...any) {
	l.in.Debug(l.toLoganArgs(msg, args)...)
}

func (l *logger) Info(msg string, args ...any) {
	l.in.Info(l.toLoganArgs(msg, args)...)
}

func (l *logger) Warn(msg string, args ...any) {
	l.in.Warn(l.toLoganArgs(msg, args)...)
}

func (l *logger) Error(msg string, args ...any) {
	l.in.Error(l.toLoganArgs(msg, args)...)
}

func (l *logger) toLoganArgs(msg string, args []any) []any {
	return append([]any{msg}, args...)
}
