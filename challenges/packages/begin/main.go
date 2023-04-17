// challenges/packages/begin/proverbs.go
package main

// import the proverbs package
import (
	"fmt"

	proverb "github.com/jboursiquot/go-proverbs"
)

// getRandomProverb returns a random proverb from the proverbs package
func getRandomProverb() string {
	return proverb.Random().Saying
}

func main() {
	// print the result of calling your getRandomProverb function
	fmt.Println(getRandomProverb())
}
