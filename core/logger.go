package core

import (
	"golog/filter"
	"golog/handler"
	golog "golog/log"
)

// 日志主体
type Logger struct {
	// 级别
	Level golog.Level
	// 异常回调
	ErrorCallBack func(err error)
	// 过滤器
	Filters []filter.Filter
	// 处理器
	Handlers []handler.Handler
}

// 获取日志级别
func (logger *Logger) GetLevel() golog.Level {
	return logger.Level
}

// 设置日志级别
func (logger *Logger) SetLevel(level golog.Level) *Logger {
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
func (logger *Logger) doFilter(level golog.Level, log string, param golog.Param) (bool, error) {
	levelFilter := filter.LevelFilter{}
	b, err := levelFilter.Verification(level, log, param)
	if err != nil {
		return false, err
	}
	if !b {
		return false, nil
	}
	for i := range logger.Filters {
		item := logger.Filters[i]
		b, err = item.Verification(level, log, param)
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
func (logger *Logger) Log(level golog.Level, log string, param ...golog.KV) {
	// 组装参数
	params := golog.Param{
		Level:        logger.Level,
		CustomParams: param,
	}
	// 判断过滤
	b, err := logger.doFilter(level, log, params)
	if err != nil && logger.ErrorCallBack != nil {
		logger.ErrorCallBack(err)
	}
	if !b {
		return
	}
	for i := range logger.Handlers {
		item := logger.Handlers[i]
		err := item.Log(level, log, params)
		if err != nil && logger.ErrorCallBack != nil {
			logger.ErrorCallBack(err)
		}
	}
}

// 记录日志
func (logger *Logger) Debug(log string, param ...golog.KV) {
	logger.Log(golog.Debug, log, param...)
}
func (logger *Logger) Info(log string, param ...golog.KV) {
	logger.Log(golog.Info, log, param...)
}
func (logger *Logger) Warning(log string, param ...golog.KV) {
	logger.Log(golog.Warning, log, param...)
}
func (logger *Logger) Error(log string, param ...golog.KV) {
	logger.Log(golog.Error, log, param...)
}
func (logger *Logger) Critical(log string, param ...golog.KV) {
	logger.Log(golog.Critical, log, param...)
}
