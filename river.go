package main

import (
	"fmt"
	"github.com/EirikBHA/Rivercrossing/state"
)

func main() {
	fmt.Println(state.ViewState());
	state.PutInBoat();
	fmt.Println(state.ViewState());
	state.Crossriver();
	fmt.Println(state.ViewState());

}