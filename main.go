package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/davemachado/enigma/rotor"
	"github.com/eiannone/keyboard"
)

func printStr(rotors []*rotor.Rotor, reflector *rotor.Reflector, clear bool) {
	if clear {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
	}
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
		printStr(rotors, reflector, true)
		fmt.Printf("\n%s\n", str)
		c, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyEsc {
			break
		}
		if key == keyboard.KeyBackspace2 {
			if len(str) == 0 {
				continue
			}
			rotor.ClickBack(rotors)
			str = str[:len(str)-1]
			continue
		}
		if key == keyboard.KeyEnter {
			str += "\n"
			rotor.Click(rotors)
			continue
		}
		if key == keyboard.KeySpace {
			str += " "
			rotor.Click(rotors)
			continue
		}
		if int(c) < 65 || int(c) > 122 {
			str += fmt.Sprintf("%s", string(c))
			rotor.Click(rotors)
			continue
		}
		str += rotor.GetOutput(rotors, reflector, string(c))
		rotor.Click(rotors)
	}
}
