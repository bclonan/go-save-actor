package promises

import "fmt"

// Promise struct represents a promise in the system
type Promise struct {
	ActorID string
	Task    func() (string, error)
}

// NewPromise creates a new promise
func NewPromise(actorID string, task func() (string, error)) *Promise {
	return &Promise{
		ActorID: actorID,
		Task:    task,
	}
}

// Resolve executes the task in the promise and returns the result
func (p *Promise) Resolve() {
	result, err := p.Task()
	if err != nil {
		fmt.Println("Error resolving promise:", err)
		return
	}
	fmt.Println("Promise resolved with result:", result)
}