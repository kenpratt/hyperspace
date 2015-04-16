package main

type FireEvent struct {
	PlayerId     string
	Time         uint64
	ProjectileId string
}

type ChangeAccelerationEvent struct {
	PlayerId  string
	Time      uint64
	Direction int8
}

type ChangeRotationEvent struct {
	PlayerId  string
	Time      uint64
	Direction int8
}
