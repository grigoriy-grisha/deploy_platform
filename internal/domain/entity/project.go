package entity

type Project struct {
	ID      string `json:"id,omitempty"`
	Command string `json:"command,omitempty"`
	Name    string `json:"name,omitempty"`
}
