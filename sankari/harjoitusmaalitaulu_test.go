package sankari

import (
	"github.com/joonazan/bossit/harjoitusmaalitaulu"
	"testing"
)

func TestSparraus(t *testing.T) {
	hm := harjoitusmaalitaulu.New()
	Sparraa(hm)
	if !hm.Voitettu() {
		t.Fatal("Harjoitusmaalitaulu katsoo sinua odottavasti.")
	}
}
