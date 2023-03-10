package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

const TYPE_PASSWORD = "p"
const TYPE_PRIVATEKEY = "i"

func getConfigFile() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("获取配置文件失败")
	}
	return home + "/servers.json"
}

type Server struct {
	Name        string `json:"name"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	PrivateKey  string `json:"private_key"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type Configs struct {
	Servers map[string]*Server
}

func GetConfig() *Configs {
	config, err := os.Open(getConfigFile())
	if err != nil {
		if os.IsNotExist(err) {
			return &Configs{

			}
		}
	}
	defer config.Close()
	var servers []*Server
	err = json.NewDecoder(config).Decode(&servers)
	HandlerErr(err, "unmarshaling servers")
	sm := make(map[string]*Server, len(servers))
	for _, s := range servers {
		sm[s.Name] = s
	}
	return &Configs{
		Servers: sm,
	}
}

func AddConfig(server *Server) {
	oldServers := GetConfig()

	servers := make([]*Server, 0)

	for _, s := range oldServers.Servers {
		servers = append(servers, s)
	}
	servers = append(servers, server)

	err := WriteServerToFile(servers)
	if err != nil {
		fmt.Println("添加失败")
	} else {
		fmt.Println("添加成功")
	}

}

func WriteServerToFile(servers []*Server) error {
	// 创建或打开服务器文件
	file, err := os.Create(getConfigFile())
	if err != nil {
		return err
	}
	defer file.Close()

	// 编码服务器信息为 JSON 格式
	data, err := json.Marshal(servers)
	if err != nil {
		return err
	}

	// 写入 JSON 数据到文件中
	_, err = io.WriteString(file, string(data))
	if err != nil {
		return err
	}

	return nil
}
