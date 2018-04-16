# subst - substitute $variables in text

`subst` parses its arguments as `KEY=VALUE` pairs, substitutes `$KEY`s by `VALUE`s in stdin, and writes the result to stdout.

- [Get it](#get-it)
- [Use it](#use-it)
    - [JSON output](#json-output)
    - [Template output](#template-output)
    - [Setting URL components](#setting-subst-components)
- [Comments](https://github.com/sgreben/subst/issues/1)

## Get it

Using go get:

```bash
go get -u github.com/sgreben/subst/cmd/subst
```

Or [download the binary](https://github.com/sgreben/subst/releases/latest) from the releases page. 

```bash
# Linux
curl -LO https://github.com/sgreben/subst/releases/download/${VERSION}/url_${VERSION}_linux_x86_64.zip
unzip url_${VERSION}_linux_x86_64.zip

# OS X
curl -LO https://github.com/sgreben/subst/releases/download/${VERSION}/url_${VERSION}_osx_x86_64.zip
unzip url_${VERSION}_osx_x86_64.zip

# Windows
curl -LO https://github.com/sgreben/subst/releases/download/${VERSION}/url_${VERSION}_windows_x86_64.zip
unzip url_${VERSION}_windows_x86_64.zip
```

## Use it

`subst` reads definitions from CLI arguments, reads from stdin, and writes to stdout. The stdin input is line-buffered.

```text

subst [OPTIONS] [KEY=VAL [KEY=VAL...]]

Usage of subst:
  -q    suppress all logs
  -unknown value
        handling of unknown keys, one of [ignore empty error] (default ignore)
```

### Examples

```shell
$ cat template.yaml
name: $NAME
url: $URL
comment: $COMMENT

$ cat template.yaml | subst -q NAME=jp  URL=github.com/sgreben/jp COMMENT=
name: jp
url: github.com/sgreben/jp
comment: 

$ cat template.yaml | subst -q NAME=jp 
name: jp
url: $URL
comment: $COMMENT

$ cat template.yaml | subst -q -unknown=empty NAME=jp 
name: jp
url: 
comment: 
```

Without the `-q` flag, `subst` also prints to stderr:

```shell
2018/04/16 15:18:37 undefined key: $URL, using value "$URL"
2018/04/16 15:18:37 undefined key: $COMMENT, using value "$COMMENT"
```

You can substitute environment variables (like `envsubst`) by letting the shell substitute their values:

```shell
$ subst USER="$USER" SHELL="$SHELL"
$USER
sgreben
$SHELL
/bin/zsh
```

## Comments

Feel free to [leave a comment](https://github.com/sgreben/subst/issues/1) or create an issue.