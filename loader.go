package main

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
)

type Loader struct {
	nextID   AssetID
	Registry map[string]AssetInfo

	image map[AssetID]*ebiten.Image
	raw   map[AssetID][]byte
}

func NewLoader() *Loader {
	return &Loader{
		nextID:   AssetID(0),
		Registry: make(map[string]AssetInfo),
		image:    make(map[AssetID]*ebiten.Image),
		raw:      make(map[AssetID][]byte),
	}
}

// TODO: Improve how this works with asset paths and just system paths in general

func (l *Loader) SystemPathToAssetPath(rootPath string, filePath string) string {

	rootPath = strings.ToLower(rootPath)
	filePath = strings.ToLower(filePath)

	rootPath = filepath.Clean(rootPath)
	path := strings.TrimPrefix(filePath, rootPath+"/")
	assetPath := strings.ReplaceAll(path[0:len(path)-len(filepath.Ext(path))], "\\", "/")
	return assetPath

}

func (l *Loader) Register(rootPath, path string, typ AssetType) AssetID {
	assetPath := l.SystemPathToAssetPath(rootPath, path)
	osPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}
	id := l.nextID
	l.nextID++
	l.Registry[assetPath] = AssetInfo{ID: id, Type: typ, OsPath: osPath}
	return id
}

func (l *Loader) Scan(rootPath string) (err error) {
	err = filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error { // root undefined, should be rootPath
		if err != nil {
			return err
		}

		if !d.IsDir() {
			ext := strings.ToLower(filepath.Ext(path))

			if ext == ".png" || ext == ".jpg" {
				l.Register(rootPath, path, AssetImage)
			} else if ext == ".wav" || ext == ".mp3" {
				l.Register(rootPath, path, AssetAudio)
			}

		}

		return nil
	})

	return err
}

func (l *Loader) loadImage(assetPath string, cache bool) (*ebiten.Image, error) {
	assetInfo, ok := l.Registry[assetPath]

	if !ok {
		return nil, errors.New("asset not found")
	}

	if assetInfo.Type != AssetImage {
		return nil, errors.New("asset is not image")
	}

	image, ok := l.image[assetInfo.ID]
	if !cache {
		delete(l.image, assetInfo.ID)
	}

	if ok {
		return image, nil
	}
	image, err := CreateImage(assetInfo.OsPath)
	if err != nil {
		return nil, err
	}

	if cache {
		l.image[assetInfo.ID] = image
	}
	return image, nil

}

func (l *Loader) GetImage(assetPath string) (*ebiten.Image, error) {
	return l.loadImage(assetPath, true)
}

func (l *Loader) PopImage(assetPath string) (*ebiten.Image, error) {
	return l.loadImage(assetPath, false)
}
