#!/bin/bash

assert() {
    expected="$1"
    input="$2"
    go run main.go $2 > out.s
    gcc -o out out.s
    ./out
    actual="$?"
    if [ "$actual" != "$expected" ]; then
        echo "$input := $actual != $expected"
        exit 1
    fi
}

assert 42 42
assert 0 0

echo 'Done'