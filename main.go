package main

import "github.com/01-edu/z01"

// PrintMemory takes an array of 10 bytes and displays their memory in hexadecimal
// format followed by their ASCII representation, with non-printable characters replaced by dots.
func PrintMemory(arr [10]byte) {
    // Loop through each byte in the array
    for i := 0; i < len(arr); i++ {
        // Print the hexadecimal representation of the byte, padded to 2 characters
        hexDigit1 := arr[i] / 16
        hexDigit2 := arr[i] % 16
        z01.PrintRune(rune('0' + hexDigit1))
        z01.PrintRune(rune('0' + hexDigit2))
        z01.PrintRune(' ')

        // Determine the ASCII character, replace non-printable characters with '.'
        if arr[i] >= 32 && arr[i] <= 126 {
            z01.PrintRune(rune(arr[i]))
        } else {
            z01.PrintRune('.')
        }
    }

    // Newline for better readability
    z01.PrintRune('\n')

    // Loop again to print the ASCII line, replacing non-printable characters with dots
    for i := 0; i < len(arr); i++ {
        if arr[i] >= 32 && arr[i] <= 126 {
            z01.PrintRune(rune(arr[i]))
        } else {
            z01.PrintRune('.')
        }
    }

    // Print a newline at the end for clean output
    z01.PrintRune('\n')
}

func main() {
    // Initialize an array of 10 bytes
    arr := [10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*', 0, 127}

    // Call PrintMemory with the array
    PrintMemory(arr)
}
