package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem2353(t *testing.T) {
	obj := Constructor([]string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"}, []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"}, []int{9, 12, 8, 15, 14, 7})
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "highestRated(korean)", obj.HighestRated("korean"), "kimchi")
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "highestRated(japanese)", obj.HighestRated("japanese"), "ramen")
	obj.ChangeRating("sushi", 16)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "highestRated(japanese)", obj.HighestRated("japanese"), "sushi")
	obj.ChangeRating("ramen", 16)
	fmt.Printf("[input]: %v        [output]: %v    [answer]: %v\n", "highestRated(japanese)", obj.HighestRated("japanese"), "ramen")
}
