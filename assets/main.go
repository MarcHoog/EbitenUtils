package assets

type AssetType string

type Asset interface {
	Load() error
	Type() AssetType
	Info() string
}

type AssetManager interface {
	Get() (*Asset, error)
	GetMulti() ([]*Asset, error)
	Pop() (*Asset, error)
	Delete() error
	Scan(rootPath string) error
}
