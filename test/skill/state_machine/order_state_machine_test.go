package machine_test

import (
	"testing"

	machine "github.com/kingson4wu/gooooooooo/test/skill/state_machine"
)

type OrderStatus int

const (
	_ OrderStatus = iota
	StatusOK
	StatusFreeze
	StatusRevert
)

func (o OrderStatus) get() int {
	return int(o)
}

type OrderEvent int

const (
	_ OrderEvent = iota
	Freeze
	Revert
	Confirm
)

func (o OrderEvent) get() int {
	return int(o)
}

func save(fromState, toState, event int, order machine.Order) bool {
	return true
}
func do(fromState, toState, event int, order machine.Order) {
}

type ConsumeOrder struct {
}

func (o *ConsumeOrder) GetCurrentState() int {
	return int(Freeze)
}

func TestMachine(t *testing.T) {
	m := machine.NewStateMachine([]*machine.Config{
		{Event: Freeze.get(), FromState: 0, ToState: StatusFreeze.get(), Guard: save, Action: do},
		{Event: Revert.get(), FromState: StatusFreeze.get(), ToState: StatusRevert.get(), Guard: save, Action: do},
		{Event: Confirm.get(), FromState: StatusFreeze.get(), ToState: StatusOK.get(), Guard: save, Action: do},
	})

	m.Fire(Confirm.get(), &ConsumeOrder{})

}
