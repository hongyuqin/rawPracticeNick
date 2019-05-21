package main

import "fmt"

//日志写入器接口
type LogWriter interface {
	Write(data interface{}) error
}

//日志器
type Logger struct {
	//这个日志器用到的日志写入器
	writerList []LogWriter
}

//注册一个日志写入器
func (l *Logger) RegisterWriter(writer LogWriter) {
	l.writerList = append(l.writerList, writer)
}

//将一个data类型的数据写入日志
func (l *Logger) Log(data interface{}) {
	//遍历所有的注册写入器
	for _, writer := range l.writerList {
		//将日志输出到每个写入器
		writer.Write(data)
	}
}

//创建日志器实例
func NewLogger() *Logger {
	return &Logger{}
}

//实现日志写入器
type cmdLogWriter struct {
}

func (c *cmdLogWriter) Write(data interface{}) error {
	fmt.Println(data)
	return nil
}
func main() {
	writer := new(cmdLogWriter)
	logger := new(Logger)
	logger.RegisterWriter(writer)
	logger.Log("dsfsdfsdfs")
}
