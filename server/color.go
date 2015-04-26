package main

import (
	"github.com/lucasb-eyer/go-colorful"
	"math/rand"
)

func RandomBrightColor() string {
	c := colorful.Hcl(rand.Float64()*360.0, rand.Float64(), 0.8 + rand.Float64()*0.2)
	return c.Hex()
}
