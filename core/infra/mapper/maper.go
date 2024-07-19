package mapper

import "encoding/json"

func ConvertStructToMap(req interface{}) (map[string]interface{}, error) {
	// Marshal the struct to JSON
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON back into a map
	var result map[string]interface{}
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
