package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
)

// 阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
func ExecShell(s string) (string, error) {
	fmt.Println("exec shell file :", s)
	isExit, _ := PathExists(s)
	if isExit == true {
		// 函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
		cmd := exec.Command("/bin/bash", "-c", `. s`)

		output, err := cmd.Output()

		if err != nil {
			fmt.Println("exec shell error: " + err.Error())
		} else {
			fmt.Println("exec shell result: ", string(output))
		}

		return string(output), err

	} else {
		fmt.Println("file not found!")
		return "false", errors.New("file not found")
	}

}

func ExecShellCmd(s string) (string, error) {
	fmt.Println("ExecShellCmd file :", s)
	isExit, _ := PathExists(s)
	if isExit == true {
		cmd := exec.Command("/bin/bash", s)

		stdout, _ := cmd.StdoutPipe()

		if err := cmd.Start(); err != nil {
			fmt.Println("Execute failed when Start:" + err.Error())
			return "false", err
		}

		out_bytes, _ := ioutil.ReadAll(stdout)
		stdout.Close()

		if err := cmd.Wait(); err != nil {
			fmt.Println("Execute failed when Wait:" + err.Error())
			return "false", err
		}

		fmt.Println("Execute finished:" + string(out_bytes))
		return string(out_bytes), nil

	} else {
		fmt.Println("file not found!")
		return "false", errors.New("file not found")
	}
}
