# nwc

**nwc** (newline, word count) is a Go implementation of the Unix `wc` command that counts lines, words, and bytes in files or standard input.

## Features

- **Line counting** (`-l`, `-lines`): Count newlines
- **Word counting** (`-w`, `-words`): Count words
- **Character counting** (`-m`, `-chars`): Count characters
- **Byte counting** (`-c`, `-bytes`): Count bytes  
- **Max line length** (`-L`, `-max-line-length`): Find longest line
- **Totals** (`-t`, `-total`): Show totals for multiple files
- **File list input** (`-f`, `-filenames-nul-sep-from`): Read null-terminated filenames from file
- Standard input support when no files specified
- Multiple file processing with optional totals

## Requirements

- Go 1.23.4 or later

## Building

```bash
git clone https://github.com/sparshpriyadarshi/nwc.git
cd nwc
go build -o nwc
```

## Installation

```bash
# Install from source
go install .

# Or copy binary to PATH
sudo cp nwc /usr/local/bin/
```

## Usage

```
nwc [flags] [path ...]
```

### Flags

| Short | Long | Description |
|-------|------|-------------|
| `-l` | `-lines` | Print number of newlines |
| `-w` | `-words` | Print number of words |
| `-m` | `-chars` | Print number of characters |
| `-c` | `-bytes` | Print number of bytes |
| `-L` | `-max-line-length` | Print longest line length |
| `-t` | `-total` | Print totals for multiple files |
| `-f` | `-filenames-nul-sep-from` | Read filenames from file |

### Examples

Basic usage:
```bash
# Default behavior (lines, words, bytes)
nwc file.txt

# Count specific metrics
nwc -l file.txt                    # Lines only
nwc -w file.txt                    # Words only
nwc -c file.txt                    # Bytes only
nwc -m file.txt                    # Characters only
nwc -L file.txt                    # Longest line only

# Multiple metrics
nwc -l -w -c file.txt

# Multiple files with totals
nwc -t file1.txt file2.txt file3.txt

# From standard input
echo "hello world" | nwc

# From file list (null-terminated)
printf "file1.txt\0file2.txt\0" | nwc -f -
```

## Testing

```bash
go test -v
```

## License

Similar to standard Unix utilities.

