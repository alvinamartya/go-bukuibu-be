package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func GetEnvVar(key string) (string, error) {
	dirname, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(path.Join(dirname, "/config/env.local.json"))
	if err != nil {
		log.Fatal(err)
	}

	blob, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	doc := make(map[string]interface{})
	if err := json.Unmarshal(blob, &doc); err != nil {
		log.Fatal(err)
	}

	if value, contains := doc[key].(string); contains {
		return value, nil
	} else {
		return "", errors.New("Unable to access key from config file ")
	}
}
