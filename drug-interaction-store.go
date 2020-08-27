package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var severities = map[string]int{
	"major":    3,
	"moderate": 2,
	"minor":    1,
}

type Key struct {
	X, Y string
}

type Interaction struct {
	Drugs        []string `json:"drugs"`
	SeverityStr  string   `json:"severity"`
	Description  string   `json:"description"`
	Severity     int
	FilePosition int
}

type Interactions struct {
	list  *[]Interaction
	store *map[Key]int
}

func buildKey(drug1 string, drug2 string) Key {
	// key has drugs in lexographical order for easy lookups
	d1 := strings.ToLower(drug1)
	d2 := strings.ToLower(drug2)
	var key Key
	if d1 >= d2 {
		key = Key{d1, d2}
	} else {
		key = Key{d2, d1}
	}
	return key
}

func (interactions *Interactions) BuildStore(filePath string) error {
	// Open and read file into in-memory buffer
	fmt.Printf("Opening %v\n", filePath)
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Successfully Opened %v\n", filePath)
	defer jsonFile.Close()
	buf, _ := ioutil.ReadAll(jsonFile)

	// write memory buffer into desired data structures
	interactionList := make([]Interaction, 0)
	if err := json.Unmarshal(buf, &interactionList); err != nil {
		return err
	}

	// Update the in-memory store
	interactions.list = &interactionList
	listLen := len(*interactions.list)
	fmt.Printf("%d interactions read from file\n", listLen)
	store := make(map[Key]int, listLen)
	for i := 0; i < listLen; i++ {
		// Set severity for priority
		sev := strings.ToLower(interactionList[i].SeverityStr)
		interactionList[i].Severity = severities[sev]

		interactionList[i].FilePosition = i

		// key has drugs in lexographical order for easy lookups
		drug1 := interactionList[i].Drugs[0]
		drug2 := interactionList[i].Drugs[1]
		key := buildKey(drug1, drug2)
		store[key] = i
	}
	interactions.store = &store
	return nil
}

func (interactions *Interactions) BuildAndPrint(filePath string) error {
	if err := interactions.BuildStore(filePath); err != nil {
		return err
	}
	fmt.Printf("Listing %d interactions :\n", len(*interactions.list))
	for key, pos := range *interactions.store {
		interaction := (*interactions.list)[pos]
		fmt.Printf("Key: %v, ", key)
		fmt.Printf("Position: %d\n", pos)
		fmt.Printf("Drugs: %v, ", interaction.Drugs)
		fmt.Printf("Severity: %d-%v, ", interaction.Severity, interaction.SeverityStr)
		fmt.Printf("Description: %v\n", interaction.Description)
	}
	return nil
}

func (interactions *Interactions) FindDrugsImpact(drugs []string) (Interaction, error) {
	var finalInteraction = Interaction{}

	drugLen := len(drugs)
	if drugLen < 2 {
		return finalInteraction, errors.New("Insufficient drugs passed in")
	}

	for i := 0; i < drugLen-1; i++ {
		drug1 := drugs[i]
		for j := i + 1; j < drugLen; j++ {
			drug2 := drugs[j]
			key := buildKey(drug1, drug2)
			if pos, ok := (*interactions.store)[key]; ok {
				interaction := (*interactions.list)[pos]
				if finalInteraction.Severity < interaction.Severity {
					finalInteraction = interaction
				} else if finalInteraction.Severity == interaction.Severity && finalInteraction.FilePosition > interaction.FilePosition {
					finalInteraction = interaction
				}
			}
		}
	}
	return finalInteraction, nil
}

func GetImpactString(interaction Interaction) string {
	if interaction.Severity > 0 {
		return fmt.Sprintf("%v: %v", strings.ToUpper(interaction.SeverityStr), interaction.Description)
	} else {
		return "No interaction"
	}
}

func main() {
	const filePath = "./interactions.json"
	var interactions = Interactions{}
	interactions.BuildStore(filePath)

	fmt.Println("Please enter drug names separated by space to continue (^C to exit) .... ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		drugs := strings.Split(scanner.Text(), " ")
		interaction, _ := interactions.FindDrugsImpact(drugs)
		fmt.Println(GetImpactString(interaction))
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
