package domain

type SubmissionDataModelPG struct {
	SubmissionUID string `json:"submissionUid" gorm:"column:submissionUid"`
	EInvoiceID    string `json:"eInvoiceId" gorm:"column:eInvoiceId"`
	SubmittedAt   string `json:"submittedAt" gorm:"column:submittedAt"`
	Status        string `json:"status" gorm:"column:status"`
	UUID          string `json:"uuid" gorm:"column:uuid" gorm:"primary key"`
	LongID        string `json:"longId" gorm:"column:longId"`
	Error         string `json:"error" gorm:"column:error"`
	TraceID       string `json:"traceId" gorm:"column:traceId"`
}

type DataDetailsPG struct {
	FetchingTime string                  `json:"fetchingTime"`
	Data         []SubmissionDataModelPG `json:"data"`
}
