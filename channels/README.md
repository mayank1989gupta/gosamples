Go Rounines

    Whenever we launch a go program, we automatically start one goroutine
    To launch a new "go routine" --> we add "go" keyword in front of the function call will create a child routine
    Go be default will try to run with only one CPU, if you have multiple core systems
    One CPU Core <--> Go Scheduler <--> Go Routines 1, ......, Go Routines n

    Go Scheduler decides which Go Routine would be executed as Go by default Go will use only one core which could be overridden ->
    Multiple CPU Cores <--> Go Scheduler <--> Go Routine 1, ......, Go Routine n
    where, Go Scheduler runs one thread on each "logical" core

    Concurrency --> We can have multiple threads executing the same code. If one is blocked, another one is picked up and workd on. we can schedule work together but they dont get executed simaltaneously
    Paralleism --> Multiple threads get executed at the same time which could be achieved with multiple cores with Go

Channels

    Channels are the only way we can communicate between "go routines" including communication between 
    Main Routine &, Child Between
    Channels are type specific
