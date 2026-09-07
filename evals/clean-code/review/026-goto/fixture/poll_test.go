package fixture

import "testing"

type countdownJob struct {
	pollsUntilDone int
}

func (j *countdownJob) Done() bool {
	j.pollsUntilDone--
	return j.pollsUntilDone < 0
}

func TestWaitUntilDoneReturnsTrueWithinBudget(t *testing.T) {
	if !WaitUntilDone(&countdownJob{pollsUntilDone: 2}, 5) {
		t.Fatal("WaitUntilDone = false, want true")
	}
}

func TestWaitUntilDoneGivesUpAfterBudget(t *testing.T) {
	if WaitUntilDone(&countdownJob{pollsUntilDone: 10}, 3) {
		t.Fatal("WaitUntilDone = true past the poll budget")
	}
}
