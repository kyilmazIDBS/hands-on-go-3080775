// challenges/packages/begin/proverbs.go
package main
// import the proverbs package
import "fmt"
import "github.com/jboursiquot/go-proverbs"

// getRandomProverb returns a random proverb from the proverbs package

func getRandomProverb(){
	fmt.Println(proverbs.Random().Saying)
}

func main() {
	// print the result of calling your getRandomProverb function
	getRandomProverb()
}
