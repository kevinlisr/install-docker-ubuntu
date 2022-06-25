package main

func main(){

}

func Exec(name string, args ...string) error { 

	cmd := exec.Command(name, args...) 
	stderr, _ := cmd.StderrPipe() 
	stdout, _ := cmd.StdoutPipe() 

	if err := cmd.Start(); err != nil {
		log.Println("exec the cmd ", name, " failed") 
		return err 
	} 

	// 正常日志 logScan := bufio.NewScanner(stdout) 
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
