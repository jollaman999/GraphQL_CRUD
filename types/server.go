package types

import "time"

type Server struct {
	Uuid 		string `json: "uuid"`
	Server_name	string `json: "server_name"`
	Server_disc	string `json: "server_disc"`
	Cpu			int `json: "cpu"`
	Memory		int `json: "memory"`
	Disk_size	int `json: "disk_size"`
	Created time.Time `json: "created"`
}

type Servers struct {
	Servers []Server `json: "server"`
}