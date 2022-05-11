package modules

import (
	"api/src/modules/auth"
	"api/src/modules/user"
)

/**
This function inits all modules.
Modules provide their services in IoC container.
Modules can be resolved by inner function Get<ModuleName>Module()
Modules services can be resolved via ioc.Pick[T](name string) T
*/
func InitModules() {
	um := user.UserModule{}
	um.Init()

	am := auth.AuthModule{}
	am.Init()
}
