package animal

import (
	"Diatics/model"
	"math/rand"
)

// --- Type Definitions (embedded Entity) ---
type Sheep struct{ model.Entity }
type Cow struct{ model.Entity }
type Chicken struct{ model.Entity }
type Rooster struct{ model.Entity }
type Wolf struct{ model.Entity }
type Lion struct{ model.Entity }
type Hunter struct{ model.Entity }

// --- Shared constructor for any entity ---
func newEntity(species model.Species, gender model.Gender, speed float64) model.Entity {
	return model.Entity{
		Species: species,
		Gender:  gender,
		X:       rand.Float64() * model.FieldWidth,
		Y:       rand.Float64() * model.FieldHeight,
		Speed:   speed,
		IsAlive: true,
	}
}

// --- Factory functions for each species ---
func NewSheep(g model.Gender) Sheep { return Sheep{Entity: newEntity(model.SHEEP, g, 2.0)} }
func NewCow(g model.Gender) Cow     { return Cow{Entity: newEntity(model.COW, g, 2.0)} }
func NewChicken() Chicken           { return Chicken{Entity: newEntity(model.CHICKEN, model.Female, 1.0)} }
func NewRooster() Rooster           { return Rooster{Entity: newEntity(model.ROOSTER, model.Male, 1.0)} }
func NewWolf(g model.Gender) Wolf   { return Wolf{Entity: newEntity(model.WOLF, g, 3.0)} }
func NewLion(g model.Gender) Lion   { return Lion{Entity: newEntity(model.LION, g, 4.0)} }
func NewHunter() Hunter             { return Hunter{Entity: newEntity(model.HUNTER, model.Male, 1.0)} }
