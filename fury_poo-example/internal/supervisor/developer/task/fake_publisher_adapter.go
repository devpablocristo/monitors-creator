package task

import "fmt"

type fakePublisher struct{}

func NewFakePublisher() *fakePublisher {
	return &fakePublisher{}
}

func (f fakePublisher) Publish(ownerId string, t Task) error {
	fmt.Printf("fake publishing: owner_id:%s, task:%#v", ownerId, t)
	return nil
}
