package events

import "time"

func (e *Events) AddToQueue(username string, eve string) {
	e.mu.Lock()
	e.Queue[username] = eve
	e.mu.Unlock()

	time.AfterFunc(360*time.Second, func() {
		e.mu.Lock()
		defer e.mu.Unlock()
		if _, exists := e.Queue[username]; exists {
			delete(e.Queue, username)
		}
	})
}

func (e *Events) GetEvent(username string) (string, error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if storedCode, exists := e.Queue[username]; exists {
		return storedCode, nil
	}

	return "", nil
}

func (e *Events) DeleteEve(username string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if _, exists := e.Queue[username]; exists {
		delete(e.Queue, username)
	}
}
