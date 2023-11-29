// go function that reads a text file from http URL that contains
// a number per line and calculates the number of non-empty lines
// Declare a variable to hold the line count
var lineCount int

// Open the text file
file, err := os.Open("/path/to/file.txt")
if err != nil {
  log.Fatal(err)
}
defer file.Close()

// Scan the file line-by-line  
scanner := bufio.NewScanner(file)
for scanner.Scan() {

  // Increment the line count  
  lineCount++

  // Get the current line
  line := scanner.Text()

  // Check if the line is empty
  if line != "" {

    // TODO: Add code here to process non-empty line

  }

}

// After the loop, lineCount will contain 
// the number of non-empty lines
fmt.Println("Non-empty line count:", lineCount)
