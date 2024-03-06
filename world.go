package sk

import "slices"

// -----------------------------------------------------------------------------
// WORLD
// -----------------------------------------------------------------------------
type World struct {
	entities []*EntityInstance
	queries  []IQuery
}

func NewWorld() *World {
	return &World{
		entities: make([]*EntityInstance, 0),
		queries:  make([]IQuery, 0),
	}
}

func (w *World) Entities() []*EntityInstance {
	return w.entities
}

func (w *World) Spawn(def *EntityDefinition, fn ...func(*EntityInstance)) {
	entity := def.New()
	w.entities = append(w.entities, entity)
	for _, f := range fn {
		f(entity)
	}

	if def.OnSpawned != nil {
		def.OnSpawned(entity)
	}

	for _, query := range w.queries {
		query.addInstance(entity)
	}
}

func (w *World) SpawnMulti(n int, def *EntityDefinition, fn ...func(*EntityInstance)) {
	for i := 0; i < n; i++ {
		w.Spawn(def, fn...)
	}
}

func (w *World) Despawn(entity *EntityInstance) {
	idx := slices.Index(w.entities, entity)
	if idx == -1 {
		return
	}

	w.entities = slices.Delete(w.entities, idx, idx+1)
	for _, query := range w.queries {
		query.removeInstance(entity)
	}

	if entity.definition.OnDespawned != nil {
		entity.definition.OnDespawned(entity)
	}
}

func (w *World) AddQuery(query IQuery) {
	if slices.Contains(w.queries, query) {
		return
	}

	w.queries = append(w.queries, query)
	for _, entity := range w.entities {
		query.addInstance(entity)
	}
}

func (w *World) RemoveQuery(query IQuery) {
	idx := slices.Index(w.queries, query)
	if idx == -1 {
		return
	}
	w.queries = slices.Delete(w.queries, idx, idx+1)
}
