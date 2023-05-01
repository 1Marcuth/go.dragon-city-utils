package orb_recall

import "fmt"

var OrbRecallConfig = map[string][]int{
	"per_levels": []int{40, 2, 2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2},
	"per_stars":  []int{120, 200, 320, 560, 800},
}

const (
	DragonMinLevel = 1
	DragonMaxLevel = 70
	DragonMinStars = 0
	DragonMaxStars = 5
)

func CalculateOrbRecallGain(dragonLevel int, dragonStars int) (int, error) {
	orbRecallGain := 0

	if dragonLevel < DragonMinLevel || dragonLevel > DragonMaxLevel {
		return 0, fmt.Errorf("'%d' is not a valid level for a dragon, choose a level between '%d' and '%d'", dragonLevel, DragonMinLevel, DragonMaxLevel)
	}

	if dragonStars < DragonMinStars || dragonStars > DragonMaxStars {
		return 0, fmt.Errorf("'%d' is not a valid number of stars for a dragon, choose a number of stars between '%d' and '%d'", dragonStars, DragonMinStars, DragonMaxStars)
	}

	for i := 0; i < dragonLevel && i < 30; i++ {
		orbRecallGain += OrbRecallConfig["per_levels"][i]
	}

	for i := 0; i < dragonStars; i++ {
		orbRecallGain += OrbRecallConfig["per_stars"][i]
	}

	return orbRecallGain, nil
}
