package main

import (
	"fmt"
	"github.com/NuyoahCh/jvmgo/ch02/classpath"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("Version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

// startJVM 启动 JVM，打印 classpath、class 名称和参数
func startJVM(cmd *Cmd) {
	// 解析 classpath 和 class 名称
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n",
		cp, cmd.class, cmd.args)
	// 将 class 名称转换为路径格式，并读取类数据
	className := strings.Replace(cmd.class, ".", "/", -1)
	// 读取类数据
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}
