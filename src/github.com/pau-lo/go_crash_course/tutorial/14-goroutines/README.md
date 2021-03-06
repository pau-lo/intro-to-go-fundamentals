- Best Practices 
    - Don't create go-routines in libraries
    - Let consumer control concurrency
    - When creating a go-routine know how it will end
        - avoids subtle memory leaks
    - Check for race conditions at compile time


- Creating go-rooutines
    - use go keyword in front of function call
    - when using anonymous functions, pass data as local variables

- Synchronization
    - use sync.WaitGroup to wait for groups of goroutines to complete
    - use sync.Mutex and sync.RWMutex to protect data access

- Parallelism
    - By default, Go will use CPU threads equal to available cores
    - change with runtime.GOMAXPROCS
    - more threads can increase performance, but too many can slow it down

