package events

import "time"

func (e *Events) AddToQueue(username string, code string) {
	e.mu.Lock()
	e.Queue[username] = code
	e.mu.Unlock()

	time.AfterFunc(60*time.Second, func() {
		e.mu.Lock()
		defer e.mu.Unlock()
		if _, exists := e.Queue[username]; exists {
			delete(e.Queue, username)
		}
	})
}

func (e *Events) CheckInQueue(username string, eve string) (string, error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if storedCode, exists := e.Queue[username]; exists && storedCode == eve {
		delete(e.Queue, username)
		return storedCode, nil
	}

	return "", nil
}
