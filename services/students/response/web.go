package response

type Response struct {
	Code     int         `json:"code"`
	Status   string      `json:"status"`
	Data     interface{} `json:"data,omitempty"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

type ValidationErrors struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Errors interface{} `json:"errors,omitempty"`
}

// singleResponse for post,update,delete
type SingleResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"errors,omitempty"`
	Message string      `json:"message"`
}
