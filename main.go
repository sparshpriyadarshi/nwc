package main

import (
	"flag"
	"fmt"
	"io"
	"log" //slog next time ?
	"os"
	"strconv"
	"strings"
)

const VERSION = "0.1.0"
const PROGRAM_NAME = "nwc"
const USAGE_TEMPLATE = `Print newline, word and byte counts for files and more. 
If no file is specified, standard input is processed.

Usage:
	%s [flags] [path ...]

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
	%s

`

var (
	LPL = log.Println
	LPF = log.Printf
	LP  = log.Print
	LFF = log.Fatalf
)

func isNl(c rune) bool {
	return c == '\n'
}

func isNUL(c rune) bool {
	return c == '\000'
}

const WCNumOfOptions = 5

// note: output order is always newline, word, character, byte, maximum line length; filename
type WCOption int

const (
	NewlineCount WCOption = iota
	WordCount
	CharacterCount
	ByteCount
	MaxLinelength
	//these feel like meta ops, not including it in methods...
	Totals       //behaves differently from wc
	FilesFromNul //behaves differently from wc
)

type WC struct {
	buff           []byte
	enabledOptions map[WCOption]bool
	resultMap      map[WCOption]int // cache result values
	totalsMap      map[WCOption]int // accumulate totals as needed
	filename       string
	maxCount       int
	countWidth     int
}

func newWC() *WC {
	wc := WC{}
	//defaults
	wc.enabledOptions = make(map[WCOption]bool)
	wc.enabledOptions[NewlineCount] = true
	wc.enabledOptions[WordCount] = true
	wc.enabledOptions[CharacterCount] = false
	wc.enabledOptions[ByteCount] = true
	wc.enabledOptions[MaxLinelength] = false
	wc.enabledOptions[Totals] = false

	wc.resultMap = make(map[WCOption]int)
	wc.resultMap[NewlineCount] = -1
	wc.resultMap[WordCount] = -1
	wc.resultMap[CharacterCount] = -1
	wc.resultMap[ByteCount] = -1
	wc.resultMap[MaxLinelength] = -1
	wc.maxCount = 0

	wc.totalsMap = make(map[WCOption]int)
	wc.totalsMap[NewlineCount] = 0
	wc.totalsMap[WordCount] = 0
	wc.totalsMap[CharacterCount] = 0
	wc.totalsMap[ByteCount] = 0
	wc.totalsMap[MaxLinelength] = 0

	return &wc
}

func (wc *WC) setOptions(newlineFlag, wordFlag, charFlag, byteFlag, maxLinelenFlag bool) {
	wc.enabledOptions[NewlineCount] = newlineFlag
	wc.enabledOptions[WordCount] = wordFlag
	wc.enabledOptions[CharacterCount] = charFlag
	wc.enabledOptions[ByteCount] = byteFlag
	wc.enabledOptions[MaxLinelength] = maxLinelenFlag
}

func (wc WC) String() string {
	return fmt.Sprintf("filename=%v, enabledOptions=%v\n", wc.filename, wc.enabledOptions)
}

func (wc *WC) loadFromStdin() (int, error) {
	//LPL("DEBUG: wc::loadFromStdin()")

	var err error
	wc.buff, err = io.ReadAll(os.Stdin)
	if err != nil {
		LFF("FATAL: Error reading from stdin: %v\n", err)
	}

	wc.filename = "" //no name for stdin

	return len(wc.buff), nil
}

func (wc *WC) loadFromFile(filepath string) (int, error) {
	//LPL("DEBUG: wc::loadFromFile() File = ", filepath)

	var err error
	wc.buff, err = os.ReadFile(filepath)
	if err != nil {
		LFF("FATAL: Error reading from file: %v\n", err)
	}

	wc.filename = filepath

	return len(wc.buff), nil
}

func (wc *WC) loadString(s string) {
	//LPL("DEBUG: wc::loadString() = ", s)

	wc.buff = []byte(s)
	wc.filename = ""
}

func (wc *WC) nlCount() (count int) {

	//LPL("DEBUG: wc::nlCount() ... counting newlines")
	for _, v := range wc.buff {
		if isNl(rune(v)) {
			count++
		}
	}
	wc.resultMap[NewlineCount] = count
	wc.totalsMap[NewlineCount] += wc.resultMap[NewlineCount]
	wc.maxCount = max(wc.maxCount, wc.resultMap[NewlineCount])
	return wc.resultMap[NewlineCount]
}

func (wc *WC) wCount() int {
	//LPL("DEBUG: wc::wCount() ... counting words")
	filestring := string(wc.buff)
	count := len(strings.Fields(filestring))
	wc.resultMap[WordCount] = count
	wc.totalsMap[WordCount] += wc.resultMap[WordCount]
	wc.maxCount = max(wc.maxCount, wc.resultMap[WordCount])
	return wc.resultMap[WordCount]
}

func (wc *WC) cCount()(count int) { // since some chars can by multibyte
	//LPL("DEBUG: wc::cCount() ... counting chars")
	// TODO clean this up
	var discardme rune // WTH
	for _, v := range string(wc.buff) {
		count++
		discardme = v // WTH
	}
	_ = discardme

	wc.resultMap[CharacterCount] = count
	wc.totalsMap[CharacterCount] += wc.resultMap[CharacterCount]
	wc.maxCount = max(wc.maxCount, wc.resultMap[CharacterCount])
	return wc.resultMap[CharacterCount]
}

func (wc *WC) bCount() int {
	//LPL("DEBUG: wc::bCount() ... counting bytes")
	wc.resultMap[ByteCount] = len(wc.buff)
	wc.totalsMap[ByteCount] += wc.resultMap[ByteCount]
	wc.maxCount = max(wc.maxCount, wc.resultMap[ByteCount])
	return wc.resultMap[ByteCount]
}

func (wc *WC) maxLinelen() int {
	//LPL("DEBUG: wc::maxLinelen() ... calculating max line width")
	currLen := 0
	maxLen := 0
	lines := strings.FieldsFunc(string(wc.buff), isNl)
	for _, l := range lines {
		currLen = len(l)
		maxLen = max(currLen, maxLen)
	}
	wc.resultMap[MaxLinelength] = maxLen
	wc.totalsMap[MaxLinelength] = max(wc.totalsMap[MaxLinelength], wc.resultMap[MaxLinelength])
	wc.maxCount = max(wc.maxCount, wc.resultMap[MaxLinelength])
	return wc.resultMap[MaxLinelength]
}

func (wc *WC) calculateOp(op WCOption) {
	switch op {
	case NewlineCount:
		wc.resultMap[NewlineCount] = wc.nlCount()
	case WordCount:
		wc.resultMap[WordCount] = wc.wCount()
	case CharacterCount:
		wc.resultMap[CharacterCount] = wc.cCount()
	case ByteCount:
		wc.resultMap[ByteCount] = wc.bCount()
	case MaxLinelength:
		wc.resultMap[MaxLinelength] = wc.maxLinelen()

	case Totals:
	case FilesFromNul:

	default:
		LFF("FATAL: op = %v; is an unrecognized operation !\n", op)
	}
}

func (wc *WC) run() {
	for opK, opV := range wc.enabledOptions {
		if opV {
			wc.calculateOp(opK)
		}
	}
}

func (wc *WC) calculateCountWidth() int {
	maxValue := wc.maxCount
	maxValueDigits := strconv.Itoa(maxValue)
	wc.countWidth = len(maxValueDigits)
	return wc.countWidth
}

func (wc *WC) getPrint() (result string) {
	var b strings.Builder
	width := wc.calculateCountWidth()
	for i := range WCNumOfOptions {
		op := WCOption(i)
		if wc.enabledOptions[op] {
			v := wc.resultMap[op]
			fmt.Fprintf(&b, "%*v ", width, v)
		}
	}
	fmt.Fprintf(&b, "%v\n", wc.filename)
	return b.String()
}

func (wc *WC) getTotalPrint() (result string) {
	var b strings.Builder
	width := wc.calculateCountWidth()
	for i := range WCNumOfOptions {
		op := WCOption(i)
		if wc.enabledOptions[op] {
			v := wc.totalsMap[op]
			fmt.Fprintf(&b, "%*v ", width, v)
		}
	}
	fmt.Fprintf(&b, "%v\n", "total")
	return b.String()
}

func runApp(lOption, wOption, mOption, cOption, LOption, tOption bool, fOption string, flagArgs []string) (string, error) {
	//LPF("DEBUG: invoked runApp(%v,%v,%v,%v,%v,%v,%v,%v)\n", lOption, wOption, mOption, cOption, LOption, tOption, fOption, flagArgs)

	var resultB strings.Builder

	wc := newWC()
	//LPF("DEBUG: in runApp() new wc=%s\n", wc)

	if lOption || wOption || mOption || cOption || LOption { // note: omiting tOption ON purpose
		wc.setOptions(lOption, wOption, mOption, cOption, LOption)
	} else {
		wc.setOptions(true, true, false, true, false)
	}

	if len(flagArgs) == 0 && len(fOption) == 0 {
		LPL("INFO: no input file, reading stdin...")
		_, err := wc.loadFromStdin()
		if err!= nil{
			return "", err	
		}
		wc.run()
		resultB.WriteString(wc.getPrint())
	} else {

		fileNames := flagArgs //init

		/* files-from -- */
		if fOption == "-" {
			nullTermBytes, err := io.ReadAll(os.Stdin)
			if err != nil {
				LPF("ERROR: Error reading from stdin: %v\n", err)
				return "", err
			}
			//LPF("DEBUG: in runApp() %v read from stdin as fOption\n", string(nullTermBytes))

			fileNames = strings.FieldsFunc(string(nullTermBytes), isNUL)

		} else if len(fOption) > 0 {
			nullTermBytes, err := os.ReadFile(fOption)
			if err != nil {
				LPF("Error reading from files from file: %v\n", err)
				return "", err
			}
			//LPF("DEBUG: in runApp() %v read from file as fOption\n", string(nullTermBytes))

			fileNames = strings.FieldsFunc(string(nullTermBytes), isNUL)
		}
		/* -- files-from */

		for _, filename := range fileNames {
			_, err := wc.loadFromFile(filename)
			if err!= nil{
				return "", err	
			}
			wc.run()
			resultB.WriteString(wc.getPrint())
		}

	}

	if tOption {
		resultB.WriteString(wc.getTotalPrint())
	}

	//LPF("DEBUG: runApp() returning with wc=%v\n", wc)
	return resultB.String(), nil
}

func main() {

	// an overwrite
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), USAGE_TEMPLATE, PROGRAM_NAME, VERSION)
		//flag.PrintDefaults()
	}

	var lOption, wOption, mOption, cOption, LOption bool
	var tOption bool
	var fOption string
	var flagArgs []string

	flag.BoolVar(&lOption, "lines", false, "print the number of newlines")
	flag.BoolVar(&lOption, "l", false, "print the number of newlines")

	flag.BoolVar(&wOption, "words", false, "print the number of words")
	flag.BoolVar(&wOption, "w", false, "print the number of words")

	flag.BoolVar(&mOption, "chars", false, "print the number of characters")
	flag.BoolVar(&mOption, "m", false, "print the number of characters")

	flag.BoolVar(&cOption, "bytes", false, "print the number of bytes")
	flag.BoolVar(&cOption, "c", false, "print the number of bytes")

	flag.BoolVar(&LOption, "max-line-length", false, "print the widest line length")
	flag.BoolVar(&LOption, "L", false, "print the widest line length")

	flag.BoolVar(&tOption, "total", false, "print the totals")
	flag.BoolVar(&tOption, "t", false, "print the totals")

	flag.StringVar(&fOption, "filenames-nul-sep-from", "", "read NULL terminated input filenames from this file (or stdin if option is \"-\")")
	flag.StringVar(&fOption, "f", "", "read NULL terminated input filenames from this file (or stdin if option is \"-\")")

	flag.Parse()
	flagArgs = flag.Args()
	//LPF("DEBUG: in main() - flag args %v \n", flagArgs)

	appResult, appError := runApp(lOption, wOption, mOption, cOption, LOption, tOption, fOption, flagArgs)
	if appError != nil {
		LFF("DEBUG: application error!; info = %v\n", appError)
	}

	fmt.Print(appResult)

}

