package core

import (
	"github.com/Lyo-Shur/golog/filter"
	"github.com/Lyo-Shur/golog/handler"
	"github.com/Lyo-Shur/golog/log"
)

// 日志主体
type Logger struct {
	// 级别
	Level log.Level
	// 异常回调
	ErrorCallBack func(err error)
	// 过滤器
	Filters []filter.Filter
	// 处理器
	Handlers []handler.Handler
}

// 获取日志级别
func (logger *Logger) GetLevel() log.Level {
	return logger.Level
}

// 设置日志级别
func (logger *Logger) SetLevel(level log.Level) *Logger {
	logger.Level = level
	return logger
}

// 添加过滤器
func (logger *Logger) AddFilter(filter filter.Filter) *Logger {
	logger.Filters = append(logger.Filters, filter)
	return logger
}

// 添加处理器
func (logger *Logger) AddHandler(handler handler.Handler) *Logger {
	logger.Handlers = append(logger.Handlers, handler)
	return logger
}

// 设置异常回调
func (logger *Logger) SetErrorCallBack(function func(err error)) *Logger {
	logger.ErrorCallBack = function
	return logger
}

// 记录日志前置处理
func (logger *Logger) doFilter(level log.Level, message string, param log.Param) (bool, error) {
	levelFilter := filter.LevelFilter{}
	b, err := levelFilter.Verification(level, message, param)
	if err != nil {
		return false, err
	}
	if !b {
		return false, nil
	}
	for i := range logger.Filters {
		item := logger.Filters[i]
		b, err = item.Verification(level, message, param)
		if err != nil {
			return false, err
		}
		if !b {
			return false, nil
		}
	}
	return true, nil
}

// 记录日志
func (logger *Logger) Log(level log.Level, message string, param ...log.KV) {
	// 组装参数
	params := log.Param{
		Level:        logger.Level,
		CustomParams: param,
	}
	// 判断过滤
	b, err := logger.doFilter(level, message, params)
	if err != nil && logger.ErrorCallBack != nil {
		logger.ErrorCallBack(err)
	}
	if !b {
		return
	}
	for i := range logger.Handlers {
		item := logger.Handlers[i]
		err := item.Log(level, message, params)
		if err != nil && logger.ErrorCallBack != nil {
			logger.ErrorCallBack(err)
		}
	}
}

// 记录日志
func (logger *Logger) Debug(message string, param ...log.KV) {
	logger.Log(log.Debug, message, param...)
}
func (logger *Logger) Info(message string, param ...log.KV) {
	logger.Log(log.Info, message, param...)
}
func (logger *Logger) Warning(message string, param ...log.KV) {
	logger.Log(log.Warning, message, param...)
}
func (logger *Logger) Error(message string, param ...log.KV) {
	logger.Log(log.Error, message, param...)
}
func (logger *Logger) Critical(message string, param ...log.KV) {
	logger.Log(log.Critical, message, param...)
}
