package handler

import (
	"github.com/Lyo-Shur/golog/filter"
	"github.com/Lyo-Shur/golog/formatter"
	"github.com/Lyo-Shur/golog/log"
)

// 日志处理器
type Handler struct {
	// 级别
	Level log.Level
	// 格式化器
	Formatter formatter.Formatter
	// 过滤器组
	Filters []filter.Filter
	// 写方法
	Write func(string) error
	// 其他自定义参数
	Params map[string]interface{}
}

// 获取日志级别
func (handler *Handler) GetLevel() log.Level {
	return handler.Level
}

// 设置日志级别
func (handler *Handler) SetLevel(level log.Level) *Handler {
	handler.Level = level
	return handler
}

// 设置格式化器
func (handler *Handler) SetFormatter(formatter formatter.Formatter) *Handler {
	handler.Formatter = formatter
	return handler
}

// 添加过滤器
func (handler *Handler) AddFilters(filter filter.Filter) *Handler {
	handler.Filters = append(handler.Filters, filter)
	return handler
}

// 记录日志前置处理
func (handler *Handler) doFilter(level log.Level, message string, param log.Param) (bool, error) {
	levelFilter := filter.LevelFilter{}
	b, err := levelFilter.Verification(level, message, param)
	if err != nil {
		return false, err
	}
	if !b {
		return false, nil
	}
	for i := range handler.Filters {
		item := handler.Filters[i]
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

// 对外Log接口
func (handler *Handler) Log(level log.Level, message string, param log.Param) error {
	// 执行过滤
	b, err := handler.doFilter(level, message, param)
	if err != nil {
		return err
	}
	if !b {
		return nil
	}
	message = handler.Formatter.Execute(level, message, param)
	return handler.Write(message)
}
