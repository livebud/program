# Program

[![Go Reference](https://pkg.go.dev/badge/github.com/livebud/program.svg)](https://pkg.go.dev/github.com/livebud/program)

Program provides a common entrypoint for CLI-based applications using [di](https://github.com/livebud/di).

## Install

```
go get github.com/livebud/program
```

## Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/livebud/cli"
	"github.com/livebud/di"
	"github.com/livebud/program"
)

// Provide an arbitrary CLI that implements `cli.Parser`.
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
```

## Contributors

- Matt Mueller ([@mattmueller](https://twitter.com/mattmueller))

## License

MIT
