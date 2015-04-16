package main

type FireEvent struct {
	PlayerId     string
	Time         uint64
	ProjectileId string
	Created      uint64
}

type ChangeAccelerationEvent struct {
	PlayerId  string
	Time      uint64
	Direction int8 // TODO: try to use an int8 here (and try to switch other types to better types instead of float64 as well)
}

type ChangeRotationEvent struct {
	PlayerId  string
	Time      uint64
	Direction int8 // TODO: try to use an int8 here (and try to switch other types to better types instead of float64 as well)
}
