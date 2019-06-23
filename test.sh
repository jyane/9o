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

assert 4 'ab = 4; return ab;'
assert 42 '42;'
assert 0 'return 0;'
assert 8 ' 11 * 1 - 3 ;'
assert 3 ' 2 * 3 / 2;'
assert 45 '1 + 2 + 3+4+5+6+7+8+9;'
assert 12 '2 * (1 + 5);'
assert 22 '10-(-12);'
assert 1 '+1;'
assert 1 '1 <= 2;'
assert 0 '2 <= 1;'
assert 1 '1 < 2;'
assert 1 '2 > 1;'
assert 1 '1 == 1;'
assert 1 '1 != 2;'
assert 6 'foo = 1; bar = 2; return (foo + bar) * 2;'
assert 1 'foo = 1; if (foo == 1) return 1; return 0;'
assert 0 'foo = 1; if (foo == 2) return 1; else return 0;'
assert 10 'foo = 1; while (foo < 10) foo = foo + 1; return foo;'
assert 3 'foo = 1; while (foo < 3) foo = foo + 1; return foo;'

echo 'Done'
