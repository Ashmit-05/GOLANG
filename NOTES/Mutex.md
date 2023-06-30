In Go, a mutex is a synchronization primitive used to control access to shared resources to ensure that only one goroutine (a lightweight thread in Go) can access the resource at any given time. Mutex stands for "mutual exclusion."

Imagine you have a room with a valuable object inside, and multiple people want to access and use that object. To prevent them from entering the room simultaneously and causing conflicts or potential damage to the object, you use a lock on the door.

Similarly, in Go, a mutex acts like a lock that allows only one goroutine to access a critical section of code or a shared resource at a time. When a goroutine wants to access the shared resource, it first checks if the mutex is locked. If it is locked by another goroutine, the requesting goroutine waits until the mutex is released. Once the mutex is released, the requesting goroutine locks the mutex, gains access to the shared resource, performs its operations, and then unlocks the mutex, allowing other [[Goroutines]] to acquire it.

This way, a mutex ensures that concurrent(see [[Concurrency vs Parallelism]]) access to shared resources is controlled, preventing data races and ensuring that only one goroutine modifies the shared resource at a time, maintaining consistency and integrity.

In summary, a mutex is like a lock on a room, allowing only one goroutine to access a shared resource at a time, preventing conflicts and ensuring orderly access.