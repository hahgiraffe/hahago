/*
 * @Author: haha_giraffe
 * @Date: 2020-02-06 15:28:25
 * @Description: 日志
 */
package hahautils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type LogLevel = uint32

type LogMethod = uint32

// const (
// 	Debug LogLevel = iota
// 	Info
// 	Warn
// 	Error
// 	Fatal
// 	None
// )

// func ToString(log LogMethod) string {
// 	if log == Debug {
// 		return "Debug"
// 	} else if log == Info {
// 		return "Info"
// 	} else if log == Warn {
// 		return "Warn"
// 	} else if log == Error {
// 		return "Error"
// 	} else if log == Fatal {
// 		return "Fatal"
// 	} else {
// 		return "None"
// 	}
// }

const (
	StdoutOnly LogMethod = iota
	FileOnly
	Stdout_File
)

type Logger struct {
	Name string
	// Level    LogLevel
	Method   LogMethod
	FileName string
}

func (logger *Logger) SetOutputFile(n string) {
	logger.FileName = n
	if logger.Method != Stdout_File {
		logger.Method = Stdout_File
	}
	file, err := os.OpenFile(logger.FileName, os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("OpenFile error")
	}
	defer file.Close()

}

//输出println形式
func (logger *Logger) myprintln(logtype string, args ...interface{}) {
	var funcname string
	pc, filename, line, ok := runtime.Caller(2)
	if ok {
		funcname = runtime.FuncForPC(pc).Name()
		funcname = filepath.Ext(funcname)
		funcname = strings.TrimPrefix(funcname, ".")
		filename = filepath.Base(filename)
	}
	if logger.Method == StdoutOnly || logger.Method == Stdout_File {
		fmt.Printf("[%s]\t[%s]\t[%v]\t[%s:%d:%s]\t", logtype, logger.Name, time.Now().Format("2006-01-02 15:04:05"), filename, line, funcname)
		fmt.Println(args...)
	}
	if logger.Method == FileOnly || logger.Method == Stdout_File {
		conf := fmt.Sprintf("[%s]\t[%s]\t[%v]\t[%s:%d:%s]\t", logtype, logger.Name, time.Now().Format("2006-01-02 15:04:05"), filename, line, funcname)
		content := fmt.Sprintln(args...)
		file, err := os.OpenFile(logger.FileName, os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal("OpenFile erorr ", err)
		}
		defer file.Close()
		writer := bufio.NewWriter(file)
		writer.WriteString(conf + content)
		writer.Flush()
	}
}

//printf形式
func (logger *Logger) myprintf(logtype string, format string, args ...interface{}) {
	var funcname string
	pc, filename, line, ok := runtime.Caller(2)
	if ok {
		funcname = runtime.FuncForPC(pc).Name()
		funcname = filepath.Ext(funcname)
		funcname = strings.TrimPrefix(funcname, ".")
		filename = filepath.Base(filename)
	}
	if logger.Method == StdoutOnly || logger.Method == Stdout_File {
		fmt.Printf("[%s]\t[%s]\t[%v]\t[%s:%d:%s]\t", logtype, logger.Name, time.Now().Format("2006-01-02 15:04:05"), filename, line, funcname)
		fmt.Printf(format, args...)
		// fmt.Println()
	}
	if logger.Method == FileOnly || logger.Method == Stdout_File {
		conf := fmt.Sprintf("[%s]\t[%s]\t[%v]\t[%s:%d:%s]\t", logtype, logger.Name, time.Now().Format("2006-01-02 15:04:05"), filename, line, funcname)
		content := fmt.Sprintf(format, args...)
		file, err := os.OpenFile(logger.FileName, os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal("OpenFile erorr ", err)
		}
		defer file.Close()
		writer := bufio.NewWriter(file)
		writer.WriteString(conf + content)
		writer.Flush()
	}
}

//Debug
func (logger *Logger) Debug(args ...interface{}) {
	logger.myprintln("Debug", args...)
}

//Info
func (logger *Logger) Info(args ...interface{}) {
	logger.myprintln("Info", args...)
}

//Warn
func (logger *Logger) Warn(args ...interface{}) {
	logger.myprintln("Warn", args...)
}

//Error
func (logger *Logger) Error(args ...interface{}) {
	logger.myprintln("Error", args...)
}

//Fatal
func (logger *Logger) Fatal(args ...interface{}) {
	logger.myprintln("Fatal", args...)
}

//Debugf
func (logger *Logger) Debugf(format string, args ...interface{}) {
	logger.myprintf("Debug", format, args...)
}

//Infof
func (logger *Logger) Infof(format string, args ...interface{}) {
	logger.myprintf("Info", format, args...)
}

//Warnf
func (logger *Logger) Warnf(format string, args ...interface{}) {
	logger.myprintf("Warn", format, args...)
}

//Errorf
func (logger *Logger) Errorf(format string, args ...interface{}) {
	logger.myprintf("Error", format, args...)
}

//Fatalf
func (logger *Logger) Fatalf(format string, args ...interface{}) {
	logger.myprintf("Fatal", format, args...)
}
func NewLogger(name string) *Logger {
	return &Logger{
		Name: name,
		// Level:    None,
		Method:   StdoutOnly,
		FileName: "",
	}
}

var HaHalog *Logger

func init() {
	HaHalog = NewLogger("chs")
}
