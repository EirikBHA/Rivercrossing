package state

/*import (
	"testing"
)
*/
var state = "[kylling rev korn HS ---\\ \\--/---------------/---]"

func PutInBoat() {
	state = "[kylling rev korn ---\\ \\_HS_/-----------------/---]"
}

func Crossriver() {
	 state = "[kylling rev korn ---\\ \\------------------\\-HS-/---]"
}

func CheckState() string {
	return state
}