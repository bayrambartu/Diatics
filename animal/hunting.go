package animal

import "Diatics/model"

// Wolf can hunt Sheep, Chicken, Rooster
func (predator *Wolf) Hunt(target LivingBeing) bool {
	targetEntity := target.GetEntity()
	if !targetEntity.IsAlive {
		return false
	}
	t := targetEntity.Species
	if t == model.SHEEP || t == model.CHICKEN || t == model.ROOSTER {
		if model.CalculateDistance(&predator.Entity, targetEntity) <= 4.0 {
			targetEntity.IsAlive = false
			return true
		}
	}
	return false
}

// Lion can hunt Cow and Sheep
func (predator *Lion) Hunt(target LivingBeing) bool {
	targetEntity := target.GetEntity()
	if !targetEntity.IsAlive {
		return false
	}
	t := targetEntity.Species
	if t == model.COW || t == model.SHEEP {
		if model.CalculateDistance(&predator.Entity, targetEntity) <= 5.0 {
			targetEntity.IsAlive = false
			return true
		}
	}
	return false
}

// Hunter can hunt any other creature (except itself)
func (predator *Hunter) Hunt(target LivingBeing) bool {
	targetEntity := target.GetEntity()
	if !targetEntity.IsAlive || &predator.Entity == targetEntity {
		return false
	}
	if model.CalculateDistance(&predator.Entity, targetEntity) <= 8.0 {
		targetEntity.IsAlive = false
		return true
	}
	return false
}
