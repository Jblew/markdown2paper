package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var defaultConfigPath string = "markdown2paper.config.yml"
var verbose bool = false

func main() {
  app := &cli.App{
    Name: "build",
		Usage: "builds ",
    Flags: []cli.Flag {
      &cli.BoolFlag{
        Name: "verbose",
        Usage: "Turn on verbosity",
        Destination: &verbose,
      },
    },
    Action: func(c *cli.Context) error {
      config, err := loadConfigFromFile(defaultConfigPath)
      if err != nil {
        return err
      }
      return Build(config)
		},
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
