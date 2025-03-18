package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

func printDefaults() {
	fmt.Println("Usage: yesno [options...] <prompt>")
	fmt.Println("Example: yesno -y '^ok' 'OK?'")
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	pattern := flag.String("y", "^[Yy]([Ee][Ss])?", "A regular expression to detect 'yes/no'.")
	h := flag.Bool("h", false, "Show this message.")
	help := flag.Bool("help", false, "Show this message.")
	flag.Parse()
	args := flag.Args()
	if *h || *help {
		printDefaults()
	}
	if len(args) != 1 {
		fmt.Println("Set a prompt string.")
		fmt.Println("")
		printDefaults()
	}
	re := regexp.MustCompile(*pattern)
	fmt.Print(args[0] + " ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() && re.MatchString(scanner.Text()) {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
	fmt.Println("Failed to read answer")
	os.Exit(2)
}
