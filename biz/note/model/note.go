package model

import (
	"encoding/json"
)

type Note struct {
	Title string          `json:"title"`
	Type  string          `json:"type"`
	Tags  []string        `json:"tags"`
	Note  json.RawMessage `json:"note"`
}

type UpdateNote struct {
	Note
	ID      string `json:"id"`
	Version int64  `json:"version"`
}
