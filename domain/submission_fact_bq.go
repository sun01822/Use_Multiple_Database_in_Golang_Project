package domain

import (
	"cloud.google.com/go/bigquery"
	"time"
)

type SubmissionDataModelBQ struct {
	SubmissionUID bigquery.NullString `json:"submissionUid"`
	EInvoiceID    bigquery.NullString `json:"eInvoiceId"`
	SubmittedAt   time.Time           `json:"submittedAt"`
	Status        bigquery.NullString `json:"status"`
	UUID          bigquery.NullString `json:"uuid"`
	LongID        bigquery.NullString `json:"longId"`
	Error         bigquery.NullString `json:"error"`
	TraceID       bigquery.NullString `json:"traceId"`
}

type DataDetailsBQ struct {
	FetchingTime string                  `json:"fetchingTime"`
	Data         []SubmissionDataModelBQ `json:"data"`
}
