![[Pasted image 20230627012742.png]]
Concurrency and parallelism are related but distinct concepts in the context of Go programming. Here's a simple explanation of both:

Concurrency: Concurrency refers to the ability of a program to execute multiple tasks simultaneously, where tasks can start, run, and complete independently of each other. In Go, concurrency is achieved through [[Goroutines]] and channels. [[Goroutines]] are lightweight threads that enable concurrent execution, allowing multiple functions to run concurrently within the same program. Channels provide a way for [[Goroutines]] to communicate and synchronize their actions. Concurrency in Go enables efficient utilization of resources and can enhance the responsiveness and performance of a program.

Parallelism: Parallelism, on the other hand, involves the simultaneous execution of multiple tasks, where each task is executed in parallel on separate processors or cores. It implies the actual execution of multiple computations at the same time. In Go, parallelism can be achieved by utilizing multiple [[Goroutines]] that run simultaneously across different CPUs or CPU cores. However, it's important to note that parallelism requires multiple processors or cores to achieve true parallel execution. If the underlying hardware doesn't support parallel execution, [[Goroutines]] will still be executed concurrently but not in parallel.

In essence, concurrency in Go allows for the design of programs that can efficiently manage multiple tasks and achieve responsiveness, while parallelism leverages multiple processors or cores to execute computations simultaneously, further improving performance when available.


![[Pasted image 20230627013155.png]]
![[Pasted image 20230627013345.png]]

**We use the 'go' keyword to create a goroutine**

See [[Mutex]] for more info
