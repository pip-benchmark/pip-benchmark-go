package results

type ResultsManager struct {
	results          []*BenchmarkResult
	updatedListeners []*ResultCallback
	messageListeners []*MessageCallback
	errorListeners   []*ErrorCallback
}

func NewResultsManager() *ResultsManager {
	return &ResultsManager{
		results:          make([]*BenchmarkResult, 0),
		updatedListeners: make([]*ResultCallback, 0),
		messageListeners: make([]*MessageCallback, 0),
		errorListeners:   make([]*ErrorCallback, 0),
	}
}

func (c *ResultsManager) All() []*BenchmarkResult {
	return c.results
}

func (c *ResultsManager) Add(result *BenchmarkResult) {
	c.results = append(c.results, result)
}

func (c *ResultsManager) Clear() {
	c.results = make([]*BenchmarkResult, 0)
}

func (c *ResultsManager) AddUpdatedListener(listener *ResultCallback) {
	c.updatedListeners = append(c.updatedListeners, listener)
}

func (c *ResultsManager) RemoveUpdatedListener(listener *ResultCallback) {
	for index := len(c.updatedListeners) - 1; index >= 0; index-- {
		if c.updatedListeners[index] == listener {
			if index == len(c.updatedListeners) {
				c.updatedListeners = c.updatedListeners[:index-1]
			} else {
				c.updatedListeners = append(c.updatedListeners[:index], c.updatedListeners[index+1:]...)
			}
		}
	}
}

func (c *ResultsManager) NotifyUpdated(result BenchmarkResult) {
	for index := 0; index < len(c.updatedListeners); index++ {
		listener := c.updatedListeners[index]
		(*listener)(result)
	}
}

func (c *ResultsManager) AddMessageListener(listener *MessageCallback) {
	c.messageListeners = append(c.messageListeners, listener)
}

func (c *ResultsManager) RemoveMessageListener(listener *MessageCallback) {
	for index := len(c.messageListeners) - 1; index >= 0; index-- {
		if c.messageListeners[index] == listener {
			if index == len(c.messageListeners) {
				c.messageListeners = c.messageListeners[:index-1]
			} else {
				c.messageListeners = append(c.messageListeners[:index], c.messageListeners[index+1:]...)
			}
		}
	}
}

func (c *ResultsManager) NotifyMessage(message string) {
	for index := 0; index < len(c.messageListeners); index++ {
		listener := c.messageListeners[index]
		(*listener)(message)
	}
}

func (c *ResultsManager) AddErrorListener(listener *ErrorCallback) {
	c.errorListeners = append(c.errorListeners, listener)
}

func (c *ResultsManager) RemoveErrorListener(listener *ErrorCallback) {
	for index := len(c.errorListeners) - 1; index >= 0; index-- {
		if c.errorListeners[index] == listener {
			if index == len(c.errorListeners) {
				c.errorListeners = c.errorListeners[:index-1]
			} else {
				c.errorListeners = append(c.errorListeners[:index], c.errorListeners[index+1:]...)
			}
		}
	}
}

func (c *ResultsManager) NotifyError(err error) {
	for index := 0; index < len(c.errorListeners); index++ {
		listener := c.errorListeners[index]
		(*listener)(err)
	}
}
