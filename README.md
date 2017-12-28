# su8s

filter command. shift-jis to utf-8. utf-8 to shift-jis. 

```
$ cat sjis.txt | su8 > utf8.txt
$ cat utf8.txt | u8s > sjis.txt
```

# Installation

```
$ go get -u github.com/kuked/su8s/cmd/su8
$ go get -u github.com/kuked/su8s/cmd/u8s
```

# License

MIT
