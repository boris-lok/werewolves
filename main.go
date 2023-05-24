package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"werewolves/lib"
)

func main() {
	cfg, err := lib.GetConfiguration("configuration/base.yaml")
	if err != nil {
		log.Fatal("Can't get the configuration")
	}

	fmt.Printf("%v", cfg)

	roles := []lib.Role{lib.Wolf, lib.Wolf, lib.Seer, lib.Witch, lib.Civilian, lib.Civilian}
	lib.Shuffle(roles, rand.NewSource(time.Now().UnixNano()))
	fmt.Printf("%v", roles)
}
