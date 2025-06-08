package assets

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

type AssetFactory func(osPath string) Asset

type FileBasedAssetManager struct {
	AssetFactories map[string]AssetFactory
	Assets         map[string]Asset
}

func CreateFileBasedAssetManager(assetFactory AssetFactory) *FileBasedAssetManager {
	am := &FileBasedAssetManager{
		AssetFactories: make(map[string]AssetFactory),
		Assets:         make(map[string]Asset),
	}
	am.RegisterAssetFactory("png", CreateImage)
	return am
}

func (am *FileBasedAssetManager) RegisterAssetFactory(ext string, factory AssetFactory) {
	ext = strings.ToLower(ext)
	am.AssetFactories[ext] = factory
}

func (am *FileBasedAssetManager) Get() (Asset, error) { return nil, nil }

func (am *FileBasedAssetManager) Scan(rootPath string) error {
	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error { // root undefined, should be rootPath
		if err != nil {
			return err
		}

		assetPath := strings.ReplaceAll(path[0:len(path)-len(filepath.Ext(path))], "\\", "/")
		osPath := "." + "/" + path

		if !d.IsDir() {
			ext := strings.ToLower(filepath.Ext(osPath))
			if factory, ok := am.AssetFactories[ext]; ok {
				am.Assets[assetPath] = factory(osPath)
			} else {
				fmt.Println("Not found asset factory", ext)
			}
		}

		return nil // missing in your code snippet
	})

	return err
}
