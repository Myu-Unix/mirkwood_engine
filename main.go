/* Started 12 April 2021
(C) Myu-Unix, 2021 - MIT Licensed - Assets used with fair use in mind, don't sue me */

package main

import (
	"embed"
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type adventurer struct {
	name           string
	class          string
	race           string
	item1          string // Usually the "main" weapon
	item2          string
	item3          string
	item4          string
	item5          string
	posx           float64
	posy           float64
	hp_max         string
	STR            string
	DEX            string
	CON            string
	INT            string
	WIS            string
	CHA            string
	alignment      string
	ac_armor_class string
}

type enemy struct {
	name           string
	race           string
	hp_max         string
	ac_armor_class string
	item1          string // Usually the enemy weapon
	item2          string
	item3          string
	item4          string
	posx           float64
	posy           float64
	alive          bool
}

var (
	//go:embed fonts
	//go:embed images
	assetFS embed.FS

	MyConfig       adventurer // TEST
	player         [2]adventurer
	npc            [4]enemy
	keyStates      = map[ebiten.Key]int{}
	cmd_run        []byte
	engine_version = "Mirkwood Engine 0.7.0 (Prototype)"
	engine_text    = "Written in Go + Ebiten // Not all those who wander are lost"
)

type config struct {
	fullscreen        bool
	hidden            bool
	showInventory     bool
	link              bool
	dm                bool
	debug             bool
	splash            bool
	header_posx       float64
	notification_posx float64
	rand              *rand.Rand
}

type state struct {
	playerSelected int
	enemySelected  int
	round          int
	d20            int
	d4             int
	d6             int
	d8             int
}

func newConfig() config {
	randSrc := rand.NewSource(time.Now().UnixNano())
	return config{
		fullscreen:        true,
		hidden:            true,
		showInventory:     true,
		dm:                false,
		debug:             false,
		link:              false,
		splash:            true,
		header_posx:       0,
		notification_posx: 1920,
		rand:              rand.New(randSrc),
	}
}

func newState() state {
	return state{
		playerSelected: 1,
		enemySelected:  1,
		round:          0,
		d20:            20,
		d4:             4,
		d6:             6,
		d8:             8,
	}
}

func init() {
	// TODO : Replace by json file config
	player[0] = adventurer{name: "Myu", class: "Level 1 Ranger", race: "Elf", item1: "Elven Shortbow +1 (45m/1d6)", item2: "Elvish Dagger +1 (1d4)", item3: "Leather Armor (AC11)", item4: "Lembas (5)", item5: "Camping supplies", posx: 630, posy: 210, hp_max: "15 HP", STR: "STR 12", DEX: "DEX 14", CON: "CON 13", INT: "INT 12", WIS: "WIS 13", CHA: "CHA 10", alignment: "Chaotic good", ac_armor_class: "AC 13"}
	player[1] = adventurer{name: "Dolph", class: "Level 1 Druid", race: "Elf", item1: "Staff of Adornment (1d6 - 1d8)", item2: "Rope", item3: "Healing Herbs", posx: 560, posy: 280, hp_max: "12 HP", STR: "STR 8", DEX: "DEX 10", CON: "CON 7", INT: "INT 15", WIS: "WIS 14", CHA: "CHA 12", alignment: "Lawful good", ac_armor_class: "AC 10"}
	npc[0] = enemy{name: "Ghaz", race: "Level 1 Goblin", posx: 1200, posy: 700, hp_max: "8 HP", ac_armor_class: "AC 5", item1: "Club (1d4)", alive: true}
	npc[1] = enemy{name: "Dhurg", race: "Level 2 Goblin Warg Rider", posx: 1100, posy: 750, hp_max: "10 HP", ac_armor_class: "AC 7", item1: "Hand-Axe (1d6)", alive: true}
	npc[2] = enemy{name: "Dorg", race: "Level 1 Skeleton Archer", posx: 1150, posy: 800, hp_max: "5 HP", ac_armor_class: "AC 6", item1: "Longbow (1d6)", alive: true}
	npc[3] = enemy{name: "Dakh", race: "Level 1 Skeleton", posx: 1150, posy: 700, hp_max: "6 HP", ac_armor_class: "AC 5", item1: "Hand-Axe (1d6)", alive: true}
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Images options
	opAdventurer1 := &ebiten.DrawImageOptions{}
	opAdventurer2 := &ebiten.DrawImageOptions{}
	opEnemy1 := &ebiten.DrawImageOptions{}
	opEnemy2 := &ebiten.DrawImageOptions{}
	opEnemy3 := &ebiten.DrawImageOptions{}
	opEnemy4 := &ebiten.DrawImageOptions{}
	opBackground := &ebiten.DrawImageOptions{}
	opInventory := &ebiten.DrawImageOptions{}
	opHeader := &ebiten.DrawImageOptions{}
	opDice20 := &ebiten.DrawImageOptions{}
	opDice4 := &ebiten.DrawImageOptions{}
	opDice6 := &ebiten.DrawImageOptions{}
	opDice8 := &ebiten.DrawImageOptions{}
	opHide := &ebiten.DrawImageOptions{}
	opNotification := &ebiten.DrawImageOptions{}
	opAdventurer1.GeoM.Translate(player[0].posx, player[0].posy)
	opAdventurer2.GeoM.Translate(player[1].posx, player[1].posy)
	opEnemy1.GeoM.Translate(npc[0].posx, npc[0].posy)
	opEnemy2.GeoM.Translate(npc[1].posx, npc[1].posy)
	opEnemy3.GeoM.Translate(npc[2].posx, npc[2].posy)
	opEnemy4.GeoM.Translate(npc[3].posx, npc[3].posy)
	opHeader.GeoM.Translate(g.config.header_posx, 32)
	opInventory.GeoM.Translate(g.config.header_posx, 220)
	opDice20.GeoM.Translate(16, 120)
	opDice4.GeoM.Translate(16, 230)
	opDice6.GeoM.Translate(16, 340)
	opDice8.GeoM.Translate(16, 450)
	opHide.GeoM.Translate(1041, 629)
	opNotification.GeoM.Translate(g.config.notification_posx, 16)

	// Draw images
	if g.config.splash { // This shows the splashscreen
		screen.DrawImage(g.assets.images.splashImage, opBackground)
		text.Draw(screen, "~ Into Mirkwood ~", g.assets.fonts.mplusTitleFont, 730, 400, color.White)
		text.Draw(screen, "A short tabletop tutorial campaign", g.assets.fonts.mplusNormalFont, 725, 575, color.White)
		text.Draw(screen, "Myu & Dolph <3", g.assets.fonts.mplusNormalFont, 865, 650, color.White)
		text.Draw(screen, "Press 'p' to start", g.assets.fonts.mplusSmallFont, 1700, 1000, color.White)
	} else {
		// Map background handler
		screen.DrawImage(g.assets.images.background1Image, opBackground)
		if g.config.notification_posx > 32 {
			g.config.notification_posx -= 128
		}
		screen.DrawImage(g.assets.images.notificationImage, opNotification)

		// Draw a line between selected player and target (if alive)
		if g.config.link {
			if npc[g.state.enemySelected-1].alive {
				ebitenutil.DrawLine(screen, player[g.state.playerSelected-1].posx+16, player[g.state.playerSelected-1].posy+32, npc[g.state.enemySelected-1].posx+16, npc[g.state.enemySelected-1].posy+32, color.RGBA{255, 128, 0, 255})
				ebitenutil.DrawLine(screen, player[g.state.playerSelected-1].posx+17, player[g.state.playerSelected-1].posy+33, npc[g.state.enemySelected-1].posx+17, npc[g.state.enemySelected-1].posy+33, color.RGBA{255, 128, 0, 255})
				a := int(npc[g.state.enemySelected-1].posx) - int(player[g.state.playerSelected-1].posx)
				b := int(npc[g.state.enemySelected-1].posy) - int(player[g.state.playerSelected-1].posy)
				// Rough distance in "ft" from pixels
				distance := math.Sqrt(float64((a*a))+float64((b*b))) / 10
				text.Draw(screen, string(strconv.Itoa(int(distance))), g.assets.fonts.mplusSmallFont, int(distance*5+player[g.state.playerSelected-1].posx), int(distance*5+player[g.state.playerSelected-1].posy), color.White)
				text.Draw(screen, "ft", g.assets.fonts.mplusSmallFont, int(distance*5+player[g.state.playerSelected-1].posx+30), int(distance*5+player[g.state.playerSelected-1].posy), color.White)
			}
		}
		// Drawing dices and values
		screen.DrawImage(g.assets.images.dice20Image, opDice20)
		screen.DrawImage(g.assets.images.dice4Image, opDice4)
		screen.DrawImage(g.assets.images.dice6Image, opDice6)
		screen.DrawImage(g.assets.images.dice8Image, opDice8)
		text.Draw(screen, string(strconv.Itoa(g.state.d20)), g.assets.fonts.mplusNormalFont, 140, 200, color.White)
		text.Draw(screen, string(strconv.Itoa(g.state.d4)), g.assets.fonts.mplusNormalFont, 140, 300, color.White)
		text.Draw(screen, string(strconv.Itoa(g.state.d6)), g.assets.fonts.mplusNormalFont, 140, 400, color.White)
		text.Draw(screen, string(strconv.Itoa(g.state.d8)), g.assets.fonts.mplusNormalFont, 140, 500, color.White)
		// Drawing adventurers/players
		screen.DrawImage(g.assets.images.adventurer1Image, opAdventurer1)
		screen.DrawImage(g.assets.images.adventurer2Image, opAdventurer2)
		// Player "token" data
		text.Draw(screen, string(player[g.state.playerSelected-1].name), g.assets.fonts.mplusSmallFont, int(player[g.state.playerSelected-1].posx+48), int(player[g.state.playerSelected-1].posy), color.White)
		// TEST - JSON gathered
		//text.Draw(screen, string(MyConfig.name), mplusSmallFont, int(player[g.state.playerSelected-1].posx+48), int(player[g.state.playerSelected-1].posy), color.White)
		text.Draw(screen, string(player[g.state.playerSelected-1].hp_max), g.assets.fonts.mplusMiniFont, int(player[g.state.playerSelected-1].posx+64), int(player[g.state.playerSelected-1].posy+18), color.White)
		text.Draw(screen, string(player[g.state.playerSelected-1].ac_armor_class), g.assets.fonts.mplusMiniFont, int(player[g.state.playerSelected-1].posx+72), int(player[g.state.playerSelected-1].posy+32), color.White)
		text.Draw(screen, string(player[g.state.playerSelected-1].item1), g.assets.fonts.mplusMiniFont, int(player[g.state.playerSelected-1].posx+72), int(player[g.state.playerSelected-1].posy+46), color.White)

		if g.config.debug {
			text.Draw(screen, engine_version, g.assets.fonts.mplusNormalFont, 40, 960, color.White)
			text.Draw(screen, engine_text, g.assets.fonts.mplusMiniFont, 40, 982, color.White)
			text.Draw(screen, "PLAYER : ", g.assets.fonts.mplusSmallFont, 32, 560, color.White)
			text.Draw(screen, "ENEMY : ", g.assets.fonts.mplusSmallFont, 32, 600, color.White)
			text.Draw(screen, "ROUND : ", g.assets.fonts.mplusSmallFont, 32, 640, color.White)
			text.Draw(screen, strconv.Itoa(g.state.playerSelected), g.assets.fonts.mplusSmallFont, 156, 560, color.White)
			text.Draw(screen, strconv.Itoa(g.state.enemySelected), g.assets.fonts.mplusSmallFont, 156, 600, color.White)
			text.Draw(screen, strconv.Itoa(g.state.round), g.assets.fonts.mplusSmallFont, 156, 640, color.White)
		}

		// If NPC is alive, draw it
		if npc[0].alive {
			screen.DrawImage(g.assets.images.enemy1Image, opEnemy1)
		}
		if npc[1].alive {
			screen.DrawImage(g.assets.images.enemy2Image, opEnemy2)
		}
		if npc[2].alive {
			screen.DrawImage(g.assets.images.enemy3Image, opEnemy3)
		}
		if npc[3].alive {
			screen.DrawImage(g.assets.images.enemy4Image, opEnemy4)
		}

		// INVENTORY CARD
		if g.config.showInventory {
			// Show header animation
			if g.config.header_posx < 1450 {
				g.config.header_posx += 290
			}
			// Show player header image
			if g.state.playerSelected == 1 {
				screen.DrawImage(g.assets.images.header1Image, opHeader)
			} else {
				screen.DrawImage(g.assets.images.header2Image, opHeader)
			}
			screen.DrawImage(g.assets.images.inventoryImage, opInventory)

			text.Draw(screen, string(player[g.state.playerSelected-1].name), g.assets.fonts.mplusNormalFont, 1480, 82, color.White)
			text.Draw(screen, string(player[g.state.playerSelected-1].class), g.assets.fonts.mplusSmallFont, 1480, 114, color.White)
			text.Draw(screen, string(player[g.state.playerSelected-1].hp_max), g.assets.fonts.mplusSmallFont, 1480, 146, color.White)
			text.Draw(screen, string(player[g.state.playerSelected-1].ac_armor_class), g.assets.fonts.mplusSmallFont, 1540, 146, color.White)
			text.Draw(screen, string(player[g.state.playerSelected-1].alignment), g.assets.fonts.mplusSmallFont, 1490, 178, color.White)
			text.Draw(screen, "-- INVENTORY --", g.assets.fonts.mplusNormalFont, 1500, 232, color.White)
			//text.Draw(screen, "Range 3-18", mplusMiniFont, 1720, 50, color.White)
			text.Draw(screen, string(player[g.state.playerSelected-1].STR), g.assets.fonts.mplusMiniFont, 1770, 70, color.White)
			text.Draw(screen, string(player[g.state.playerSelected-1].DEX), g.assets.fonts.mplusMiniFont, 1770, 90, color.White)
			text.Draw(screen, string(player[g.state.playerSelected-1].CON), g.assets.fonts.mplusMiniFont, 1770, 110, color.White)
			text.Draw(screen, string(player[g.state.playerSelected-1].INT), g.assets.fonts.mplusMiniFont, 1770, 130, color.White)
			text.Draw(screen, string(player[g.state.playerSelected-1].WIS), g.assets.fonts.mplusMiniFont, 1770, 150, color.White)
			text.Draw(screen, string(player[g.state.playerSelected-1].CHA), g.assets.fonts.mplusMiniFont, 1770, 170, color.White)
			text.Draw(screen, string(player[g.state.playerSelected-1].item1), g.assets.fonts.mplusSmallFont, 1532, 270, color.White)
			text.Draw(screen, string(player[g.state.playerSelected-1].item2), g.assets.fonts.mplusSmallFont, 1532, 310, color.White)
			text.Draw(screen, string(player[g.state.playerSelected-1].item3), g.assets.fonts.mplusSmallFont, 1532, 350, color.White)
			text.Draw(screen, string(player[g.state.playerSelected-1].item4), g.assets.fonts.mplusSmallFont, 1532, 390, color.White)
			text.Draw(screen, string(player[g.state.playerSelected-1].item5), g.assets.fonts.mplusSmallFont, 1532, 430, color.White)
			//text.Draw(screen, string(player[g.state.playerSelected-1].item6), mplusSmallFont, 1532, 470, color.White)
		} // INVENTORY CARD END

		// Show/hide enemy data
		if npc[g.state.enemySelected-1].alive {
			text.Draw(screen, string(npc[g.state.enemySelected-1].race), g.assets.fonts.mplusSmallFont, int(npc[g.state.enemySelected-1].posx+48), int(npc[g.state.enemySelected-1].posy-10), color.White)
			text.Draw(screen, string(npc[g.state.enemySelected-1].hp_max), g.assets.fonts.mplusMiniFont, int(npc[g.state.enemySelected-1].posx+64), int(npc[g.state.enemySelected-1].posy+18), color.White)
			text.Draw(screen, string(npc[g.state.enemySelected-1].ac_armor_class), g.assets.fonts.mplusMiniFont, int(npc[g.state.enemySelected-1].posx+72), int(npc[g.state.enemySelected-1].posy+32), color.White)
			text.Draw(screen, string(npc[g.state.enemySelected-1].item1), g.assets.fonts.mplusMiniFont, int(npc[g.state.enemySelected-1].posx+72), int(npc[g.state.enemySelected-1].posy+46), color.White)
		}

		// "For of war"/hidden roof for map 1
		if g.config.hidden {
			screen.DrawImage(g.assets.images.hideImage, opHide)
		}

		// Notification for round
		if g.state.round == 0 {
			text.Draw(screen, "Setting the scene !", g.assets.fonts.mplusNotificationFont, 72, 72, color.White)
			text.Draw(screen, "DM explains the scene and/or what happens next.", g.assets.fonts.mplusNotificationFont, 72, 94, color.White)
		} else if g.state.round == 1 {
			text.Draw(screen, "Movement - Up to your speed", g.assets.fonts.mplusNotificationFont, 72, 72, color.White)
			text.Draw(screen, "Interaction - i.e opening a door, sheathing a weapon", g.assets.fonts.mplusNotificationFont, 72, 94, color.White)
		} else if g.state.round == 2 {
			text.Draw(screen, "Action - Attack, Dash, Improvise, Hide, Search, ...", g.assets.fonts.mplusNotificationFont, 72, 72, color.White)
			text.Draw(screen, "Combat resolution", g.assets.fonts.mplusNotificationFont, 72, 94, color.White)
		}

		// DM cheat sheet
		if g.config.dm {
			screen.DrawImage(g.assets.images.dmImage, opBackground)
			text.Draw(screen, "--- SUPERSIMPLIFIED COMBAT RULES (WIP) ---", g.assets.fonts.mplusSmallFont, 32, 32, color.RGBA{255, 128, 0, 255})
			text.Draw(screen, "Is anyone surprised ? If you surprise an enemy, you'll have an additional turn.", g.assets.fonts.mplusSmallFont, 32, 82, color.White)
			text.Draw(screen, "Everyone rolls initiative (1d20 + initiative modifier) and the one with highest start first", g.assets.fonts.mplusSmallFont, 32, 132, color.White)
			text.Draw(screen, "On your turn, you can move a distance up to your speed and take 1 action", g.assets.fonts.mplusSmallFont, 32, 182, color.White)
			text.Draw(screen, "To attack, roll a d20 and add weapons modifiers and check that against AC value", g.assets.fonts.mplusSmallFont, 32, 214, color.White)
			text.Draw(screen, "Then to it, roll the dice from you weapon (i.e 1d6)", g.assets.fonts.mplusSmallFont, 32, 246, color.White)
			text.Draw(screen, "--- SKILL CHECKS/SAVINGS THROWS (1d20) ---", g.assets.fonts.mplusSmallFont, 32, 320, color.RGBA{255, 128, 0, 255})
			text.Draw(screen, "DM can ask for a skill check before a player can process with an action. This is resolved with a D20 roll +/- modifiers", g.assets.fonts.mplusSmallFont, 32, 370, color.White)
			text.Draw(screen, "DM can ask for a saving throw based on abilities. Must resolve the difficulty (DC) set by the DM or else fail", g.assets.fonts.mplusSmallFont, 32, 420, color.White)
			text.Draw(screen, "DC : 5 = very easy / 10 = Easy / 15 = Moderate / 20 = Hard / 25 = Very Hard", g.assets.fonts.mplusSmallFont, 32, 470, color.White)
			text.Draw(screen, "--- MIRKWOOD ENGINE KEYBOARD SHORTCUTS ---", g.assets.fonts.mplusSmallFont, 32, 520, color.RGBA{255, 128, 0, 255})
			text.Draw(screen, "P to switch player - R to roll a dice - Z/S/Q/D to move player - K to quit - up/down/left/right to move ennemies - e to switch enemies", g.assets.fonts.mplusSmallFont, 32, 570, color.White)
			text.Draw(screen, "K to quit - up/down/left/right to move ennemies - e to switch enemies - L to link", g.assets.fonts.mplusSmallFont, 32, 620, color.White)
			text.Draw(screen, "I - Show inventory/character panel - U DM info - KP1/KP2/KP3/KP4 to 'kill' enemies 1/2/3/4 - N for next round - G debug info", g.assets.fonts.mplusSmallFont, 32, 670, color.White)
			text.Draw(screen, "PRESS 'U' to open/close this panel :)", g.assets.fonts.mplusLargeFont, 500, 900, color.White)
		}
	}
	return
}

func (g *Game) Update() error {
	// Handle single keypress with Ebiten
	handleKeyboard()

	// Handle keypress and set states
	g.handleState()

	return nil
}

type Game struct {
	config *config
	state  *state
	assets *assets
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1920, 1080
}

func main() {
	cfg := newConfig()
	state := newState()

	a, err := loadAssets()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(engine_version)

	// TEST
	err = readConfigPlayer1()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(MyConfig.name))

	ebiten.SetFullscreen(true)
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle(engine_version)
	game := Game{
		config: &cfg,
		state:  &state,
		assets: a,
	}
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
