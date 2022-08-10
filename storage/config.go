package storage

type FileConfig struct {
	CommonInfoFile   string
	MenuFile         string
	ProductsListFile string
}

func GetFileConfig() FileConfig {
	return FileConfig{
		CommonInfoFile:   "storage/json/general_info.json",
		MenuFile:         "storage/json/menu.json",
		ProductsListFile: "storage/json/products.json",
	}
}
