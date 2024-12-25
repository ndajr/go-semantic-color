package semantic_color

import (
	"errors"
	"math"
)

var ErrColorNotFound = errors.New("color not found")

type (
	RGB        [3]uint8
	ColorMatch struct {
		Name     string
		RGB      RGB
		Distance float64
	}
)

func CosineSimilarity(a, b RGB) float64 {
	color1 := [3]float64{float64(a[0]), float64(a[1]), float64(a[2])}
	color2 := [3]float64{float64(b[0]), float64(b[1]), float64(b[2])}

	// Calculate dot product
	dotProduct := color1[0]*color2[0] + color1[1]*color2[1] + color1[2]*color2[2]

	// Calculate norms (magnitudes) of each vector
	norm1 := math.Sqrt(color1[0]*color1[0] + color1[1]*color1[1] + color1[2]*color1[2])
	norm2 := math.Sqrt(color2[0]*color2[0] + color2[1]*color2[1] + color2[2]*color2[2])

	// Avoid division by zero
	if norm1 == 0 || norm2 == 0 {
		return 0
	}

	// Calculate cosine similarity
	return dotProduct / (norm1 * norm2)
}

func NearestColor(given RGB) (ColorMatch, error) {
	minDistanceSq := math.Inf(1)
	var bestMatch string

	for name := range colorMap {
		distanceSq := math.Pow(float64(given[0]-colorMap[name][0]), 2) +
			math.Pow(float64(given[1]-colorMap[name][1]), 2) +
			math.Pow(float64(given[2]-colorMap[name][2]), 2)

		if distanceSq < minDistanceSq {
			minDistanceSq = distanceSq
			bestMatch = name
		}
	}

	rgb, found := colorMap[bestMatch]
	if !found {
		return ColorMatch{}, ErrColorNotFound
	}
	if bestMatch != "" {
		return ColorMatch{
			Name:     bestMatch,
			RGB:      rgb,
			Distance: math.Sqrt(minDistanceSq),
		}, nil
	}

	return ColorMatch{
		RGB:      rgb,
		Distance: math.Sqrt(minDistanceSq),
	}, nil
}

func TextToRGB(colorName string) (RGB, error) {
	rgb, found := colorMap[colorName]
	if !found {
		return RGB{0, 0, 0}, ErrColorNotFound
	}
	return rgb, nil
}
