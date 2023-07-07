package gen

import (
	"math/rand"
	"time"
)

const BeginPhrase = `Select the minutage and click in the Start Button
to begin your test. Gotta GO Fastypo!!!`

const (
	gen1 = `Did you know that the sun isn't burning
whats really happens is called fission reaction`
	gen2 = `The first one even programmer was a lady
and your name is Ada Lovelace`
	gen3 = `If is at night and you go to one place that there aren't
light around, you will see a starlight sky`
	gen4 = `The phenomenon known as "ball lightning" is a rare and	
unexplained occurrence where luminous spheres of electricity appear during thunderstorms`
	gen5 = `Did you know that the Earth's magnetic field is
constantly shifting and changing its polarity`
	gen6 = `The world's largest flower, the Rafflesia arnoldii, can grow up
to three feet in diameter and emits a foul odor to attract insects for pollination`
	gen7 = `The "smell of rain" that people often notice is caused by a compound
called geosmin, released by soil-dwelling bacteria when it rains`
	gen8 = `The world's oldest known living organism is a 5,000-year-old
tree named "Methuselah," located in the White Mountains of California.`
)

type GeneratePhrases struct{}

func (g *GeneratePhrases) Generate() (string, []string) {
	rand.New(rand.NewSource(time.Now().Unix()))
	phraseArr := make([]string, 0, 8)
	phraseArr = append(phraseArr, gen1, gen2, gen3, gen4, gen5, gen6, gen7, gen8)
	return phraseArr[rand.Intn(len(phraseArr))], phraseArr
}
