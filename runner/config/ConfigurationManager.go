package config

type ConfigurationManager struct {
	measurementType MeasurementType
	nominalRate     int
	executionType   ExecutionType
	duration        int64
	forceContinue   bool
	changeListeners []*ConfigurationCallback
}

func NewConfigurationManager() *ConfigurationManager {
	c := ConfigurationManager{
		measurementType: Peak,
		nominalRate:     1,
		executionType:   Proportional,
		duration:        60,
		forceContinue:   false,
	}
	c.changeListeners = make([]*ConfigurationCallback, 0)
	return &c
}

func (c *ConfigurationManager) GetMeasurementType() MeasurementType {
	return c.measurementType
}

func (c *ConfigurationManager) SetMeasurementType(value MeasurementType) {
	c.measurementType = value
	c.notifyChanged()
}

func (c *ConfigurationManager) GetNominalRate() int {
	return c.nominalRate
}

func (c *ConfigurationManager) SetNominalRate(value int) {
	c.nominalRate = value
	c.notifyChanged()
}

func (c *ConfigurationManager) GetExecutionType() ExecutionType {
	return c.executionType
}

func (c *ConfigurationManager) SetExecutionType(value ExecutionType) {
	c.executionType = value
	c.notifyChanged()
}

func (c *ConfigurationManager) GetDuration() int64 {
	return c.duration
}

func (c *ConfigurationManager) SetDuration(value int64) {
	c.duration = value
	c.notifyChanged()
}

func (c *ConfigurationManager) GetForceContinue() bool {
	return c.forceContinue
}

func (c *ConfigurationManager) SetForceContinue(value bool) {
	c.forceContinue = value
	c.notifyChanged()
}

func (c *ConfigurationManager) AddChangeListener(listener *ConfigurationCallback) {
	c.changeListeners = append(c.changeListeners, listener)
}

func (c *ConfigurationManager) RemoveChangeListener(listener *ConfigurationCallback) {
	for index := len(c.changeListeners) - 1; index >= 0; index-- {
		if c.changeListeners[index] == listener {

			if index == len(c.changeListeners) {
				c.changeListeners = c.changeListeners[:index-1]
			} else {
				c.changeListeners = append(c.changeListeners[:index], c.changeListeners[index+1:]...)
			}

		}
	}
}

func (c *ConfigurationManager) notifyChanged() {
	for index := 0; index < len(c.changeListeners); index++ {
		listener := c.changeListeners[index]
		(*listener)()
	}
}
