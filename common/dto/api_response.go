package dto

// APIResponseDto api 반환으로 공통적인 응답을 사용합니다
type APIResponseDto struct {
	Status bool        `json:"status"`
	Result interface{} `json:"result"`
}

func newAPIResponse(status bool, value interface{}) *APIResponseDto {
	a := APIResponseDto{
		Status: status,
		Result: value,
	}
	return &a
}

// NewOKAPIResponse 정상적인 API 응답값을 생성해서 반환합니다.
func NewOKAPIResponse(value interface{}) *APIResponseDto {
	return newAPIResponse(true, value)
}

// NewErrorAPIResponse 에러 API 응답값을 생성해서 반환합니다.
func NewErrorAPIResponse(value interface{}) *APIResponseDto {
	return newAPIResponse(false, value)
}
