package helpers

import "encoding/json"

func JSON_stringify(data any) ([]byte, error) {
	jsonString, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return jsonString, nil
}

func JSON_parse(data []byte, dest any) error {
	if err := json.Unmarshal(data, &dest); err != nil {
		return err
	}
	return nil
}
