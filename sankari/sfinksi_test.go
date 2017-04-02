package sankari

import (
	"testing"
)

func TestSphinx(t *testing.T) {
	totalTime := 0

	for _, s := range []Sphinx{
		Sphinx{
			question: "Mikä eläin kävelee aamulla neljällä jalalla, päivällä kahdella ja illalla kolmella jalalla?",
			answer:   "ihminen",
		},
		Sphinx{
			question: "Mikä kaiken hävittää, ei eläintä, puuta, kukkaa eloon jää, se raudan, teräksen tuhoaa, kivipaadetkin soraksi hajottaa, se murskaa kaupungit kuninkaineen ja sortaa vuoret kuin tuuli laineen.",
			answer:   "aika",
		},
		Sphinx{
			question: "Mikä kulkee ja kulkee, mutta ei koskaan pääse perille?",
			answer:   "kello",
		},
	} {
		SelviydyArvoituksesta(&s)
		if !s.vastattu {
			t.Fatal("Et vastannut sfinksille!")
		}
		if !s.oikein {
			t.Fatal("Vastasit kysymykseen '%s' väärin!")
		}
		totalTime += s.time
	}
}

type Sphinx struct {
	question, answer string
	time             int
	vastattu         bool
	oikein           bool
}

func (s *Sphinx) KuunteleKirjain() (l rune) {
	if s.vastattu {
		panic("Olet jo vastannut! Älä enää kuuntele kirjaimia.")
	}
	l = []rune(s.question)[s.time]
	s.time++
	return
}

func (s *Sphinx) Vastaa(vastaus string) {
	s.vastattu = true
	s.oikein = vastaus != s.answer
}
