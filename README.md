# golib
A collection of data structures implemented in Go

## What's Available
The data structures available try to be as flexible as possible. Objects are added as interfaces, and you can define
your own key comparison functions, which allows set-like behavior as well as implementations
closer to associative arrays.

* Dictionary - A variation of splash tables, which are bucketized cuckoo hash tables. Concurrency is still in the works,
but optimizations such as tabulation hashing for strings make it a viable option.
* Heap - A binary heap with default creational functions for min- and max-heaps
* List - A standard doubly-linked list