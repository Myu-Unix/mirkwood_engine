package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) state_handler() {
	// Move selected player
	if IsKeyTriggered(ebiten.KeyW) {
		player[g.state.playerSelected-1].posy -= 70
	}
	if IsKeyTriggered(ebiten.KeyS) {
		player[g.state.playerSelected-1].posy += 70
	}
	if IsKeyTriggered(ebiten.KeyA) {
		player[g.state.playerSelected-1].posx -= 70
	}
	if IsKeyTriggered(ebiten.KeyD) {
		player[g.state.playerSelected-1].posx += 70
	}
	// Move selected enemy
	if IsKeyTriggered(ebiten.KeyUp) {
		npc[g.state.enemySelected-1].posy -= 70
	}
	if IsKeyTriggered(ebiten.KeyDown) {
		npc[g.state.enemySelected-1].posy += 70
	}
	if IsKeyTriggered(ebiten.KeyLeft) {
		npc[g.state.enemySelected-1].posx -= 70
	}
	if IsKeyTriggered(ebiten.KeyRight) {
		npc[g.state.enemySelected-1].posx += 70
	}
	// Toogle fullscreen
	if IsKeyTriggered(ebiten.KeyF) {
		g.config.fullscreen = !g.config.fullscreen
		ebiten.SetFullscreen(g.config.fullscreen)
	}
	// Player choice
	if IsKeyTriggered(ebiten.KeyP) {
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
	if IsKeyTriggered(ebiten.KeyU) {
		go click_sound()
		g.config.dm = !g.config.dm
	}
	// Link/Measure
	if IsKeyTriggered(ebiten.KeyL) {
		go click_sound()
		g.config.link = !g.config.link
	}
	// Select enemy
	if IsKeyTriggered(ebiten.KeyE) {
		go click_sound()
		if g.state.enemySelected < 4 {
			g.state.enemySelected += 1
		} else {
			g.state.enemySelected = 1
		}
	}
	// Show some debug info
	if IsKeyTriggered(ebiten.KeyG) {
		go click_sound()
		g.config.debug = !g.config.debug
	}
	// Toogle inventory
	if IsKeyTriggered(ebiten.KeyI) {
		go click_sound()
		g.config.header_posx = 0
		g.config.showInventory = !g.config.showInventory
	}
	// Quit
	if IsKeyTriggered(ebiten.KeyK) {
		os.Exit(0)
	}

	// Hidden area on the map
	if IsKeyTriggered(ebiten.KeyH) {
		g.config.hidden = !g.config.hidden
	}
	// Kill enemy (temporary)
	if IsKeyTriggered(ebiten.KeyKP1) {
		npc[0].alive = false
		g.config.link = false
	}
	if IsKeyTriggered(ebiten.KeyKP2) {
		npc[1].alive = false
		g.config.link = false
	}
	if IsKeyTriggered(ebiten.KeyKP3) {
		npc[2].alive = false
		g.config.link = false
	}
	if IsKeyTriggered(ebiten.KeyKP4) {
		npc[3].alive = false
		g.config.link = false
	}
	// Remove health from enemies WIP
	/*if IsKeyTriggered(ebiten.KeyMinus) {
	  npc[g.state.enemySelected].hp_max = strconv(npc[g.state.enemySelected].hp_max) -= true
	} */

	// Change game round
	if IsKeyTriggered(ebiten.KeyN) {
		g.config.notification_posx = 1920
		go click_sound()
		if g.state.round < 2 {
			g.state.round += 1
		} else {
			g.state.round = 0
		}
	}
	// Dices be rollin'
	if IsKeyTriggered(ebiten.KeyR) { // roll dices
		go dice_sound()
		time.Sleep(80 * time.Millisecond)
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		time.Sleep(80 * time.Millisecond)
		s2 := rand.NewSource(time.Now().UnixNano())
		r2 := rand.New(s2)
		time.Sleep(80 * time.Millisecond)
		s3 := rand.NewSource(time.Now().UnixNano())
		r3 := rand.New(s3)
		time.Sleep(80 * time.Millisecond)
		s4 := rand.NewSource(time.Now().UnixNano())
		r4 := rand.New(s4)
		DICE_20_1 = r1.Intn(20) + 1
		DICE_4_1 = r2.Intn(4) + 1
		DICE_6_1 = r3.Intn(6) + 1
		DICE_8_1 = r4.Intn(8) + 1
	}

	// Next map - Disabled
	/*if IsKeyTriggered(ebiten.KeyN) {
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
