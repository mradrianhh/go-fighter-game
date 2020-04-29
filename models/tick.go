package models

// Tick is a type which encapsulates the value of the current gameloop-cycle and a method to be executed while ticking.
type Tick struct {
	value      int
	tasks      []func()
	extraTasks []func()
}

// Tick executes the tick action in which it increases the value of the tick object, and performs the action specified.
func (tick *Tick) Tick() {
	for _, task := range tick.extraTasks {
		task()
	}
	tick.extraTasks = []func(){}

	for _, task := range tick.tasks {
		task()
	}

	tick.value++
}

// NewTick creates and returns a new tick-object.
func NewTick(tasks []func()) Tick {
	return Tick{tasks: tasks}
}

// QueueExtraTasks adds extra "tasks" to the extraTasks field of the tick to be performed the next "Tick()".
func (tick *Tick) QueueExtraTasks(tasks ...func()) {
	tick.extraTasks = append(tasks, tasks...)
}

// GetValue returns the current value of the tick.
func (tick *Tick) GetValue() int {
	return tick.value
}
