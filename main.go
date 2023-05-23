package main

import (
	"fmt"
	"math/rand"
	"time"
	"werewolves/lib"
)

func main() {
  roles := []lib.Role{lib.Wolf, lib.Wolf, lib.Seer, lib.Witch, lib.Civilian, lib.Civilian}
  lib.Shuffle(roles, rand.NewSource(time.Now().UnixNano()))
  fmt.Printf("%v", roles)
}
