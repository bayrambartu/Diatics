# Zoo Simülation

This project is a basic ecosystem simulation written in Go, where various animals interact in a virtual environment.
Creatures move randomly, hunt, and reproduce under defined biological rules.
---

## Features

-Creatures move randomly across a 2D field.
-Predators (Wolf, Lion, Hunter) can hunt specific species.
-Reproducible species (Sheep, Cow, Chicken, Wolf, Lion) can give birth to offspring under correct conditions.
-Each entity has a reproduction cooldown to avoid instant repeated births.
-Detailed simulation reports are printed every 100 steps.
---

## Project Structure
```plaintext
Diatics/
├── main.go               # Entry point
├── model/
│   └── entity.go         # Entity definition, species/gender enums, movement logic
├── animal/               # Behaviors and types of all animals
│   ├── interface.go      # LivingBeing, Predator, Reproducible interfaces
│   ├── factory.go        # Constructors for each animal (NewSheep, NewLion, etc.)
│   ├── movement.go       # Move() and GetEntity() implementations
│   ├── hunting.go        # Hunting behavior for predators
│   └── reproduction.go   # Shared reproduction logic and dummy methods
├── sim/
│   └── simloop.go        # Simulation loop and report logic
└── go.mod                # Go module definition

---
```
## How to Run

```bash
go run main.go
```

## Sample Output
```bash
Starting Zoo Simulation...
Initial population count: 79

========== Step 100 Report ==========
Events:
- HUNT: Lion -> Cow: 1 times
- HUNT: Lion -> Sheep: 1 times
- BIRTH: Sheep: 2 times
- BIRTH: Chicken: 1 times
- BIRTH: Cow: 1 times

- SHEEP  : 29
- COW    : 9
- CHICKEN: 10
- ROOSTER: 10
- WOLF   : 10
- LION   : 8
- HUNTER : 1
Total living beings: 78
```
