package main

import (
	"fmt"
	"wiredemo/bootstrap"
)

func main() {
	//service.NormalMission()

	//WireVersionMission()

	UserController()
}

func WireVersionMission() {
	mission := bootstrap.InitMission("yong shi")
	mission.Start()

	mission2 := bootstrap.InitMission2("Dragon")
	mission2.Start()
}

func UserController() {
	userImpl := bootstrap.InitUserImpl()
	user := userImpl.GetByID(233)
	fmt.Printf("user: %v", user)
}
