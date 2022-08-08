package responses

type BaseResponse struct {
	APIResponse
}

type APIResponse struct {
	Version float64 `db:"name" json:"api_version" validate:"required"`
}

type AppErrorResponse struct {
	Message string `json:"message" validate:"required"`
	Code    int    `json:"code" validate:"required"`
	APIResponse
}
