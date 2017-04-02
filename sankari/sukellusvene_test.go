package sankari

import (
	"testing"
)

func TestUpotus(t *testing.T) {

	for paikka := 1; paikka <= 100; paikka++ {
		löydetty := false
		LöydäSukellusvene(func(arvaus int) int {
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
			t.Fatalf("Näillä etsinnöillä sukellusvenettä ei löytyisi jos se olisi laivan %d alla.\n", paikka)
		}
	}
}
