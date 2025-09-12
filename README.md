# PassDotGo
A small command-line password generator written in Go.

## Usage
The program can be used to generate passwords up to 80 characters long, including letters, numbers, and symbols like `@$&!#%&*`.

```bash
passdotgo <length> [flags]
```

### Flags
| Flag            | Description                        |
| --------------- | ---------------------------------- |
| -l, --letters   | Letters only password.             |
| -n, --numbers   | Letters and numbers password.      |
| -s, --symbols   | All characters password (Default). |
| -e, --eye-candy | Outputs a colorful password.       |
| -c, --count     | Generates multiple passwords.      |
| -h, --help      | Helps.                             |

> [!IMPORTANT]
> The `--eye-candy` flag will not function properly if your terminal doesn't support ANSI colors or if you're exporting the output using the `>` operator.

## Building
```bash
go build -o ./bin/passdotgo
```
The binary will appear inside the `/bin` subfolder.

## Examples
```bash
# A single 42-char password
./bin/passdotgo 42

# Three 42-char passwords
./bin/passdotgo 42 --count 3

# Two alphanumeric passwords, with a colorful output
./bin/passdotgo 42 -n -c 2 --eye-candy
```