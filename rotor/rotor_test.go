package rotor

import "testing"

func TestGetContact(t *testing.T) {
	rotors := GetRotors()
	for i, r := range rotors {
		actual := r.GetContact(r.Input[i], false)
		expected := Alphabet[i]
		if actual != expected {
			t.Errorf("expected %s, got %s", expected, actual)
		}
	}
}

func TestGetOutput(t *testing.T) {
	rotors := GetRotors()[:3]
	reflector := GetReflectors()[0]
	if out := GetOutput(rotors, reflector, "X"); out != "F" {
		t.Errorf("expected F, got %s", out)
	}
}
