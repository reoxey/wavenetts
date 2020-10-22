// To use the Google Cloud API, you need to set the environment variable of GOOGLE_APPLICATION_CREDENTIALS.
// This is required for authentication before calling the Cloud API. You need to have an account of GCP to this code.

package wavenetts

import (
	"context"
	"os"

	cloudtts "cloud.google.com/go/texttospeech/apiv1"
	tts "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

// WithLang initialise Wave with language.
// Gender and voice will be randomly picked if
// more than 1 voices are available.
func WithLang(lang string) *Wave {
	return &Wave{
		Lang: lang,
		Rate: 1.0,
	}
}

// WithVoice initialise Wave with VoiceKey values as per Cloud TTS values.
// Specified voice will be picked.
func WithVoice(voice VoiceKey) *Wave {

	m := getVoice(voice)

	return &Wave{
		Voice: m.Name,
		Lang:  m.Code,
		Rate:  1.0,
	}
}

// WithLangGender initialise Wave with language and gender.
// voice will be picked randomly picked if more than 1 voices are available
// for the specified gender
func WithLangGender(lang string, gender tts.SsmlVoiceGender) *Wave {
	return &Wave{
		Lang:   lang,
		Rate:   1.0,
		Gender: gender,
	}
}

// SetLang sets language ISO 2 type
func (w *Wave) SetLang(l string) *Wave {
	w.Lang = l

	return w
}

// SetText sets text to be converted to audio.
// Make sure to keep it well punctuated, correct words and proper grammer.
func (w *Wave) SetText(t string) *Wave {
	w.Text = t

	return w
}

// SetGender sets the gender of the speaker
func (w *Wave) SetGender(g tts.SsmlVoiceGender) *Wave {
	w.Gender = g

	return w
}

// SetRate sets the speed of the voice of the speaker
func (w *Wave) SetRate(r float64) *Wave {
	w.Rate = r

	return w
}

// Fetch the response from the Cloud TTS server
// If the specified voice is not available then error will returned.
// Response will further used to retreive audio data.
func (w *Wave) Fetch() (Res, error) {

	ctx := context.Background()

	c, e := cloudtts.NewClient(ctx)
	if e != nil {
		return Res{}, e
	}

	var vSel *tts.VoiceSelectionParams
	if w.Voice != "" {
		vSel = &tts.VoiceSelectionParams{
			LanguageCode: w.Lang,
			Name:         w.Voice,
		}
	} else if w.Gender != 0 {
		vSel = &tts.VoiceSelectionParams{
			LanguageCode: w.Lang,
			SsmlGender:   w.Gender,
		}
	} else {
		vSel = &tts.VoiceSelectionParams{
			LanguageCode: w.Lang,
		}
	}

	req := &tts.SynthesizeSpeechRequest{
		Input: &tts.SynthesisInput{
			InputSource: &tts.SynthesisInput_Text{
				Text: w.Text,
			},
		},
		Voice: vSel,
		AudioConfig: &tts.AudioConfig{
			AudioEncoding: tts.AudioEncoding_MP3,
			SpeakingRate:  w.Rate,
		},
	}

	res, e := c.SynthesizeSpeech(ctx, req)

	return Res{res}, e

}

// Save will create an audio file in the specified path
// error will return if the operation failed.
func (r *Res) Save(path string) error {

	f, e := os.Create(path)

	if e != nil {
		return e
	}
	defer f.Close()

	_, e = f.Write(r.SynthesizeSpeechResponse.GetAudioContent())

	if e != nil {
		return e
	}

	return nil
}
