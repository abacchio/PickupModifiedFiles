package services

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/PickupModifiedFiles/models"
)

// OutputJSON outputs Json file
func OutputJSON(dirContent models.DirContent, logfilename string) {
	jsonBytes, err := json.Marshal(dirContent)
	if err != nil {
		panic(err)
	}

	logfile, err := os.Create(logfilename)
	if err != nil {
		panic(err)
	}
	defer logfile.Close()

	logfile.WriteString(string(jsonBytes))
}

// ReadJSON read Json file
func ReadJSON(logfilename string) models.DirContent {
	jsonFile, err := ioutil.ReadFile(logfilename)
	if err != nil {
		panic(err)
	}

	var dirContent models.DirContent
	if err := json.Unmarshal(jsonFile, &dirContent); err != nil {
		panic(err)
	}

	return dirContent
}
