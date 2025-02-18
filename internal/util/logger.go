package util

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/donnie4w/go-logger/logger"
)

var (
	Logger *logger.Logging
)

func ExcutePath() string {
	excutePath, err := os.Executable()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	return filepath.Dir(excutePath)
}

func InitLogger() {
	Logger = logger.NewLogger()
	logFilePath := ExcutePath() + "/logs"
	//按日期分割
	Logger.SetRollingDaily(logFilePath, "log.txt")
	//可按 小时，天，月 分割日志
	// Logger.SetRollingByTime()
	//指定文件大小分割日志
	// Logger.SetRollingFile("/var/logs", "log.txt", 300, 10)
	//指定文件大小分割日志，并指定保留最大日志文件数
	Logger.SetRollingFileLoop(logFilePath, "log.txt", 30, logger.MB, 50)
	//压缩分割的日志文件
	Logger.SetGzipOn(true)
	//监控日志打印，并输出日志到 logMonitor.log
	// Logger.SetRollingDaily("", "logMonitor.log")
	Logger.Debug("this is monitor debug log")
}
