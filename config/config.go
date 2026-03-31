package config

type Config struct {
	ScreenWidth     int32
	ScreenHeight    int32
	EditPanelHeight int32
	MaxInputChars   int32
}

func GetConfig() *Config {
	return &Config{
		ScreenWidth:     1000,
		ScreenHeight:    600,
		EditPanelHeight: 100,
		MaxInputChars:   5,
	}
}
