package sankari

import (
	"github.com/joonazan/bossit/cyberdemon"
	"testing"
)

func TestTapaDemoni(t *testing.T) {
	d := cyberdemon.New()
	TapaDemoni(d)
	if !d.Kuollut() {
		t.Fatal("TapaDemoni:n jälkeen demoni oli yhä hengissä!")
	}
}
