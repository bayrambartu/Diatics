package animal

import "Diatics/model"

// All living beings must implement basic movement and data access
type LivingBeing interface {
	Move()
	GetEntity() *model.Entity
}

// Only predators implement this interface
type Predator interface {
	Hunt(target LivingBeing) bool
}

// Only species that can reproduce implement this interface
type Reproducible interface {
	Reproduce(partner LivingBeing) (LivingBeing, bool)
}
