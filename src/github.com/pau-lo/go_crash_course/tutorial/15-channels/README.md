- channel basic
    - create a channel with make command 
        - make(chan int)
    - Send message into channel
        - ch <- val
    - receiver message from channel
        - val := <-ch

- restricting data flow
    - channel can be cast into send-only or receive only versions
        - send-only: chan int
        - receive-only: <- chan int

- buffered channels
    - channels block sender side till receiver is available  
    - Block receiver side till message is available
    - can decouple sender and receiver with buffered channels
        - make(chan int, 50)
    - use buffered channels when send and receiver have assymmetric loading

- for .... range loops with channels
    - use the monitor channel and process messages as they arrive
    - loop exits when channel is closed

- select statements
    - allow a go-routine to monitor several channels at once
    - BLocks if all channels block
    - if multiple channels receive value simultaneously, behavior is undefined










