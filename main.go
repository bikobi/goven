package main

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {

	// Variables that hold flags' values.
	var (
		numWords     int
		wordlistPath string
		separator    string
		pascalCase   bool
		showHelp     bool
	)

	// Define command-line flags.
	flag.IntVar(&numWords, "length", 4, "Number of words to generate")
	flag.StringVar(&wordlistPath, "wordlist", "", "Path to wordlist file")
	flag.StringVar(&separator, "separator", "-", "Separator between words")
	flag.BoolVar(&pascalCase, "pascalcase", false, "Use PascalCase instead of lowercase")
	flag.BoolVar(&showHelp, "help", false, "Show help message")

	// Parse command-line flags
	flag.Parse()

	if showHelp {
		// Print help message and exit
		flag.PrintDefaults()
		os.Exit(0)
	}

	if wordlistPath != "" {
		// Load wordlist from file
		userWordlist, err := loadWordlist(wordlistPath)
		if err != nil {
			fmt.Println("Error loading wordlist:", err)
			os.Exit(1)
		}

		wordlist = userWordlist
	}

	// Generate passphrase
	passphrase := generatePassphrase(numWords, wordlist, separator, pascalCase)

	// Print passphrase to stdout
	fmt.Println(passphrase)

}

// loadWordlist loads a list of words from a file at the specified path.
func loadWordlist(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		words = append(words, word)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

// generatePassphrase generates a passphrase with the specified number of words,
// using words from the provided wordlist separated by the specified separator.
// If pascalCase is true, the words will be capitalized.
func generatePassphrase(numWords int, wordlist []string, separator string, pascalCase bool) string {
	var words []string
	for i := 0; i < numWords; i++ {
		word := getWord(wordlist)
		if pascalCase {
			word = capitalize(word)
		}
		words = append(words, word)
	}

	return strings.Join(words, separator)
}

// getWord returns one truely random word from the provided wordlist.
func getWord(wordlist []string) string {
	max := big.NewInt(int64(len(wordlist)))

	r, err := rand.Int(rand.Reader, max)
	if err != nil {
		log.Fatal(err)
	}

	return wordlist[r.Int64()]
}

// capitalize takes a string and returns it Capitalized.
func capitalize(s string) string {
	return cases.Title(language.Und).String(s)
}
