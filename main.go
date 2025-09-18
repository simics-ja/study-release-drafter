package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "0.1.0"

func main() {
	var (
		showVersion = flag.Bool("version", false, "Show version")
		name        = flag.String("name", "World", "Name to greet")
		help        = flag.Bool("help", false, "Show help")
	)

	flag.Parse()

	if *help {
		fmt.Println("Simple Go CLI - A demonstration CLI for Release Drafter")
		fmt.Println("\nUsage:")
		fmt.Println("  simple-cli [options]")
		fmt.Println("\nOptions:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *showVersion {
		fmt.Printf("simple-cli version %s\n", version)
		os.Exit(0)
	}

	fmt.Printf("Hello, %s\n", *name)
	fmt.Printf("Bye, %s\n", *name)
}
