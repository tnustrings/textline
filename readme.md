# textline

languages word by word in go

```
$ echo "Dzień dobry" | textline pl
Dzień [[day]] dobry [[good]]
```

## install and try

**on linux with apt**

download the [textline-\<version\>.deb](https://github.com/tnustrings/textline/releases), then say

```
sudo apt install ./textline-<version>.deb
```

**on any os with go**

```
go install github.com/tnustrings/textline
```

**try**

try feeding it some text:

```
echo "Dzień dobry" | textline pl
```

there's also a textline [browser
extension](https://chromewebstore.google.com/detail/tagid-with-google-dict/aacfmkdpcdadjcpohbjfcddomedmfdai).

## code

the code is written in codetext, install from
[here](https://github.com/tnustrings/ct).

then say:

```
ct textline.ct
```

to generate the go code.