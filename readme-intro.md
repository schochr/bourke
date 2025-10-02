## Ordered Data Structures

Disclaimer: The _Red-Black Tree_ implementation is a port of Java's [TreeMap](https://docs.oracle.com/en/java/javase/24/docs/api/java.base/java/util/TreeMap.html) implementation in the standard library.

`bourke` provides implementations of a [_Red-Black Tree_](https://en.wikipedia.org/wiki/Red%E2%80%93black_tree), a self-balancing binary search tree and a [_Trie_](https://en.wikipedia.org/wiki/Trie), also 
called _Prefix Tree_.

They provide the standard map operations _Put_, _Remove_ and _Get_ with the expected runtime complexity for the _Red-Black Tree_ 
structure of _O(log N)_ where _N_ is the number of vertices and _O(L)_ for the _Trie_ structure where _L_ corresponds to the length 
of the longest key in the tree.<br>
[Standard iterators](https://pkg.go.dev/iter) can be retrieved from both structures for in-order traversal. Further, "Bounded" iterators can be obtained 
to in-order traverse contiguous parts of the trees, bounded by either a lower bound, upper bound or both.

The basic structures `Tree` and `Trie` are not thread-safe. Synchronization needs to be provided if they are to be used concurrently. 
The [kernel structure for Linux](https://www.kernel.org/doc/html/latest/core-api/rbtree.html) for instance, is not thread-safe and neither
does it provide any means of locking. It is generally too expensive to do fine-grained locking for connected structures like this and 
may not be needed by users.
Thread-safe versions of the structures are provided, `TreeConcurrent` and `TrieConcurrent` (**_WORK-IN-PROGRESS_**). They manage a slice of their non-thread-safe 
building blocks (partitions). A partition is selected based on the output of consistent non-cryptographic hashing of the key modulo the
number of partitions. All **write** and **single-result read traversals** like `Get` or `Successor` are _**always**_ synchronized.
Iterators are not synchronized for two reasons: _a)_ **All** partitions would have to be locked and _b)_ no statement can be made about
how they will be used and latency could be introduced in between iterations. Hence, read-locking over all partitions is provided but its 
usage is left to the user.