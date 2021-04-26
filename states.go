package main

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) handleState() {
	// Move selected player
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		player[g.state.playerSelected-1].posy -= 70
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		player[g.state.playerSelected-1].posy += 70
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		player[g.state.playerSelected-1].posx -= 70
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		player[g.state.playerSelected-1].posx += 70
	}
	// Move selected enemy
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		npc[g.state.enemySelected-1].posy -= 70
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		npc[g.state.enemySelected-1].posy += 70
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		npc[g.state.enemySelected-1].posx -= 70
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		npc[g.state.enemySelected-1].posx += 70
	}
	// Toogle fullscreen
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		g.config.fullscreen = !g.config.fullscreen
		ebiten.SetFullscreen(g.config.fullscreen)
	}
	// Player choice
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		g.config.splash = false
		g.config.header_posx = 0
		go click_sound()
		if g.state.playerSelected < 2 {
			g.state.playerSelected += 1
		} else {
			g.state.playerSelected = 1
		}
	}
	// DM screen
	if inpututil.IsKeyJustPressed(ebiten.KeyU) {
		go click_sound()
		g.config.dm = !g.config.dm
	}
	// Link/Measure
	if inpututil.IsKeyJustPressed(ebiten.KeyL) {
		go click_sound()
		g.config.link = !g.config.link
	}
	// Select enemy
	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		go click_sound()
		if g.state.enemySelected < 4 {
			g.state.enemySelected += 1
		} else {
			g.state.enemySelected = 1
		}
	}
	// Show some debug info
	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		go click_sound()
		g.config.debug = !g.config.debug
	}
	// Toogle inventory
	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		go click_sound()
		g.config.header_posx = 0
		g.config.showInventory = !g.config.showInventory
	}
	// Quit
	if inpututil.IsKeyJustPressed(ebiten.KeyK) {
		os.Exit(0)
	}

	// Hidden area on the map
	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		g.config.hidden = !g.config.hidden
	}
	// Kill enemy (temporary)
	if inpututil.IsKeyJustPressed(ebiten.KeyKP1) {
		npc[0].alive = false
		g.config.link = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyKP2) {
		npc[1].alive = false
		g.config.link = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyKP3) {
		npc[2].alive = false
		g.config.link = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyKP4) {
		npc[3].alive = false
		g.config.link = false
	}
	// Remove health from enemies WIP
	/*if inpututil.IsKeyJustPressed(ebiten.KeyMinus) {
	  npc[g.state.enemySelected].hp_max = strconv(npc[g.state.enemySelected].hp_max) -= true
	} */

	// Change game round
	if inpututil.IsKeyJustPressed(ebiten.KeyN) {
		g.config.notification_posx = 1920
		go click_sound()
		if g.state.round < 2 {
			g.state.round += 1
		} else {
			g.state.round = 0
		}
	}
	// Dices be rollin'
	if inpututil.IsKeyJustPressed(ebiten.KeyR) { // roll dices
		go dice_sound()
		g.state.d20 = g.config.rand.Intn(20) + 1
		g.state.d4 = g.config.rand.Intn(4) + 1
		g.state.d6 = g.config.rand.Intn(6) + 1
		g.state.d8 = g.config.rand.Intn(8) + 1
	}

	// Next map - Disabled
	/*if inpututil.IsKeyJustPressed(ebiten.KeyN) {
	  go click_sound()
	  g.config.header_posx = 0
	  STATE_MAP=2
	  npc[0].alive=1
	  npc[1].alive=1
	  npc[2].alive=1
	  npc[3].alive=1
	  player[0].posx = 240
	  player[0].posy = 250
	  player[1].posx = 180
	  player[1].posy = 340
	  npc[0].posx = 1240
	  npc[0].posy = 650
	  npc[1].posx = 1180
	  npc[1].posy = 840
	  npc[2].posx = 1240
	  npc[2].posy = 250
	  npc[3].posx = 1580
	  npc[3].posy = 340
	} */
}
