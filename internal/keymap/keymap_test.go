package keymap

import (
	"testing"
)

func TestDefaultKeyMapBindings(t *testing.T) {
	if len(DefaultKeyMap.Text.Submit.Keys()) == 0 {
		t.Error("Submit binding should have at least one key")
	}
	if len(DefaultKeyMap.Text.NewLine.Keys()) == 0 {
		t.Error("NewLine binding should have at least one key")
	}

	wantSubmitKeys := []string{"alt+enter"}
	gotSubmitKeys := DefaultKeyMap.Text.Submit.Keys()
	if len(gotSubmitKeys) != len(wantSubmitKeys) || gotSubmitKeys[0] != wantSubmitKeys[0] {
		t.Errorf("Submit keys = %v; want %v", gotSubmitKeys, wantSubmitKeys)
	}

	wantNewLineKeys := []string{"enter"}
	gotNewLineKeys := DefaultKeyMap.Text.NewLine.Keys()
	if len(gotNewLineKeys) != len(wantNewLineKeys) || gotNewLineKeys[0] != wantNewLineKeys[0] {
		t.Errorf("NewLine keys = %v; want %v", gotNewLineKeys, wantNewLineKeys)
	}
}
