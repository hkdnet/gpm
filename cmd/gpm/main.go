package main

import (
	"fmt"
	"os"

	"golang.org/x/net/context"

	"github.com/hkdnet/gpm"
)

func main() {
	os.Exit(run())
}

func run() int {
	if len(os.Args) < 1 {
		help()
		return 0
	}
	token := os.Getenv("GITHUB_TOKEN")
	client := gpm.NewClient(token)
	ctx, cancel := context.WithCancel(context.Background())
	// TODO: handle signals
	_ = cancel
	projects, err := client.ListRepoProjects(ctx, "hkdnet/gpm")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	for _, project := range projects {
		fmt.Printf("%s", project.Name)
	}

	return 0
}

func help() {
	// TODO: write help
	fmt.Println("help")
}
