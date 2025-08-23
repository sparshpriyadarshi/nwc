# nwc

**nwc** (newline, word count) is a modern Go implementation of the Unix `wc` command that prints newline, word, and byte counts for files and standard input. It provides enhanced functionality while maintaining compatibility with the traditional `wc` behavior.

## Features

- **Line counting** (`-l`): Count the number of newlines
- **Word counting** (`-w`): Count the number of words
- **Character counting** (`-m`): Count the number of characters
- **Byte counting** (`-c`): Count the number of bytes
- **Max line length** (`-L`): Find the length of the longest line
- **Totals** (`-t`): Display totals when processing multiple files
- **File list input** (`-f`): Read NULL-terminated filenames from a file
- **Standard input support**: Process data from stdin when no files are specified
- **Multiple file processing**: Handle multiple files with aggregate totals

## Requirements

- Go 1.23.4 or later
- Unix-like environment (Linux, macOS, WSL on Windows)

## Building

### From Source

1. **Clone the repository:**
   ```bash
   git clone https://github.com/sparshpriyadarshi/nwc.git
   cd nwc
   ```

2. **Build the binary:**
   ```bash
   go build -o nwc
   ```

3. **Verify the build:**
   ```bash
   ./nwc --help
   ```

### Using Go toolchain

```bash
go build -v
```

This will create an executable named `nwc` in the current directory.

## Installation

### Option 1: Manual Installation

After building from source:

```bash
# Copy to a directory in your PATH
sudo cp nwc /usr/local/bin/

# Or copy to your personal bin directory
mkdir -p ~/bin
cp nwc ~/bin/
export PATH="$HOME/bin:$PATH"  # Add to your shell profile
```

### Option 2: Go Install

```bash
go install github.com/sparshpriyadarshi/nwc@latest
```

Note: Ensure your `GOPATH/bin` is in your PATH for the installed binary to be accessible.

### Option 3: Development Installation

For development purposes:

```bash
go install .
```

This installs the binary to `$GOPATH/bin/nwc` (or `$HOME/go/bin/nwc` if GOPATH is not set).

## Usage

### Basic Syntax

```
nwc [flags] [path ...]
```

If no file is specified, standard input is processed.

### Command-Line Flags

| Flag | Long Form | Description |
|------|-----------|-------------|
| `-l` | `--lines` | Print the number of newlines |
| `-w` | `--words` | Print the number of words |
| `-m` | `--chars` | Print the number of characters |
| `-c` | `--bytes` | Print the number of bytes |
| `-L` | `--max-line-length` | Print the length of the longest line |
| `-t` | `--total` | Print totals for multiple files |
| `-f` | `--filenames-nul-sep-from` | Read NULL-terminated filenames from file |

### Examples

#### Basic Usage

**Count lines, words, and bytes (default behavior):**
```bash
nwc filename.txt
# Output: 42  310 2048 filename.txt
#         ^   ^    ^    ^
#      lines words bytes filename
```

**Process standard input:**
```bash
echo "Hello, World!" | nwc
# Output: 1  2 14
```

**Count specific metrics:**
```bash
# Count only lines
nwc -l file.txt

# Count only words  
nwc -w file.txt

# Count only bytes
nwc -c file.txt

# Count characters (may differ from bytes for UTF-8)
nwc -m file.txt

# Find longest line length
nwc -L file.txt
```

#### Multiple Files

**Process multiple files:**
```bash
nwc file1.txt file2.txt file3.txt
# Output shows counts for each file individually
```

**Process multiple files with totals:**
```bash
nwc -t file1.txt file2.txt file3.txt
# Shows individual counts plus a total line
```

#### Advanced Usage

**Combine multiple counting options:**
```bash
nwc -l -w -c file.txt          # Lines, words, and bytes
nwc -m -L file.txt             # Characters and max line length
nwc -l -w -m -c -L file.txt    # All available counts
```

**Read filenames from a NULL-terminated file:**
```bash
# Create a file with NULL-terminated filenames
find . -name "*.txt" -print0 > filelist.txt

# Process all files listed in filelist.txt
nwc -f filelist.txt

# Read from stdin for file list
find . -name "*.go" -print0 | nwc -f -
```

**Complex pipeline example:**
```bash
# Count lines in all Go files in current directory tree
find . -name "*.go" -print0 | nwc -l -t -f -
```

#### Practical Examples

**Count lines of code in a project:**
```bash
find . -name "*.go" -print0 | nwc -l -f -
```

**Compare file sizes:**
```bash
nwc -c *.txt
```

**Find files with long lines:**
```bash
nwc -L *.md
```

**Get comprehensive statistics:**
```bash
nwc -l -w -m -c -L -t *.txt
```

## Output Format

The output format matches the standard `wc` command:

```
[lines] [words] [characters] [bytes] [max-line-length] filename
```

Only the requested counts are displayed. When processing multiple files with `-t`, a "total" line is added at the end.

## Testing

Run the test suite:

```bash
# Run all tests
go test -v

# Run specific test
go test -run TestStdinUsage

# Run tests with coverage
go test -cover
```

## Development

### Project Structure

```
.
├── main.go           # Main application code
├── main_test.go      # Test suite
├── go.mod           # Go module definition
├── go.sum           # Go module checksums
├── testdata/        # Test data files
├── testscripts/     # Integration test scripts
└── README.md        # This file
```

### Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature-name`
3. Make your changes and add tests
4. Run tests: `go test -v`
5. Commit your changes: `git commit -am 'Add feature'`
6. Push to the branch: `git push origin feature-name`
7. Submit a pull request

## Version

Current version: **0.1.0**

## License

This project follows the same principles as standard Unix utilities. See the source code for specific licensing information.

## Compatibility

`nwc` aims to be compatible with the standard Unix `wc` command while providing additional features. The default behavior (counting lines, words, and bytes) matches `wc` exactly.

