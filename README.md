# Go Semantic Color

Compute the similarity between two colors in RGB.

Example:

```go
import color "github.com/ndajr/go-semantic-color"

red := color.RGB{255, 0, 0}
lightRed := color.RGB{255, 100, 100}
fmt.Printf("%.2f\n", color.CosineSimilarity(red, lightRed))
// 0.87 or 87%
```
