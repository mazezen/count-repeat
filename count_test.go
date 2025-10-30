package count_repeat

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	fmt.Println(Count([]interface{}{"10", "20", "600", "20", "30"}))
	fmt.Println(Count([]interface{}{"20", "20", "600", "20", "30"}))
	fmt.Println(Count([]interface{}{"600", "20", "600", "20", "30"}))
	fmt.Println(Count([]interface{}{"20", "20", "600", "20", "30"}))
	fmt.Println(Count([]interface{}{"80", "20", "600", "20", "30"}))
	fmt.Println(Count([]interface{}{"100", "20", "600", "20", "30"}))
}
