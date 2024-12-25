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

func ExampleNearestColor() {
	rgb, _ := NearestColor(RGB{253, 0, 0})
	fmt.Printf("%+v", rgb)
	// Output: {Name:Lemon Glacier RGB:[253 255 0] Distance:1}
}

func ExampleTextToRGB() {
	rgb, _ := TextToRGB("Red")
	fmt.Printf("%v", rgb)
	// Output: [255 0 0]
}
