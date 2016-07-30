package main

import (
	"github.com/antmanler/docker-machine-driver-aliyunecs/aliyunecs"
	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(aliyunecs.NewDriver("", ""))
}
