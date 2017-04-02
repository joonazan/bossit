package harjoitusmaalitaulu

type Harjoitusmaalitaulu interface {
	Lyö()
	Tervehdi()
	Katso() string
	Karkaa()

	Voitettu() bool
}

type ukko bool

func New() Harjoitusmaalitaulu {
	u := ukko(false)
	return &u
}

func (u ukko) Voitettu() bool {
	return bool(u)
}

func (u ukko) Katso() string {
	return "Oljella täytetty nukke. Se seisoo paikallaan. Mitäköhän se tekee seuraavaksi?"
}

func (u *ukko) Lyö() {
	*u = true
}

func (u *ukko) Tervehdi() {
	*u = true
}

func (u *ukko) Karkaa() {
	*u = true
}
