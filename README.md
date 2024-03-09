# SKALD ENGINE

This is a 2D ECS game engine based on [Ebitenengine](https://ebitengine.org/) for Golang. I created Skald for personal usage, but I will keep it open hoping that would be useful for other people.

Features:

- **Performatic ECS**: although my goal is not performance, the ECS implementation is quite lightweight and can handle dozens of thousands individual entities at same time.
- **Command-Based Input System**: use commands to abstract the actions of the game, using multiple devices at same time for complex combinations. Change the shortcuts in realtime as needed.
- **Asset Loading**: load assets asynchrounously.

In development:

- Camera
- Integrated physics
- Integrated simple collisions
- Plugin systems
- Tilemaps
- Texture Atlas
- Asset pipeline
- UI
- Particle System
- Event System

Roadmap:

1. Example: pong Game

	 The idea is to create a single game to feel the development flow of the engine and work on real necessities, moreover, working as documentation and showcase.

   Dependencies:

	 - input system
	 - simple collisions
	 - text
	 - mobile build
	 - web build
	 - PC build

