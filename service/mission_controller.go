package service

// 非依赖注入方式
func NormalMission() {
	monster := NewMonster()
	player := NewPlayer("dj")
	mission := NewMission(player, monster)

	mission.Start()
}
