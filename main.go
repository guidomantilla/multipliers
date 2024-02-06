package main

import (
	"multipliers/cmd"
	"multipliers/pkg/core"
)

func main() {
	for i := 1; i <= 100; i++ {
		println(core.GetMultiplierType(i))
	}
	cmd.ExecuteAppCmd()
}
