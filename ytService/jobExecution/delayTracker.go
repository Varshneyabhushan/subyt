package jobexecution

import "time"

type DelayTracker struct {
	defaultDelay time.Duration
	currentDelay time.Duration
	delayCount   int
}

func NewDelayTracker(defaultDelay time.Duration) *DelayTracker {
	return &DelayTracker{defaultDelay: defaultDelay, currentDelay: defaultDelay}
}

func (tracker *DelayTracker) Delay() time.Duration {
	return tracker.currentDelay
}

func (tracker *DelayTracker) Reset() {
	tracker.currentDelay = tracker.defaultDelay
	tracker.delayCount = 0
}

// ProportionalDelay ProportionDelay increases delay proportional to totalDelays
func (tracker *DelayTracker) ProportionalDelay() time.Duration {
	tracker.delayCount += 1
	tracker.currentDelay += tracker.currentDelay
	return tracker.Delay()
}

// ExponentialBackOff doubles delay by 2
func (tracker *DelayTracker) ExponentialBackOff() time.Duration {
	tracker.delayCount += 1
	tracker.currentDelay *= 2
	return tracker.Delay()
}
