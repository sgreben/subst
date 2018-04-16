# subst - substitute $variables in text

`subst` parses its arguments as `KEY=VALUE` pairs, substitutes `$KEY`s by `VALUE`s in stdin, and writes the result to stdout.

<!-- TOC -->

- [Get it](#get-it)
- [Use it](#use-it)
    - [Examples](#examples)
- [Comments](#comments)

<!-- /TOC -->

## Get it

Using go get:

```bash
go get -u github.com/sgreben/subst
```

Or [download the binary](https://github.com/sgreben/subst/releases/latest) from the releases page. 

```bash
# Linux
curl -LO https://github.com/sgreben/subst/releases/download/1.1.0/subst_1.1.0_linux_x86_64.zip
unzip subst_1.1.0_linux_x86_64.zip

# OS X
curl -LO https://github.com/sgreben/subst/releases/download/1.1.0/subst_1.1.0_osx_x86_64.zip
unzip subst_1.1.0_osx_x86_64.zip

# Windows
curl -LO https://github.com/sgreben/subst/releases/download/1.1.0/subst_1.1.0_windows_x86_64.zip
unzip subst_1.1.0_windows_x86_64.zip
```

## Use it

`subst` uses key=value mappings given as CLI arguments, reads from stdin, and writes to stdout. The stdin input is line-buffered.

```text

subst [OPTIONS] [KEY=VAL [KEY=VAL...]]

Usage of subst:
  -q    suppress all logs
  -escape string
        set the escape string. if set, a '$' preceded by the escape string does not lead to expansion
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

To ignore occurrences of defined variables, you can define an escape string using the `-escape` flag:

```shell
$ echo 'unescaped: $X, ${X}. escaped: \$X, \${X}.' | subst -escape '\' X=value
unescaped: value, value. escaped: \$X, \${X}.
```

## Comments

Feel free to [leave a comment](https://github.com/sgreben/subst/issues/1) or create an issue.
