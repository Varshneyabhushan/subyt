package periodicjob

import (
	"time"
)

type jobSignal struct{}
var CloseSignal = jobSignal{}

type periodicJob struct {
	endSignal chan jobSignal
	period time.Duration
	job func()
}

/*
for every timePeriod given, repeat the job
util a closeSignal is passed
*/
func (p periodicJob) Start() {
	go func ()  {
		for {
			jobChannel := make(chan jobSignal)
			timeout := time.After(p.period)
			go func() {
				p.job()
				jobChannel <- CloseSignal
			}()

			select {
			//when closeSignal is passed, return from the func (process complete) 
			case <- p.endSignal:
				return

			//if the job is over, wait till the time is up
			case <- timeout:
				<- jobChannel
			}

			close(jobChannel)
		}
	}()
}

//end the periodic job
func (p periodicJob) End() {
	p.endSignal <- CloseSignal
}

func New(job func(), period time.Duration) periodicJob {
	return periodicJob{
		period: period,
		job: job,
		endSignal: make(chan jobSignal),
	}
}
