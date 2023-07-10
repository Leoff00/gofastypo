package gen

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

const BeginPhrase = `Select the minutage and click in the Start Button
to begin your test. Gotta GO Fastypo!!!`

type GeneratePhrases struct{}

func (g *GeneratePhrases) readPhrases() []string {
	path := "phrases.txt"
	file, err := os.Open(path)

	if err != nil {
		log.Fatal("File doesn't exist ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	phrase := ""

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		phrase += line + "\n"

		if len(line) > 0 && line[len(line)-1] == '.' {
			lines = append(lines, phrase)
			phrase = ""
		}
	}

	if err := scanner.Err(); err != nil {
		log.Default().Fatal("Probably cannot read the content of txt...", err)
	}

	return lines
}

func (g *GeneratePhrases) Generate() (string, []string) {
	rand.New(rand.NewSource(time.Now().Unix()))
	phraseArr := g.readPhrases()
	phrase := phraseArr[rand.Intn(len(phraseArr))]
	return phrase, phraseArr
}
