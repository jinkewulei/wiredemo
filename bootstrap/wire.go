//go:build wireinject

// go build时可传入一些选项，根据这个选项决定某些文件是否编译。
// wire工具只会处理有wireinject的文件，所以我们的wire.go文件要加上这个。生成的wire_gen.go是给我们来使用的，wire不需要处理，故有!wireinject。
// 注意build和package之间必须要有一个空行

package bootstrap

import (
	"wiredemo/service"
	"wiredemo/service/hunt_error"
	"wiredemo/service/impl"

	"github.com/google/wire"
)

// -----------------------------基础用法----------------------------- //
// 初始化mission
// 此处通过wire初始化，对比mission_controller.NormalMission()
// 返回值：就是我们要创建的类型，wire只知道要创建的类型是什么即可，return返回的是什么不重要
func InitMission(name string) service.Mission {
	// 调用wire.Build()将创建Mission所依赖的类型的构造器传进去，NewMission所依赖的类型也要传进去
	wire.Build(service.NewMonster, service.NewPlayer, service.NewMission)
	return service.Mission{}
}

// 错误处理：provider（构造器）中出现错误的处理
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

// -----------------------------基础用法----------------------------- //

// -----------------------------高阶用法----------------------------- //
// ProviderSet
// 上述用法，针对同一个provider被多处注入时，需要反复声明，因此有了简化版；集中声明
var ProviderSet = wire.NewSet(
	service.NewMonster,
	service.NewPlayer,
)

// 之后就可以批量使用这个ProviderSet
// 注意其中的string入参是传递给了player使用
func InitMission2(name string) service.Mission {
	wire.Build(ProviderSet, service.NewMission)
	return service.Mission{}
}

// Bind: 绑定接口和实现
// 多用于声明了持久层接口，再指定实现的场景
var ImplProvider = wire.NewSet(
	impl.NewUserImpl,
)

func InitUserImpl() {
	wire.Bind(new(impl.UserService), new(impl.UserImpl))
}

// -----------------------------高阶用法----------------------------- //
