package sankari

// Mikä eläin kävelee aamulla neljällä jalalla,
// päivällä kahdella
// ja illalla kolmella jalalla?
// - ihminen -- Oidipus
//
// Sfinksit ovat harvinaisen harmillisia olentoja. Jos niiden
// arvoituksiin vastaa väärin, tulee syödyksi. Onneksi tiedät,
// että ne osaavat vain kolme eri kysymystä. Ensimmäinen niistä
// näkyy yllä ja toiset kaksi alla.
//
// Mikä kulkee ja kulkee, mutta ei koskaan pääse perille?
// - kello
//
// Mikä kaiken hävittää,
// ei eläintä, puuta, kukkaa eloon jää,
// se raudan, teräksen tuhoaa,
// kivipaadetkin soraksi hajottaa,
// se murskaa kaupungit kuninkaineen
// ja sortaa vuoret kuin tuuli laineen.
// - aika
//
// Koska sfinksit ovat pelottavia, haluat tietää vastauksen mahdollisimman
// nopeasti. Onneksi et panikoi, vaan kuuntelet äärimmäisen tarkasti.

type Sfinksi interface {

	// odottaa, että Sfinksi sanoo yhden kirjaimen ja palauttaa sen kirjaimen.
	KuunteleKirjain() rune

	// vastaa Sfinksille.
	Vastaa(string)
}

func SelviydyArvoituksesta(sfinksi Sfinksi) {

}
