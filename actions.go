package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/urfave/cli"
)

// actionMain is reading CSV file, and output JSON formatted string.
func actionMain(c *cli.Context) error {
	// check CSV file arguments.
	if c.NArg() == 0 {
		return cli.NewExitError("[ERROR] CSV file is not selected", 1)
	}

	// read CSV format config file.
	formatFile := c.String("format")
	if formatFile == "" {
		formatFile = "format.json"
	}
	b, err := ioutil.ReadFile(formatFile)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	// convert to csvFormat type from CSV format config JSON.
	r := strings.NewReader(string(b))
	d := json.NewDecoder(r)
	var confFmt csvFormat
	d.Decode(&confFmt)

	// read CSV file.
	var fp *os.File
	fp, err = os.OpenFile(c.Args().First(), os.O_RDONLY, 0)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	defer func() {
		fp.Close()
	}()

	reader := csv.NewReader(fp)
	// specify CSV field separator character (rune).
	// defalt is ','.
	if confFmt.Separator == "tab" {
		reader.Comma = '\t'
	}
	// check double quote.
	// default is true.
	reader.LazyQuotes = confFmt.LazyQuotes

	// output JSON map type.
	var data = []interface{}{}

	rowIndex := 0

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		// skip first line, if readAll is true.
		if confFmt.isSkip(rowIndex) {
			rowIndex++
			continue
		}
		rowIndex++

		fieldData := make(map[string]interface{})

		for i := 0; i < len(record); i++ {

			if confFmt.CsvFields[i].Name != "_" {
				fd := confFmt.CsvFields[i]
				switch fd.Type {
				case "int":
					fieldData[fd.Name], err = string2Int(record[i])
					if err != nil {
						return cli.NewExitError(err.Error(), 1)
					}
				case "float":
					fieldData[fd.Name], err = string2Float(record[i])
					if err != nil {
						return cli.NewExitError(err.Error(), 1)
					}
				default:
					fieldData[fd.Name] = strings.TrimSpace(record[i])
				}
			}
		}
		data = append(data, fieldData)
	}

	b, err = json.MarshalIndent(data, "", "  ")
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	os.Stdout.Write(b)

	return nil
}

// actionInit output format.json sample.
func actionInit(c *cli.Context) error {
	format := csvFormat{
		ReadAll:    false,
		Separator:  "comma",
		LazyQuotes: true,
		CsvFields: []csvField{
			csvField{
				Name: "id",
				Type: "int",
			},
			csvField{
				Name: "name",
				Type: "string",
			},
		},
	}

	b, _ := json.MarshalIndent(format, "", "  ")
	os.Stdout.Write(b)

	return nil
}
