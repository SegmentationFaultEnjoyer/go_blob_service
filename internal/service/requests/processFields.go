package requests

import (
	"testService/internal/service/helpers"
	"testService/resources"
)

func FieldsToJSON(blob *resources.Blob) ([]byte, []byte, []byte, error) {
	attributes, err := helpers.JSON_stringify(blob.Attributes)
	if err != nil {
		return nil, nil, nil, err
	}
	relationships, err := helpers.JSON_stringify(blob.Relationships)
	if err != nil {
		return nil, nil, nil, err
	}
	key, err := helpers.JSON_stringify(blob.Key)
	if err != nil {
		return nil, nil, nil, err
	}

	return attributes, relationships, key, nil
}

func JSONToFields(blob *resources.Blob, attributes []byte, relationships []byte, key []byte) error {
	if err := helpers.JSON_parse(attributes, &blob.Attributes); err != nil {
		return err
	}
	if err := helpers.JSON_parse(relationships, &blob.Relationships); err != nil {
		return err
	}
	if err := helpers.JSON_parse(key, &blob.Key); err != nil {
		return err
	}

	return nil
}
