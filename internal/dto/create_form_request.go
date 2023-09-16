package dto

type CreateFormRequest struct {
	Name       string `json:"name"`
	Definition string `json:"definition"`
}
