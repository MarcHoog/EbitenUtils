package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"os"
)

type AssetType int

const (
	AssetImage AssetType = iota
	AssetAudio
)

func (t AssetType) String() string {
	switch t {
	case AssetImage:
		return "image"
	case AssetAudio:
		return "audio"
	default:
		return "unknown"
	}
}

type AssetID int

type AssetInfo struct {
	ID     AssetID
	OsPath string
	Type   AssetType
}

func CreateImage(osPath string) (*ebiten.Image, error) {
	f, err := os.Open(osPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	ebtImg := ebiten.NewImageFromImage(img)
	return ebtImg, nil
}
