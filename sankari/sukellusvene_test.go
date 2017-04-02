package sankari

import (
	"testing"
)

func TestUpotus(t *testing.T) {

	max_aika := 0

	for paikka := 1; paikka <= 100; paikka++ {
		löydetty := false
		aika := 0
		LöydäSukellusvene(func(arvaus int) int {
			aika++
			if arvaus < paikka {
				return 1
			} else if arvaus > paikka {
				return -1
			} else {
				löydetty = true
				return 0
			}
		})

		if löydetty == false {
			t.Fatalf("Näillä etsinnöillä sukellusvenettä ei löytyisi, jos se olisi laivan %d alla.\n", paikka)
		}

		if aika > max_aika {
			max_aika = aika
		}
	}

	if max_aika > 10 {
		t.Logf("Onnistuit, mutta jos se ei olisi ollut zombi, niin se olisi kyllä jo Atlantilla.")
	}

}
