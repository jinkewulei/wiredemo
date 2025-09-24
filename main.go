package main

import "wiredemo/bootstrap"

func main() {
	//service.NormalMission()

	WireVersionMission()
}

func WireVersionMission() {
	mission, _ := bootstrap.InitMission("yong shi")
	mission.Start()
}

//func NormalInject() {
//	// 非依赖注入
//	db, err := gorm.Open(mysql.Open("dsn"))
//	if err != nil {
//		panic(err)
//	}
//	ud := dao.NewUserDAO(db)
//	repo := repository.NewUserRepository(ud)
//	fmt.Println(repo)
//
//	// 依赖注入
//	bootstrap.InitRepository()
//}
