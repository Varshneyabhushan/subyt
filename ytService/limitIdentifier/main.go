package limitidentifier

import "encoding/json"

//tracks if videos limit reached or not
type LimitIdentifier struct {
	From string
	Limit string
}

type storageFormat struct {
	Limit string
}

func (i LimitIdentifier) MarshalJSON() ([]byte, error){
	return json.Marshal(storageFormat{ Limit: i.Limit })
}

func (tracker *LimitIdentifier) AdvanceLimit() {
	tracker.Limit = tracker.From
	tracker.From = ""
}

type IdProvider interface {
	GetId() string
}

func (tracker *LimitIdentifier) Scrutinise(idProviders []IdProvider) int {
	totalValidProviders := 0

	if len(idProviders) != 0 && len(tracker.From) == 0 {
		tracker.From = idProviders[0].GetId()
	}

	for _, idProvider := range idProviders {
		if idProvider.GetId() == tracker.Limit {
			break
		}

		totalValidProviders += 1
	}

	return totalValidProviders
}