package sim

import (
	"fmt"
	"math/rand"
	"time"

	"Diatics/animal"
	"Diatics/model"
)

func RunSimulation() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Starting Zoo Simulation...")

	var beings []animal.LivingBeing

	// Constructor map
	breedingModels := map[model.Species]func(model.Gender) animal.LivingBeing{
		model.SHEEP: func(g model.Gender) animal.LivingBeing { s := animal.NewSheep(g); return &s },
		model.COW:   func(g model.Gender) animal.LivingBeing { c := animal.NewCow(g); return &c },
		model.WOLF:  func(g model.Gender) animal.LivingBeing { w := animal.NewWolf(g); return &w },
		model.LION:  func(g model.Gender) animal.LivingBeing { l := animal.NewLion(g); return &l },
		model.CHICKEN: func(g model.Gender) animal.LivingBeing {
			if g == model.Male {
				r := animal.NewRooster()
				return &r
			}
			ch := animal.NewChicken()
			return &ch
		},
	}
	breedingModels[model.ROOSTER] = breedingModels[model.CHICKEN]

	// Initial Population
	for i := 0; i < 15; i++ {
		m, f := animal.NewSheep(model.Male), animal.NewSheep(model.Female)
		beings = append(beings, &m, &f)
	}
	for i := 0; i < 5; i++ {
		m, f := animal.NewCow(model.Male), animal.NewCow(model.Female)
		beings = append(beings, &m, &f)
	}
	for i := 0; i < 10; i++ {
		ch := animal.NewChicken()
		r := animal.NewRooster()
		beings = append(beings, &ch, &r)
	}
	for i := 0; i < 5; i++ {
		m, f := animal.NewWolf(model.Male), animal.NewWolf(model.Female)
		beings = append(beings, &m, &f)
	}
	for i := 0; i < 4; i++ {
		m, f := animal.NewLion(model.Male), animal.NewLion(model.Female)
		beings = append(beings, &m, &f)
	}
	h := animal.NewHunter()
	beings = append(beings, &h)

	fmt.Printf("Initial population count: %d\n", len(beings))

	eventCounters := make(map[string]int)

	for step := 1; step <= model.StepLimit; step++ {
		newborns := []animal.LivingBeing{}
		reproduced := make(map[animal.LivingBeing]bool)

		// 1. Move and cooldown tick
		for _, b := range beings {
			entity := b.GetEntity()
			if entity.ReproductionCooldown > 0 {
				entity.ReproductionCooldown--
			}
			b.Move()
		}

		// 2. Interactions
		for i, b1 := range beings {
			if !b1.GetEntity().IsAlive {
				continue
			}

			// Hunting
			if hunter, ok := b1.(animal.Predator); ok {
				for _, b2 := range beings {
					if b1 != b2 && b2.GetEntity().IsAlive {
						if hunter.Hunt(b2) {
							key := fmt.Sprintf("HUNT: %s -> %s", b1.GetEntity().Species, b2.GetEntity().Species)
							eventCounters[key]++
							break
						}
					}
				}
			}

			// Reproduction
			if _, ok := b1.(animal.Reproducible); ok {
				if reproduced[b1] {
					continue
				}
				for j := i + 1; j < len(beings); j++ {
					b2 := beings[j]
					if reproduced[b2] {
						continue
					}
					if baby, success := animal.CommonReproduction(b1, b2, breedingModels); success {
						newborns = append(newborns, baby)
						reproduced[b1] = true
						reproduced[b2] = true
						key := fmt.Sprintf("BIRTH: %s", baby.GetEntity().Species)
						eventCounters[key]++
						break
					}
				}
			}
		}

		// 3. Update list
		survivors := []animal.LivingBeing{}
		for _, b := range beings {
			if b.GetEntity().IsAlive {
				survivors = append(survivors, b)
			}
		}
		beings = append(survivors, newborns...)

		// 4. Reporting
		if step%100 == 0 || step == model.StepLimit {
			fmt.Printf("\n========== Step %d Report ==========\n", step)
			if len(eventCounters) > 0 {
				fmt.Println("Events:")
				for key, count := range eventCounters {
					fmt.Printf("- %s: %d times\n", key, count)
				}
			} else {
				fmt.Println("No significant events this round.")
			}
			eventCounters = make(map[string]int)

			// Summary
			summary := make(map[model.Species]int)
			for _, b := range beings {
				summary[b.GetEntity().Species]++
			}
			order := []model.Species{model.SHEEP, model.COW, model.CHICKEN, model.ROOSTER, model.WOLF, model.LION, model.HUNTER}
			for _, sp := range order {
				fmt.Printf("- %-7s: %d\n", sp, summary[sp])
			}
			fmt.Printf("Total living beings: %d\n", len(beings))
		}
	}

	fmt.Println("\n====================================")
	fmt.Println("       SIMULATION COMPLETE")
	fmt.Println("====================================")
}
