package periodicjob

import (
	"time"
)

type jobSignal struct{}

var CloseSignal = jobSignal{}

type PeriodicJob struct {
	endSignal chan jobSignal
	period    time.Duration
	job       func()
}

/*
for every timePeriod given, repeat the job
util a closeSignal is passed
*/
func (p PeriodicJob) Start() {
	go func() {
		for {
			jobChannel := make(chan jobSignal)
			timeout := time.After(p.period)
			go func() {
				p.job()
				jobChannel <- CloseSignal
			}()

			select {
			//when closeSignal is passed, return from the func (process complete)
			case <-p.endSignal:
				return

			//if the job is over, wait till the time is up
			case <-timeout:
				<-jobChannel
			}

			close(jobChannel)
		}
	}()
}

// end the periodic job
func (p PeriodicJob) End() {
	p.endSignal <- CloseSignal
}

func New(job func(), period time.Duration) PeriodicJob {
	return PeriodicJob{
		period:    period,
		job:       job,
		endSignal: make(chan jobSignal),
	}
}
