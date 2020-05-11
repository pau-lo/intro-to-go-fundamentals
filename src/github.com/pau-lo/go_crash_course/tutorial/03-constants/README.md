summary

- Constants 
    - immutable, but can be shadowed
    - replaced by the compiler at compile time
        - value must be calculable at compile time
    - are named like variables 
        - PascalCase for exported constants 
        - camelCase for internal constants
    - typed constants work like immutable variables
        - can interoperate only with same type
    - untype constants work like literals 
        - can interoperate  with similar types

    - Enumerated constants
    - special symbol iota allows related constants to be created easily
    - iota starts at 0 in each const block and increment by one
    - watch out of constant values that match zero values for variables

    - Enumerated expressions
        - Operations that can be determined at compile time are allowed
            - arithmetic
            - bitwise Operations
            - bitshifting 
