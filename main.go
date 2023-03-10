package main

import (
	"flag"
	"fmt"
	"ssh-manager/command"
	"ssh-manager/lib"
)

func main() {
	// 定义命令行参数
	add := flag.Bool("add", false, "添加一台服务器")
	help := flag.Bool("help", false, "查看帮助")

	// 解析命令行参数
	flag.Parse()

	// 如果命令行参数中包含 -add 参数，则进入交互模式
	if *add {
		server := command.ReadAddInput()
		lib.AddConfig(&server)
		return
	}

	// 如果命令行参数中包含 -help 参数，则输出帮助信息
	if *help {
		c := lib.GetConfig()
		lib.ShowHelp(c)
		return
	}

	// 否则输出错误信息和帮助信息
	name := flag.Arg(0)
	if name == "" {
		fmt.Println("请指定服务器名称")
		c := lib.GetConfig()
		lib.ShowHelp(c)
		return
	}

	c := lib.GetConfig()
	if server, ok := c.Servers[name]; !ok {
		fmt.Println("服务器不存在")
		lib.ShowHelp(c)
	} else {
		lib.NewShell(server)
	}
}


