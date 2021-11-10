package main

import (
	"fmt"
	"hexagonal-template/cmd/cmds"
)

func main() {
	fmt.Printf(`
    ___       ______    _____
   / _  \    |   _  \  |_   _|
  /  __  \   |   __ /   _ | _
 /__/  \__\  |__|      |_____|

   hexagonal-template %s, built with Go %s
 `, cmds.Version, cmds.GoVersion)
	cmds.Execute()
}
