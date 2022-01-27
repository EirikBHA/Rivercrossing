package main

import (
	"fmt"
	"github.com/EirikBHA/Rivercrossing/state"
)

func main() {
	fmt.Println(state.CheckState())

	state.PutInBoat();
	fmt.Println(state.CheckState())

	state.Crossriver()
	fmt.Println(state.CheckState())

}