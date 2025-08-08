package main

import (
	"flag"
	"fmt"
	"os"
)

// Cmd 结构体用于存储命令行参数
type Cmd struct {
	helpFlag    bool     // 是否显示帮助信息
	versionFlag bool     // 是否显示版本信息
	cpOption    string   // classpath 选项的值
	class       string   // 要执行的主类名
	args        []string // 传递给主类的参数
}

// parseCmd 解析命令行参数并返回 Cmd 结构体指针
func parseCmd() *Cmd {
	cmd := &Cmd{} // 创建 Cmd 结构体实例

	// 设置自定义的 Usage 打印函数
	flag.Usage = printUsage

	// 绑定 helpFlag 到 --help 和 -? 选项，用于显示帮助信息
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")

	// 绑定 versionFlag 到 --version 选项，用于显示版本信息
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")

	// 绑定 cpOption 到 --classpath 和 -cp 选项，用于指定 classpath
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")

	// 解析命令行参数
	flag.Parse()

	// 获取剩余的非选项参数
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0] // 第一个参数为 class 名称
		cmd.args = args[1:] // 剩余参数为传递给 class 的参数
	}
	return cmd
}

// printUsage 打印命令行使用说明
func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
