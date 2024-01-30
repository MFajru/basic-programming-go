package solution

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var dictionary = map[string]string{
	"makan": "eat",
}

func userInput() string {
	input := bufio.NewReader(os.Stdin)
	text, err := input.ReadString('\n')
	text = strings.TrimSpace(text)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return text
}

func translate() {

	fmt.Print("Word to translate: ")
	word := userInput()
	fmt.Println()

	wordIDNLower := strings.ToLower(word)
	wordIDNTitle := cases.Title(language.Indonesian).String(word)
	wordEng, exist := dictionary[wordIDNLower]

	if !exist {
		fmt.Printf("sorry, \"%s\" is not found in dictionary\n", wordIDNTitle)
		return
	}

	wordENGTitle := cases.Title(language.English).String(wordEng)

	fmt.Printf("ID: %s\n", wordIDNTitle)
	fmt.Printf("EN: %s\n", wordENGTitle)
}

func addWord() {
	fmt.Print("Word to be added in dict: ")
	word := userInput()

	if !strings.Contains(word, "#") {
		fmt.Println("word must be separated with # char")
		return
	}

	splittedString := strings.Split(word, "#")

	wordIDNLower := strings.ToLower(splittedString[0])
	wordENGLower := strings.ToLower(splittedString[1])

	_, exist := dictionary[wordIDNLower]

	if exist {
		fmt.Printf("cannot add existing word \"%s\"\n", wordIDNLower)
		return
	}

	dictionary[wordIDNLower] = wordENGLower
	fmt.Println("new word successfully added")
	fmt.Println()

}

func removeWord() {
	fmt.Print("Word to be removed: ")
	word := userInput()

	wordIDNLower := strings.ToLower(word)
	wordIDNTitle := cases.Title(language.Indonesian).String(wordIDNLower)

	_, exist := dictionary[wordIDNLower]

	if !exist {
		fmt.Printf("sorry, \"%s\" is not found in dictionary\n", wordIDNTitle)
		return
	}

	delete(dictionary, wordIDNLower)
	fmt.Printf("\"%s\" has been removed\n", wordIDNLower)
}

func sortKeyDict() []string {
	var keySlice []string
	for k := range dictionary {
		keySlice = append(keySlice, k)
	}

	return keySlice
}

func printDict() {
	sortedKey := sortKeyDict()
	fmt.Println("\nID to EN Dictionary:")
	for i := 0; i < len(sortedKey); i++ {
		wordENGTitle := cases.Title(language.English).String(dictionary[sortedKey[i]])
		fmt.Printf("%d. %s: %s\n", i+1, sortedKey[i], wordENGTitle)
	}

}

func DisplayMenu() {
	var menu int

	for {
		fmt.Println("ID to EN Dictionary")
		fmt.Println("Menu:")
		fmt.Println("1. Translate")
		fmt.Println("2. Add Word")
		fmt.Println("3. Remove Word")
		fmt.Println("4. Print Dictionary")
		fmt.Println("0. Exit")

		fmt.Print("Input: ")
		_, err := fmt.Scanln(&menu)

		switch menu {
		case 1:
			translate()
		case 2:
			addWord()
		case 3:
			removeWord()
		case 4:
			printDict()
		case 0:
			return
		default:
			fmt.Println(err)
		}
		fmt.Print("\nPress Enter key to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		fmt.Println()
	}
}
