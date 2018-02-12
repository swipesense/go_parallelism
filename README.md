# Go Parallelism

This repo is a contrived example of how io-heavy code can be optimized in go.

## Setup

This requires that both go and node are installed. Node is used to run our test server.

```
yarn install # Install package.json node dependencies.
```

## Running Examples

First, start the node server in another console:

```
node server.js
```

There are binaries for each example included in this repo, but they will only run on OSX. If you're on another os, you
can recompile:

```
# Example for example1.go
go build -o example1 example1.go
```

To execute, simply run the binary: `./example1`

## Explanation

The node server in server.js simply accepts connections, waits 1 second, then sends a "Hello World!" response. Since a
webrequest is all network io, the cpu making the request is free to execute other code while waiting for a response.
Keep this 1s network io wait time in mind while reading the examples.

### Example1 - No Parallelism

This example demonstrates hitting the server without parallelism. You can see the webRequest funcion is called
sequentially 5 times, each time waiting for the previous request to complete. If you time this program, you'll see that
it takes just over 5s (1s for each request). Most of this 5s span is wasted waiting on responses. Let's see how we can
do better.

### Example2 - Parallelism via Goroutines

In this example, we run webRequest asynchronously as a goroutine using the "go" keyword. We also introduce the concept
of wait groups. The main function of a go program determines when the program returns, so we use a waitGroup to ensure
that we wait until all requests are complete. This example runs all 5 requests in parallel. It should take just over 1s
since all 5 requests are run at the same time. The runtime of our program should now be roughly the length of our
longest web request. That's much better than the summation of every request and will scale as the number of requests
increase.

### Example3 - Workers and Channels

In this final example, we demonstrate how to limit parallelism by creating a set number of asynchronous "workers". This
might be important if you want to avoid DOSing your remote resouce (imagine if example2 made 1000 requests - that might
be too much for our server to handle simultaneously). This example also introduces the concept of channels - a simple
mechanism for communicating between asynchronous routines. We create 5 workers, and 10 jobs. In the output, you can see
each worker start up and pull jobs from the channel. Once all jobs are complete, the program exits. Because there a 10
jobs, but only 5 workers, we can expect this program to run just over 2s (10 1s jobs / 5 workers).
