package requestStatus

type RequestStatus string

const (
	Rejected RequestStatus = "Rejected"
	Pending  RequestStatus = "Pending"
	Accepted RequestStatus = "Accepted"
)
