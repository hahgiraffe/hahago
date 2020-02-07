/*
 * @Author: haha_giraffe
 * @Date: 2020-02-06 15:44:57
 * @Description: logTest
 */
package hahautils

import (
	"errors"
	"testing"
)

func TestLog(t *testing.T) {
	//output stdout
	HaHalog.Info("hello, good moring ", 123123)
	HaHalog.Debug("hello, good moring ", 123123)
	HaHalog.Error("hello, good moring ", 123123)
	HaHalog.Fatal("hello, good moring ", 123123)
	HaHalog.Warn("hello, good moring ", 123123)
	//output File
	HaHalog.SetOutputFile("Mylog.log")
	HaHalog.Debug("Wrong Write")
	HaHalog.Fatal("kiill")
	HaHalog.Info("my info")
	HaHalog.Warn("my warn", errors.New("my error"))
	HaHalog.Error("my error")
}
