package dto

type StrResourceAuthAction struct {
	Value   string   `json:"value,omitempty"`
	Actions []string `json:"actions,omitempty"`
}
