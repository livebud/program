package main

import (
	"context"
	"fmt"
	"os"

	"github.com/livebud/cli"
	"github.com/livebud/di"
	"github.com/livebud/program"
)

func provideCLI(in di.Injector) (cli.Parser, error) {
	cli := cli.New("cli", "some cli")
	cmd := cli.Command("say", "say hello")
	cmd.Run(func(ctx context.Context) error {
		fmt.Println("hello")
		return nil
	})
	return cli, nil
}

func main() {
	in := di.New()
	di.Provide[cli.Parser](in, provideCLI)
	os.Exit(program.Run(in))
}
