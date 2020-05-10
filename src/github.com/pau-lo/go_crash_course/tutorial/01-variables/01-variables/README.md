### 3 different ways to declare a variable

- var foo int

- var foo int = 42

let the compiler decide 
foo := 42

- Can't redeclare a variables within the same scope, but can shadow them

- All variables must be used because it will triggered a compiler error

### Visibility rules

- lower case first letter variables for package scope
- upper case fist letter to export to the package all have access
- no private scope only to scope it to a block

### Naming conventions

- Pascal or camelCase

    - CAPITALIZE ACRONYMS (HTTP, URL)

- As short as reasonable

    - longer names for longer lives
    - keep names as brief and concise

- Type conversions

    - destinationType(variable) 
    - use strconv package for strings conversions 

