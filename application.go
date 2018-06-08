package stripe

import "encoding/json"

type Application struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// UnmarshalJSON handles deserialization of an Application.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (out *Application) UnmarshalJSON(data []byte) error { // remi-done
	if id, ok := ParseID(data); ok {
		out.ID = id
		return nil
	}

	type application Application
	var v application
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*out = Application(v)
	return nil
}
