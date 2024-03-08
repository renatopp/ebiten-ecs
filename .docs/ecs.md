# ECS

- [ECS](#ecs)
	- [Basics](#basics)
	- [Components](#components)
	- [Entities](#entities)
	- [Archetypes](#archetypes)
	- [Queries](#queries)
	- [Systems](#systems)


ECS, or Entity-Component-System, is an architectural pattern used to represent and process objects within the game world. It separates data and behavior into three main elements:

- **Components**: essentially data containers that hold specific aspects of an object, like its position, health, or appearance.

- **Entities**: individual objects within the game, formed by combining various components. For instance, an enemy entity might have "position," "health," and "enemy behavior" components.

- **Systems**: act like processors, operating on entities that possess specific combinations of components. For example, a "render system" might look for entities with "position" and "appearance" components to draw them on the screen, while a "movement system" might update the positions of entities with "position" and "movement behavior" components.

The ECS architecture provides a series of benefits for our games:

- **Reusability**: Components can be combined and reused across different entities, reducing code duplication and promoting modularity.
- **Flexibility**: New features can be added by creating new components or systems, simplifying the modification and extension of the game.
- **Performance**: ECS can improve performance by allowing for efficient data access and manipulation through components.
- **Maintainability**: Separating data, behavior, and logic can make code easier to understand, maintain, and debug.
- **Parallelization**: ECS can facilitate parallel processing, potentially improving performance in situations where multiple systems can operate concurrently.

## Basics

Creates the component purse, notice that we gotta encapsulate the struct into a `ComponentDefinition` object.

```go
type DPurse struct {
	Coins int
}

var Purse = sk.NewComponent[DPurse]()
```

With a component created, now we can create an entity containing it:

```go
var Player = sk.NewEntity(Health)
```

Notice that this `Player` is now an `EntityDefinition` type, which serve us as a "template" for every player instance in the game.

Before developing a system, we can create a query object, so we can cache its results and run the system at its maximum performance.

```go
var DebitQuery = sk.NewQuery[struct {
	Purse *DPurse
}]()
```

Defining the system, adding the query as dependency

```go
var Debit = sk.NewSystem(func (g *Game) error {
	for _, r := DebitQuery.Query() {
		r.Purse.Coins++
	}
}, DebitQuery)
```

Register everything and the start the game

```go
var game = sk.NewGame()
game.AddSystem(Debit)
game.World.Spawn(Player)
game.Play()
```

Notice that we spawn the player (i.e. we create an instance of the Player entity) in the world, which serve us as a container for everything that encompases the game world. We can also understand the world object as a kind of database that we can `Query` over.

## Components

Components are the building blocks of an ECS architecture. Usually they contain only data, which can be easily serializable. However, **Skald Engine** allow you to add any kind of method with it.

```go
type DPurse struct {
	Coins int
}

func (d *DPurse) Debit(value int) {
	d.Coins += value
}

func (d *DPurse) Credit(value int) {
	d.Coins = math.max(Coins - value, 0)
}
```

> **Important!** Only one component of each type is allowed in the same entity. Duplicated components are ignored.

Now imagine that you want to initialize all purses with a predefined number of coins, you can do it by passing a default object as template:

```go
type DPurse struct {
	Coins int
}

var Purse = sk.NewComponent(DPurse {
	Coins: 10
})
```

Alternatively, you can provide a "constructor" function that will be called as the component is created within the entity:

```go
type DPurse struct {
	Coins int
}

var Purse = sk.NewComponent[DPurse]()
Purse.OnCreated = func(c *DPurse) {
	c.Coins = 10
}
```

You may use them interchangeable, but the `OnCreated` hook will be called after, thus, overriding any default value provided in the template.

Aiming to provide a flexible solution, you may also define some hooks in the components that will be called automatically after certain events, like attaching the component to an entity.

- `OnCreated(c *T)` called when the component is created, within the entity instantiation or standalone.
- `OnAttached(c *T, e *EntityInstance)` called when the component is attached to an entity instance.
- `OnDettached(c *T, e *EntityInstance)` called when the component is removed from an entity instance.

## Entities

Entities are simple objects that contain components. You can create a new entity definition as:

```go
Player := sk.NewEntity(Transform, Sprite, Collision)
```

where `Transform`, `Sprite` and `Collision` are all components. Notice that at this moment, `Player` is an entity "template", or, as it is called internally, an `EntityDefinition`. This object will help us instantiate all Player entities in the game:

```go
game.world.Spawn(Player)
```

The spawn function will instantiate the entity and include it in the world. You may want to customize the individual entity upon creation (instead of changing the definition of all of them), to this, you can provide a setup function:

```go
game.world.SpawnMulti(10, Player, func (e *EntityInstance) {
	transform := Transform.Get(e)
	transform.position.x = rand.Int()
	transform.position.y = rand.Int()
})
```

Alternatively, you may provide the hooks for:

- `OnSpawned(e *EntityInstance)` called when the entity joins the world.
- `OnDespawned(e *EntityInstance)` called when the entity is removed from the world.

## Archetypes

Archetypes is one common topic in ECS engines. In **Skald Engine**, we don't make much use of them. Instead, you may simply define an archetype as a collection of Components and use them to create the entity:

```go
moverArchetype := []IComponent { Transform, Velocity, Collider }
player := sk.NewEntity(moverArchetype...)
```

## Queries

Before using system, it is important to understand the Queries. In **Skald Engine**, queries serve as index for data, just like in a database. The query will cache all components it needs every time an entity enters or leave a world, making the process of querying a specific group of components very performance. However, similar to a database, the price of this performance is memory. Internally, the query holds a reference for every component described in it.

```go
var DebitQuery = sk.NewQuery[struct {
	Purse   *DPurse
	Account *DAccount `sk:"optional"`
}]()
```

The query above will cache every entity that has the `DPurse` component and, if present, the `DAccount` component.

You can register the query manually in the world:

```go
game.world.AddQuery(DebitQuery)
```

or insert it indirectly as dependency of a system.

## Systems

Systems are processors. Simple functions that execute the game logic upon a group of entities. For example, you can write a system that just prints a `Hello, World` without manipulating any entity:

```go
helloSystem := sk.NewSystem(func (g *sk.Game) error {
	println("Hello, World")
})
```

You can also have a system that looks for the Purse component in every entity in the world:

```go
slowSystem := sk.NewSystem(func (g *sk.Game) error {
	for _, e := g.world.Entities() {
		if p := Purse.Get(e); p != nil {
			p.Coins++
		}
	}
})
```

This `slowSystem`, as informed by its name, does its job but does very poorly in terms of performance, since it must check in every object of the world for a component that just a couple objects may have. In order to improve it, you can add a Query for caching the components:

```go
fastSystem := sk.NewSystem(func (g *sk.Game) error {
	for _, r := DebitQuery.Query() {
		r.Purse.Coins++
	}
}, DebitQuery)


game.AddSystem(fastSystem)
```

By adding the `fastSystem` in the game, **Skald Engine** automatically adds all the dependent queries in the world.

