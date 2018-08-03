package npcgenerator

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"strings"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func getName(lang, gender string) string {
	firstName, err := getFirstName(lang, gender)
	if err != nil {
		fmt.Println(err)
	}

	lastName, err := getLastName(lang)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprint(firstName, lastName)
}

func getFirstName(lang, gender string) (string, error) {
	sample, err := buildSample(lang, gender)
	if err != nil {
		return "", err
	}
	names := strings.Split(strings.TrimSpace(string(sample)), "\n")
	name := names[r.Intn(len(names))]
	return name, nil
}

func getLastName(lang string) (string, error) {
	sample, err := buildSample(lang, "lastName")
	if err != nil {
		return "", err
	}
	names := strings.Split(strings.TrimSpace(string(sample)), "\n")
	name := names[r.Intn(len(names))]
	return name, nil
}

func buildSample(lang, gender string) ([]byte, error) {
	var sample []byte
	var err error

	switch gender {
	case "male":
		sample, err = readSampleFiles(lang, []string{"first_names_male", "first_names_neutral"})
	case "female":
		sample, err = readSampleFiles(lang, []string{"first_names_female", "first_names_neutral"})
	case "lastName":
		sample, err = readSampleFiles(lang, []string{"last_names"})
	default:
		sample, err = readSampleFiles(lang, []string{"first_names_female", "first_names_male", "first_names_neutral"})
	}
	if err != nil {
		return nil, err
	}
	return sample, err
}

func readSampleFiles(lang string, filenames []string) ([]byte, error) {
	var data []byte

	for _, filename := range filenames {
		nameFile := filepath.Join(".", "namedata", lang, fmt.Sprintf("%s_%s", lang, filename))
		nameData, err := ioutil.ReadFile(nameFile)
		if err != nil {
			return nil, err
		}

		data = append(data, nameData...)
	}

	return data, nil
}
