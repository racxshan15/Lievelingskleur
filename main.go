package main

import (
	"fmt"
	"os"
)

func main() {
	var kleur string
	fmt.Println("Wat is je favoriete kleur? ")
	fmt.Scanln(&kleur) //Hier wordt gescant wat je in de terminal typt. Wat je in de terminal typt wordt gezien als 'kleur'.
	zin, fout := getSentence(kleur)
	if fout != nil {
		fmt.Println("Er is een fout opgetreden:", fout)
	}

	err := writeToFile(zin)
	if err != nil {
		fmt.Println("couldn't write to file", err)
		os.Exit(1)
	}
}

func getSentence(kleur string) (string, error) {
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
	case "paars":
		tekst = "Paars als de nacht."
	case "oranje":
		tekst = "Oranje als de zonsondergang."
	case "wit":
		tekst = "Wit als de sneeuw."
	case "roze":
		tekst = "Roze als de bloesem."
	default: //Dit gebeurt er als je iets intypt wat niet bij de case hoort.
		//fmt.Println("Sorry, ik ken deze kleur niet.")
		//os.Exit(1) //Foutmelding.
		return "", fmt.Errorf("kleur %s is niet bekend", kleur) //%s is wat ik in de terminal invoer.
	}
	return tekst, nil
}

func writeToFile(sentence string) error {
	file, fout := os.Create("gedicht.txt") //Hier wordt met behulp van os. een textfile aangemaakt. En  de woorden: "file" en "fout" wordt eraan gekoppeld, zodat je er later naar kan verwijzen voor bijvoorbeeld een actie.
	if fout != nil {
		return fout //Als fout niet niks is, oftewel als er fouten inzitten print het een tekst met een foutcode.
		//fmt.Println("Kon het bestand niet maken.")
		//os.Exit(1)
	}
	defer file.Close() //Dit wordt automatisch aan het einde geplaatst als de code runt. Zodat de code altijd wordt afgesloten.

	_, fout = file.WriteString(sentence) //Hier wordt de text gewrite naar de file. De underscore is daar als een 'valse' waarde. Zodat het alleen focust op 'fout'.
	if fout != nil {                     //Als fout niet niks is, oftewel als er fouten inzitten print het een tekst met een foutcode.
		return fout
		//fmt.Println("Kon de tekst niet naar het bestand schrijven.")
		//os.Exit(1)
	}
	return nil
}
