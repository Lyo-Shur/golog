package filter

import "github.com/lyoshur/golog/log"

// 过滤器
type Filter interface {
	// 验证是否能通过 过滤
	Verification(level log.Level, message string, param log.Param) (bool, error)
}

// 等级过滤器
type LevelFilter struct{}

func (levelFilter *LevelFilter) Verification(level log.Level, message string, param log.Param) (bool, error) {
	return param.Level <= level, nil
}
