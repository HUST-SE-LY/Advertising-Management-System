package enum

type ReviewStatusType int64

const (
	UnderReview ReviewStatusType = iota
	PASSED
	REJECTED
	STOP
)
