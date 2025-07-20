package model

import (
	"math"
	"math/rand"
)

// --- Constants ---
const (
	FieldWidth  = 500
	FieldHeight = 500
	StepLimit   = 1000
)

// --- Gender Enum ---
type Gender int

const (
	Male Gender = iota
	Female
)

// --- Species Enum ---
type Species string

const (
	SHEEP   Species = "Sheep"
	COW     Species = "Cow"
	CHICKEN Species = "Chicken"
	ROOSTER Species = "Rooster"
	WOLF    Species = "Wolf"
	LION    Species = "Lion"
	HUNTER  Species = "Hunter"
)

// --- Core Entity ---
type Entity struct {
	Species              Species
	Gender               Gender
	X, Y                 float64
	Speed                float64
	IsAlive              bool
	ReproductionCooldown int
}

// --- Movement logic ---
func (e *Entity) Move() {
	if !e.IsAlive {
		return
	}

	dx := rand.Float64()*2 - 1
	dy := rand.Float64()*2 - 1
	e.X += dx * e.Speed
	e.Y += dy * e.Speed

	if e.X < 0 {
		e.X = 0
	}
	if e.X >= FieldWidth {
		e.X = FieldWidth - 1
	}
	if e.Y < 0 {
		e.Y = 0
	}
	if e.Y >= FieldHeight {
		e.Y = FieldHeight - 1
	}
}

// --- Distance calculation ---
func CalculateDistance(e1, e2 *Entity) float64 {
	return math.Sqrt(math.Pow(e1.X-e2.X, 2) + math.Pow(e1.Y-e2.Y, 2))
}
