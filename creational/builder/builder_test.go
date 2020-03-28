package builder

import "testing"

func TestBuild(t *testing.T){
	
	// <------------- Arrange -------------->
	// Create a concrete builder object
	carBuilder :=  CarBuilder{}
	// Add concrete builder to Factory
	factory := Factory{}
	// <------------- Arrange -------------->

	// <------------- Act     -------------->
	factory.SetBuilder(&carBuilder)
	factory.Build()
	car := carBuilder.GetVehicle()
	// <------------- Act     -------------->

	// <------------- Assert  -------------->
	if car.EngineNo != "E20205422J" {
		t.Errorf("Engine number %s does not match with expected engine no 'E20205422J'\n", car.EngineNo)
	}

	if car.ChasisNo != "F23J4543K" {
		t.Errorf("Chasis number %s does not match with expected chasis no 'F23J4543K'", car.ChasisNo)
	}

	if car.Wheels != 4 {
		t.Error("Produced car doesn't have 4 wheels")
	}

	if car.Seats != 5 {
		t.Error("Produced car doesn't have 5 seats")
	}
	// <------------- Assert  -------------->
}