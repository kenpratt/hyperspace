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
	Direction int8
}

type ChangeRotationEvent struct {
	PlayerId  string
	Time      uint64
	Direction int8
}
