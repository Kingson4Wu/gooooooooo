package machine

type StateMachine struct {
	actionConfigs map[int]*Config
}

func NewStateMachine(configs []*Config) *StateMachine {
	m := make(map[int]*Config, len(configs))
	for _, c := range configs {
		m[c.Event] = c
	}
	return &StateMachine{
		actionConfigs: m,
	}
}

type Order interface {
	GetCurrentState() int
}

func (m *StateMachine) Fire(event int, order Order) bool {

	var (
		c  *Config
		ok bool
	)
	if c, ok = m.actionConfigs[event]; !ok {
		return false
	}

	if order.GetCurrentState() != c.FromState {
		return false
	}

	if c.Guard(c.FromState, c.ToState, c.Event, order) {
		c.Action(c.FromState, c.ToState, c.Event, order)
	}

	return false

}

type Config struct {
	Event     int
	FromState int
	ToState   int
	Guard     func(fromState, toState, event int, order Order) bool
	Action    func(fromState, toState, event int, order Order)
}
