package base

import (
	"fmt"

	//	"log"
	"sync"

	"github.com/cihub/seelog"
)

type mylog struct {
	File string                 //普通log
	Log  seelog.LoggerInterface // log interface
}

var (
	my_log     *mylog
	mylog_once sync.Once
)

//创建mylog单实例
func Mylog() *mylog {

	mylog_once.Do(func() {
		my_log = &mylog{} //赋值 mylog 实例
		if err := my_log.LogConfig(); err != nil {
			fmt.Println("create mylog err")
		}
	})
	return my_log
}

// LogConfig 从 seelog中读取配置
func (l *mylog) LogConfig() error {
	fmt.Println("1111")
	conf := GetConfig()
	file := conf.Log.File
	logger, err := seelog.LoggerFromConfigAsFile(file)
	if err != nil {
		fmt.Println("seelog error load xml")
		return err
	}
	l.Log = logger

	return nil
}

//Info输出
func (l *mylog) Infof(format string, v ...interface{}) {
	l.Log.Infof(format, v)
}

//Debug输出
func (l *mylog) Debugf(format string, v ...interface{}) {
	l.Log.Debugf(format, v)
}

// Warnf 输出warn信息
func (l *mylog) Warnf(format string, v ...interface{}) {
	l.Log.Warnf(format, v...)
}

// Errorf 输出error信息
func (l *mylog) Errorf(format string, v ...interface{}) {
	l.Log.Errorf(format, v...)
}

func (l *mylog) Info(v ...interface{}) {
	l.Log.Info(v...)
}

func (l *mylog) Debug(v ...interface{}) {
	l.Log.Debug(v...)
}

func (l *mylog) Warn(v ...interface{}) {
	l.Log.Warn(v...)
}

func (l *mylog) Error(v ...interface{}) {
	l.Log.Error(v...)
}

func (l *mylog) Stop() {
	l.Log.Flush()
}

func GetMylog() *mylog {
	return my_log
}
