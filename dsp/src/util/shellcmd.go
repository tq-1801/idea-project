package util

import (
	"bytes"
	"os/exec"
)

/*
 * 适用于执行普通非阻塞shell命令，且需要shell标准输出的
 */
func ExecShell(s string) (string, error) {
	LoggerInit().Debug("--------------------cmd begin-------------------")
	LoggerInit().Debug(s)
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", s)

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()
	LoggerInit().Debug("Stdout: ", out.String())
	if err != nil {
		LoggerInit().Debug("Stderr: ", stderr.String())
		//return out.String(), err
	}

	LoggerInit().Debug("--------------------cmd end-------------------")
	return out.String(), err
}
