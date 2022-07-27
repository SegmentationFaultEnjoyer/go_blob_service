package helpers

import (
	"encoding/json"
	"testService/internal/data"
	"testService/resources"
)

func FieldsToJSON(blob *resources.Blob) (*data.Blob, error) {
	blobData := &data.Blob{}
	var err error

	blobData.Attributes, err = json.Marshal(blob.Attributes)
	if err != nil {
		return nil, err
	}
	blobData.Relationships, err = json.Marshal(blob.Relationships)
	if err != nil {
		return nil, err
	}
	blobData.Key, err = json.Marshal(blob.Key)
	if err != nil {
		return nil, err
	}

	return blobData, nil
}

func JSONToFields(blob *resources.Blob, blobData *data.Blob) error {
	if err := json.Unmarshal(blobData.Attributes, &blob.Attributes); err != nil {
		return err
	}
	if err := json.Unmarshal(blobData.Relationships, &blob.Relationships); err != nil {
		return err
	}
	if err := json.Unmarshal(blobData.Key, &blob.Key); err != nil {
		return err
	}

	blob.Key.ID = blobData.ID

	return nil
}
