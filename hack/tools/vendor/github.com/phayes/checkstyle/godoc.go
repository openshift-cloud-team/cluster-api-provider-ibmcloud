/*
Package checkstyle allows the parsing of generation of checkstyle XML files.

Checkstyle XML files are a standard file format for reporting errors in source code, and is often generated by static analysis tools.

Example usage:
  // Print XML into human readable format
  checkSyle, err := checkstyle.ReadFile("checkstyle_report.xml")
    if err != nil {
    log.Fatal(err)
  }
  for _, file := range checkStyle.File {
    fmt.Println(File.Name)
    for _, codingError := range file.Error {
      fmt.Println("\t", codingError.Line, codingError.Message)
    }
  }

  // Create a new XML file from scratch
  check := checkstyle.New()

  // Ensure that a file has been added
  file := check.EnsureFile("/path/to/file")

  // Create an error on line 10, column 5
  codingError := checkstyle.NewError(10, 5, checkstyle.SeverityWarning, "format", "line must end with a full stop")

  // Add the error to the file
  file.AddError(codingError)

  // Output XML
  fmt.Print(check)

For more information on checkstyle XML see: http://checkstyle.sourceforge.net/checks.html
*/
package checkstyle
