package wavenetts_test

import (
	"testing"

	"github.com/reoxey/wavenetts"
)

func TestWithLang(t *testing.T) {

	w := wavenetts.WithLang(wavenetts.English_US)

	w.SetGender(wavenetts.FEMALE).SetRate(1.1)

	r, e := w.Fetch()
	if e != nil {
		t.Errorf("TestWithLang failed: %v", e)
	}

	if e = r.Save("withlang.mp3"); e != nil {
		t.Errorf("TestWithLang Save failed: %v", e)
	}
}

func TestWithVoice(t *testing.T) {

	w := wavenetts.WithVoice(wavenetts.English_Australia_FEMALE_2)

	w.SetRate(1.1)

	r, e := w.Fetch()
	if e != nil {
		t.Errorf("TestWithVoice failed: %v", e)
	}

	if e = r.Save("withvoice.mp3"); e != nil {
		t.Errorf("TestWithVoice Save failed: %v", e)
	}
}

func TestWithLangGender(t *testing.T) {

	w := wavenetts.WithLangGender(wavenetts.English_UK, wavenetts.FEMALE)

	w.SetRate(1.1)

	r, e := w.Fetch()
	if e != nil {
		t.Errorf("TestWithLangGender failed: %v", e)
	}

	if e = r.Save("withlangender.mp3"); e != nil {
		t.Errorf("TestWithLangGender Save failed: %v", e)
	}
}