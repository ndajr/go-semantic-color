package semantic_color

import "fmt"

func ExampleCosineSimilarity() {
	color1 := RGB{255, 0, 0}     // Red
	color2 := RGB{0, 255, 0}     // Green
	color3 := RGB{255, 100, 100} // Light Red

	fmt.Printf("%.2f\n", CosineSimilarity(color1, color2)) // Red vs Green
	fmt.Printf("%.2f\n", CosineSimilarity(color1, color3)) // Red vs Light Red
	// Output:
	// 0.00
	// 0.87
}
