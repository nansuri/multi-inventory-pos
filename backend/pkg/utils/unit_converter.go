package utils

import (
	"errors"
	"strings"
)

// Unit factors relative to base unit (GR for mass, ML for volume)
var unitFactors = map[string]float64{
	"KG":    1000,
	"GR":    1,
	"LITRE": 1000,
	"ML":    1,
	"PIECE": 1,
}

func ConvertUnit(amount float64, fromUnit, toUnit string) (float64, error) {
	from := strings.ToUpper(fromUnit)
	to := strings.ToUpper(toUnit)

	if from == to {
		return amount, nil
	}

	// Check if units are compatible
	isMassFrom := from == "KG" || from == "GR"
	isMassTo := to == "KG" || to == "GR"
	isVolumeFrom := from == "LITRE" || from == "ML"
	isVolumeTo := to == "LITRE" || to == "ML"

	if (isMassFrom && !isMassTo) || (isVolumeFrom && !isVolumeTo) || (from == "PIECE" || to == "PIECE") {
		return 0, errors.New("incompatible units: " + fromUnit + " and " + toUnit)
	}

	fromFactor, ok1 := unitFactors[from]
	toFactor, ok2 := unitFactors[to]

	if !ok1 || !ok2 {
		return 0, errors.New("unsupported unit")
	}

	// Convert to base unit then to target unit
	baseAmount := amount * fromFactor
	return baseAmount / toFactor, nil
}
