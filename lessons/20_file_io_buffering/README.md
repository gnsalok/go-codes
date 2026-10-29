# Lesson 20: File I/O, Streaming, and Buffering

## Learning goals

- Keep resource ownership and close responsibility explicit.
- Process large inputs with bounded memory through `io.Reader` and `io.Writer`.
- Explain what application buffering changes and what the operating system already
  buffers.
- Choose line scanning, byte buffering, chunked reads, copying, or atomic replacement
  from the data contract.

## Mental model

`io.Reader` and `io.Writer` describe streaming boundaries without owning resource
lifecycle. A read may return data and an error together, so process `n > 0` bytes
before handling the error. `io.EOF` means the stream ended normally; it is not an
operational failure.

Small application reads against an `os.File` can cause many system calls. A
`bufio.Reader` fetches larger blocks and serves small reads from memory. If the
algorithm naturally processes blocks, reading directly into a reusable chunk is
usually clearer. Measure with the actual storage and access pattern before choosing
a buffer size.

## Run and test

The experiment creates an 8 MiB temporary file, compares three strategies, and
removes the file when it exits:

```bash
go run ./lessons/20_file_io_buffering/cmd/read_strategies
go test ./lessons/20_file_io_buffering
```

## Production tradeoffs and common mistakes

- `bufio.Scanner` has a token-size limit; configure it or use `bufio.Reader` for
  potentially large records.
- `defer file.Close()` is clear for one resource, but deferring inside a long loop
  can keep many descriptors open.
- Check `Flush`, `Close`, and final write errors when durability matters.
- For atomic replacement, write a temporary file in the destination directory,
  flush and close it, then rename it; define the required durability separately.
- Never read an unbounded file or response entirely into memory without a size limit.

## Interview prompts

1. Why may `Read` legally return both bytes and `io.EOF`?
2. What is the difference between `bufio.Reader`, the OS page cache, and disk cache?
3. How would you safely replace a configuration file after a crash?
4. When is `io.Copy` preferable to a manual read loop?

## Practice

Add an injected reader that returns data and a non-EOF error in the same call. Prove
the returned statistics include the data before the error is reported.
