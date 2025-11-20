# lingobin

languages word by word in go

```
$ echo "Dzień dobry" | lingobin pl
Dzień [[day]] dobry [[good]]
```

## install and try

**on linux with apt**

download the [lingobin-\<version\>.deb](https://github.com/tnustrings/lingobin/releases), then say

```
sudo apt install ./lingobin-<version>.deb
```

**on any os with go**

```
go install github.com/tnustrings/lingobin
```

**try**

try feeding it some text:

```
echo "Dzień dobry" | lingobin pl
```

there's also a lingobin [browser
extension](https://chromewebstore.google.com/detail/tagid-with-google-dict/aacfmkdpcdadjcpohbjfcddomedmfdai).

## code

the code is written in codetext, install from
[here](https://github.com/tnustrings/ct).

then say:

```
ct lingobin.ct
```

to generate the go code.