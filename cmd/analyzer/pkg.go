// +build js,wasm

package main

import (
	"fmt"
	"github.com/myitcvscratch/depanalysis/jsdep"
	"github.com/myitcvscratch/depanalysis/normdep"
)

func main() {
	fmt.Println(jsdep.Pi, normdep.Pi)
}
