package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func handleKeyboard() {
	keys := []ebiten.Key{
		ebiten.KeyW,
		ebiten.KeyS,
		ebiten.KeyA,
		ebiten.KeyD,
		ebiten.KeyUp,
		ebiten.KeyDown,
		ebiten.KeyLeft,
		ebiten.KeyF,
		ebiten.KeyE,
		ebiten.KeyK,
		ebiten.KeyP,
		ebiten.KeyH,
		ebiten.KeyR,
		ebiten.KeyI,
		ebiten.KeyKP1,
		ebiten.KeyKP2,
		ebiten.KeyKP3,
		ebiten.KeyKP4,
		ebiten.KeyN,
		ebiten.KeyG,
		ebiten.KeyU,
		ebiten.KeyL,
		ebiten.KeyMinus,
	}

	for _, key := range keys {
		if ebiten.IsKeyPressed(key) {
			keyStates[key]++
		} else {
			keyStates[key] = 0
		}
	}
}

func IsKeyTriggered(key ebiten.Key) bool {
	return keyStates[key] == 1
}
