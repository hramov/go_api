package ioc

import (
	"api/src/core/logger"

	"github.com/golobby/container/v3"
)

func Pick[T comparable](name string) T {
	var instance T
	if err := container.NamedResolve(&instance, name); err != nil {
		logger.Error(err.Error())
	}
	return instance
}

func Put[T comparable](name string, value T) {
	if err := container.NamedSingleton(name, func() T { return value }); err != nil {
		logger.Error(err.Error())
	}
}
