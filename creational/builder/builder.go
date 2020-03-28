package builder

// Builder ... Public interface implementing the builder design pattern.
// The build process undergoes a number of steps as part of build pipeline.
type Builder interface {
	MountChasis(string) Builder
	MountEngine(string) Builder
	SetWheels() Builder
	SetSeats() Builder
	GetVehicle() Vehicle
}

// Factory ... The build controller that will manage the building process
// for different types of vehicles
type Factory struct {
	builder Builder
}

// Build ... The method that actually triggers the manufacturing process.
func (f *Factory) Build() {
	f.builder.MountChasis("F23J4543K")
	f.builder.MountEngine("E20205422J")
	f.builder.SetWheels()
	f.builder.SetSeats()
}

// SetBuilder ... A method to set the concrete builder to the factory
func (f *Factory) SetBuilder(b Builder) {
	f.builder = b
}

// Vehicle ... The final porduct that is produced after the build process
type Vehicle struct {
	EngineNo string
	ChasisNo string
	Wheels   int
	Seats    int
}

// CarBuilder ... The builder for object tpy Car
type CarBuilder struct {
	vehicle Vehicle
}

// MountEngine ... The build process to mount engine
func (c *CarBuilder) MountEngine(engineNo string) Builder {
	c.vehicle.EngineNo = engineNo
	return c
}

// MountChasis ... The build process to mount chasis
func (c *CarBuilder) MountChasis(chasisNo string) Builder {
	c.vehicle.ChasisNo = chasisNo
	return c
}

// SetWheels ... The build process to fit wheels
func (c *CarBuilder) SetWheels() Builder {
	c.vehicle.Wheels = 4
	return c
}

// SetSeats ... The build process to fit seats
func (c *CarBuilder) SetSeats() Builder {
	c.vehicle.Seats = 5
	return c
}

// GetVehicle ... Returns the build product.
func (c *CarBuilder) GetVehicle() Vehicle {
	return c.vehicle
}


