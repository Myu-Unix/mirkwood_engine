package main

import (
  "github.com/hajimehoshi/ebiten"
  "math/rand"
  "os"
  "time"
)

func state_handler() {
	   // ---- PLAYER MOVE ----
       if IsKeyTriggered(ebiten.KeyW) == true {
         player[STATE_PLAYER_SELECTED-1].posy = player[STATE_PLAYER_SELECTED-1].posy - 70
       }
       if IsKeyTriggered(ebiten.KeyS) == true {
        player[STATE_PLAYER_SELECTED-1].posy = player[STATE_PLAYER_SELECTED-1].posy + 70
   	   }
       if IsKeyTriggered(ebiten.KeyA) == true {
        player[STATE_PLAYER_SELECTED-1].posx = player[STATE_PLAYER_SELECTED-1].posx - 70
   	   }
       if IsKeyTriggered(ebiten.KeyD) == true {
        player[STATE_PLAYER_SELECTED-1].posx = player[STATE_PLAYER_SELECTED-1].posx + 70
   	   }
   	   // ---- ENEMY MOVE ----
   	   if IsKeyTriggered(ebiten.KeyUp) == true {
        npc[STATE_ENEMY_SELECTED-1].posy = npc[STATE_ENEMY_SELECTED-1].posy - 70
   	   }
       if IsKeyTriggered(ebiten.KeyDown) == true {
        npc[STATE_ENEMY_SELECTED-1].posy = npc[STATE_ENEMY_SELECTED-1].posy + 70
   	   }
       if IsKeyTriggered(ebiten.KeyLeft) == true {
        npc[STATE_ENEMY_SELECTED-1].posx = npc[STATE_ENEMY_SELECTED-1].posx - 70
   	   }
       if IsKeyTriggered(ebiten.KeyRight) == true {
        npc[STATE_ENEMY_SELECTED-1].posx = npc[STATE_ENEMY_SELECTED-1].posx + 70
   	   }
       if IsKeyTriggered(ebiten.KeyF) == true {
       	if STATE_FULLSCREEN == 0 {
         ebiten.SetFullscreen(true)
         STATE_FULLSCREEN=1
        } else {
        	ebiten.SetFullscreen(false)
        	STATE_FULLSCREEN=0
        }
       }
       if IsKeyTriggered(ebiten.KeyP) == true {
       	STATE_SHOW_SPLASH = 0
        header_posx = 0
       	go click_sound()
       if STATE_PLAYER_SELECTED==1 {
       	STATE_PLAYER_SELECTED=2
       } else {
       	STATE_PLAYER_SELECTED=1
       }
       }
       if IsKeyTriggered(ebiten.KeyU) == true {
        go click_sound()
       if STATE_DM==1 {
        STATE_DM=0
       } else {
        STATE_DM=1
       }
       }
       if IsKeyTriggered(ebiten.KeyL) == true {
         go click_sound()
         if STATE_LINK==1 {
           STATE_LINK=0
         } else {
           STATE_LINK=1
         }
       }
       if IsKeyTriggered(ebiten.KeyE) == true {
       	go click_sound()
        if STATE_ENEMY_SELECTED==1 {
       	  STATE_ENEMY_SELECTED=2
        } else if STATE_ENEMY_SELECTED==2 {
          STATE_ENEMY_SELECTED=3
        } else if STATE_ENEMY_SELECTED==3 {
          STATE_ENEMY_SELECTED=4
        } else {
       	  STATE_ENEMY_SELECTED=1
        }
      }
       if IsKeyTriggered(ebiten.KeyG) == true {
       	go click_sound()
        if STATE_SHOW_DEBUG==1 {
       	  STATE_SHOW_DEBUG = 0
        } else {
       	  STATE_SHOW_DEBUG = 1
        }
       }

       if IsKeyTriggered(ebiten.KeyI) == true {
         go click_sound()
         header_posx = 0
         if STATE_SHOW_INVENTORY==1 {
          STATE_SHOW_INVENTORY = 0
         } else {
          STATE_SHOW_INVENTORY = 1
         }
       }
       // Quit
       if IsKeyTriggered(ebiten.KeyK) == true {
         os.Exit(0)
       }

       // Hidden area on the map
       if IsKeyTriggered(ebiten.KeyH) == true {
         STATE_HIDDEN=0
       }

       if IsKeyTriggered(ebiten.KeyKP1) == true {
         npc[0].alive = 0
         STATE_LINK=0
       }
       if IsKeyTriggered(ebiten.KeyKP2) == true {
         npc[1].alive = 0
         STATE_LINK=0
       }
       if IsKeyTriggered(ebiten.KeyKP3) == true {
         npc[2].alive = 0
         STATE_LINK=0
       }
       if IsKeyTriggered(ebiten.KeyKP4) == true {
         npc[3].alive = 0
         STATE_LINK=0
       }
       // Remove health from enemies WIP
       /*if IsKeyTriggered(ebiten.KeyMinus) == true {
         npc[STATE_ENEMY_SELECTED].hp_max = strconv(npc[STATE_ENEMY_SELECTED].hp_max) -= 1
       } */

       if IsKeyTriggered(ebiten.KeyN) == true {
        notification_posx = 1920 
         go click_sound()
         if STATE_ROUND < 2 {
          STATE_ROUND +=1
         } else {
          STATE_ROUND = 0
         }
       }

       // Next map
       /*if IsKeyTriggered(ebiten.KeyN) == true {
         go click_sound()
         header_posx = 0
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

       if IsKeyTriggered(ebiten.KeyR) == true { // roll dices
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
}