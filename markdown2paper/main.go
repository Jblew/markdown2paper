package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
  app := &cli.App{
    Name: "build",
		Usage: "builds ",
		Flags: []cli.Flag {
      &cli.StringFlag{
        Name: "bib",
        Value: "",
        Usage: "bibliography file",
			},
			&cli.StringFlag{
        Name: "outline",
        Value: "",
        Usage: "markdown file with the outline section",
			},
			&cli.StringFlag{
        Name: "out",
        Value: "",
        Usage: "output file",
      },
    },
    Action: func(c *cli.Context) error {
			params := BuildParams{
				BibFile: c.String("bib"),
				OutlineFile: c.String("outline"),
				OutFile: c.String("out"),
			}
      return Build(params)
		},
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}

// BuildParams — parameters for building
type BuildParams struct {
	BibFile string
	OutlineFile string
	OutFile string
}

// Build actually builds the paper
func Build(params BuildParams) error {
  log.Printf("%+v", params)
  outlineContents, err := ReadFileToText(params.OutlineFile)
  if err != nil {
    return err
  }
  sections, err := ParseTextToMarkdown("", outlineContents, 0)
  if err != nil {
    return err
  }

  log.Printf("%+v", sections)

  out, err := json.MarshalIndent(sections, "", "    ")
  if err != nil {
    return err
  }
  log.Printf("%s", string(out))

	return nil
}
