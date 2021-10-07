/*
 * Read file to get variable values
 * For example a configuration.php with database variables configuration
 *
 * @Author	Eduardo Jiménez (eduk2)
 */

package govarsfromfiles

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type fileController struct {
	file        *os.File //file for reading
	pathFile    string   //file path included the name file
	err         error    //error control when you try to open the file
	initialized bool     //true if the file is opened successfully
	regularExp  string   //regular expression to find the value in a line of the file
}

//The Public struct to work with this module
var File fileController

//This type is used to return the values of all variables that you need to find in a file
type fileValues map[string]string

// Set file path (included the file name) where you need to get variable values
// Return true if the file is opened successfully
func (f *fileController) SetPath(pathFile string) bool {

	f.pathFile = pathFile

	f.file, f.err = os.Open(f.pathFile)
	if f.err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", f.err)
		os.Exit(1)
	}
	f.initialized = true

	fmt.Println("File opened successfully:", f.pathFile)
	return f.initialized
}

// To get the variable values from a file defined with SetPath
// varsIn variables you need to get. Examples: $db_host, $db_username...
// Return variable values from file defined in SetFile
func (f *fileController) GetValues(varsIn ...string) fileValues {

	if f.initialized != true {
		fmt.Println("Error: File not initialized. You must use SetFile")
		os.Exit(1)
	}

	defer f.closeFile()

	scanner := bufio.NewScanner(f.file)

	var values = make(fileValues)

	var line string
	for scanner.Scan() {
		line = scanner.Text()

		for _, searching := range varsIn {
			if strings.Contains(line, searching) {

				re := regexp.MustCompile(File.regularExp)
				match := re.FindStringSubmatch(line)
				if len(match) > 1 {
					values[searching] = match[1]
				} else {
					fmt.Println(searching, "match not found")
				}
			}
		}
	}
	return values
}

func (f *fileController) closeFile() {
	err := f.file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("File closed", f.pathFile)
}

// Set regular expression to find values in a line of the the file
// By default to find string between " "  => `"(.*)"`
// Other typical regular expression is ' ' =>  `'(.*)'`
func (f *fileController) SetRegularExp(regularExp string) {
	f.regularExp = regularExp
}

// Set the Regular expression to find values in a line y default
// The regular expression is string between " "
func init() {
	File.regularExp = `"(.*)"`
}
