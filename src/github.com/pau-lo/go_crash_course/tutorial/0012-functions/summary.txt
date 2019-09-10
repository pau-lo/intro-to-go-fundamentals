 - Basic Syntax 

 - Parameters
    - comma delimited list of variables and types
        - func foo(bar string, cost int) 
    - when pointers are passed in, the function can change the value in the caller
        - this always tre for dat of slices and maps 
    - variadic parameters to send list of same types in
        - Must bve last parameters
        - Received as a slices
        - func foo (bar string, cost in)

    - Return values
        - single return values just list types
        - multiple return values list types surrounded by parentheses
            - func foo() (int, error)
        - the (result type, error) paradigm is a very common idiom 
        - can use named return values
            - using return keyword on its own
        - can return addresses of local variables
            - automatically promoted from local memory (stack) to shared memory (heap)

 - Anonymous functions
    - functions don't have name if they are immediately invoke
    - assigned to a variable or passed as argument to a function 

 - Functions as types 
    can assigned functions to variables or use as arguments and return values in functions


 - Methods
    - special type of function
    - function that executes in context of a type
    - receiver can be value or pointers
        - value receiver gets copy of type
        - pointer receiver gets pointer to type
