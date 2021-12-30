package operation
import (
 "fmt"
 "math"
)
type Init struct {
 A, B float64
}
func (e *Init) GetSquareSum() float64 {
	fmt.Println(e)
 return (math.Pow(e.A,2) + math.Pow(e.B,2))
}