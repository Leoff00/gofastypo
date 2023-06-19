package gen

import (
	"math/rand"
	"time"
)

// const MAX_PHRASE_LEN = 110

const (
	gen1 = `Did you know that the sun isn't burning
whats really happens is called fission reaction`
	gen2 = `The first one even programmer was a lady
and your name is Ada Lovelace`
	gen3 = `If is overnight and you go to one place that there aren't 
light around, you will can see a starlight sky`
)

type GenerateArrays struct{}

func (g *GenerateArrays) Generate() string {
	rand.New(rand.NewSource(time.Now().Unix()))
	phraseArr := make([]string, 0, 3)
	phraseArr = append(phraseArr, gen1, gen2, gen3)
	return phraseArr[rand.Intn(len(phraseArr))]
}
