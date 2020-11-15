package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	const i=9000234234
	p := message.NewPrinter(language.English)
	p.Printf("Hello: %d\n",i)

	v := message.NewPrinter(language.Russian)
	v.Printf("Hello: %d\n",i)
}
