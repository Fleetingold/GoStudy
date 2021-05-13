package main

import "fmt"

func main() {
	printAll("GoLand is a cross-platform integrated development environment (IDE) for Go developers. GoLand includes such features as context-dependent code completion and refactoring, debugging, profiling, type and declaration navigation, and error analysis. In addition to tools for core Go development, supports JavaScript, TypeScript, Node.js, SQL, Databases, Docker, Kubernetes, and Terraform. You can always extend current functionality by installing additional plugins from the plugin repository.", 2147483647, false, 1<<64-1, 4, '♡', 3.14, 1+2i)
}

func printAll(s string,
	i int,
	b bool,
	u uint,
	by byte,
	r rune,
	f float64,
	c complex64) {
	fmt.Println(s, i, b, u, by, string(r), f, c)
}
