package sk

import (
	"log"
	"reflect"
	"slices"
)

var nextQueryDefinitionId ID = 0

type IQuery interface {
	Id() ID
	Name() string
	addInstance(entity *EntityInstance)
	removeInstance(entity *EntityInstance)
}

type queryPart struct {
	tp       reflect.Type
	name     string
	optional bool
}

// -----------------------------------------------------------------------------
// QUERY DEFINITION
// -----------------------------------------------------------------------------
type QueryDefinition[T any] struct {
	id        ID
	name      string
	tp        reflect.Type
	parts     []queryPart
	cache     []*T
	entityMap map[ID]*T
}

func NewQuery[T any]() *QueryDefinition[T] {
	var t T
	tp := reflect.TypeOf(t)
	name := tp.String()
	id := nextQueryDefinitionId
	parts := make([]queryPart, 0)

	for i := 0; i < tp.NumField(); i++ {
		field := tp.Field(i)
		tag := field.Tag.Get("sk")
		parts = append(parts, queryPart{
			tp:       field.Type,
			name:     field.Name,
			optional: tag == "optional",
		})
	}

	nextQueryDefinitionId++
	return &QueryDefinition[T]{
		id:        id,
		name:      name,
		tp:        tp,
		parts:     parts,
		cache:     make([]*T, 0),
		entityMap: make(map[ID]*T),
	}
}

func (c *QueryDefinition[T]) Id() ID {
	return c.id
}

func (c *QueryDefinition[T]) Name() string {
	return c.name
}

func (c *QueryDefinition[T]) Query() []*T {
	return c.cache
}

func (c *QueryDefinition[T]) addInstance(entity *EntityInstance) {
	if !c.match(entity.definition.components) {
		return
	}
	inst := c.build(entity)
	c.cache = append(c.cache, inst)
	c.entityMap[entity.id] = inst
}

func (c *QueryDefinition[T]) removeInstance(entity *EntityInstance) {
	inst, ok := c.entityMap[entity.id]
	if !ok {
		return
	}

	delete(c.entityMap, entity.id)
	idx := slices.Index(c.cache, inst)
	c.cache = slices.Delete(c.cache, idx, idx+1)
}

func (c *QueryDefinition[T]) match(components []IComponent) bool {
	for _, part := range c.parts {
		found := false
		for _, component := range components {
			if compareTypes(component.Type(), part.tp) {
				found = true
				break
			}
		}
		if !found && !part.optional {
			return false
		}
	}

	return true
}

func (c *QueryDefinition[T]) build(entity *EntityInstance) *T {
	inst := reflect.New(c.tp).Interface().(*T)

	for _, part := range c.parts {
		for _, component := range entity.components {
			if compareTypes(reflect.TypeOf(component), part.tp) {
				setValue(inst, part.name, component)
				break
			}
		}
	}

	return inst
}

func compareTypes(a, b reflect.Type) bool {
	ak := a.Kind()
	bk := b.Kind()

	for ak == reflect.Ptr {
		a = a.Elem()
		ak = a.Kind()
	}

	for bk == reflect.Ptr {
		b = b.Elem()
		bk = b.Kind()
	}

	return a == b
}

func setValue(obj any, field string, value any) {
	ref := reflect.ValueOf(obj)

	// if its a pointer, resolve its value
	if ref.Kind() == reflect.Ptr {
		ref = reflect.Indirect(ref)
	}

	if ref.Kind() == reflect.Interface {
		ref = ref.Elem()
	}

	// should double check we now have a struct (could still be anything)
	if ref.Kind() != reflect.Struct {
		log.Fatal("unexpected type")
	}

	prop := ref.FieldByName(field)
	prop.Set(reflect.ValueOf(value))
}
