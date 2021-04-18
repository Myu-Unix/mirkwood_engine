package main

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
)

func load_images() {

	var err error
	background1Image, _, err = ebitenutil.NewImageFromFile("images/background.png")
	if err != nil {
		log.Fatal(err)
	}
	hideImage, _, err = ebitenutil.NewImageFromFile("images/hide.png")
	if err != nil {
		log.Fatal(err)
	}
	adventurer1Image, _, err = ebitenutil.NewImageFromFile("images/avatar.png")
	if err != nil {
		log.Fatal(err)
	}
	adventurer2Image, _, err = ebitenutil.NewImageFromFile("images/avatar2.png")
	if err != nil {
		log.Fatal(err)
	}
	header1Image, _, err = ebitenutil.NewImageFromFile("images/player1_header2.png")
	if err != nil {
		log.Fatal(err)
	}
	header2Image, _, err = ebitenutil.NewImageFromFile("images/player2_header2.png")
	if err != nil {
		log.Fatal(err)
	}
	enemy1Image, _, err = ebitenutil.NewImageFromFile("images/goblin.png")
	if err != nil {
		log.Fatal(err)
	}
	enemy2Image, _, err = ebitenutil.NewImageFromFile("images/warg.png")
	if err != nil {
		log.Fatal(err)
	}
	enemy3Image, _, err = ebitenutil.NewImageFromFile("images/skeleton64.png")
	if err != nil {
		log.Fatal(err)
	}
	enemy4Image, _, err = ebitenutil.NewImageFromFile("images/skeleton64_axe.png")
	if err != nil {
		log.Fatal(err)
	}
	SplashImage, _, err = ebitenutil.NewImageFromFile("images/splash.png")
	if err != nil {
		log.Fatal(err)
	}
	inventoryImage, _, err = ebitenutil.NewImageFromFile("images/inventory.png")
	if err != nil {
		log.Fatal(err)
	}
	dice20Image, _, err = ebitenutil.NewImageFromFile("images/dice20.png")
	if err != nil {
		log.Fatal(err)
	}
	dice4Image, _, err = ebitenutil.NewImageFromFile("images/dice4.png")
	if err != nil {
		log.Fatal(err)
	}
	dice6Image, _, err = ebitenutil.NewImageFromFile("images/dice6.png")
	if err != nil {
		log.Fatal(err)
	}
	dice8Image, _, err = ebitenutil.NewImageFromFile("images/dice8.png")
	if err != nil {
		log.Fatal(err)
	}
	dmImage, _, err = ebitenutil.NewImageFromFile("images/dm.png")
	if err != nil {
		log.Fatal(err)
	}
	notificationImage, _, err = ebitenutil.NewImageFromFile("images/notification.png")
	if err != nil {
		log.Fatal(err)
	}
}
