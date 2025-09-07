# PassDotGo
A small command-line password generator written in Go.

## Usage
The program can be used to generate passwords up to 80 characters long, including letters, numbers, and symbols like `@$&!#%&*`.

```bash
passdotgo [size] [...options]
```

### Options
| Option          | Description                        |
| --------------- | ---------------------------------- |
| -l, --letters   | Letters only password.             |
| -n, --numbers   | Letters and numbers password.      |
| -s, --symbols   | All characters password (Default). |
| -u, --uncolored | Outputs an uncolored password.     |

> [!IMPORTANT]
> If your terminal doesn't render colored characters properly, or if you're exporting the output using `> password.txt`, use the `--uncolored` option.

## Building
```bash
# Compiling
go build -o ./bin/passdotgo ./src/

# Running (42-char all chars password)
./bin/passdotgo 42 --symbols
```