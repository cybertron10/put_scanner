# put_scanner

`put_scanner` is a simple Go-based tool to check if the `PUT` method is allowed for a list of domains by sending an `OPTIONS` request. It helps to identify potentially vulnerable endpoints where PUT requests are enabled.

## Installation

### For Go 1.18 and Above (Recommended)
If you have Go 1.18 or above, you should use the `go install` command to install the tool:

```bash
go install github.com/cybertron10/put_scanner@latest
```
This will download and install the latest version of put_scanner into your $GOBIN directory.

### For Go 1.17 and Below
If you're using Go 1.17 or below, you'll need to use the older go get command:

```
go get github.com/cybertron10/put_scanner
```
This will download and install the tool in your $GOPATH/bin directory.

Make sure that $GOPATH/bin is added to your PATH environment variable to run put_scanner from anywhere.

## Usage

Once installed, you can run the tool with the following flags:

- `-l` (required): Path to the file containing a list of domains.
- `-o`: Path to the output file where the results will be saved. Default is `output.txt`.
- `-c`: The number of concurrent workers. Default is 10.

### Example

To run the tool with an input file `domains.txt` and save the results in `vulnerable_domains.txt`:

```bash
put_scanner -l domains.txt -o vulnerable_domains.txt -c 20
```
This will check each domain in domains.txt, and any domain where the PUT method is enabled will be saved to vulnerable_domains.txt.

### License
This tool is open-source and licensed under the MIT License.

```
You can copy and paste this into your `README.md` file. Let me know if you need any changes!
```

