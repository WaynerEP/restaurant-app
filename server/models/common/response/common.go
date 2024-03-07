package response

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

type PageRecord struct {
	RecordID interface{} `json:"recordId"`
	Record   interface{} `json:"record"`
}
