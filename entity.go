package sk

import "slices"

// -----------------------------------------------------------------------------
// ENTITY DEFINITION
// -----------------------------------------------------------------------------
var nextEntityDefinitionId ID = 0
var nextEntityInstanceId ID = 0

type EntityDefinition struct {
	id         ID
	components []IComponent

	OnSpawned   func(*EntityInstance)
	OnDespawned func(*EntityInstance)
}

type EntityOptions struct {
	Components  []IComponent
	OnSpawned   func(*EntityInstance)
	OnDespawned func(*EntityInstance)
}

func NewEntity(components ...IComponent) *EntityDefinition {
	e := &EntityDefinition{
		id:         nextEntityDefinitionId,
		components: components,
	}

	nextEntityDefinitionId++
	return e
}

func NewEntityWithOptions(options EntityOptions) *EntityDefinition {
	e := NewEntity(options.Components...)
	e.OnSpawned = options.OnSpawned
	e.OnDespawned = options.OnDespawned
	return e
}

func (e *EntityDefinition) New() *EntityInstance {
	instance := &EntityInstance{
		id:         nextEntityInstanceId,
		definition: e,
		components: make(map[ID]any),
	}

	for _, compDef := range e.components {
		instance.AddComponent(compDef)
	}

	nextEntityInstanceId++
	return instance
}

func (e *EntityDefinition) Id() ID {
	return e.id
}

func (e *EntityDefinition) Components() []IComponent {
	return e.components
}

func (e *EntityDefinition) AddComponent(c IComponent) bool {
	if e.HasComponent(c) {
		return false
	}

	e.components = append(e.components, c)
	return true
}

func (e *EntityDefinition) RemoveComponent(c IComponent) {
	idx := slices.Index(e.components, c)
	if idx == -1 {
		return
	}

	e.components = slices.Delete(e.components, idx, idx+1)
}

func (e *EntityDefinition) HasComponent(c IComponent) bool {
	return slices.Contains(e.components, c)
}

// -----------------------------------------------------------------------------
// ENTITY INSTANCE
// -----------------------------------------------------------------------------
type EntityInstance struct {
	id         ID
	definition *EntityDefinition
	components map[ID]any
}

func (e *EntityInstance) Id() ID {
	return e.id
}

func (e *EntityInstance) Definition() *EntityDefinition {
	return e.definition
}

func (e *EntityInstance) GetComponent(id ID) any {
	return e.components[id]
}

func (e *EntityInstance) HasComponent(id ID) bool {
	_, ok := e.components[id]
	return ok
}

func (e *EntityInstance) AddComponent(c IComponent) bool {
	if e.HasComponent(c.Id()) {
		return false
	}

	for _, dep := range c.getDependencies() {
		e.AddComponent(dep)
	}

	comp := c.New()
	e.components[c.Id()] = comp
	c.onAttached(comp, e)

	return true
}

func (e *EntityInstance) RemoveComponent(c IComponent) {
	if !e.HasComponent(c.Id()) {
		return
	}

	comp := e.components[c.Id()]
	delete(e.components, c.Id())
	c.onDetached(comp, e)

}
