# MediaMagic

A handy command-line tool to rename and organize media files.

## Usage

```shell
$ go build .

$ ./media_magic --label eurotrip
No source directory provided, using default: $HOME/Downloads
No output directory provided, using default: $HOME/Desktop/2023-07-21_renamed
Output directory created successfully.

Renaming files and moving to output directory: $HOME/Desktop/2023-07-21_renamed
- Renamed 'IMG_001.jpg' to '2023-07-21-eurotrip-bYsrI2RGIx.jpg'...
- Renamed 'IMG_002.jpg' to '2023-07-21-eurotrip-okLbi8NDf8.jpg'...

Number of files found: 2
Number of files renamed: 2
Number of files skipped: 0
```

### Help menu

```shell
$ ./media_magic --help

Options:
  -l string
        Shorthand for label (default "pic")
  -label string
        Label for renamed files
  -o string
        Shorthand for output directory
  -output string
        Path to output directory for renamed media files
  -s string
        Shorthand for source directory
  -source string
        Path to source directory with unorganized media files
```