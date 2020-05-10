Summary

- Boolean type  
    - values are true or false
    - not an alias fo other types (e.g. int)
    - zero values is false

- Numeric types
    - integers
        - signed integers
            - int type has varying size, but min 32 bits
            - 8 bit(int8) through 64 bit (int64)
        - unsigned integers
            - 8 bit (byte and uint8) through 32 bit(uint32)
        - Arithmetic operations
            - addition, subtraction, multiplications, division, remainder
        - bitwise operations
            - and, or, xor, and not
        - zero value is 0
        - can't mix types in same family! (uint16 + uint32 = error)
    
    - floating point numbers
        - follow IEEE-754 standard
        - zero value is 0
        - 32 and 64 bit versions
        - Literal styles
            - decimal
            - exponential
            - mixed 
        - Arithmetic operations
            - addition, subtraction, multiplications, division

    - Complex numbers
        - zero value is (0+0i)
        - 64 and 128 bits versions
        - built in functions
            - complex - make complex number from two floats
            - real - get real part as float
            - imag - get imaginary part as a float
        - Arithmetic operations
            - addition, subtraction, multiplications, division
  
- Text types
    - strings
        - UTF-8
        - immutable
        - can be concatenated with plus operator
        - can be converted to []byte

    - Rune
        - UTF-32 
        - alias for int32
        - special methods normally required to process
            - e.g. strings, reader
