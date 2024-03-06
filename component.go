package sk

import "reflect"

type IComponent interface {
	Id() ID
	Name() string
	Type() reflect.Type
	New() any

	onCreated(any)
	onAttached(any, *EntityInstance)
	onDetached(any, *EntityInstance)
}

// -----------------------------------------------------------------------------
// COMPONENT DEFINITION
// -----------------------------------------------------------------------------
var nextComponentDefinitionId ID = 0

type ComponentDefinition[T any] struct {
	id   ID
	name string
	tp   reflect.Type
	def  *T

	OnCreated  func(*T)
	OnAttached func(*T, *EntityInstance)
	OnDetached func(*T, *EntityInstance)
}

func NewComponent[T any](default_ ...T) *ComponentDefinition[T] {
	var t T
	tp := reflect.TypeOf(t)
	name := tp.String()
	id := nextComponentDefinitionId

	var def *T = nil
	if len(default_) > 0 {
		def = &default_[0]
	}

	nextComponentDefinitionId++
	return &ComponentDefinition[T]{
		id:   id,
		name: name,
		tp:   tp,
		def:  def,
	}
}

func (c *ComponentDefinition[T]) New() any {
	inst := reflect.New(c.tp)
	if c.def != nil {
		def := reflect.ValueOf(c.def)
		for i := 0; i < def.Elem().NumField(); i++ {
			inst.Elem().Field(i).Set(def.Elem().Field(i))
		}
	}

	instance := inst.Interface()
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

func (c *ComponentDefinition[T]) Has(entity *EntityInstance) bool {
	return entity.HasComponent(c.id)
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
