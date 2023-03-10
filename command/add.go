package command

import (
	"fmt"
	"net"
	"os"
	"ssh-manager/lib"
	"strings"
)

// 读取用户输入
func ReadAddInput() lib.Server {
	var (
		name             string
		description      string
		host             string
		port             int
		user             string
		password         string
		passwordType     string
		identifyFilePath string
	)
	for {
		fmt.Print("name: ")
		fmt.Scanln(&name)
		if name != "" {
			break
		}
	}

	fmt.Print("描述: ")
	fmt.Scanln(&description)
	for {
		fmt.Print("host: ")
		fmt.Scanln(&host)
		if host != "" {
			break
		}
		if net.ParseIP(host) == nil{
			fmt.Print("host格式不正确")
		}

	}

	fmt.Print("port: ")
	fmt.Scanln(&port)
	if port == 0 {
		port = 22
	}
	for {
		fmt.Print("user: ")
		fmt.Scanln(&user)
		if user != "" {
			break
		}
	}

	fmt.Print("验证模式(密码(p)/秘钥(i)): ")
	fmt.Scanln(&passwordType)
	if passwordType == "" {
		passwordType = "p"
	}
	if strings.TrimSpace(strings.ToLower(passwordType)) == "p" {
		for {
			fmt.Print("password: ")
			fmt.Scanln(&password)
			if password != "" {
				break
			}
		}
	} else {
		for {
			fmt.Print("秘钥文件路径: ")
			fmt.Scanln(&identifyFilePath)
			_, err := os.Stat(identifyFilePath)
			if err != nil {
				if os.IsNotExist(err) {
					fmt.Println("文件不存在")
				}
				continue
			}
			if identifyFilePath != "" {
				break
			}
		}
	}

	return lib.Server{
		Name:        name,
		Port:        port,
		Description: description,
		Host:        host,
		Username:    user,
		Type:        passwordType,
		PrivateKey:  identifyFilePath,
		Password:    password,
	}
}
