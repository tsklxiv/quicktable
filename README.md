<!-- markdownlint-configure-file {
  "MD013": {
    "code_blocks": false,
    "tables": false
  },
  "MD033": false,
  "MD041": false
} -->

<div align="center">

<img src="quicktable_logo.png" alt="Quicktable logo; A 2x2 black grid with a yellow background">

# quicktable

Generate beautiful CLI tables on the fly, from CSV files.

[Getting started](#getting-started) •
[Installation](#installation) •
[CLI](#cli)

</div>

## Getting started

[![asciicast](https://asciinema.org/a/0SGXgIst0w6SrS2CE6t8MK98r.svg)](https://asciinema.org/a/0SGXgIst0w6SrS2CE6t8MK98r)

```sh
quicktable -f data.csv                           # Generate table from a CSV file
quicktable -h                                    # Print help section
quicktable -f Sample100.csv -c "100 CSV Records" # Generate table with custom captions
quicktable -f Sample100.csv -o something.txt     # Output table to file something.txt
```

## Installation

Install directly from `go`:

```sh
$ go install github.com/HoangTuan110/quicktable
```

## CLI

```
Generate beautiful CLI tables on the fly from CSV files.

Options:

  -h, --help      display help information
  -f, --file      file to parse
  -c, --caption   caption for table
  -o, --output    Output to file
```

</div>
