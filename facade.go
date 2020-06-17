package golog

import (
	"fmt"
	"golog/core"
	"golog/filter"
	"golog/formatter"
	"golog/handler"
	"golog/log"
)

// 日志
type Logger = core.Logger

func GetLogger() *Logger {
	return &Logger{
		Level: log.Debug,
		ErrorCallBack: func(err error) {
			fmt.Println(err)
		},
		Filters:  make([]filter.Filter, 0),
		Handlers: make([]handler.Handler, 0),
	}
}

// 过滤器
type Filter = filter.Filter

// 格式化器
type Formatter = formatter.Formatter

func GetSimpleFormatter() Formatter {
	return &formatter.SimpleFormatter{}
}

// 处理器
type Handler = handler.Handler

func GetPrintHandler() Handler {
	h := Handler{}
	h.Formatter = &formatter.SimpleFormatter{}
	h.Params = make(map[string]interface{})
	h.Write = func(log string) error {
		fmt.Println(log)
		return nil
	}
	return h
}
func GetFileHandler(dir string) Handler {
	h := Handler{}
	h.Formatter = &formatter.SimpleFormatter{}
	h.Params = make(map[string]interface{})
	h.Write = func(log string) error {
		return handler.FileWrite(h, dir, log)
	}
	return h
}

// 值
type Level = log.Level

const (
	Invalid  = log.Invalid
	Debug    = log.Debug
	Info     = log.Info
	Warning  = log.Warning
	Error    = log.Error
	Critical = log.Critical
)

type KV = log.KV
type Param = log.Param
