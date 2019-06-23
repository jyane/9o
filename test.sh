#!/bin/bash -x

assert() {
    expected="$1"
    input="$2"
    go run *.go "$input" > out.s
    cat out.s
    gcc -o out out.s
    ./out
    actual="$?"
    if [ "$actual" != "$expected" ]; then
        echo "$input => $actual != $expected"
        exit 1
    fi
}

assert 42 'test/000.c'
assert 4  'test/001.c'
assert 0  'test/002.c'
assert 8  'test/003.c'
assert 3  'test/004.c'
assert 45 'test/005.c'
assert 12 'test/006.c'
assert 22 'test/007.c'
assert 1  'test/008.c'
assert 1  'test/009.c'
assert 0  'test/010.c'
assert 1  'test/011.c'
assert 1  'test/012.c'
assert 1  'test/013.c'
assert 1  'test/014.c'
assert 6  'test/015.c'
assert 1  'test/016.c'
assert 0  'test/017.c'
assert 10 'test/018.c'
assert 3  'test/019.c'
assert 11 'test/020.c'

echo 'Done'
