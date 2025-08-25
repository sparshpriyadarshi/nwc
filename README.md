# nwc

A Go implementation of the coreutils' `wc` command that prints newline, word, and byte counts for files.

## Usage


```bash
nwc --help
```

Output:
```
Print newline, word and byte counts for files and more. 
If no file is specified, standard input is processed.

Usage:
	nwc [flags] [path ...]

Flags:

  -l, -lines
    	print the number of newlines
  -w, words
    	print the number of words
  -m, -chars
    	print the number of characters
  -c, -bytes
    	print the number of bytes
  -L, -max-line-length
    	print the widest line length
  -f, -filenames-nul-sep-from <file>
    	read NULL terminated input filenames from <file> (or stdin if <file> is "-")
  -t, -total
    	print the totals

Version:
	0.1.0
```

## Clone, Test, Build, Install

```bash
git clone https://github.com/sparshpriyadarshi/nwc.git
cd nwc
```
```bash
# Run tests
go test -v -cover
```
```bash
# Build here
go build -v
```
```bash
# Install locally
go install
```




## Usage examples

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
# Find longest lines
nwc -L *.md
```

**Combine multiple counting options:**
```bash
nwc -l -w -c file.txt          # Lines, words, and bytes
nwc -m -L file.txt             # Characters and widest line
nwc -l -w -m -c -L -t *.*      # Get comprehensive statistics
```

**Compare file sizes:**
```bash
nwc -c *.txt
```

**Complex pipeline example:**
```bash
# Count lines code in current directory tree
find . -name "*.go" -print0 | nwc -l -t -f -
```


