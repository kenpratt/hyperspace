package main

type FireEvent struct {
	PlayerId     string
	Time         float64
	ProjectileId string
}

type ChangeAccelerationEvent struct {
	PlayerId  string
	Time      float64
	Direction float64 // TODO: try to use an int8 here (and try to switch other types to better types instead of float64 as well)
}

type ChangeRotationEvent struct {
	PlayerId  string
	Time      float64
	Direction float64 // TODO: try to use an int8 here (and try to switch other types to better types instead of float64 as well)
}
