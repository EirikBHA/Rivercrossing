
package state

import (
	"testing"
)


func TestViewState(t *testing.T) {
	wanted := "[kylling rev korn HS ---\\ \\--/---------------/---]"
	state := ViewState();
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}