package gtmetrix

type TestStateType string

const (
	TestStateTypeQueued    = "queued"
	TestStateTypeStarted   = "started"
	TestStateTypeCompleted = "completed"
	TestStateTypeError     = "error"
)
