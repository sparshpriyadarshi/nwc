package main

import (
	"fmt"
	"os/exec"
	"slices"
	"strings"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

// go test -v
func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags

	testscript.Main(m, map[string]func(){
		"nwc-program": main,
	})
	m.Run()
}

// refer: https://pkg.go.dev/github.com/rogpeppe/go-internal/testscript
func TestStdinUsage(t *testing.T) {

	// one way to run this stuff: go test -run=TestStdinUsage -testwork -v
	testscript.Run(t, testscript.Params{
		Dir: "testscripts",
	})

}

type TestCase struct {
	l, w, m, c, L bool
	t             bool
	f             string
	args          []string
	wcCmd         *exec.Cmd //want
}

func TestLinecounts(t *testing.T) {
	var tests = []TestCase{
		{true, false, false, false, false, false, "", []string{"testdata/tc-empty"}, exec.Command("wc", "-l", "testdata/tc-empty")},
		{true, false, false, false, false, false, "", []string{"testdata/tc-lorem.txt"}, exec.Command("wc", "-l", "testdata/tc-lorem.txt")},
		{true, false, false, false, false, false, "", []string{"testdata/tc-mangledc.c"}, exec.Command("wc", "-l", "testdata/tc-mangledc.c")},
		{true, false, false, false, false, false, "", []string{"testdata/tc-one.txt"}, exec.Command("wc", "-l", "testdata/tc-one.txt")},
		{true, false, false, false, false, false, "", []string{"testdata/tc-proverbs.md"}, exec.Command("wc", "-l", "testdata/tc-proverbs.md")},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("lines %#v", tt.args)
		t.Run(testName, func(t *testing.T) {

			cmdoutBytes, err := tt.wcCmd.Output()
			cmdoutFields := strings.Fields(string(cmdoutBytes))
			if err != nil {
				t.Errorf("wc execution failed, test = %#v, err = %#v\n", testName, err)
			}
			result, err := runApp(tt.l, tt.w, tt.m, tt.c, tt.L, tt.t, tt.f, tt.args)
			nwcoutFields := strings.Fields(result)
			if err != nil || !slices.Equal(nwcoutFields, cmdoutFields) {
				t.Errorf("got %s, want %s", nwcoutFields, cmdoutFields)
			}
		})
	}

}

func TestWordcounts(t *testing.T) {
	var tests = []TestCase{
		{false, true, false, false, false, false, "", []string{"testdata/tc-empty"}, exec.Command("wc", "-w", "testdata/tc-empty")},
		{false, true, false, false, false, false, "", []string{"testdata/tc-lorem.txt"}, exec.Command("wc", "-w", "testdata/tc-lorem.txt")},
		{false, true, false, false, false, false, "", []string{"testdata/tc-mangledc.c"}, exec.Command("wc", "-w", "testdata/tc-mangledc.c")},
		{false, true, false, false, false, false, "", []string{"testdata/tc-one.txt"}, exec.Command("wc", "-w", "testdata/tc-one.txt")},
		{false, true, false, false, false, false, "", []string{"testdata/tc-proverbs.md"}, exec.Command("wc", "-w", "testdata/tc-proverbs.md")},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("words %#v", tt.args)
		t.Run(testName, func(t *testing.T) {

			cmdoutBytes, err := tt.wcCmd.Output()
			cmdoutFields := strings.Fields(string(cmdoutBytes))
			if err != nil {
				t.Errorf("wc execution failed, test = %#v, err = %#v\n", testName, err)
			}
			result, err := runApp(tt.l, tt.w, tt.m, tt.c, tt.L, tt.t, tt.f, tt.args)
			nwcoutFields := strings.Fields(result)
			if err != nil || !slices.Equal(nwcoutFields, cmdoutFields) {
				t.Errorf("got %s, want %s", nwcoutFields, cmdoutFields)
			}
		})
	}

}

func TestCharcounts(t *testing.T) {

	var tests = []TestCase{
		{false, false, true, false, false, false, "", []string{"testdata/tc-empty"}, exec.Command("wc", "-m", "testdata/tc-empty")},
		{false, false, true, false, false, false, "", []string{"testdata/tc-lorem.txt"}, exec.Command("wc", "-m", "testdata/tc-lorem.txt")},
		{false, false, true, false, false, false, "", []string{"testdata/tc-mangledc.c"}, exec.Command("wc", "-m", "testdata/tc-mangledc.c")},
		{false, false, true, false, false, false, "", []string{"testdata/tc-one.txt"}, exec.Command("wc", "-m", "testdata/tc-one.txt")},
		{false, false, true, false, false, false, "", []string{"testdata/tc-proverbs.md"}, exec.Command("wc", "-m", "testdata/tc-proverbs.md")},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("chars %#v", tt.args)
		t.Run(testName, func(t *testing.T) {

			cmdoutBytes, err := tt.wcCmd.Output()
			cmdoutFields := strings.Fields(string(cmdoutBytes))
			if err != nil {
				t.Errorf("wc execution failed, test = %#v, err = %#v\n", testName, err)
			}
			result, err := runApp(tt.l, tt.w, tt.m, tt.c, tt.L, tt.t, tt.f, tt.args)
			nwcoutFields := strings.Fields(result)
			if err != nil || !slices.Equal(nwcoutFields, cmdoutFields) {
				t.Errorf("got %s, want %s", nwcoutFields, cmdoutFields)
			}
		})
	}
}

func TestBytecounts(t *testing.T) {

	var tests = []TestCase{
		{false, false, false, true, false, false, "", []string{"testdata/tc-empty"}, exec.Command("wc", "-c", "testdata/tc-empty")},
		{false, false, false, true, false, false, "", []string{"testdata/tc-lorem.txt"}, exec.Command("wc", "-c", "testdata/tc-lorem.txt")},
		{false, false, false, true, false, false, "", []string{"testdata/tc-mangledc.c"}, exec.Command("wc", "-c", "testdata/tc-mangledc.c")},
		{false, false, false, true, false, false, "", []string{"testdata/tc-one.txt"}, exec.Command("wc", "-c", "testdata/tc-one.txt")},
		{false, false, false, true, false, false, "", []string{"testdata/tc-proverbs.md"}, exec.Command("wc", "-c", "testdata/tc-proverbs.md")},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("bytes %#v", tt.args)
		t.Run(testName, func(t *testing.T) {

			cmdoutBytes, err := tt.wcCmd.Output()
			cmdoutFields := strings.Fields(string(cmdoutBytes))
			if err != nil {
				t.Errorf("wc execution failed, test = %#v, err = %#v\n", testName, err)
			}
			result, err := runApp(tt.l, tt.w, tt.m, tt.c, tt.L, tt.t, tt.f, tt.args)
			nwcoutFields := strings.Fields(result)
			if err != nil || !slices.Equal(nwcoutFields, cmdoutFields) {
				t.Errorf("got %s, want %s", nwcoutFields, cmdoutFields)
			}
		})
	}
}

func TestMaxLinelengths(t *testing.T) {

	var tests = []TestCase{
		{false, false, false, false, true, false, "", []string{"testdata/tc-empty"}, exec.Command("wc", "-L", "testdata/tc-empty")},
		{false, false, false, false, true, false, "", []string{"testdata/tc-lorem.txt"}, exec.Command("wc", "-L", "testdata/tc-lorem.txt")},
		{false, false, false, false, true, false, "", []string{"testdata/tc-mangledc.c"}, exec.Command("wc", "-L", "testdata/tc-mangledc.c")},
		{false, false, false, false, true, false, "", []string{"testdata/tc-one.txt"}, exec.Command("wc", "-L", "testdata/tc-one.txt")},
		{false, false, false, false, true, false, "", []string{"testdata/tc-proverbs.md"}, exec.Command("wc", "-L", "testdata/tc-proverbs.md")},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("max width %#v", tt.args)
		t.Run(testName, func(t *testing.T) {

			cmdoutBytes, err := tt.wcCmd.Output()
			cmdoutFields := strings.Fields(string(cmdoutBytes))
			if err != nil {
				t.Errorf("wc execution failed, test = %#v, err = %#v\n", testName, err)
			}
			result, err := runApp(tt.l, tt.w, tt.m, tt.c, tt.L, tt.t, tt.f, tt.args)
			nwcoutFields := strings.Fields(result)
			if err != nil || !slices.Equal(nwcoutFields, cmdoutFields) {
				t.Errorf("got %s, want %s", nwcoutFields, cmdoutFields)
			}
		})
	}
}

func TestDefaultouts(t *testing.T) {

	var tests = []TestCase{
		{false, false, false, false, false, false, "", []string{"testdata/tc-empty"}, exec.Command("wc", "testdata/tc-empty")},
		{false, false, false, false, false, false, "", []string{"testdata/tc-lorem.txt"}, exec.Command("wc", "testdata/tc-lorem.txt")},
		{false, false, false, false, false, false, "", []string{"testdata/tc-mangledc.c"}, exec.Command("wc", "testdata/tc-mangledc.c")},
		{false, false, false, false, false, false, "", []string{"testdata/tc-one.txt"}, exec.Command("wc", "testdata/tc-one.txt")},
		{false, false, false, false, false, false, "", []string{"testdata/tc-proverbs.md"}, exec.Command("wc", "testdata/tc-proverbs.md")},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("default out %#v", tt.args)
		t.Run(testName, func(t *testing.T) {

			cmdoutBytes, err := tt.wcCmd.Output()
			cmdoutFields := strings.Fields(string(cmdoutBytes))
			if err != nil {
				t.Errorf("wc execution failed, test = %#v, err = %#v\n", testName, err)
			}
			result, err := runApp(tt.l, tt.w, tt.m, tt.c, tt.L, tt.t, tt.f, tt.args)
			nwcoutFields := strings.Fields(result)
			if err != nil || !slices.Equal(nwcoutFields, cmdoutFields) {
				t.Errorf("got %s, want %s", nwcoutFields, cmdoutFields)
			}
		})
	}
}

func TestMultifileArgs(t *testing.T) {

	var tests = []TestCase{
		{false, false, false, false, false, true, "", []string{"testdata/tc-empty", "testdata/tc-lorem.txt", "testdata/tc-mangledc.c"}, exec.Command("wc", "testdata/tc-empty", "testdata/tc-lorem.txt", "testdata/tc-mangledc.c")},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("multifile default out %#v", tt.args)
		t.Run(testName, func(t *testing.T) {

			cmdoutBytes, err := tt.wcCmd.Output()
			cmdoutFields := strings.Fields(string(cmdoutBytes))
			if err != nil {
				t.Errorf("wc execution failed, test = %#v, err = %#v\n", testName, err)
			}
			result, err := runApp(tt.l, tt.w, tt.m, tt.c, tt.L, tt.t, tt.f, tt.args)
			nwcoutFields := strings.Fields(result)
			if err != nil || !slices.Equal(nwcoutFields, cmdoutFields) {
				t.Errorf("got %s, want %s", nwcoutFields, cmdoutFields)
			}
		})
	}
}

func TestTotalwhen(t *testing.T) {

	var tests = []TestCase{
		{false, false, false, false, false, true, "", []string{"testdata/tc-empty"}, exec.Command("wc", "--total=always", "testdata/tc-empty")},
		{false, false, false, false, false, true, "", []string{"testdata/tc-lorem.txt"}, exec.Command("wc", "--total=always", "testdata/tc-lorem.txt")},
		{false, false, false, false, false, false, "", []string{"testdata/tc-empty", "testdata/tc-lorem.txt", "testdata/tc-mangledc.c"}, exec.Command("wc", "--total=never", "testdata/tc-empty", "testdata/tc-lorem.txt", "testdata/tc-mangledc.c")},
		{false, false, false, false, false, true, "", []string{"testdata/tc-empty", "testdata/tc-lorem.txt", "testdata/tc-mangledc.c"}, exec.Command("wc", "--total=always", "testdata/tc-empty", "testdata/tc-lorem.txt", "testdata/tc-mangledc.c")},
		{false, false, false, false, false, true, "", []string{"testdata/tc-empty", "testdata/tc-lorem.txt", "testdata/tc-mangledc.c", "testdata/tc-one.txt", "testdata/tc-proverbs.md"}, exec.Command("wc", "--total=always", "testdata/tc-empty", "testdata/tc-lorem.txt", "testdata/tc-mangledc.c", "testdata/tc-one.txt", "testdata/tc-proverbs.md")},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("totals when %#v", tt.args)
		t.Run(testName, func(t *testing.T) {

			cmdoutBytes, err := tt.wcCmd.Output()
			cmdoutFields := strings.Fields(string(cmdoutBytes))
			if err != nil {
				t.Errorf("wc execution failed, test = %#v, err = %#v\n", testName, err)
			}
			result, err := runApp(tt.l, tt.w, tt.m, tt.c, tt.L, tt.t, tt.f, tt.args)
			nwcoutFields := strings.Fields(result)
			if err != nil || !slices.Equal(nwcoutFields, cmdoutFields) {
				t.Errorf("got %s, want %s", nwcoutFields, cmdoutFields)
			}
		})
	}

}

func TestFilesfromNulFile(t *testing.T) {
	var tests = []TestCase{
		{true, true, true, true, true, true, "testdata/filenames.list", []string{}, exec.Command("wc", "-cmlLw", "--files0-from=testdata/filenames.list")},
		//note: -filesfrom=- is covered in the test of stdin; (testscript tests)
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("filenames from file %#v", tt.args)
		t.Run(testName, func(t *testing.T) {

			cmdoutBytes, err := tt.wcCmd.Output()
			cmdoutFields := strings.Fields(string(cmdoutBytes))
			if err != nil {
				t.Errorf("wc execution failed, test = %#v, err = %#v\n", testName, err)
			}
			result, err := runApp(tt.l, tt.w, tt.m, tt.c, tt.L, tt.t, tt.f, tt.args)
			nwcoutFields := strings.Fields(result)
			if err != nil || !slices.Equal(nwcoutFields, cmdoutFields) {
				t.Errorf("got %s, want %s", nwcoutFields, cmdoutFields)
			}
		})
	}
}

func ExampleWC_printOut_internalusage() {

	// leading and trailing chars are verbatim,
	const testString = `this is a sentence with      10	words
and 2 newlines...
`

	wc := newWC()
	wc.loadString(testString)
	wc.setOptions(true, true, false, false, true)
	//TODO: chars and bytes off-by-1, investigate why this test fails... seems to work fine when in a file...

	wc.run() //important
	fmt.Print(wc.getPrint())
	// Output:
	// 2 10 37
}

//TODO: binary files,
//TODO: fix first line item width for multifile cases

