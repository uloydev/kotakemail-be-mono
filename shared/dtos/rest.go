package dtos

type RestResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data"`
	Meta    struct {
		Pagination    PaginationMeta `json:"pagination"`
		RequestID     string         `json:"request_id"`
		TraceID       string         `json:"trace_id"`
		ExecutionTime float64        `json:"execution_time"`
		RequestTime   string         `json:"request_time"`
		Error         *Error         `json:"error,omitempty"`
	} `json:"meta"`
}
