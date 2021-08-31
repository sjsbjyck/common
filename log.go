package common

import (
	"fmt"
	"path"
	"runtime"
	"strconv"
	"time"
)

//
//  WriteLog
//  @Author: ChenKun
//  @Description: 输出指定格式日志（增加当前打印代码位置及行数）
//  @param format 	string			需要格式化的模板
//  @param v 		...interface{} 	可变参数
//
func WriteLog(format string, v ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}

	_, filename := path.Split(file)

	msg := time.Now().Format("2006-01-02 15:04:05") + " " + filename + ":" + strconv.Itoa(line) + " " + fmt.Sprintf(format, v...)

	fmt.Printf(msg + "\n")
}
