package model

import "encoding/json"

type Node struct {
	Title string          `json:"title"`
	Type  string          `json:"type"`
	Tags  []string        `json:"tags"`
	Note  json.RawMessage `json:"note"`
}
