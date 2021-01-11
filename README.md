# fops
A command-line application to print line count and checksum of file. Build with [cobra](github.com/spf13/cobra) and integrate with CircleCI.

# 3rd-party libraries
- [cobra](github.com/spf13/cobra) v1.1.1
- [mimetype](github.com/gabriel-vasile/mimetype) v1.1.2
- [ghr](github.com/tcnksm/ghr) v0.13.0

# Build
```
$ make bild
```

# Run
```
$ go run main.go
```

# Features

Prepare input file
```bash
$ cat <<EOF > myfile.txt
how do
you
turn this
on
EOF
```
Print the line count of file, support both short and long options
```bash
$ fops linecount -f myfile.tx
4
$ fops linecount --file myfile.txt
4
```
Print the checksum of file, support multiple algorithms: md5, sha1 and sha256
```bash
$ fops checksum -f myfile.txt --md5
a8c5d553ed101646036a811772ffbdd8
$ fops checksum -f myfile.txt --sha1
a656582ca3143a5f48718f4a15e7df018d286521
$ fops checksum -f myfile.txt --sha256
495a3496cfd90e68a53b5e3ff4f9833b431fe996298f5a28228240ee2a25c09d
```
Handle non-existent, invalid input file
```bash
$ fops linecount -f non-exist-file.tx
error: No such file 'non-exist-file.txt'
$ fops linecount -f /tmp
error: Expected file got directory '/tmp'
$ fops linecount -f fops
error: Cannot do linecount for binary file 'fops'
$ fops checksum -f fops --sha256
f07bb6a888308db77fda750aa3739b7c643b07675c5c6a2d6de6c9e69de05ceb
```
Show version and help
```bash
$ fops version
fops v0.0.1
$ fops help
File Ops
Usage:
fops [flags]
fops [command]
Available Commands:
linecount Print line count of file
checksum Print checksum of file
version Show the version info
help Help about commands
Flags:
-h, --help help for fops
```
Show subcommand help
```bash
$ fops help linecount
Print line count of file
Usage:
fops linecount [flags]
Flags:
-f, --file the input file
```

# TODO

- [ ] Publish github release with CircleCI