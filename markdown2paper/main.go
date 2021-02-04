package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var defaultConfigPath string = "markdown2paper.config.yml"

func main() {
  app := &cli.App{
    Name: "build",
		Usage: "builds ",
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
