package npcgenerator

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"strings"
	"time"

	api "github.com/MarcvanMelle/face-tome/internal/pb/facetomeapi"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func getName(lang api.RealLanguage, gender api.Gender) []string {
	firstName, err := getFirstName(lang, gender)
	if err != nil {
		fmt.Println(err)
	}

	lastName, err := getLastName(lang)
	if err != nil {
		fmt.Println(err)
	}

	return []string{firstName, lastName}
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
		nameFile := filepath.Join(".", "namedata", mappedLang, fmt.Sprintf("%s_%s", mappedLang, filename))
		nameData, err := ioutil.ReadFile(nameFile)
		if err != nil {
			return nil, err
		}

		data = append(data, nameData...)
	}

	return data, nil
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
