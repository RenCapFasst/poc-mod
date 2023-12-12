package pocMod

import (
	"fmt"
	"github.com/RenCapFasst/poc-mod/internal/env"
	"os"
)

func init() {
	_, ok := os.LookupEnv("FOO_BAR")
	if !ok {
		panic("env var FOO_BAR not present")
	}
	env.Env()
}
func Hello() {
	fmt.Printf("Hello, World! %s", env.Env().FooBar)
}
