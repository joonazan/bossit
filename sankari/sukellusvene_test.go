package sankari

import (
	"math/rand"
	"testing"
	"time"
)

func TestUpotus(t *testing.T) {
	löydetty := false
	rand.Seed(time.Now().UnixNano())
	paikka := rand.Intn(100) + 1
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
		t.Fatal("Näillä etsinnöillä sukellusvenettä ei löytynyt.")
	}
}
