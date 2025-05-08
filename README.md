# 2D game - Space Invaders

This project is an attempt to create a simple clone of the popular game space invadors.
The game is one of the recommended projects sugested by Austin Z. Henley in his blog [Challenging projects every programmer should try](https://austinhenley.com/blog/challengingprojects.html).

## Basic Design

In his blog, Ausin suggest to the reader that they focus on implementing a "well-defined fame from start to finish" rather that "getting bogged down" by the game design and art. To avoid this, we will first implement the game in a very rudamentory way (i.e., using simple shapes to represent elements such as the player, bullets and enemies/invaders).

...

## Choosin a Graphic Library

In his blog, Austin recommends using a barebones 2D graphics library; I will be using `raylib-go`, a lightweight Go binding of the C-Based `raylib` library. This library is minimalist and easy to setup - it dos not hid any of the lov-level details behind opaic functions.

## Implementation

Below, I've detalied the steps I've taken to implement this game from scratch.

### Project Setup

To setup this project we have to create a project workspace, then install the neccessery Go libraies.

#### 1. Install Go

This can be done with a simple brew install, as shown below:

```sh
brew install go
```

#### 2. Setup workspace

Given that Go follows a strict workspace structure, we'll need to structure our project directory as follows:

```plain
space-invaders/
    |-- bin/
    |-- src/
        |-- space-invaders/
    |-- pkg/
```

#### 3. Initialise Go module

To track out project's dependencies we will need our project module (a `go.mod` file). We can initialise our poject module by running the following command inside of `src/space-invaders`:

```sh
go mod init github.com/TobiAdeniyi/space-invaders
```

#### 4. Installing Graphic Library

To install the Go binding of the C-Based library, `raylib`, we need to first install the C library followed by it's accomponying Go bindings.

```sh
brew install raylib # installing the C library
```

Then:

```sh
go get -u github.com/gen2brain/raylib-go/raylib # installing the Go-bindings
```

### Game Architecture

...

### Basic Game Feature

...

### Drawings

...

## Additional Features

...

## Testing

...
