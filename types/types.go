package types

type MathRequest struct {
	NumA float32 `json:"numA" validate:"required" conform:"trim"`
	NumB float32 `json:"numB" validate:"required" conform:"trim"`
}

type MathResponse struct {
	Result float32 `json:"result"`
}
