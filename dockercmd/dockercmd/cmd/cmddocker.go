package cmd

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
)

func cmddocker(){
	// wei le fu wu bao cuo er bu tui chu, xu yao bu huo yi chang
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover 捕获到了异常,次异常为程序退出时发生，可忽略！", err)
			//fmt.Println(err)
		}
	}()
	com := "docker"
	args := "version"
	if err := Exec(com, args); err != nil{
		fmt.Errorf("command run error")
	}

}

func Exec(name string, args ...string) error {

	cmd := exec.Command(name, args...)
	stderr, _ := cmd.StderrPipe()
	stdout, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		log.Println("exec the cmd ", name, " failed", "will install docker...")
		if err := logs("apt-get","update"); err != nil{
			return err
		}
		//
		//cmdupdate := exec.Command("apt-get", "update")
		//apterr, _ := cmdupdate.StderrPipe()
		//aptout, _ := cmdupdate.StdoutPipe()
		//if err := cmdupdate.Start();err != nil{
		//	log.Println("exec the cmd apt-get failed")
		//	return err
		//}
		//
		//// 正常日志
		//logScan := bufio.NewScanner(aptout)
		//go func() {
		//	for logScan.Scan() {
		//		log.Println(logScan.Text())
		//	}
		//}()
		//
		//// 错误日志
		//errBuf := bytes.NewBufferString("")
		//scan := bufio.NewScanner(apterr)
		//for scan.Scan() {
		//	s := scan.Text()
		//	log.Println("build error: ", s)
		//	errBuf.WriteString(s)
		//	errBuf.WriteString("\n")
		//}

		//exec.Command("apt-get", "install", "-y", "apt-transport-https")
		if err := logs("apt-get", "install", "-y", "apt-transport-https"); err != nil{
			return err
		}

		//exec.Command("curl", "https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg","|", "apt-key add", "-")
		if err := logs("/bin/bash", "-c", `curl https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | apt-key add -`); err != nil{
			return err
		}

		//exec.Command("echo", "deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main", ">",  "/etc/apt/sources.list.d/kubernetes.list")
		if err := logs("/bin/bash", "-c", `echo deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main > /etc/apt/sources.list.d/kubernetes.list`); err != nil{
			return err
		}

		//exec.Command("apt-get", "install", "-y", "docker.io")
		if err := logs("/bin/bash", "-c", `apt install -y docker.io`);err != nil{
			return err
		} else {
			fmt.Println("SUCCESSFUL!  Docker Install Complete !")
		}

	}

	// 正常日志
	logScan := bufio.NewScanner(stdout)
	go func() {
		for logScan.Scan() {
			log.Println(logScan.Text())
		}
	}()

	// 错误日志
	errBuf := bytes.NewBufferString("")
	scan := bufio.NewScanner(stderr)
	for scan.Scan() {
		s := scan.Text()
		log.Println("build error: ", s)
		errBuf.WriteString(s)
		errBuf.WriteString("\n")
	}
	// 等待命令执行完
	cmd.Wait()
	if !cmd.ProcessState.Success() {
		// 执行失败，返回错误信息
		return errors.New(errBuf.String())
	}
	return nil
}



func logs (cmdname string, para ...string) error {
	cmd := exec.Command(cmdname, para...)

	stderr, _ := cmd.StderrPipe()
	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start();err != nil{
		log.Println("exec the cmd", cmdname, para, "failed")
		fmt.Println("==========================start exec command... ===========================")
		return err
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover 捕获到了异常", err)
			fmt.Println(err)
		}
	}()


	// 正常日志
	logScan := bufio.NewScanner(stdout)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recover 捕获到了异常", err)
				fmt.Println(err)
			}
		}()
		for logScan.Scan() {
			fmt.Println("===================normal logs...========================")
			log.Println(logScan.Text())

		}
	}()

	// 错误日志
	errBuf := bytes.NewBufferString("")
	scan := bufio.NewScanner(stderr)
	for scan.Scan() {
		fmt.Println("===================false logs...=====================")
		s := scan.Text()
		log.Println("build error: ", s)
		errBuf.WriteString(s)
		errBuf.WriteString("\n")
	}
	// 等待命令执行完
	cmd.Wait()
	if !cmd.ProcessState.Success() {
		// 执行失败，返回错误信息
		fmt.Println("=====================On failure ===================")
		return errors.New(errBuf.String())
	}

	return  nil
}
