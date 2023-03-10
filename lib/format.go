package lib

import "fmt"

func ShowHelp(s *Configs) {
	help:= fmt.Sprintf("%-10s \r\n","服务器连接管理器")
	help += "-add  添加一台服务器\r\n"
	help += "基本用法 sm 服务器名称\r\n"
	help += fmt.Sprintf("| %-8s | %-16s | %-20s \r\n", "名称", "IP", "描述")

	for _, server := range s.Servers {
		help += fmt.Sprintf("| %-10s | %-16s | %-20s \r\n", server.Name, server.Host, server.Description)
	}

	fmt.Println(help)
}
