package main

import "wiredemo/bootstrap"

func main() {
	//service.NormalMission()

	WireVersionMission()
}

func WireVersionMission() {
	mission := bootstrap.InitMission("yong shi")
	mission.Start()

	mission2 := bootstrap.InitMission2("Dragon")
	mission2.Start()
}
