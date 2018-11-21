package main

import (
	"github.com/dayan-be/id-srv-snow/logic"
	"github.com/dayan-be/id-srv-snow/proto"
	"github.com/dayan-be/kit"
)


func main(){
	kit.Init()
	platform_id_srv_snow.RegisterSnowHandler(kit.DefaultService.Server(),&logic.Handle{})
	kit.Run()
}