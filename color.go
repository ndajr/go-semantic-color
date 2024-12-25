package semantic_color

import "math"

type RGB [3]float64

func CosineSimilarity(color1, color2 RGB) float64 {
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
