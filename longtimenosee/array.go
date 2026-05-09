package main
import (
	"fmt"
)

// arrays and slices

func main() {
	usernames := []string {"shubhdevs", "abhinav", "utk", "tarun"};
	for i, username := range usernames {
		fmt.Printf("%v -> index -> (%v) \n", username, i);
	}
}