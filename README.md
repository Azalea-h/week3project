# TCP Server

This project takes everything you have learned and applies it to a TCP Server
which you will write. This allows you to demonstrate your understanding of Go.
If time permits you can also attempt the optional tasks.

A test harness is provided to allow you to perform _basic_ checks on your 
solution, but part of the project is providing a full suite of tests.

## Requirements

A TCP Server that will act as an in memory key value store. Clients can connect
to the server and read/write/delete data. The server must have a full test suite
written by you. The project should be idiomatic Go.

### Protocol

The protocol utilise 3 byte command strings. Some command strings may require
arguments. An argument is provided in 3 parts.

Part 1 is a single byte `[0-9]` as indicates how many bytes are used to describe 
the argument size.

Part 2 is an `n` byte size string of digits where `n` is the length specified
by Part 1. This indicates how big the argument actually is.

Part 3 is `n` bytes where `n` was defined by part 2. Part 3 will always be 
representable as a `string` type.

For example:

```
put13key212stored value
```

`put`, argument 1 has a 1 byte size component, and is 3 bytes long. Argument
1 is `key`. Argument 2 has a 2 byte size component and is 12 bytes long.
Argument 2 is `stored value`.

The server will respond to any invalid request with `err`.

### Put

Put a value into the store under the given key overwriting it if it already 
exists:

```
put<key><value>
```

Both `key` and `value` are required arguments.

Response:

```
ack
```

Example:
```
-> put11a13foo
<- ack
```

### Get

Get a value from the store under the give key:

```
get<key>
```

`key` is a required argument.

Responses:

```
val<value>
nil
```

Where `ack` will be sent with the value if the key was found, and `nil` will be
sent if the key was not found.

Example:
```
-> get11a
<- val13foo
-> get11b
<- nil
```

### Delete

Delete a value from the store under the given key. Deleting a key that doesn't
exist is a no-op.


```
del<key>
```

`key` is a required argument.

Response:

```
ack
```

Example:
```
-> del11a
<- ack
-> del11a
<- ack
```

### Terminate the connection

```
bye
```

No response is given, the server should hang up.

## Stretch Goals

### Variable length get

Modify the get command to take a second, special argument. This would be a
single byte indicating the size of the next block, then `n` bytes indicating
how many bytes should be returned. If `n == 0` then all the bytes should be
returned. If `n > len(value)` then `len(value)` bytes should be returned.

For example:

```
-> put11k15vvvvv
<- ack
-> get11k0
<- val15vvvvv
-> get11k12
<- val11vv
-> get11k299
<- val15vvvvv
```

### Distributed store

Make the store distributed. Allow a connection to any one of `n` servers and
read/write will propagate through the cluster. You can make `put` block while 
the data propagates to avoid consistency errors for a single client. Write a 
test harness for your implementation.

### UDP store

Use UDP to broadcast updates, remembering that UDP is not a guaranteed protocol. 
Servers should not need to know about each other and `put` should no longer 
block. Write a test harness for your implementation. Since the whole system is
now eventually consistent you will likely need to put waits into the test
harness to allow data to propagate after writes.

### Go Nuts!

Add your own objectives to show off your new skills.
