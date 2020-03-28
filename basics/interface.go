package main

import "fmt"

// RailGauge ...
// Describes a rail gauge
type RailGauge interface {
	GetGauge() int
}

// RailRoute ...
// Describes a rail route
type RailRoute struct {
	Gauge int
}

// IsCompatibleGauge ...
// Check if the train can run on the rail route
func (r *RailRoute) IsCompatibleGauge(p RailGauge) bool {
	return p.GetGauge() == r.Gauge
}

// Train ...
// A simple struct hold train details
type Train struct {
	TrainGauge int
}

// GetGauge ...
// Returns the guage length of weels
func (p *Train) GetGauge() int {
	return p.TrainGauge
}

func runInterface() {
	railRoute := RailRoute{Gauge: 10}

	passengerTrain := &Train{TrainGauge: 10}
	goodsTrain := &Train{TrainGauge: 15}

	canPassengerTrainRun := railRoute.IsCompatibleGauge(passengerTrain)
	canGoodsTrainRun := railRoute.IsCompatibleGauge(goodsTrain)

	fmt.Println("Can passenger train run on new route: ", canPassengerTrainRun)
	fmt.Println("Cang googs train run on new route: ", canGoodsTrainRun)

	fmt.Printf("Type of passengerTrain : %T\n", passengerTrain)
}
