package myvelo

type Interval struct {
	Start int64 // in ms
	End   int64
}

type IntervalStr struct {
	Start string `json:"start"` // in correct format
	End   string `json:"end"`   // in correct format
}

type GetEdgeLinkMetricsRequest struct {
	EdgeId       int         `json:"edgeId"`
	EnterpriseId int         `json:"enterpriseId"`
	IntervalStr  IntervalStr `json:"interval"`
}
