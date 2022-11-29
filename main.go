package main
import (
"bufio"
"fmt"
"io"
"os"
"strings"
)
/*
* Complete the 'ModifyString' function below and add imports if
needed.
*
* The function is expected to return a STRING.
* The function accepts STRING str as parameter.
*/
func ModifyString(s string) string {
	rns := []rune(s) // convert to rune
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
  
        // swap the letters of the string,
        // like first with last and so on.
        rns[i], rns[j] = rns[j], rns[i]
    }
  
    // return the reversed string.
    return string(rns)
}
func main() {

str := readLine(reader)

result := ModifyString(str)

fmt.Fprintf(writer, "%s\n", result)

