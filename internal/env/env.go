package env

import (
	"os"
	"sync"
)

var env EnvType
var envOnce sync.Once

type EnvType struct {
	FooBar string
}

func Env() EnvType {
	envOnce.Do(func() {
		env = EnvType{FooBar: os.Getenv("FOO_BAR")}
	})
	return env
}
