# id

The `id` package provides a robust and secure way to generate random strings of a specified length using a provided alphabet. It leverages cryptographic randomness to ensure high-quality randomness.

## Usage

### Generating Random Strings

You can generate random strings of a specific length using the `Generate` function. This function requires an alphabet and the desired length of the string.

- `Generate(alphabet string, length int)`: Generates a random string of the specified length using the provided alphabet.

### Creating a Generator

For convenience, you can create a reusable string generator function with the `Generator` function. This is useful when you need to generate multiple strings with the same alphabet and length.

- `Generator(alphabet string, length int)`: Returns a function that generates random strings using the provided alphabet and length.

## Example

Below is an example of how to use the `id` package to generate random strings.

```go
package main

import (
	"fmt"
	"github.com/lance-free/id"
)

func main() {
	// Define the alphabet and length
	alphabet := "abcdefghijklmnopqrstuvwxyz0123456789"
	length := 10

	// Generate a random string
	randomString, err := id.Generate(alphabet, length)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Random String:", randomString)

	// Create a reusable generator function
	generator := id.Generator(alphabet, length)

	// Generate multiple random strings
	for i := 0; i < 5; i++ {
		randomString, err := generator()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Random String:", randomString)
	}
}
```

## License

This package is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.