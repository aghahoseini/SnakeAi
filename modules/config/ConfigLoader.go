package config

import (
	"fmt"
	"image/color"
	"github.com/BurntSushi/toml"
)

type ConfigLoader struct {
	filepath string
	settings map[string]interface{}
}

func NewConfigLoader(filepath string) *ConfigLoader {
	return &ConfigLoader{
		filepath: filepath,
		settings: make(map[string]interface{}),
	}
}

func (cl *ConfigLoader) LoadSettings() (map[string]interface{}, error) {
	if _, err := toml.DecodeFile(cl.filepath, &cl.settings); err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}

	for key, val := range cl.settings {
		switch v := val.(type) {
		case []interface{}:
			if len(v) == 4 { 
				if isColorSlice(v) {
					cl.settings[key] = interfaceToColor(v)
				}
			}
		case int64, float64:
			cl.settings[key] = interfaceToInt(v)
		}
	}

	return cl.settings, nil
}


func isColorSlice(slice []interface{}) bool {
	for _, item := range slice {
		switch item.(type) {
		case int64, float64:
			
		default:
			return false
		}
	}
	return true
}

func interfaceToColor(colorSlice []interface{}) color.RGBA {
	r, _ := toUint8(colorSlice[0])
	g, _ := toUint8(colorSlice[1])
	b, _ := toUint8(colorSlice[2])
	a, _ := toUint8(colorSlice[3])

	return color.RGBA{R: r, G: g, B: b, A: a}
}

func interfaceToInt(val interface{}) int {
	switch v := val.(type) {
	case int64:
		return int(v)
	case float64:
		return int(v)
	default:
		return 0
	}
}

func toUint8(val interface{}) (uint8, error) {
	switch v := val.(type) {
	case int64:
		if v < 0 || v > 255 {
			return 0, fmt.Errorf("value out of range")
		}
		return uint8(v), nil
	case float64:
		if v < 0 || v > 255 {
			return 0, fmt.Errorf("value out of range")
		}
		return uint8(v), nil
	default:
		return 0, fmt.Errorf("invalid type")
	}
}