package dto

type PutAnswerRequest struct {
	Field string `json:"field"`
	Value string `json:"value"`
}
