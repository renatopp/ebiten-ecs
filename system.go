package sk

type SystemFn func(*Game) error

// -----------------------------------------------------------------------------
// SYSTEM ENTRY
// -----------------------------------------------------------------------------
type systemEntry struct {
	definition *SystemDefinition
	options    *SystemOptions
}

// -----------------------------------------------------------------------------
// SYSTEM OPTIONS
// -----------------------------------------------------------------------------
type SystemOptions struct {
	once     bool
	priority int
}

func NewSystemOptions() *SystemOptions {
	return &SystemOptions{
		once:     false,
		priority: 100,
	}
}

func (sb *SystemOptions) Once() *SystemOptions {
	sb.once = true
	return sb
}

func (sb *SystemOptions) Priority(priority int) *SystemOptions {
	sb.priority = priority
	return sb
}

// -----------------------------------------------------------------------------
// SYSTEM DEFINITION
// -----------------------------------------------------------------------------
var nextSystemDefinitionId ID = 0

type SystemDefinition struct {
	id      ID
	fn      SystemFn
	queries []IQuery
}

func NewSystem(fn SystemFn, queries ...IQuery) *SystemDefinition {
	s := &SystemDefinition{
		id:      nextSystemDefinitionId,
		fn:      fn,
		queries: queries,
	}

	nextSystemDefinitionId++
	return s
}

func (s *SystemDefinition) Id() ID {
	return s.id
}
