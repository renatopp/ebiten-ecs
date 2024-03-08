package sk

import "reflect"

type IService interface {
	Id() ID
	Name() string
	Type() reflect.Type
	New() any
}

// -----------------------------------------------------------------------------
// SERVICE DEFINITION
// -----------------------------------------------------------------------------
var nextServiceDefinitionId ID = 0

type ServiceDefinition[T any] struct {
	id   ID
	name string
	tp   reflect.Type
	val  *T

	OnCreated func(*T)
}

func NewService[T any](default_ ...T) *ServiceDefinition[T] {
	var t T
	tp := reflect.TypeOf(t)
	name := tp.String()
	id := nextServiceDefinitionId

	var val *T = nil
	if len(default_) > 0 {
		val = &default_[0]
	}

	nextServiceDefinitionId++
	return &ServiceDefinition[T]{
		id:   id,
		name: name,
		tp:   tp,
		val:  val,
	}
}

func (s *ServiceDefinition[T]) New() any {
	var r *T = nil
	if s.val != nil {
		r = s.val
	} else {
		r = reflect.New(s.tp).Interface().(*T)
	}

	if s.OnCreated != nil {
		s.OnCreated(r)
	}

	return r
}

func (s *ServiceDefinition[T]) Id() ID {
	return s.id
}

func (s *ServiceDefinition[T]) Name() string {
	return s.name
}

func (s *ServiceDefinition[T]) Type() reflect.Type {
	return s.tp
}

func (s *ServiceDefinition[T]) Get(g *Game) *T {
	return g.services[s.id].(*T)
}

func (s *ServiceDefinition[T]) In(g *Game) bool {
	_, ok := g.services[s.id]
	return ok
}
