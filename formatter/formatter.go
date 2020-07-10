package formatter

import (
	"github.com/Lyo-Shur/golog/log"
	"strconv"
	"time"
)

// 格式化器
type Formatter interface {
	// 执行格式化
	Execute(level log.Level, message string, param log.Param) string
}

// 简单格式化器
type SimpleFormatter struct{}

func (simpleFormatter *SimpleFormatter) Execute(level log.Level, message string, param log.Param) string {
	return time.Now().Format("2006-01-02 15:04:05") + " [级别" + strconv.Itoa(level) + "] " + message
}
