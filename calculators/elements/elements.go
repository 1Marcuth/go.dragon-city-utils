package elements

import (
	"fmt"
	"github.com/fatih/structs"
)

type ElementConfig struct {
	strongs    []string
	weaknesses []string
}

type ElementsConfig struct {
	terra    ElementConfig
	flame    ElementConfig
	sea      ElementConfig
	nature   ElementConfig
	electric ElementConfig
	ice      ElementConfig
	metal    ElementConfig
	dark     ElementConfig
	light    ElementConfig
	war      ElementConfig
	pure     ElementConfig
	legend   ElementConfig
	primal   ElementConfig
	wind     ElementConfig
	time     ElementConfig
	happy    ElementConfig
	chaos    ElementConfig
	magic    ElementConfig
	soul     ElementConfig
	beauty   ElementConfig
	dream    ElementConfig
}

var elementsConfig = ElementsConfig{
	terra: ElementConfig{
		strongs:    []string{"electric", "flame"},
		weaknesses: []string{"metal", "war"},
	},
	flame: ElementConfig{
		strongs:    []string{"nature", "ice"},
		weaknesses: []string{"sea", "terra"},
	},
	sea: ElementConfig{
		strongs:    []string{"flame", "war"},
		weaknesses: []string{"nature", "electric"},
	},
	nature: ElementConfig{
		strongs:    []string{"sea", "light"},
		weaknesses: []string{"flame", "ice"},
	},
	electric: ElementConfig{
		strongs:    []string{"sea", "metal"},
		weaknesses: []string{"terra", "light"},
	},
	ice: ElementConfig{
		strongs:    []string{"nature", "war"},
		weaknesses: []string{"flame", "metal"},
	},
	metal: ElementConfig{
		strongs:    []string{"terra", "ice"},
		weaknesses: []string{"electric", "dark"},
	},
	dark: ElementConfig{
		strongs:    []string{"metal", "light"},
		weaknesses: []string{"war"},
	},
	light: ElementConfig{
		strongs:    []string{"electric", "dark"},
		weaknesses: []string{"nature"},
	},
	war: ElementConfig{
		strongs:    []string{"terra", "dark"},
		weaknesses: []string{"sea", "ice"},
	},
	pure: ElementConfig{
		strongs:    []string{"wind"},
		weaknesses: []string{"primal"},
	},
	legend: ElementConfig{
		strongs:    []string{"primal"},
		weaknesses: []string{"pure"},
	},
	primal: ElementConfig{
		strongs:    []string{"pure"},
		weaknesses: []string{"time"},
	},
	wind: ElementConfig{
		strongs:    []string{"time"},
		weaknesses: []string{"legend"},
	},
	time: ElementConfig{
		strongs:    []string{"legend"},
		weaknesses: []string{"wind"},
	},
	happy: ElementConfig{
		strongs:    []string{"chaos", "magic"},
		weaknesses: []string{},
	},
	chaos: ElementConfig{
		strongs:    []string{"magic", "soul"},
		weaknesses: []string{},
	},
	magic: ElementConfig{
		strongs:    []string{"soul", "beauty"},
		weaknesses: []string{},
	},
	soul: ElementConfig{
		strongs:    []string{"dream", "beauty"},
		weaknesses: []string{},
	},
	beauty: ElementConfig{
		strongs:    []string{"dream", "happy"},
		weaknesses: []string{},
	},
	dream: ElementConfig{
		strongs: []string{"happy", "chaos"},
		weaknesses: []string{},
	},
}

func contains(arr []string, val string) bool {
    for _, v := range arr {
        if v == val {
            return true
        }
    }
    return false
}

func getElementConfig(elementName string) (ElementConfig, error) {
    switch elementName {
		case "terra":
			return elementsConfig.terra, nil

		case "flame":
			return elementsConfig.flame, nil

		case "sea":
			return elementsConfig.sea, nil

		case "nature":
			return elementsConfig.nature, nil

		case "electric":
			return elementsConfig.electric, nil

		case "ice":
			return elementsConfig.ice, nil

		case "metal":
			return elementsConfig.metal, nil

		case "dark":
			return elementsConfig.dark, nil

		case "light":
			return elementsConfig.light, nil

		case "war":
			return elementsConfig.war, nil

		case "pure":
			return elementsConfig.pure, nil

		case "legend":
			return elementsConfig.legend, nil

		case "primal":
			return elementsConfig.primal, nil

		case "wind":
			return elementsConfig.wind, nil

		case "time":
			return elementsConfig.time, nil

		case "happy":
			return elementsConfig.happy, nil

		case "chaos":
			return elementsConfig.chaos, nil

		case "magic":
			return elementsConfig.magic, nil

		case "soul":
			return elementsConfig.soul, nil

		case "beauty":
			return elementsConfig.beauty, nil

		case "dream":
			return elementsConfig.dream, nil

		default:
			return ElementConfig{}, fmt.Errorf("Invalid element name: %s", elementName)
    }
}

func CalculateStrongs(elements []string) ([]string, error) {
    var strongs []string

    for _, element := range elements {
        config, err := getElementConfig(element)
        if err != nil {
            return nil, err
        }

        for _, e := range config.strongs {
            if !contains(strongs, e) {
                strongs = append(strongs, e)
            }
        }
    }

    return strongs, nil
}

func CalculateWeaknesses(firstElement string) []string {
    weaknesses := []string{}
	elementNames := structs.Names(elementsConfig)
	
    for i := 0; i < len(elementNames); i++ {
		elementName := elementNames[i]
		elementConfig, _ := getElementConfig(elementName)

        if contains(elementConfig.strongs, firstElement) {
            weaknesses = append(weaknesses, elementName)
        }
    }

    return weaknesses
}