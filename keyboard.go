package main

import (
  "github.com/hajimehoshi/ebiten/v2"
)

func keyboard_handler() {
        // Handle keypress with Ebiten
        if ebiten.IsKeyPressed(ebiten.KeyW) {
            keyStates[ebiten.KeyW]++
        } else {
            keyStates[ebiten.KeyW] = 0
        }
        if ebiten.IsKeyPressed(ebiten.KeyS) {
            keyStates[ebiten.KeyS]++
        } else {
            keyStates[ebiten.KeyS] = 0
        }
        if ebiten.IsKeyPressed(ebiten.KeyA) {
            keyStates[ebiten.KeyA]++
        } else {
            keyStates[ebiten.KeyA] = 0
        }
        if ebiten.IsKeyPressed(ebiten.KeyD) {
            keyStates[ebiten.KeyD]++
        } else {
            keyStates[ebiten.KeyD] = 0
        }
        if ebiten.IsKeyPressed(ebiten.KeyUp) {
            keyStates[ebiten.KeyUp]++
        } else {
            keyStates[ebiten.KeyUp] = 0
        }
        if ebiten.IsKeyPressed(ebiten.KeyDown) {
            keyStates[ebiten.KeyDown]++
        } else {
            keyStates[ebiten.KeyDown] = 0
        }
        if ebiten.IsKeyPressed(ebiten.KeyLeft) {
            keyStates[ebiten.KeyLeft]++
        } else {
            keyStates[ebiten.KeyLeft] = 0
        }
        if ebiten.IsKeyPressed(ebiten.KeyRight) {
            keyStates[ebiten.KeyRight]++
        } else {
            keyStates[ebiten.KeyRight] = 0
        }
        if ebiten.IsKeyPressed(ebiten.KeyF) {
            keyStates[ebiten.KeyF]++
        } else {
            keyStates[ebiten.KeyF] = 0
        }
        if ebiten.IsKeyPressed(ebiten.KeyE) {
            keyStates[ebiten.KeyE]++
        } else {
            keyStates[ebiten.KeyE] = 0
        }
        if ebiten.IsKeyPressed(ebiten.KeyK) {
            keyStates[ebiten.KeyK]++
        } else {
            keyStates[ebiten.KeyK] = 0
        }
        if ebiten.IsKeyPressed(ebiten.KeyP) { // Switch Player
            keyStates[ebiten.KeyP]++
        } else {
            keyStates[ebiten.KeyP] = 0
        }
        if ebiten.IsKeyPressed(ebiten.KeyH) {
            keyStates[ebiten.KeyH]++
        } else {
            keyStates[ebiten.KeyH] = 0
        }
        if ebiten.IsKeyPressed(ebiten.KeyR) {
            keyStates[ebiten.KeyR]++
        } else {
            keyStates[ebiten.KeyR] = 0
        }
        if ebiten.IsKeyPressed(ebiten.KeyI) {
            keyStates[ebiten.KeyI]++
        } else {
            keyStates[ebiten.KeyI] = 0
        }
        if ebiten.IsKeyPressed(ebiten.KeyKP1) {
            keyStates[ebiten.KeyKP1]++
        } else {
            keyStates[ebiten.KeyKP1] = 0
        }
        if ebiten.IsKeyPressed(ebiten.KeyKP2) {
            keyStates[ebiten.KeyKP2]++
        } else {
            keyStates[ebiten.KeyKP2] = 0
        }
        if ebiten.IsKeyPressed(ebiten.KeyKP3) {
            keyStates[ebiten.KeyKP3]++
        } else {
            keyStates[ebiten.KeyKP3] = 0
        }
        if ebiten.IsKeyPressed(ebiten.KeyKP4) {
            keyStates[ebiten.KeyKP4]++
        } else {
            keyStates[ebiten.KeyKP4] = 0
        }
        if ebiten.IsKeyPressed(ebiten.KeyN) {
            keyStates[ebiten.KeyN]++
        } else {
            keyStates[ebiten.KeyN] = 0
        }
        if ebiten.IsKeyPressed(ebiten.KeyG) {
            keyStates[ebiten.KeyG]++
        } else {
            keyStates[ebiten.KeyG] = 0
        }
        if ebiten.IsKeyPressed(ebiten.KeyU) {
            keyStates[ebiten.KeyU]++
        } else {
            keyStates[ebiten.KeyU] = 0
        }
        if ebiten.IsKeyPressed(ebiten.KeyL) {
            keyStates[ebiten.KeyL]++
        } else {
            keyStates[ebiten.KeyL] = 0
        }
        if ebiten.IsKeyPressed(ebiten.KeyMinus) {
            keyStates[ebiten.KeyMinus]++
        } else {
            keyStates[ebiten.KeyMinus] = 0
        }
      }

func IsKeyTriggered(key ebiten.Key) bool {
  return keyStates[key] == 1
}