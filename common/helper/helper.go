package helper

import (
	"encoding/json"
	"log"
)

func PrettyPrint(data interface{}) string {
	JSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	return string(JSON)
}
