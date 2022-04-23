# Command Line Vigenere Encoder/Decoder

For use as a basic cipher, and/or in CTF where [Cyberchef](https://gchq.github.io/CyberChef/) is just too much effort :)

## Usage

The arguments usable are:

- `-k`, which is required, and sets the key to use for encoding/decoding
- `-e`, optional, specifies that the input should be encoded. If omitted decoding is assumed.
- `-f`, optional, specify the file to use as input. If not provided standard input is read.

## Example of use

Assuming a text file called test.txt containing `This is a test of my vigenere encoder. Let'ts see how it Works!`:

```
PS C:\dev\vigenere-cmd> go run .\main.go -e -f test.txt -k "test"
Mlal bw s mxwl hy qq obkwgxvw xgggwxv. Dxm'xk lxi zhp ml Phvcl!
```

Or, via echoing in data and decoding by omitting `-e`:

```
PS C:\dev\vigenere-cmd> echo "Mlal bw s mxwl hy qq obkwgxvw xgggwxv." | go run .\main.go -k "test"
This is a test of my vigenere encoder.
```

Note a single character as key turns this into a rot/caesar cipher. E.g. key `n` makes it implement rot13.