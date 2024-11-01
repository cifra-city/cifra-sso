package action_permission

import "time"

func (a *ActionPermission) AddToQueue(username string, eve string) {
	a.mu.Lock()
	a.Queue[username] = eve
	a.mu.Unlock()

	time.AfterFunc(180*time.Second, func() {
		a.mu.Lock()
		defer a.mu.Unlock()
		if _, exists := a.Queue[username]; exists {
			delete(a.Queue, username)
		}
	})
}

func (a *ActionPermission) GetEvent(username string) (string, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if storedCode, exists := a.Queue[username]; exists {
		return storedCode, nil
	}

	return "", nil
}

func (a *ActionPermission) DeleteEve(username string) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if _, exists := a.Queue[username]; exists {
		delete(a.Queue, username)
	}
}
