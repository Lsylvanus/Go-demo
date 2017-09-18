package main

import (
	"fmt"
	"bufio"
	"io"
	"os/exec"
	"time"
	"os"
	"runtime/pprof"
	"log"
	"runtime/trace"
)

func main() {
	cpuProfile()
	var str string = "world, go..."
	fmt.Printf("hallo %s", str)

	BasicPrint()
	heapProfile()

	command := "ls"
	params := []string{"-l"}
	//执行cmd命令: ls -l
	execCommand(command, params)

	traceProfile()
}

func BasicPrint() {
	str := `Be what u wanna be ...` +
		`This is a very nice song ` +
			`which can make you feel comfortable.`
			fmt.Println(str)
}

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}

	cmd.Start()

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}

	cmd.Wait()
	return true
}

// 生成 CPU 报告
func cpuProfile() {
	f, err := os.OpenFile("cpu.prof", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.Println("CPU Profile started")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	time.Sleep(3 * time.Second)
	fmt.Println("CPU Profile stopped")
}

// 生成堆内存报告
func heapProfile() {
	f, err := os.OpenFile("heap.prof", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	time.Sleep(2 * time.Second)

	pprof.WriteHeapProfile(f)
	fmt.Println("Heap Profile generated")
}

// 生成追踪报告
func traceProfile() {
	f, err := os.OpenFile("trace.out", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.Println("Trace started")
	trace.Start(f)
	defer trace.Stop()

	time.Sleep(3 * time.Second)
	fmt.Println("Trace stopped")
}
