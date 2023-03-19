package jobexecution

import (
	"time"
)

// RecurringJob
// when executed, same job is called recursively until it throws an error
// job itself determines how long to wait to be executed it once again
type RecurringJob func() (time.Duration, error)

func (job RecurringJob) Execute() error {
	jobStartTime := time.Now()
	waitingTime, err := job()
	if err != nil {
		return err
	}

	timeElapsed := time.Now().Sub(jobStartTime)
	waitingTime -= timeElapsed

	//don't wait when waitTime is negative (i.e, time taken is more than waitTime)
	if waitingTime > 0 {
		<-time.After(waitingTime)
	}

	return job.Execute()
}
