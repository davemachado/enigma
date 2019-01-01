package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/davemachado/enigma/rotor"
	"github.com/eiannone/keyboard"
)

func printStr(rotors []*rotor.Rotor, reflector *rotor.Reflector) {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	for _, r := range rotors {
		fmt.Printf("%v\n", r.Input)
		fmt.Printf("%v\n", r.Output)
		fmt.Println()
	}
	fmt.Println()
	fmt.Printf("%v\n", reflector.Input)
	fmt.Printf("%v\n", reflector.Output)
	fmt.Println()
}

func main() {
	rotors := rotor.GetRotors()[:3]
	reflector := rotor.GetReflectors()[0]

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	var str string

	for {
		printStr(rotors, reflector)
		fmt.Printf("\n%s\n", str)
		c, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyEsc {
			break
		}
		if key == keyboard.KeySpace {
			str += " "
			continue
		}

		str += rotor.GetOutput(rotors, reflector, string(c))
		rotor.Click(rotors)
	}
}
