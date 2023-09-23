package dto

type CreateFormRequest struct {
	Name string                 `json:"name"`
	Data map[string]interface{} `json:"data"`
}
