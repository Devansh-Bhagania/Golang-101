// Overview
// Understanding Go's memory model helps write efficient applications.

// Stack vs. Heap Allocation
// 			Stack: Automatically managed, faster access.
// 			Heap: Managed by the garbage collector, used for dynamic allocation.

// Garbage Collector Mechanics
// 			Go uses a concurrent, mark-and-sweep garbage collector.
// 			Minimize allocations to reduce GC overhead.

// Optimizing Memory Usage

// 			Reuse objects when possible.
// 			Use pools (sync.Pool) for expensive allocations.
// 			Avoid unnecessary pointers.