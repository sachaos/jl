# jl (JSON by line)

Convert to [JSON Lines](https://jsonlines.org/) from STDOUT.

The main purpose is pre-process output from UNIX command to query by jq.

## Example

```
$ ps aux | jl --header | jq -r ".USER" | sort | uniq
...
root
sachaos
```

## Install

```shell
$ git clone https://github.com/sachaos/jl.git
$ cd jl
$ go install
```
