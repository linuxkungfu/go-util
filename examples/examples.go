package main

import (
	"fmt"

	"github.com/linuxkungfu/go-util"
)

func main() {
	serverId := util.CreateServerId("123")
	fmt.Printf("server id:%s\n", serverId)
}
