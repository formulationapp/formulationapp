package dto

type UpdateFormRequest struct {
	Name string                 `json:"name"`
	Data map[string]interface{} `json:"data"`
}
