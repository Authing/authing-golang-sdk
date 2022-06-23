package dto

type PolicyCondition struct {
	Param    string      `json:"param"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}
