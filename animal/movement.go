package animal

import "Diatics/model"

// Forwarding methods for all species to satisfy LivingBeing interface

func (s *Sheep) GetEntity() *model.Entity { return &s.Entity }
func (s *Sheep) Move()                    { s.Entity.Move() }

func (c *Cow) GetEntity() *model.Entity { return &c.Entity }
func (c *Cow) Move()                    { c.Entity.Move() }

func (c *Chicken) GetEntity() *model.Entity { return &c.Entity }
func (c *Chicken) Move()                    { c.Entity.Move() }

func (r *Rooster) GetEntity() *model.Entity { return &r.Entity }
func (r *Rooster) Move()                    { r.Entity.Move() }

func (w *Wolf) GetEntity() *model.Entity { return &w.Entity }
func (w *Wolf) Move()                    { w.Entity.Move() }

func (l *Lion) GetEntity() *model.Entity { return &l.Entity }
func (l *Lion) Move()                    { l.Entity.Move() }

func (h *Hunter) GetEntity() *model.Entity { return &h.Entity }
func (h *Hunter) Move()                    { h.Entity.Move() }
