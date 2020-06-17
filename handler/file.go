package handler

import (
	"os"
	"sync"
	"time"
)

// 检查文件是否存在
func checkFileIsExist(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	panic(err)
}

// 创建或打开文件
func createOrOpenFile(filename string) (*os.File, error) {
	if !checkFileIsExist(filename) {
		return os.Create(filename)
	}
	return os.OpenFile(filename, os.O_APPEND, 0666)
}

// 获取当前日期
func getDataTime(layout string) string {
	return time.Now().Format(layout)
}

// 锁 避免多线程打印 重复打开文件或重复分割
var lock = sync.Mutex{}

// 写文件
func FileWrite(handler Handler, dir string, message string) error {
	lock.Lock()
	// 创建或打开文件
	if _, ok := handler.Params["osFile"]; !ok {
		b := checkFileIsExist(dir)
		if !b {
			err := os.Mkdir(dir, os.ModePerm)
			if err != nil {
				return err
			}
		}
		fileName := getDataTime("20060102150405") + ".log"
		osFile, err := createOrOpenFile(dir + "\\" + fileName)
		if err != nil {
			return err
		}
		handler.Params["osFile"] = osFile
	}
	// 处理日志日期
	ymd := getDataTime("20060102")
	if _, ok := handler.Params["datatime"]; !ok {
		handler.Params["datatime"] = ymd
	}
	// 处理日志大小
	if _, ok := handler.Params["datasize"]; !ok {
		var size int64 = 10485760
		handler.Params["datasize"] = size
	}
	// 获取当前日志文件信息
	osFile := handler.Params["osFile"].(*os.File)
	fileInfo, err := osFile.Stat()
	if err != nil {
		return err
	}
	// 判断是否需要切割日志文件
	datatime := handler.Params["datatime"].(string)
	datasize := handler.Params["datasize"].(int64)
	if ymd != datatime || fileInfo.Size() > datasize {
		fileName := getDataTime("20060102150405") + ".log"
		osFile, err := createOrOpenFile(dir + "\\" + fileName)
		if err != nil {
			return err
		}
		handler.Params["osFile"] = osFile
	}
	lock.Unlock()
	// 记录日志
	osFile = handler.Params["osFile"].(*os.File)
	_, err = osFile.WriteString(message + "\n")
	if err != nil {
		return err
	}
	return nil
}
