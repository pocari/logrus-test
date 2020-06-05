package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	initLog()
}

func initLog() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

type MyData struct {
	key1 string
	key2 int
}

func (d *MyData) toLogFields() log.Fields {
	return map[string]interface{}{
		"key1": d.key1,
		"key2": d.key2,
	}
}

func main() {
	logger := log.WithFields(log.Fields{"default_field": "hogehoge"})

	logger2 := logger.WithFields(log.Fields{"hoge": "piyo"})
	logger.Info("hoge")
	logger2.Info("hoge")

	md := &MyData{
		key1: "value1",
		key2: 123,
	}

	logger2.WithFields(md.toLogFields()).Info()
}
