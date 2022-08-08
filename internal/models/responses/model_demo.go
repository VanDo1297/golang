package responses

type DemoResponses struct {
	Data struct {
		Demos []*Demo `json:"demos" validate:"required"`
	} `json:"data" validate:"required"`
	APIResponse
}

type Demo struct {
	Base
	Message string `json:"message" validate:"required"`
}
