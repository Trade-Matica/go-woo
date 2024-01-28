package shared

import "encoding/json"

func UnmarshalTo[T any](raw []byte) (*T, error) {
	var resp T
	err := json.Unmarshal(raw, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
