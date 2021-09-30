package common

import "encoding/json"


func PareJson(str string) []map[string]string {
	var vs []map[string]string
	if err := json.Unmarshal([]byte(str), &vs); err != nil {
		panic(err)
	}

	return vs
}
