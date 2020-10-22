package wavenetts

import (
	tts "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

// SsmlGender
const (
	MALE    tts.SsmlVoiceGender = 1
	FEMALE  tts.SsmlVoiceGender = 2
	NEUTRAL tts.SsmlVoiceGender = 3
)

// Wave struct
type Wave struct {

	// Voice predefined voice string as per Wavenet doc
	Voice string

	// Lang language - country ISO 2 type
	// eg. en-GB
	Lang string

	// Text to be converted to audio file.
	// Make sure to keep it well punctuated, correct words and proper grammer.
	Text string

	// Rate at which the speaker will speak it.
	// increase the value to increase speed
	Rate float64

	// Gender of the speaker.
	Gender tts.SsmlVoiceGender

	// Pitch  float64 //TODO
}

//Res Response from the Cloud API
type Res struct {
	*tts.SynthesizeSpeechResponse
}

// Language values
const (
	Arabic               = "ar-XA"
	Czech_Czech_Republic = "cs-CZ"
	Danish_Denmark       = "da-DK"
	Dutch_Netherlands    = "nl-NL"
	English_Australia    = "en-AU"
	English_India        = "en-IN"
	English_UK           = "en-GB"
	English_US           = "en-US"
	Filipino_Philippines = "fil-PH"
	Finnish_Finland      = "fi-FI"
	French_Canada        = "fr-CA"
	French_France        = "fr-FR"
	German_Germany       = "de-DE"
	Greek_Greece         = "el-GR"
	Hindi_India          = "hi-IN"
	Hungarian_Hungary    = "hu-HU"
	Indonesian_Indonesia = "id-ID"
	Italian_Italy        = "it-IT"
	Japanese_Japan       = "ja-JP"
	Korean_South_Korea   = "ko-KR"
	Mandarin_China       = "cmn-CN"
	Mandarin_Taiwan      = "cmn-TW"
	Norwegian_Norway     = "nb-NO"
	Polish_Poland        = "pl-PL"
	Portuguese_Brazil    = "pt-BR"
	Portuguese_Portugal  = "pt-PT"
	Russian_Russia       = "ru-RU"
	Slovak_Slovakia      = "sk-SK"
	Swedish_Sweden       = "sv-SE"
	Turkish_Turkey       = "tr-TR"
	Ukrainian_Ukraine    = "uk-UA"
	Vietnamese_Vietnam   = "vi-VN"
)

// VoiceKey type used to set predefined voice values
type VoiceKey string

// Voice values
const (
	Arabic_FEMALE_1               VoiceKey = "Arabic_FEMALE_1"
	Arabic_MALE_1                 VoiceKey = "Arabic_MALE_1"
	Arabic_MALE_2                 VoiceKey = "Arabic_MALE_2"
	Czech_Czech_Republic_FEMALE_1 VoiceKey = "Czech_Czech_Republic_FEMALE_1"
	Danish_Denmark_FEMALE_1       VoiceKey = "Danish_Denmark_FEMALE_1"
	Danish_Denmark_MALE_1         VoiceKey = "Danish_Denmark_MALE_1"
	Danish_Denmark_FEMALE_2       VoiceKey = "Danish_Denmark_FEMALE_2"
	Danish_Denmark_FEMALE_3       VoiceKey = "Danish_Denmark_FEMALE_3"
	Dutch_Netherlands_FEMALE_1    VoiceKey = "Dutch_Netherlands_FEMALE_1"
	Dutch_Netherlands_MALE_1      VoiceKey = "Dutch_Netherlands_MALE_1"
	Dutch_Netherlands_MALE_2      VoiceKey = "Dutch_Netherlands_MALE_2"
	Dutch_Netherlands_FEMALE_2    VoiceKey = "Dutch_Netherlands_FEMALE_2"
	Dutch_Netherlands_FEMALE_3    VoiceKey = "Dutch_Netherlands_FEMALE_3"
	English_Australia_FEMALE_1    VoiceKey = "English_Australia_FEMALE_1"
	English_Australia_MALE_1      VoiceKey = "English_Australia_MALE_1"
	English_Australia_FEMALE_2    VoiceKey = "English_Australia_FEMALE_2"
	English_Australia_MALE_2      VoiceKey = "English_Australia_MALE_2"
	English_India_FEMALE_1        VoiceKey = "English_India_FEMALE_1"
	English_India_MALE_1          VoiceKey = "English_India_MALE_1"
	English_India_MALE_2          VoiceKey = "English_India_MALE_2"
	English_India_FEMALE_2        VoiceKey = "English_India_FEMALE_2"
	English_UK_FEMALE_1           VoiceKey = "English_UK_FEMALE_1"
	English_UK_MALE_1             VoiceKey = "English_UK_MALE_1"
	English_UK_FEMALE_2           VoiceKey = "English_UK_FEMALE_2"
	English_UK_MALE_2             VoiceKey = "English_UK_MALE_2"
	English_UK_FEMALE_3           VoiceKey = "English_UK_FEMALE_3"
	English_US_MALE_1             VoiceKey = "English_US_MALE_1"
	English_US_MALE_2             VoiceKey = "English_US_MALE_2"
	English_US_FEMALE_1           VoiceKey = "English_US_FEMALE_1"
	English_US_MALE_3             VoiceKey = "English_US_MALE_3"
	English_US_FEMALE_2           VoiceKey = "English_US_FEMALE_2"
	English_US_FEMALE_3           VoiceKey = "English_US_FEMALE_3"
	English_US_FEMALE_4           VoiceKey = "English_US_FEMALE_4"
	English_US_FEMALE_5           VoiceKey = "English_US_FEMALE_5"
	English_US_MALE_4             VoiceKey = "English_US_MALE_4"
	English_US_MALE_5             VoiceKey = "English_US_MALE_5"
	Filipino_Philippines_FEMALE_1 VoiceKey = "Filipino_Philippines_FEMALE_1"
	Filipino_Philippines_FEMALE_2 VoiceKey = "Filipino_Philippines_FEMALE_2"
	Filipino_Philippines_MALE_1   VoiceKey = "Filipino_Philippines_MALE_1"
	Filipino_Philippines_MALE_2   VoiceKey = "Filipino_Philippines_MALE_2"
	Finnish_Finland_FEMALE_1      VoiceKey = "Finnish_Finland_FEMALE_1"
	French_Canada_FEMALE_1        VoiceKey = "French_Canada_FEMALE_1"
	French_Canada_MALE_1          VoiceKey = "French_Canada_MALE_1"
	French_Canada_FEMALE_2        VoiceKey = "French_Canada_FEMALE_2"
	French_Canada_MALE_2          VoiceKey = "French_Canada_MALE_2"
	French_France_FEMALE_1        VoiceKey = "French_France_FEMALE_1"
	French_France_MALE_1          VoiceKey = "French_France_MALE_1"
	French_France_FEMALE_2        VoiceKey = "French_France_FEMALE_2"
	French_France_MALE_2          VoiceKey = "French_France_MALE_2"
	French_France_FEMALE_3        VoiceKey = "French_France_FEMALE_3"
	German_Germany_FEMALE_1       VoiceKey = "German_Germany_FEMALE_1"
	German_Germany_MALE_1         VoiceKey = "German_Germany_MALE_1"
	German_Germany_FEMALE_2       VoiceKey = "German_Germany_FEMALE_2"
	German_Germany_MALE_2         VoiceKey = "German_Germany_MALE_2"
	German_Germany_MALE_3         VoiceKey = "German_Germany_MALE_3"
	German_Germany_FEMALE_3       VoiceKey = "German_Germany_FEMALE_3"
	Greek_Greece_FEMALE_1         VoiceKey = "Greek_Greece_FEMALE_1"
	Hindi_India_FEMALE_1          VoiceKey = "Hindi_India_FEMALE_1"
	Hindi_India_MALE_1            VoiceKey = "Hindi_India_MALE_1"
	Hindi_India_MALE_2            VoiceKey = "Hindi_India_MALE_2"
	Hindi_India_FEMALE_2          VoiceKey = "Hindi_India_FEMALE_2"
	Hungarian_Hungary_FEMALE_1    VoiceKey = "Hungarian_Hungary_FEMALE_1"
	Indonesian_Indonesia_FEMALE_1 VoiceKey = "Indonesian_Indonesia_FEMALE_1"
	Indonesian_Indonesia_MALE_1   VoiceKey = "Indonesian_Indonesia_MALE_1"
	Indonesian_Indonesia_MALE_2   VoiceKey = "Indonesian_Indonesia_MALE_2"
	Indonesian_Indonesia_FEMALE_2 VoiceKey = "Indonesian_Indonesia_FEMALE_2"
	Italian_Italy_FEMALE_1        VoiceKey = "Italian_Italy_FEMALE_1"
	Italian_Italy_FEMALE_2        VoiceKey = "Italian_Italy_FEMALE_2"
	Italian_Italy_MALE_1          VoiceKey = "Italian_Italy_MALE_1"
	Italian_Italy_MALE_2          VoiceKey = "Italian_Italy_MALE_2"
	Japanese_Japan_FEMALE_1       VoiceKey = "Japanese_Japan_FEMALE_1"
	Japanese_Japan_FEMALE_2       VoiceKey = "Japanese_Japan_FEMALE_2"
	Japanese_Japan_MALE_1         VoiceKey = "Japanese_Japan_MALE_1"
	Japanese_Japan_MALE_2         VoiceKey = "Japanese_Japan_MALE_2"
	Korean_South_Korea_FEMALE_1   VoiceKey = "Korean_South_Korea_FEMALE_1"
	Korean_South_Korea_FEMALE_2   VoiceKey = "Korean_South_Korea_FEMALE_2"
	Korean_South_Korea_MALE_1     VoiceKey = "Korean_South_Korea_MALE_1"
	Korean_South_Korea_MALE_2     VoiceKey = "Korean_South_Korea_MALE_2"
	Mandarin_Chinese_FEMALE_1     VoiceKey = "Mandarin_Chinese_FEMALE_1"
	Mandarin_Chinese_MALE_1       VoiceKey = "Mandarin_Chinese_MALE_1"
	Mandarin_Chinese_MALE_2       VoiceKey = "Mandarin_Chinese_MALE_2"
	Mandarin_Chinese_FEMALE_2     VoiceKey = "Mandarin_Chinese_FEMALE_2"
	Mandarin_Taiwan_FEMALE_1      VoiceKey = "Mandarin_Taiwan_FEMALE_1"
	Mandarin_Taiwan_MALE_1        VoiceKey = "Mandarin_Taiwan_MALE_1"
	Mandarin_Taiwan_MALE_2        VoiceKey = "Mandarin_Taiwan_MALE_2"
	Norwegian_Norway_FEMALE_1     VoiceKey = "Norwegian_Norway_FEMALE_1"
	Norwegian_Norway_MALE_1       VoiceKey = "Norwegian_Norway_MALE_1"
	Norwegian_Norway_FEMALE_2     VoiceKey = "Norwegian_Norway_FEMALE_2"
	Norwegian_Norway_MALE_2       VoiceKey = "Norwegian_Norway_MALE_2"
	Norwegian_Norway_FEMALE_3     VoiceKey = "Norwegian_Norway_FEMALE_3"
	Polish_Poland_FEMALE_1        VoiceKey = "Polish_Poland_FEMALE_1"
	Polish_Poland_MALE_1          VoiceKey = "Polish_Poland_MALE_1"
	Polish_Poland_MALE_2          VoiceKey = "Polish_Poland_MALE_2"
	Polish_Poland_FEMALE_2        VoiceKey = "Polish_Poland_FEMALE_2"
	Polish_Poland_FEMALE_3        VoiceKey = "Polish_Poland_FEMALE_3"
	Portuguese_Brazil_FEMALE_1    VoiceKey = "Portuguese_Brazil_FEMALE_1"
	Portuguese_Portugal_FEMALE_1  VoiceKey = "Portuguese_Portugal_FEMALE_1"
	Portuguese_Portugal_MALE_1    VoiceKey = "Portuguese_Portugal_MALE_1"
	Portuguese_Portugal_MALE_2    VoiceKey = "Portuguese_Portugal_MALE_2"
	Portuguese_Portugal_FEMALE_2  VoiceKey = "Portuguese_Portugal_FEMALE_2"
	Russian_Russia_FEMALE_1       VoiceKey = "Russian_Russia_FEMALE_1"
	Russian_Russia_MALE_1         VoiceKey = "Russian_Russia_MALE_1"
	Russian_Russia_FEMALE_2       VoiceKey = "Russian_Russia_FEMALE_2"
	Russian_Russia_MALE_2         VoiceKey = "Russian_Russia_MALE_2"
	Russian_Russia_FEMALE_3       VoiceKey = "Russian_Russia_FEMALE_3"
	Slovak_Slovakia_FEMALE_1      VoiceKey = "Slovak_Slovakia_FEMALE_1"
	Swedish_Sweden_FEMALE_1       VoiceKey = "Swedish_Sweden_FEMALE_1"
	Turkish_Turkey_FEMALE_1       VoiceKey = "Turkish_Turkey_FEMALE_1"
	Turkish_Turkey_MALE_1         VoiceKey = "Turkish_Turkey_MALE_1"
	Turkish_Turkey_FEMALE_2       VoiceKey = "Turkish_Turkey_FEMALE_2"
	Turkish_Turkey_FEMALE_3       VoiceKey = "Turkish_Turkey_FEMALE_3"
	Turkish_Turkey_MALE_2         VoiceKey = "Turkish_Turkey_MALE_2"
	Ukrainian_Ukraine_FEMALE_1    VoiceKey = "Ukrainian_Ukraine_FEMALE_1"
	Vietnamese_Vietnam_FEMALE_1   VoiceKey = "Vietnamese_Vietnam_FEMALE_1"
	Vietnamese_Vietnam_MALE_1     VoiceKey = "Vietnamese_Vietnam_MALE_1"
	Vietnamese_Vietnam_FEMALE_2   VoiceKey = "Vietnamese_Vietnam_FEMALE_2"
	Vietnamese_Vietnam_MALE_2     VoiceKey = "Vietnamese_Vietnam_MALE_2"
)

type voice struct {
	Code string
	Name string
}

// getVoice get values like language and voice from map
func getVoice(v VoiceKey) voice {

	m := map[VoiceKey]voice{}

	m["Arabic_FEMALE_1"] = voice{"ar-XA", "ar-XA-Wavenet-A"}
	m["Arabic_MALE_1"] = voice{"ar-XA", "ar-XA-Wavenet-B"}
	m["Arabic_MALE_2"] = voice{"ar-XA", "ar-XA-Wavenet-C"}
	m["Czech_Czech_Republic_FEMALE_1"] = voice{"cs-CZ", "cs-CZ-Wavenet-A"}
	m["Danish_Denmark_FEMALE_1"] = voice{"da-DK", "da-DK-Wavenet-A"}
	m["Danish_Denmark_MALE_1"] = voice{"da-DK", "da-DK-Wavenet-C"}
	m["Danish_Denmark_FEMALE_2"] = voice{"da-DK", "da-DK-Wavenet-D"}
	m["Danish_Denmark_FEMALE_3"] = voice{"da-DK", "da-DK-Wavenet-E"}
	m["Dutch_Netherlands_FEMALE_1"] = voice{"nl-NL", "nl-NL-Wavenet-A"}
	m["Dutch_Netherlands_MALE_1"] = voice{"nl-NL", "nl-NL-Wavenet-B"}
	m["Dutch_Netherlands_MALE_2"] = voice{"nl-NL", "nl-NL-Wavenet-C"}
	m["Dutch_Netherlands_FEMALE_2"] = voice{"nl-NL", "nl-NL-Wavenet-D"}
	m["Dutch_Netherlands_FEMALE_3"] = voice{"nl-NL", "nl-NL-Wavenet-E"}
	m["English_Australia_FEMALE_1"] = voice{"en-AU", "en-AU-Wavenet-A"}
	m["English_Australia_MALE_1"] = voice{"en-AU", "en-AU-Wavenet-B"}
	m["English_Australia_FEMALE_2"] = voice{"en-AU", "en-AU-Wavenet-C"}
	m["English_Australia_MALE_2"] = voice{"en-AU", "en-AU-Wavenet-D"}
	m["English_India_FEMALE_1"] = voice{"en-IN", "en-IN-Wavenet-A"}
	m["English_India_MALE_1"] = voice{"en-IN", "en-IN-Wavenet-B"}
	m["English_India_MALE_2"] = voice{"en-IN", "en-IN-Wavenet-C"}
	m["English_India_FEMALE_2"] = voice{"en-IN", "en-IN-Wavenet-D"}
	m["English_UK_FEMALE_1"] = voice{"en-GB", "en-GB-Wavenet-A"}
	m["English_UK_MALE_1"] = voice{"en-GB", "en-GB-Wavenet-B"}
	m["English_UK_FEMALE_2"] = voice{"en-GB", "en-GB-Wavenet-C"}
	m["English_UK_MALE_2"] = voice{"en-GB", "en-GB-Wavenet-D"}
	m["English_UK_FEMALE_3"] = voice{"en-GB", "en-GB-Wavenet-F"}
	m["English_US_MALE_1"] = voice{"en-US", "en-US-Wavenet-A"}
	m["English_US_MALE_2"] = voice{"en-US", "en-US-Wavenet-B"}
	m["English_US_FEMALE_1"] = voice{"en-US", "en-US-Wavenet-C"}
	m["English_US_MALE_3"] = voice{"en-US", "en-US-Wavenet-D"}
	m["English_US_FEMALE_2"] = voice{"en-US", "en-US-Wavenet-E"}
	m["English_US_FEMALE_3"] = voice{"en-US", "en-US-Wavenet-F"}
	m["English_US_FEMALE_4"] = voice{"en-US", "en-US-Wavenet-G"}
	m["English_US_FEMALE_5"] = voice{"en-US", "en-US-Wavenet-H"}
	m["English_US_MALE_4"] = voice{"en-US", "en-US-Wavenet-I"}
	m["English_US_MALE_5"] = voice{"en-US", "en-US-Wavenet-J"}
	m["Filipino_Philippines_FEMALE_1"] = voice{"fil-PH", "fil-PH-Wavenet-A"}
	m["Filipino_Philippines_FEMALE_2"] = voice{"fil-PH", "fil-PH-Wavenet-B"}
	m["Filipino_Philippines_MALE_1"] = voice{"fil-PH", "fil-PH-Wavenet-C"}
	m["Filipino_Philippines_MALE_2"] = voice{"fil-PH", "fil-PH-Wavenet-D"}
	m["Finnish_Finland_FEMALE_1"] = voice{"fi-FI", "fi-FI-Wavenet-A"}
	m["French_Canada_FEMALE_1"] = voice{"fr-CA", "fr-CA-Wavenet-A"}
	m["French_Canada_MALE_1"] = voice{"fr-CA", "fr-CA-Wavenet-B"}
	m["French_Canada_FEMALE_2"] = voice{"fr-CA", "fr-CA-Wavenet-C"}
	m["French_Canada_MALE_2"] = voice{"fr-CA", "fr-CA-Wavenet-D"}
	m["French_France_FEMALE_1"] = voice{"fr-FR", "fr-FR-Wavenet-A"}
	m["French_France_MALE_1"] = voice{"fr-FR", "fr-FR-Wavenet-B"}
	m["French_France_FEMALE_2"] = voice{"fr-FR", "fr-FR-Wavenet-C"}
	m["French_France_MALE_2"] = voice{"fr-FR", "fr-FR-Wavenet-D"}
	m["French_France_FEMALE_3"] = voice{"fr-FR", "fr-FR-Wavenet-E"}
	m["German_Germany_FEMALE_1"] = voice{"de-DE", "de-DE-Wavenet-A"}
	m["German_Germany_MALE_1"] = voice{"de-DE", "de-DE-Wavenet-B"}
	m["German_Germany_FEMALE_2"] = voice{"de-DE", "de-DE-Wavenet-C"}
	m["German_Germany_MALE_2"] = voice{"de-DE", "de-DE-Wavenet-D"}
	m["German_Germany_MALE_3"] = voice{"de-DE", "de-DE-Wavenet-E"}
	m["German_Germany_FEMALE_3"] = voice{"de-DE", "de-DE-Wavenet-F"}
	m["Greek_Greece_FEMALE_1"] = voice{"el-GR", "el-GR-Wavenet-A"}
	m["Hindi_India_FEMALE_1"] = voice{"hi-IN", "hi-IN-Wavenet-A"}
	m["Hindi_India_MALE_1"] = voice{"hi-IN", "hi-IN-Wavenet-B"}
	m["Hindi_India_MALE_2"] = voice{"hi-IN", "hi-IN-Wavenet-C"}
	m["Hindi_India_FEMALE_2"] = voice{"hi-IN", "hi-IN-Wavenet-D"}
	m["Hungarian_Hungary_FEMALE_1"] = voice{"hu-HU", "hu-HU-Wavenet-A"}
	m["Indonesian_Indonesia_FEMALE_1"] = voice{"id-ID", "id-ID-Wavenet-A"}
	m["Indonesian_Indonesia_MALE_1"] = voice{"id-ID", "id-ID-Wavenet-B"}
	m["Indonesian_Indonesia_MALE_2"] = voice{"id-ID", "id-ID-Wavenet-C"}
	m["Indonesian_Indonesia_FEMALE_2"] = voice{"id-ID", "id-ID-Wavenet-D"}
	m["Italian_Italy_FEMALE_1"] = voice{"it-IT", "it-IT-Wavenet-A"}
	m["Italian_Italy_FEMALE_2"] = voice{"it-IT", "it-IT-Wavenet-B"}
	m["Italian_Italy_MALE_1"] = voice{"it-IT", "it-IT-Wavenet-C"}
	m["Italian_Italy_MALE_2"] = voice{"it-IT", "it-IT-Wavenet-D"}
	m["Japanese_Japan_FEMALE_1"] = voice{"ja-JP", "ja-JP-Wavenet-A"}
	m["Japanese_Japan_FEMALE_2"] = voice{"ja-JP", "ja-JP-Wavenet-B"}
	m["Japanese_Japan_MALE_1"] = voice{"ja-JP", "ja-JP-Wavenet-C"}
	m["Japanese_Japan_MALE_2"] = voice{"ja-JP", "ja-JP-Wavenet-D"}
	m["Korean_South_Korea_FEMALE_1"] = voice{"ko-KR", "ko-KR-Wavenet-A"}
	m["Korean_South_Korea_FEMALE_2"] = voice{"ko-KR", "ko-KR-Wavenet-B"}
	m["Korean_South_Korea_MALE_1"] = voice{"ko-KR", "ko-KR-Wavenet-C"}
	m["Korean_South_Korea_MALE_2"] = voice{"ko-KR", "ko-KR-Wavenet-D"}
	m["Mandarin_Chinese_FEMALE_1"] = voice{"cmn-CN", "cmn-CN-Wavenet-A"}
	m["Mandarin_Chinese_MALE_1"] = voice{"cmn-CN", "cmn-CN-Wavenet-B"}
	m["Mandarin_Chinese_MALE_2"] = voice{"cmn-CN", "cmn-CN-Wavenet-C"}
	m["Mandarin_Chinese_FEMALE_2"] = voice{"cmn-CN", "cmn-CN-Wavenet-D"}
	m["Mandarin_Taiwan_FEMALE_1"] = voice{"cmn-TW", "cmn-TW-Wavenet-A"}
	m["Mandarin_Taiwan_MALE_1"] = voice{"cmn-TW", "cmn-TW-Wavenet-B"}
	m["Mandarin_Taiwan_MALE_2"] = voice{"cmn-TW", "cmn-TW-Wavenet-C"}
	m["Norwegian_Norway_FEMALE_1"] = voice{"nb-NO", "nb-NO-Wavenet-A"}
	m["Norwegian_Norway_MALE_1"] = voice{"nb-NO", "nb-NO-Wavenet-B"}
	m["Norwegian_Norway_FEMALE_2"] = voice{"nb-NO", "nb-NO-Wavenet-C"}
	m["Norwegian_Norway_MALE_2"] = voice{"nb-NO", "nb-NO-Wavenet-D"}
	m["Norwegian_Norway_FEMALE_3"] = voice{"nb-NO", "nb-no-Wavenet-E"}
	m["Polish_Poland_FEMALE_1"] = voice{"pl-PL", "pl-PL-Wavenet-A"}
	m["Polish_Poland_MALE_1"] = voice{"pl-PL", "pl-PL-Wavenet-B"}
	m["Polish_Poland_MALE_2"] = voice{"pl-PL", "pl-PL-Wavenet-C"}
	m["Polish_Poland_FEMALE_2"] = voice{"pl-PL", "pl-PL-Wavenet-D"}
	m["Polish_Poland_FEMALE_3"] = voice{"pl-PL", "pl-PL-Wavenet-E"}
	m["Portuguese_Brazil_FEMALE_1"] = voice{"pt-BR", "pt-BR-Wavenet-A"}
	m["Portuguese_Portugal_FEMALE_1"] = voice{"pt-PT", "pt-PT-Wavenet-A"}
	m["Portuguese_Portugal_MALE_1"] = voice{"pt-PT", "pt-PT-Wavenet-B"}
	m["Portuguese_Portugal_MALE_2"] = voice{"pt-PT", "pt-PT-Wavenet-C"}
	m["Portuguese_Portugal_FEMALE_2"] = voice{"pt-PT", "pt-PT-Wavenet-D"}
	m["Russian_Russia_FEMALE_1"] = voice{"ru-RU", "ru-RU-Wavenet-A"}
	m["Russian_Russia_MALE_1"] = voice{"ru-RU", "ru-RU-Wavenet-B"}
	m["Russian_Russia_FEMALE_2"] = voice{"ru-RU", "ru-RU-Wavenet-C"}
	m["Russian_Russia_MALE_2"] = voice{"ru-RU", "ru-RU-Wavenet-D"}
	m["Russian_Russia_FEMALE_3"] = voice{"ru-RU", "ru-RU-Wavenet-E"}
	m["Slovak_Slovakia_FEMALE_1"] = voice{"sk-SK", "sk-SK-Wavenet-A"}
	m["Swedish_Sweden_FEMALE_1"] = voice{"sv-SE", "sv-SE-Wavenet-A"}
	m["Turkish_Turkey_FEMALE_1"] = voice{"tr-TR", "tr-TR-Wavenet-A"}
	m["Turkish_Turkey_MALE_1"] = voice{"tr-TR", "tr-TR-Wavenet-B"}
	m["Turkish_Turkey_FEMALE_2"] = voice{"tr-TR", "tr-TR-Wavenet-C"}
	m["Turkish_Turkey_FEMALE_3"] = voice{"tr-TR", "tr-TR-Wavenet-D"}
	m["Turkish_Turkey_MALE_2"] = voice{"tr-TR", "tr-TR-Wavenet-E"}
	m["Ukrainian_Ukraine_FEMALE_1"] = voice{"uk-UA", "uk-UA-Wavenet-A"}
	m["Vietnamese_Vietnam_FEMALE_1"] = voice{"vi-VN", "vi-VN-Wavenet-A"}
	m["Vietnamese_Vietnam_MALE_1"] = voice{"vi-VN", "vi-VN-Wavenet-B"}
	m["Vietnamese_Vietnam_FEMALE_2"] = voice{"vi-VN", "vi-VN-Wavenet-C"}
	m["Vietnamese_Vietnam_MALE_2"] = voice{"vi-VN", "vi-VN-Wavenet-D"}

	return m[v]
}
