# put_scanner

`put_scanner` is a simple Go-based tool to check if the `PUT` method is allowed for a list of domains by sending an `OPTIONS` request. It helps to identify potentially vulnerable endpoints where PUT requests are enabled.

## Installation

To install the tool using `go get`:

```bash
go get github.com/cybertron10/put_scanner
```
Usage
Once installed, you can run the tool with the following flags:

-l (required): Path to the file containing a list of domains.
-o: Path to the output file where the results will be saved. Default is output.txt.
-c: The number of concurrent workers. Default is 10.
