package log

import (
	"io"
	goLog "log"
	"os"
	"path/filepath"

	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger(logDir string) {
	goLog.SetOutput(io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:  filepath.Join(logDir, "server.log"),
		MaxSize:   500,
		LocalTime: true,
	}))
}
