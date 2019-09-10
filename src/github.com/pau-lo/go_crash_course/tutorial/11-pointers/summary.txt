- Pointers

    - creating Pointers
        - Pointers types use and asterisk * as a prefix to type pointed to
        - *int - a pointer to an integer

    - Use the 'addressof' operator & to get address of variable

- Dereferencing Pointers
    - Dereferencing a pointer by preceding with an asterisk *
    - complex types (e.g. structs) are automatically dereferenced

- Create pointers to objects
    - can use the addressof (&) if value type already exists     

- Use the 'addressof' operator before initializer 
- Use the new keyword
    - can't initialize fields at the same time

- Types with internal pointers
     - all assignment operations in GO are copy operations 
     - slices and maps contain internal pointers 
     - copies point to same underlying data
