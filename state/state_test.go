
package state

import (
	"testing"
)


func TestCheckState(t *testing.T) {
	wanted := "[kylling rev korn HS ---\\ \\--/---------------/---]"
	state := CheckState();
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}