package sk

import "reflect"

type Components []IComponent

type IComponent interface {
	Id() ID
	Name() string
	Type() reflect.Type
	New() any

	getDependencies() []IComponent
	onCreated(any)
	onAttached(any, *EntityInstance)
	onDetached(any, *EntityInstance)
}

// -----------------------------------------------------------------------------
// COMPONENT DEFINITION
// -----------------------------------------------------------------------------
var nextComponentDefinitionId ID = 0

type ComponentDefinition[T any] struct {
	id           ID
	name         string
	tp           reflect.Type
	dependencies []IComponent
	OnCreated    func(*T)
	OnAttached   func(*T, *EntityInstance)
	OnDetached   func(*T, *EntityInstance)
}

type ComponentOptions[T any] struct {
	Dependencies []IComponent
	OnCreated    func(*T)
	OnAttached   func(*T, *EntityInstance)
	OnDetached   func(*T, *EntityInstance)
}

func NewComponent[T any]() *ComponentDefinition[T] {
	var t T
	tp := reflect.TypeOf(t)
	name := tp.String()
	id := nextComponentDefinitionId

	nextComponentDefinitionId++
	return &ComponentDefinition[T]{
		id:           id,
		name:         name,
		tp:           tp,
		dependencies: []IComponent{},
	}
}

func NewComponentWithOptions[T any](options ComponentOptions[T]) *ComponentDefinition[T] {
	c := NewComponent[T]()
	c.dependencies = options.Dependencies
	c.OnCreated = options.OnCreated
	c.OnAttached = options.OnAttached
	c.OnDetached = options.OnDetached
	return c
}

func (c *ComponentDefinition[T]) New() any {
	typeInstance := reflect.New(c.tp)
	instance := typeInstance.Interface()
	c.onCreated(instance)
	return instance
}

func (c *ComponentDefinition[T]) Id() ID {
	return c.id
}

func (c *ComponentDefinition[T]) Name() string {
	return c.name
}

func (c *ComponentDefinition[T]) Type() reflect.Type {
	return c.tp
}

func (c *ComponentDefinition[T]) Get(entity *EntityInstance) *T {
	return entity.GetComponent(c.id).(*T)
}

func (c *ComponentDefinition[T]) In(entity *EntityInstance) bool {
	return entity.HasComponent(c.id)
}

func (c *ComponentDefinition[T]) getDependencies() []IComponent {
	return c.dependencies
}

func (c *ComponentDefinition[T]) onCreated(cp any) {
	if c.OnCreated != nil {
		c.OnCreated(cp.(*T))
	}
}

func (c *ComponentDefinition[T]) onAttached(cp any, e *EntityInstance) {
	if c.OnAttached != nil {
		c.OnAttached(cp.(*T), e)
	}
}

func (c *ComponentDefinition[T]) onDetached(cp any, e *EntityInstance) {
	if c.OnDetached != nil {
		c.OnDetached(cp.(*T), e)
	}
}
