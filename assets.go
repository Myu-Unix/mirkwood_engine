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
	var err error
	a.background1Image, err = newImageFromEmbed("images/background.png")
	if err != nil {
		return nil, err
	}
	a.hideImage, err = newImageFromEmbed("images/hide.png")
	if err != nil {
		return nil, err
	}
	a.adventurer1Image, err = newImageFromEmbed("images/avatar.png")
	if err != nil {
		return nil, err
	}
	a.adventurer2Image, err = newImageFromEmbed("images/avatar2.png")
	if err != nil {
		return nil, err
	}
	a.header1Image, err = newImageFromEmbed("images/player1_header2.png")
	if err != nil {
		return nil, err
	}
	a.header2Image, err = newImageFromEmbed("images/player2_header2.png")
	if err != nil {
		return nil, err
	}
	a.enemy1Image, err = newImageFromEmbed("images/goblin.png")
	if err != nil {
		return nil, err
	}
	a.enemy2Image, err = newImageFromEmbed("images/warg.png")
	if err != nil {
		return nil, err
	}
	a.enemy3Image, err = newImageFromEmbed("images/skeleton64.png")
	if err != nil {
		return nil, err
	}
	a.enemy4Image, err = newImageFromEmbed("images/skeleton64_axe.png")
	if err != nil {
		return nil, err
	}
	a.splashImage, err = newImageFromEmbed("images/splash.png")
	if err != nil {
		return nil, err
	}
	a.inventoryImage, err = newImageFromEmbed("images/inventory.png")
	if err != nil {
		return nil, err
	}
	a.dice20Image, err = newImageFromEmbed("images/dice20.png")
	if err != nil {
		return nil, err
	}
	a.dice4Image, err = newImageFromEmbed("images/dice4.png")
	if err != nil {
		return nil, err
	}
	a.dice6Image, err = newImageFromEmbed("images/dice6.png")
	if err != nil {
		return nil, err
	}
	a.dice8Image, err = newImageFromEmbed("images/dice8.png")
	if err != nil {
		return nil, err
	}
	a.dmImage, err = newImageFromEmbed("images/dm.png")
	if err != nil {
		return nil, err
	}
	a.notificationImage, err = newImageFromEmbed("images/notification.png")
	if err != nil {
		return nil, err
	}
	return &a, nil
}
