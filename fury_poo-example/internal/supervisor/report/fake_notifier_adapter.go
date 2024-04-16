package report

import "fmt"

type fakeNotifier struct{}

func NewFakeNotifier() *fakeNotifier {
	return &fakeNotifier{}
}

func (f fakeNotifier) Notify(report Report) error {
	fmt.Printf("reporte notificado: %#v", report)
	return nil
}
