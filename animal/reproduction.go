package animal

import (
	"Diatics/model"
	"math/rand"
)

// Dummy implementations — required to fulfill the Reproducible interface

func (s *Sheep) Reproduce(partner LivingBeing) (LivingBeing, bool)   { return nil, false }
func (s *Cow) Reproduce(partner LivingBeing) (LivingBeing, bool)     { return nil, false }
func (s *Chicken) Reproduce(partner LivingBeing) (LivingBeing, bool) { return nil, false }
func (s *Rooster) Reproduce(partner LivingBeing) (LivingBeing, bool) { return nil, false }
func (s *Wolf) Reproduce(partner LivingBeing) (LivingBeing, bool)    { return nil, false }
func (s *Lion) Reproduce(partner LivingBeing) (LivingBeing, bool)    { return nil, false }

// --- General Reproduction Logic ---

func CommonReproduction(
	parent1, parent2 LivingBeing,
	factory map[model.Species]func(model.Gender) LivingBeing,
) (LivingBeing, bool) {
	e1, e2 := parent1.GetEntity(), parent2.GetEntity()

	// Reproduction cooldown check
	if e1.ReproductionCooldown > 0 || e2.ReproductionCooldown > 0 {
		return nil, false
	}

	// Species & gender compatibility
	if e1.Species == e2.Species && e1.Gender != e2.Gender && e1.IsAlive && e2.IsAlive {
		isChickenRoosterPair := (e1.Species == model.CHICKEN && e2.Species == model.ROOSTER) ||
			(e1.Species == model.ROOSTER && e2.Species == model.CHICKEN)

		if e1.Species != e2.Species && !isChickenRoosterPair {
			return nil, false
		}

		// Distance check
		if model.CalculateDistance(e1, e2) <= 3.0 {
			newGender := model.Gender(rand.Intn(2))
			childSpecies := e1.Species
			if isChickenRoosterPair {
				childSpecies = model.CHICKEN
			}

			if builder, ok := factory[childSpecies]; ok {
				child := builder(newGender)
				child.GetEntity().X = e1.X
				child.GetEntity().Y = e1.Y

				// Cooldowns
				e1.ReproductionCooldown = 50
				e2.ReproductionCooldown = 50
				child.GetEntity().ReproductionCooldown = 100

				return child, true
			}
		}
	}

	return nil, false
}
