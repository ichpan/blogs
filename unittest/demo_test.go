package unittest

import (
	"testing"
)

func TestStore(t *testing.T) {
	monster := &Monster{
		"张三", 18, "跑步",
	}
	result := monster.Store()
	if !result {
		t.Fatalf("Error TestStore! %v", result)
	}
}

func TestReStore(t *testing.T) {
	monster := Monster{}
	result := monster.ReStore()
	if !result {
		t.Fatalf("Error TestReStore!")
	}
}
