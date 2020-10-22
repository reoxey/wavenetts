# wavenetts

Simple Text to Speech module which connects to Google Cloud TTS API and downloads the audio file based on the text provided. This module only works if you have an account on Google Cloud and Cloud TTS access credentials.

### Access with 3 ways
- **WithLang:** The gender and the voice of the speaker will be automatically selected by the Cloud TTS based on the specified language. 
- **WithLangGender:** The voice of the speaker will be selected based on specified gender and language.
- **WithVoice:** The specified voice will be used, some languages have more than 1 voice.

### Install 

`go get -v github.com/reoxey/wavenetts`

### Import
```go
import "github.com/reoxey/wavenetts"
```

### How to use

**It is required before using with Google cloud to set `GOOGLE_APPLICATION_CREDENTIALS` for authentication.**

```go
func main() {

	log.SetFlags(log.LstdFlags|log.Lshortfile)

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/my/cloud/auth.json")


	text := "What are you doing here?"

//	tts := wavenetts.WithLang(wavenetts.English_Australia).SetText(text)
//  tts := wavenetts.WithLangGender(wavenetts.English_US, wavenetts.FEMALE).SetText(text)

	tts := wavenetts.WithVoice(wavenetts.English_UK_FEMALE_2).SetText(text).SetRate(1.1)

	r, e := tts.Fetch()

	if e != nil {
		log.Fatal(e)
	}

	e = r.Save("audio.mp3")

	if e != nil {
		log.Fatal(e)
	}
	
}
```

### TODO

Done? | Features
:---:| ---
✅| Save Audio
✅| Google Cloud TTS
⬜️| Play if mp3 codecs supported
