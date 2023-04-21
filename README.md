# fwatch

fwatch is a simple tool to watch a file and execute a command when it changes.

## Build

```bash
$ go build
```

## Usage

```bash
$ ./fwatch file.txt echo 'file changed'
```

## Install

Depending on your system the installation is different.

On Linux you can just move the compiled binary to `/usr/local/bin`:

```bash
$ sudo mv fwatch /usr/local/bin
```

Installation instructions for other systems will come soon.
