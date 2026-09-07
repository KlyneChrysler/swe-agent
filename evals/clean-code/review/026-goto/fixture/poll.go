package fixture

type Job interface {
	Done() bool
}

func WaitUntilDone(job Job, maxPolls int) bool {
	polls := 0
poll:
	if job.Done() {
		return true
	}
	polls++
	if polls < maxPolls {
		goto poll
	}
	return false
}
