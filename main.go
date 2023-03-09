package main

import (
	"fmt"
	"os"
)

func main() {
	var kleur string
	fmt.Println("Wat is je favoriete kleur? ")
	fmt.Scanln(&kleur) //Hier wordt gescant wat je in de terminal typt. Wat je in de terminal typt wordt gezien als 'kleur'.
	code(kleur)
}

func code(kleur string) {
	var tekst string //String voor de tekst wat in de textfile komt.
	switch kleur {   //Een switch statement voor de verschillende mogelijkheden.
	case "rood":
		tekst = "Rood met passie."
	case "blauw":
		tekst = "Blauw zoals de lucht."
	case "groen":
		tekst = "Groen van de natuur."
	case "geel":
		tekst = "Geel als de stralen van de zon."
	default: //Dit gebeurt er als je iets intypt wat niet bij de case hoort.
		fmt.Println("Sorry, ik ken deze kleur niet.")
		os.Exit(1) //Foutmelding.
	}

	file, fout := os.Create("gedicht.txt") //Hier wordt met behulp van os. een textfile aangemaakt. En  de woorden: "file" en "fout" wordt eraan gekoppeld, zodat je er later naar kan verwijzen voor bijvoorbeeld een actie.
	if fout != nil {                       //Als fout niet niks is, oftewel als er fouten inzitten print het een tekst met een foutcode.
		fmt.Println("Kon het bestand niet maken.")
		os.Exit(1)
	}
	defer file.Close() //Dit wordt automatisch aan het einde geplaatst als de code runt. Zodat de code altijd wordt afgesloten.

	_, fout = file.WriteString(tekst) //Hier wordt de text gewrite naar de file. De underscore is daar als een 'valse' waarde. Zodat het alleen focust op 'fout'.
	if fout != nil {                  //Als fout niet niks is, oftewel als er fouten inzitten print het een tekst met een foutcode.
		fmt.Println("Kon de tekst niet naar het bestand schrijven.")
		os.Exit(1)
	}

	fmt.Printf("Gedicht geschreven naar gedicht.txt.\n") //Als het ingevulde woord bij de switch case staat, gebeurt dit.
}
