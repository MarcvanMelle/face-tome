package npcgenerator

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	api "github.com/MarcvanMelle/face-tome/internal/pb/facetomeapi"
)

type npcName struct {
	firstName  string
	lastName   string
	middleName string
	title      string
	prefix     string
	suffix     string
}

func (npc *NpcData) setName() {
	lang := npc.request.GetLanguage()
	if lang == api.RealLanguage_LANG_UNKNOWN {
		lang = api.RealLanguage_LANG_EN
	}

	gender := npc.request.GetGender()
	if gender == api.Gender_GEN_UNKNOWN {
		gender = selectWeightedGender()
	}
	npc.npcGender = gender

	firstName, err := getFirstName(lang, gender)
	if err != nil {
		fmt.Println(err)
	}

	lastName, err := getLastName(lang)
	if err != nil {
		fmt.Println(err)
	}

	npc.npcName = &npcName{
		firstName: firstName,
		lastName:  lastName,
	}
}

func getFirstName(lang api.RealLanguage, gender api.Gender) (string, error) {
	sample, err := buildNameSample(lang, gender)
	if err != nil {
		return "", err
	}
	names := strings.Split(strings.TrimSpace(string(sample)), "\n")
	name := names[r.Intn(len(names))]
	return name, nil
}

func getLastName(lang api.RealLanguage) (string, error) {
	sample, err := buildLastNameSample(lang)
	if err != nil {
		return "", err
	}
	names := strings.Split(strings.TrimSpace(string(sample)), "\n")
	name := names[r.Intn(len(names))]
	return name, nil
}

func buildNameSample(lang api.RealLanguage, gender api.Gender) ([]byte, error) {
	var sample []byte
	var err error

	switch gender {
	case api.Gender_GEN_MALE:
		sample, err = readSampleFiles(lang, []string{"first_names_male", "first_names_neutral"})
	case api.Gender_GEN_TRANSMALE:
		sample, err = readSampleFiles(lang, []string{"first_names_male", "first_names_neutral"})
	case api.Gender_GEN_FEMALE:
		sample, err = readSampleFiles(lang, []string{"first_names_female", "first_names_neutral"})
	case api.Gender_GEN_TRANSFEMALE:
		sample, err = readSampleFiles(lang, []string{"first_names_female", "first_names_neutral"})
	default:
		sample, err = readSampleFiles(lang, []string{"first_names_female", "first_names_male", "first_names_neutral"})
	}
	return sample, err
}

func buildLastNameSample(lang api.RealLanguage) ([]byte, error) {
	return readSampleFiles(lang, []string{"last_names"})
}

func readSampleFiles(lang api.RealLanguage, filenames []string) ([]byte, error) {
	var data []byte

	mappedLang := mapAPILangToISO639[lang]

	for _, filename := range filenames {
		nameFile := filepath.Join("/go/src/github.com/MarcvanMelle/face-tome", "internal", "namedata", mappedLang, fmt.Sprintf("%s_%s", mappedLang, filename))
		nameData, err := ioutil.ReadFile(nameFile)
		if err != nil {
			return nil, err
		}

		data = append(data, nameData...)
	}

	return data, nil
}

func selectWeightedGender() api.Gender {
	weightedSelector := r.Intn(99)

	for gender, intRange := range weightedGenders {
		min := intRange[0]
		max := intRange[len(intRange)-1]
		if (weightedSelector > min && weightedSelector < max) || weightedSelector == min || weightedSelector == max {
			return gender
		}
	}
	return api.Gender_GEN_ADNROGYNOUS
}

var weightedGenders = map[api.Gender][]int{
	api.Gender_GEN_ADNROGYNOUS: generateIntRange(70, 79),
	api.Gender_GEN_FEMALE:      generateIntRange(0, 34),
	api.Gender_GEN_MALE:        generateIntRange(35, 69),
	api.Gender_GEN_TRANSFEMALE: generateIntRange(87, 93),
	api.Gender_GEN_TRANSMALE:   generateIntRange(94, 99),
	api.Gender_GEN_UNGENDERED:  generateIntRange(80, 86),
}

var mapAPILangToISO639 = map[api.RealLanguage]string{
	api.RealLanguage_LANG_UNKNOWN: "unknown",
	api.RealLanguage_LANG_AF:      "af",
	api.RealLanguage_LANG_AR:      "ar",
	api.RealLanguage_LANG_CS:      "cs",
	api.RealLanguage_LANG_DE:      "de",
	api.RealLanguage_LANG_EL:      "el",
	api.RealLanguage_LANG_EN:      "en",
	api.RealLanguage_LANG_ES:      "es",
	api.RealLanguage_LANG_FI:      "fi",
	api.RealLanguage_LANG_FR:      "fr",
	api.RealLanguage_LANG_GA:      "ga",
	api.RealLanguage_LANG_HE:      "he",
	api.RealLanguage_LANG_HI:      "hi",
	api.RealLanguage_LANG_IT:      "it",
	api.RealLanguage_LANG_JA:      "ja",
	api.RealLanguage_LANG_KO:      "ko",
	api.RealLanguage_LANG_LA:      "la",
	api.RealLanguage_LANG_PL:      "pl",
	api.RealLanguage_LANG_RU:      "ru",
	api.RealLanguage_LANG_SA:      "sa",
	api.RealLanguage_LANG_SV:      "sv",
	api.RealLanguage_LANG_VI:      "vi",
	api.RealLanguage_LANG_ZH:      "zh",
}
