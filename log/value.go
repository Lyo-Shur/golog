package log

type Level = int

const (
	Invalid Level = iota
	Debug
	Info
	Warning
	Error
	Critical
)

// 获取级别名
func GetLevelName(level Level) string {
	return [6]string{"Invalid", "Debug", "Info", "Warning", "Error", "Critical"}[level]
}

// 键值对
type KV struct {
	Key   string
	Value string
}

// 日志参数
type Param struct {
	// 设置的日志级别
	Level Level
	// 自定义参数
	CustomParams []KV
}
