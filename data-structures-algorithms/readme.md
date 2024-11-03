# Data Structures & Algorithms

`Data Structure`: Structure that can store data in itself.

`Algorithm`: Problem solving flow handled every edge case and concludes with a result. The flow often contains searching or manipulating data structures and that is why they are taught together.

---

DSA (Data Structures and Algorithms) helps us to use large amount of data, store them most efficient to callback, and solve the problems most effectively.

---

`Example`: Imagine your family tree. To store your family in the computer we can use the same approach and break your each family member into a tree node.
In that example:

-   We can easily and most effectively observe relation between each member thanks to each node are connected to each other.
-   We can easily and most effectively find a member without traversing all your family member by following relations between members.

Now imagine there is a really large family, in that case it is genuinely hard to manage the data without automating its flow with algorithms.

---

In computer science, databases and internet indexing services can be considered as the same, therefore DSA has a massively and vitaly important in computer science, and most of the data is stored in `data structures` as well as managed or observed or any manipulation is done by `algorithms` in DSA.

---

There are 2 types of Data Structures:

-   Primitive Data Structures: basic data structures that struct more complex data structures. These basic structures are provided by the programming language itself by default. These are the foundation stone, minimum required data structures to create everything other. Examples: `integer, float, char, bool, ...`
-   Abstract Data Structures: higher-level data structures that are built using primitive data types and provide more complex and problem-specific, specialized operations. Examples: `array, linked list, stack, queue, tree, graph, ...`

Every data structure and algorithm is best for specific situations. If the engineer can choose the best fit data structures and algorithms for the existing situation can solve the related problems faster, using less memory, and systematically.

> DSA is fundamental in nearly every part of the software world.

> Every DSA has its own use case, there is no such DSA fits everything most efficient.

---

# Big O notation O(some_complexity), ex: O(1), O(n), ...

https://www.youtube.com/watch?v=__vX2sjlpXU

Big O notation represents an algorithm's time comlexity. It is the main comparasion metric in computer science.

If there are 2 algorithms solve the same problem with sequentially O(n) and O(1), then we say second algorithm is more efficient and better than the first one in terms of optimizations.

# Most Common Data Structures, Their Operations & Their Time Complexity

![bigocheatsheet.com](./bigocheatsheet.com.png)

### Array

Array is a data structure that stores a fixed-size sequential collection of elements of the same type. An array is used to store a collection of data, but it is often more useful to think of an array as a collection of variables of the same type.

pros:

```diff
+ fast access
+ fast search
+ fast insertion
+ fast deletion
+ less memory allocation
```

cons:

```diff
- immutable
- pre-allocated
- reallocation
```

operations:

```json
access: 1
search: "n"
deletion: "n"
deletion: "n"
space: "n"
```

-   Array can easily give the address of the indexed value due to sequentially allocated structure of it in the memory. That way it directly reaches the address.

-   Searching in array is also sequentialy, and that cause finding an array by its value but not index (finding the value by index is access) to be as much comparasion as the length of the array.

-   Arrays are immutable and pre-allocated structures in memory. That causes if there needs to be more memory allocation, then there will be need for the reallocation. This is as much movement as the length of the array.

-   Deletion has the same reason why it is O(n) with the insertion.

-   Array just allocates as much as the element size \* element count, and so that there is only as much as array's elements allocated in the memory.

use cases:

> When the size of the data is known and fixed.

> When the data is not going to be changed frequently.

> When the data is going to be accessed frequently.

---

### Stack

[Go to code >](./data-stuctures/stack/stack.go)

https://www.youtube.com/watch?v=KcT3aVgrrpU

Stack is a data structure that follows the last in, first out (LIFO) principle. It has two main operations: push, which adds an element to the collection, and pop, which removes the most recently added element that was not yet removed.

pros:

```diff
+ fast insertion
+ fast deletion
```

cons:

```diff
- slow access
- slow search
```

operations:

```json
access: "n"
search: "n"
insertion: 1
deletion: 1
space: "n"
```

-   As there is LIFO (Last in, First out) mechanism in Stack, that restricts to reach other than the most top element. That causes both the access and delete operations has O(n) complexity.

-   Insertion and deletion operations are O(1) because of the need to just change the pointers of the nodes.

use cases:

> When the data is going to be accessed in a LIFO manner.

> When the data is going to be inserted and deleted frequently.

---

### Queue:

[Go to code >](./data-stuctures/queue/queue.go)

https://www.youtube.com/watch?v=D6gu-_tmEpQ

Queue is a data structure that follows the first in, first out (FIFO) principle. It has two main operations: enqueue, which adds an element to the collection, and dequeue, which removes the oldest element that was added.

pros:

```diff
+ fast insertion
+ fast deletion
```

cons:

```diff
- slow access
- slow search
```

operations:

```json
access: "n"
search: "n"
insertion: 1
deletion: 1
space: "n"
```

-   As there is FIFO (First in, First out) mechanism in Queue, that restricts to reach other than the most latest element. That causes both the access and delete operations has O(n) complexity.

-   Insertion and deletion operations are O(1) because of the need to just change the pointers of the nodes.

use cases:

> When the data is going to be accessed in a FIFO manner.

> When the data is going to be inserted and deleted frequently.

---

### Singly-Linked-List:

[Go to code >](./data-stuctures/singly-linked-list/singly_linked_list.go)

https://www.youtube.com/watch?v=F8AbOfQwl1c

Singly-Linked-List is a data structure that has a head node and each node has a pointer to the next node. That causes the access and search operations to be O(n) because of the need to traverse all nodes to reach the desired node.

pros:

```diff
+ fast insertion
+ fast deletion
```

cons:

```diff
- slow access
- slow search
```

operations:

```json
access: "n"
search: "n"
insertion: 1
deletion: 1
space: "n"
```

-   Singly-Linked-List is a data structure that has a head node and each node has a pointer to the next node. That causes the access and search operations to be O(n) because of the need to traverse all nodes to reach the desired node.
-   Insertion and deletion operations are O(1) because of the need to just change the pointers of the nodes.

use cases:

> When the data is going to be inserted and deleted frequently.

---

### Doubly-Linked-List

No code beacuse same with the singly linked list, just add `pointer previous` to Node struct.

https://www.youtube.com/watch?v=e9NG_a6Z0mg

pros:

```diff
+ fast insertion
+ fast deletion
```

cons:

```diff
- slow access
- slow search
```

operations:

```json
access: "n"
search: "n"
insertion: 1
deletion: 1
space: "n"
```

-   Doubly-Linked-List is a data structure that has a head node and each node has a pointer to the next and previous node. That causes the access and search operations to be O(n) because of the need to traverse all nodes to reach the desired node.

-   Insertion and deletion operations are O(1) because of the need to just change the pointers of the nodes.

use cases:

> When the data is going to be inserted and deleted frequently.

> When the data is going to be accessed in both directions.

---

### Skip List

[Go to code >](./data-stuctures/skip-list/skip_list.go)

https://www.youtube.com/watch?v=NDGpsfwAaqo

Skip List is a data structure that is a probabilistic data structure that allows fast search within an ordered sequence of elements. It is a data structure that allows fast search within an ordered sequence of elements.

pros:

```diff
+ fast access
+ fast search
+ fast insertion
+ fast deletion
```

cons:

```diff
- a bit more memory allocation, due to the need of the skip pointers in each node.
```

operations:

```json
access: "log(n)"
search: "log(n)"
insertion: "log(n)"
deletion: "log(n)"
space: "nlog(n)"
```

-   Skip List is a data structure that is a probabilistic data structure that allows fast search within an ordered sequence of elements. It is a data structure that allows fast search within an ordered sequence of elements.

-   Skip List has a bit more memory allocation, due to the need of the skip pointers in each node.

-   Skip List has O(log(n)) time complexity for access, search, insertion, and deletion operations.

use cases:

> When the data is going to be accessed, searched, inserted, and deleted frequently.

---

### Hash Table

[Go to code >](./data-stuctures/hash-table/hash_table.go)

https://www.youtube.com/watch?v=knV86FlSXJ8

Hash Table is a data structure that implements an associative array abstract data type, a structure that can map keys to values. A hash table uses a hash function to compute an index into an array of buckets or slots, from which the desired value can be found.

Implementation of a `dictionary` or `map` using a hash function. This hash function is used to convert the key into the index of the array.

Sometimes, there can be a collision, and that is handled by chaining or open addressing.

-   Chaining: If there is a collision, then the values are stored in a linked list.
-   Open Addressing: If there is a collision, then the values are stored in the next available slot.

Good hash function should be:

-   Deterministic: Same input should always produce the same output.
-   Fast: The hash function should be fast to compute.
-   Uniform: The hash function should uniformly distribute the keys. To achieve this, maximum randomness and least collisions are required.

Hashing functions:

> m = size of the hash table

-   Division Method: h(k) = k mod m
-   Multiplication Method: h(k) = floor(m _ (k _ A mod 1))
-   Universal Hashing: h(k) = ((ak + b) mod p) mod m
-   Perfect Hashing: h(k) = a \* k + b
-   Consistent Hashing: h(k) = hash(k)

pros:

```diff
+ fast access
+ fast search
+ fast insertion
+ fast deletion
```

operations:

```json
access: N/A
search: 1
insertion: 1
deletion: 1
```

-   Hash Table has O(1) time complexity for search, insertion, and deletion operations.
-   Hash Table has N/A time complexity for access operations, because it is not possible to access the hash table by index. Instead, it is accessed by key and that is called search operation.

use cases:

> Real world examples: Databases, caches, sets, maps, ...

> When the data is going to be accessed, searched, inserted, and deleted frequently.

---
