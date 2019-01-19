package log

import (
	"github.com/johntdyer/slackrus"
	"github.com/sirupsen/logrus"
)

// Logging levels: Debug < Info < Warning < Error < Panic.
// Error and Panic will trigger an alarm at AWS which will send emails and slack messages to the support

func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Warningf(format string, args ...interface{}) {
	logrus.Warningf(format, args...)
}

func PrintErrorInSlack(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func Panic(args ...interface{}) {
	logrus.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	logrus.Panicf(format, args...)
}

func Init(logLevel string) {
	level, err := logrus.ParseLevel(logLevel)
	if err == nil {
		logrus.SetLevel(level)
		return
	}
	logrus.Panic(err)
}

func SetHooks(slackWebhookURL, channel string) {
	logrus.AddHook(&slackrus.SlackrusHook{
		HookURL:        slackWebhookURL,
		AcceptedLevels: slackrus.LevelThreshold(logrus.ErrorLevel),
		Channel:        channel,
	})
}
