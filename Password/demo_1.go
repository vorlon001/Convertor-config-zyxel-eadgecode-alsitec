package main

import (
	garbler "github.com/michaelbironneau/garbler/lib"
	"fmt"
)

func main() {

	p, _ := garbler.NewPassword(nil)
	fmt.Println(p)

	fmt.Printf(" %#v \n",garbler.Strong)
	fmt.Printf(" %#v \n",garbler.Paranoid)
	p, _ = garbler.NewPassword(&garbler.Strong ) //Strong)
	fmt.Println(p)

	paranoid := garbler.PasswordStrengthRequirements{ MinimumTotalLength: 32, Uppercase: 6, Digits: 12, Punctuation: 8, }
	
	p, _ = garbler.NewPassword( &paranoid ) //&garbler.Paranoid ) //Strong)
	fmt.Println(p)


	reqs := garbler.MakeRequirements("asdfGG11!")
	p, _ = garbler.NewPassword(&reqs)
	fmt.Println(p)

	reqs = garbler.PasswordStrengthRequirements{MinimumTotalLength: 20, Digits:10}
	p, e := garbler.NewPassword(&reqs)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(p)
}
