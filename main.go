package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mkideal/cli"
	"github.com/olekukonko/tablewriter"
)

type argT struct {
  cli.Helper
  FileName string `cli:"f,file" usage:"file to parse" dft:""`
  Caption  string `cli:"c,caption" usage:"caption for table" dft:""`
  Output   string `cli:"o,output" usage:"Output to file" dft:""`
}

func (argv *argT) AutoHelp() bool {
  // I found a way to add custom one-line into this thing! :D
  if argv.Help {
    fmt.Printf("Generate beautiful CLI tables on the fly from CSV files.\n\n")
  }
	return argv.Help
}

func main() {
  os.Exit(cli.Run(new(argT), func(ctx *cli.Context) error {
    argv := ctx.Argv().(*argT)
    name := argv.FileName
    capt := argv.Caption
    output := argv.Output

    if name == "" {
      log.Fatal("Please give me a file to parse")
    }

    // Generate table from the CSV data
    // I thought I will need to read the CSV, then feed it
    // to the table and anything, but it turns out tablewriter
    // has automatic support for generating table from CSV files
    // thanks maintainer! :D
    tableString := &strings.Builder{}
    table, err := tablewriter.NewCSV(tableString, name, true)
    if err != nil {
      log.Fatal(err)
    }
    table.SetAutoMergeCells(true)
    table.SetRowLine(true)
    table.SetCaption(true, capt)
    table.Render()

    // Check if we need to output to file or we can just output
    // back to the terminal
    if output == "" {
      fmt.Println(tableString.String())
    } else {
      f, err := os.Create(output)
      if err != nil {
        log.Fatal(err)
      }
      defer f.Close()

      _, err = f.WriteString(tableString.String())
      if err != nil {
        log.Fatal(err)
      }
    }

    return nil
  }))
}
