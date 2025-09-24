//go:build wireinject

// go build时可传入一些选项，根据这个选项决定某些文件是否编译。
// wire工具只会处理有wireinject的文件，所以我们的wire.go文件要加上这个。生成的wire_gen.go是给我们来使用的，wire不需要处理，故有!wireinject。
// 注意build和package之间必须要有一个空行

package bootstrap

import (
	"wiredemo/service"
	"wiredemo/service/hunt_error"

	"github.com/google/wire"
)

// 初始化mission
// 此处通过wire初始化，对比mission_controller.NormalMission()

// 返回值：就是我们要创建的类型，wire只知道要创建的类型是什么即可，return返回的是什么不重要
func InitMission(name string) service.Mission {
	// 调用wire.Build()将创建Mission所依赖的类型的构造器传进去，NewMission所依赖的类型也要传进去
	wire.Build(service.NewMonster, service.NewPlayer, service.NewMission)
	return service.Mission{}
}

// 错误：
// wire遵循fail-fast的原则，错误必须被处理。如果我们的注入器不返回错误，但构造器返回错误，wire工具会报错！
func InitHunt(name string) (hunt_error.Hunt, error) {
	cat, err := hunt_error.NewCat(name)
	if err != nil {
		return hunt_error.Hunt{}, err
	}

	//wire.Build(cat, hunt_error.NewHunt)		// 包含错误处理的情况，无法使用wire.Build
	hunt := hunt_error.NewHunt(cat)
	return hunt, nil
}

//func InitRepository() *repository.UserRepository {
//	// 我只在这里声明我要用的各种东西，但是具体怎么构造，怎么编排顺序
//	// 这个方法里面传入各个组件的初始化方法
//	wire.Build(repository.InitDB, repository.NewUserRepository, dao.NewUserDAO)
//	return new(repository.UserRepository)
//}
