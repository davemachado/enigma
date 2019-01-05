package rotor

import "strings"

var Alphabet = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

type Rotor struct {
	Name           string
	Input          []string
	Output         []string
	TurnoverLetter []string
}

type Reflector struct {
	Name   string
	Input  []string
	Output []string
}

func GetRotors() []*Rotor {
	rotors := []*Rotor{
		&Rotor{
			Name:           "1",
			Input:          strings.Split("EKMFLGDQVZNTOWYHXUSPAIBRCJ", ""),
			TurnoverLetter: []string{"Q"},
		},
		&Rotor{
			Name:           "2",
			Input:          strings.Split("AJDKSIRUXBLHWTMCQGZNPYFVOE", ""),
			TurnoverLetter: []string{"E"},
		},
		&Rotor{
			Name:           "3",
			Input:          strings.Split("BDFHJLCPRTXVZNYEIWGAKMUSQO", ""),
			TurnoverLetter: []string{"V"},
		},
		&Rotor{
			Name:           "4",
			Input:          strings.Split("ESOVPZJAYQUIRHXLNFTGKDCMWB", ""),
			TurnoverLetter: []string{"J"},
		},
		&Rotor{
			Name:           "5",
			Input:          strings.Split("VZBRGITYUPSDNHLXAWMJQOFECK", ""),
			TurnoverLetter: []string{"Z"},
		},
		&Rotor{
			Name:           "6",
			Input:          strings.Split("JPGVOUMFYQBENHZRDKASXLICTW", ""),
			TurnoverLetter: []string{"Z", "M"},
		},
		&Rotor{
			Name:           "7",
			Input:          strings.Split("NZJHGRCXMYSWBOUFAIVLPEKQDT", ""),
			TurnoverLetter: []string{"Z", "M"},
		},
		&Rotor{
			Name:           "8",
			Input:          strings.Split("FKQHTLXOCBJSPDZRAMEWNIUYGV", ""),
			TurnoverLetter: []string{"Z", "M"},
		},
	}
	for _, r := range rotors {
		var output []string
		for i := range r.Input {
			output = append(output, Alphabet[i])
		}
		r.Output = output
	}
	return rotors
}

func GetReflectors() []*Reflector {
	reflectors := []*Reflector{
		&Reflector{
			Name:  "A",
			Input: strings.Split("EJMZALYXVBWFCRQUONTSPIKHGD", ""),
		},
		&Reflector{
			Name:  "B",
			Input: strings.Split("YRUHQSLDPXNGOKMIEBFZCWVJAT", ""),
		},
		&Reflector{
			Name:  "C",
			Input: strings.Split("FVPJIAOYEDRZXWGCTKUQSBNMHL", ""),
		},
	}
	for _, r := range reflectors {
		var output []string
		for i := range r.Input {
			output = append(output, Alphabet[i])
		}
		r.Output = output
	}
	return reflectors
}

func (r *Rotor) GetContact(inputChar string, reverseFlow bool) string {
	var input, output []string
	if reverseFlow {
		input = r.Output
		output = r.Input
	} else {
		input = r.Input
		output = r.Output
	}
	inputCharCheck := strings.ToUpper(inputChar)
	for i, letter := range input {
		if letter == inputCharCheck {
			return output[i]
		}
	}
	panic("rotor input did not match alphabet output")
}

func (r *Rotor) Rotate() bool {
	idx := 0
	first := r.Output[0]
	for idx < len(r.Output) {
		r.Output[idx] = r.Output[(idx+1)%len(r.Output)]
		idx = idx + 1
	}
	r.Output[len(r.Output)-1] = first
	for _, l := range r.TurnoverLetter {
		if r.Output[0] == l {
			return true
		}
	}
	return false
}

func (r *Rotor) RotateBack() bool {
	idx := len(r.Output) - 1
	last := r.Output[idx]
	for idx > 0 {
		r.Output[idx] = r.Output[idx-1]
		idx--
	}
	r.Output[0] = last
	for _, l := range r.TurnoverLetter {
		if r.Output[0] == l {
			return true
		}
	}
	return false
}

func Click(rotors []*Rotor) {
	t := true
	for _, r := range rotors {
		if t {
			t = r.Rotate()
		}
	}
}

func ClickBack(rotors []*Rotor) {
	t := true
	for _, r := range rotors {
		if t {
			t = r.RotateBack()
		}
	}
}

func GetOutput(rotors []*Rotor, reflector *Reflector, input string) string {
	signal := input
	for _, r := range rotors {
		signal = r.GetContact(signal, false)
	}
	success := false
	for i, l := range reflector.Input {
		if l == signal {
			success = true
			signal = reflector.Output[i]
			break
		}
	}
	if !success {
		panic("missed an output match")
	}
	for i := len(rotors) - 1; i >= 0; i-- {
		signal = rotors[i].GetContact(signal, true)
	}
	return signal
}
