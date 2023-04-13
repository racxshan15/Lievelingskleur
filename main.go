package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Color struct {
	Kleur   string `json:"name"`
	Gedicht string `json:"text"`
}

func main() {
	kleur := getInput("Wat is je favoriete kleur?")
	tekst, err := getTekst(kleur)
	if err != nil {
		handleError("Er is een fout opgetreden:", err, 1)
	}
	err = writeToFile("gedicht.txt", tekst)
	if err != nil {
		handleError("Kon de gedicht niet naar het bestand schrijven.", err, 2)
	}
	fmt.Println("Gedicht is geschreven naar gedicht.txt")
}

func getInput(vraag string) string {
	var inputkleur string
	fmt.Println(vraag)
	fmt.Scanln(&inputkleur)
	return inputkleur
}

func getTekst(kleur string) (string, error) {
	file, err := ioutil.ReadFile("tekst.json")
	if err != nil {
		return "", err
	}
	var kleuren []Color
	err = json.Unmarshal([]byte(file), &kleuren)

	for _, color := range kleuren {
		if color.Kleur == kleur {
			return color.Gedicht, nil
		}
	}
	return "", fmt.Errorf("Kleur is niet beschikbaar")
}

func writeToFile(filename string, inhoud string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(inhoud)
	if err != nil {
		return err
	}
	return nil
}

func handleError(boodschap string, err error, exitcode int) {
	fmt.Println(boodschap, err)
	os.Exit(exitcode)
}
