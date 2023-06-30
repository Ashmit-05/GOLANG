To check for race conditions, we use the command :
```
go run --race <filename>
```

A race condition occurs when multiple concurrent processes or [[Goroutines]] access and manipulate shared data concurrently, leading to unexpected and erroneous results. It happens when the order and timing of operations are not properly coordinated, causing conflicts in data access and modification.

To understand race conditions, consider an example where two [[Goroutines]] are incrementing a shared variable `count` simultaneously. Both [[Goroutines]] read the initial value of `count`, increment it, and write the updated value back. However, if the timing and order of these operations are not properly synchronized, race conditions can occur.

For instance, if both [[Goroutines]] read the value of `count` simultaneously, they might get the same initial value. Then, both [[Goroutines]] increment it and write back the updated value. As a result, the final value of `count` may not reflect the intended increments from both [[Goroutines]] because they overwrote each other's changes.

Race conditions can lead to inconsistent and unpredictable program behavior, making debugging and understanding code behavior difficult. These issues are particularly tricky because they might not occur consistently and may depend on the timing and scheduling of the [[Goroutines]].

To prevent race conditions, Go provides synchronization primitives such as [[Deadlocks]], [[Channels]], and other mechanisms. By properly using these synchronization techniques, you can ensure that only one goroutine accesses shared data at a time, preventing conflicting modifications.

In summary, [[Goroutines]] enable concurrent(see [[Concurrency vs Parallelism]]) execution in Go, allowing multiple tasks to run simultaneously. However, if not synchronized properly, these [[Goroutines]] can lead to race conditions where concurrent access to shared data causes unexpected and incorrect results. To avoid race conditions, synchronization mechanisms should be used to coordinate access to shared resources and ensure consistent and predictable program behavior.