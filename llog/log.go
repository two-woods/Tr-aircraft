package llog

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	ConsoleMode = iota
	FileMode
)

func InitLog(businessKey string, outputMode int, level log.Level) {
	//设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	//file, err := os.OpenFile("/root/installLog.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	writers := []io.Writer{
		os.Stdout}
	var err error
	var file *os.File
	if ConsoleMode == FileMode {
		file, err = os.OpenFile("/tmp/"+businessKey+"Log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		writers = append(writers, file)
	}

	//同时写文件和屏幕
	fileAndStdoutWriter := io.MultiWriter(writers...)
	if err == nil {
		log.SetOutput(fileAndStdoutWriter)
	} else {
		log.Info("failed to log to file.")
	}
	//设置最低loglevel
	log.SetLevel(level)
}
