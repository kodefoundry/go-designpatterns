package singleton

// UnsafeSingleton ... Interface to the unsafe singleton counter
// The interface exposes a method AddOne that will return the current
// state of the unsafe counter
type UnsafeSingleton interface {
	AddOne() int
}

// counter ... A non thread safe simple counter type.
type counter struct {
	count int
}

// Create a pointer to the counter which will hold the singleton instance
var instance *counter

// GetUnsafeInstance ... The public method which will give a interface to the counter.
func GetUnsafeInstance() UnsafeSingleton {
	if nil == instance {
		instance = new(counter)
	}
	return instance
}

// AddOne ... This method returns the current counter value after increment by 1
func (c *counter) AddOne() int {
	c.count++
	return c.count
}

// AddTwo ... This method returns the current counter value after increment by 2
// This method should not be accessible from UnsafeSingleton interface.
func (c *counter) AddTwo() int {
	c.count += 2
	return c.count
}
