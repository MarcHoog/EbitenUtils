package assets

import (
	"fmt"
	"image"
	"os"
)

const (
	ImageType AssetType = "image"
)

type ImageAsset struct {
	Image  *ebiten.Image
	OsPath string
}

func CreateImage(osPath string) Asset {
	return &ImageAsset{
		Image:  nil,
		OsPath: osPath,
	}
}

func (a *ImageAsset) Load() error {
	f, err := os.Open(osPath)
	if err != nil {
		return err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	a.Image = ebiten.NewImageFromImage(img)

	return nil
}

func (a *ImageAsset) Type() AssetType {
	return ImageType
}

func (a *ImageAsset) Info() string {
	return fmt.Sprintf("Texture: %dx%d pixels", a.Image.Width, a.Image.Height)
}
