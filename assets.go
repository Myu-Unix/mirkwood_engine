package main

import (
	"bytes"
	"image"
	_ "image/png"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

type assets struct {
	images *imageAssets
	fonts  *fontAssets
}

type fontAssets struct {
	mplusTitleFont        font.Face
	mplusLargeFont        font.Face
	mplusNormalFont       font.Face
	mplusSmallFont        font.Face
	mplusNotificationFont font.Face
	mplusMiniFont         font.Face
}

type imageAssets struct {
	background1Image  *ebiten.Image
	background2Image  *ebiten.Image
	splashImage       *ebiten.Image
	adventurer1Image  *ebiten.Image
	adventurer2Image  *ebiten.Image
	header1Image      *ebiten.Image
	header2Image      *ebiten.Image
	enemy1Image       *ebiten.Image
	enemy2Image       *ebiten.Image
	enemy3Image       *ebiten.Image
	enemy4Image       *ebiten.Image
	inventoryImage    *ebiten.Image
	dice20Image       *ebiten.Image
	dice4Image        *ebiten.Image
	dice6Image        *ebiten.Image
	dice8Image        *ebiten.Image
	hideImage         *ebiten.Image
	dmImage           *ebiten.Image
	notificationImage *ebiten.Image
}

func newImageFromEmbed(path string) (*ebiten.Image, error) {
	b, err := assetFS.ReadFile(path)
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(b)
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	ebImg := ebiten.NewImageFromImage(img)
	return ebImg, nil
}

func loadAssets() (*assets, error) {
	fonts, err := loadFonts()
	if err != nil {
		return nil, err
	}

	imgs, err := loadImages()
	if err != nil {
		return nil, err
	}
	return &assets{
		fonts:  fonts,
		images: imgs,
	}, nil
}

func customFontSize(tt *truetype.Font, size float64) font.Face {
	const dpi = 72
	return truetype.NewFace(tt, &truetype.Options{Size: size, DPI: dpi, Hinting: font.HintingFull})
}

func loadFonts() (*fontAssets, error) {
	b, err := assetFS.ReadFile("fonts/harabara.ttf")
	if err != nil {
		return nil, err
	}

	tt, err := truetype.Parse(b)
	if err != nil {
		return nil, err
	}
	return &fontAssets{
		mplusTitleFont:        customFontSize(tt, 96),
		mplusLargeFont:        customFontSize(tt, 72),
		mplusNormalFont:       customFontSize(tt, 48),
		mplusSmallFont:        customFontSize(tt, 24),
		mplusNotificationFont: customFontSize(tt, 20),
		mplusMiniFont:         customFontSize(tt, 14),
	}, nil
}

func loadImages() (*imageAssets, error) {
	a := imageAssets{}
	matches := map[string]**ebiten.Image{
		"images/background.png":      &a.background1Image,
		"images/hide.png":            &a.hideImage,
		"images/avatar.png":          &a.adventurer1Image,
		"images/avatar2.png":         &a.adventurer2Image,
		"images/player1_header2.png": &a.header1Image,
		"images/player2_header2.png": &a.header2Image,
		"images/goblin.png":          &a.enemy1Image,
		"images/warg.png":            &a.enemy2Image,
		"images/skeleton64.png":      &a.enemy3Image,
		"images/skeleton64_axe.png":  &a.enemy4Image,
		"images/splash.png":          &a.splashImage,
		"images/inventory.png":       &a.inventoryImage,
		"images/dice20.png":          &a.dice20Image,
		"images/dice4.png":           &a.dice4Image,
		"images/dice6.png":           &a.dice6Image,
		"images/dice8.png":           &a.dice8Image,
		"images/dm.png":              &a.dmImage,
		"images/notification.png":    &a.notificationImage,
	}

	var err error
	for path, img := range matches {
		*img, err = newImageFromEmbed(path)
		if err != nil {
			return nil, err
		}
	}
	return &a, nil
}
