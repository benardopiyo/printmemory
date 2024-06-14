# PrintMemory

* This function is designed to display the contents of a 10-byte array in both hexadecimal and ASCII formats. 
* This is particularly useful for debugging and examining binary data in a human-readable form.
Function Signature

## Description

* The PrintMemory function takes an array of 10 bytes as input and prints each byte's hexadecimal representation alongside its corresponding ASCII character. 
* Non-printable ASCII characters are replaced with a dot (.) to enhance readability.

* The output format is as follows:

   1. Each byte's hexadecimal value is printed first.
    2. If the byte corresponds to a printable ASCII character, it is displayed next to the hexadecimal value.
   3. Non-printable characters are represented by a dot (.).
    4. Bytes are separated by spaces for clarity, and each byte's output is clearly separated.