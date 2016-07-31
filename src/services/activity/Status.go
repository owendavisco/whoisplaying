package activity

type Status string

const (
	Completed Status = "Completed"
	Running Status = "Running"
	NotStarted Status = "NotStarted"
	Failed Status = "Failed"
)
