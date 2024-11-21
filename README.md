# diff-color

Quick and easy string diffing function for Golang. Mainly for diffing strings in tests.

Get it:

    go get github.com/KoNekoD/diff-color

Example:

    import (
        "strings"
        "testing"
        "github.com/KoNekoD/diff"
    )

    const expected = `
    ...
    `

    func TestFoo(t *testing.T) {
        actual := Foo(...)
        if a, e := strings.TrimSpace(actual), strings.TrimSpace(expected); a != e {
            t.Errorf("Result not as expected:\n%v", diff.LineDiff(e, a))
        }
    }


Motivated by https://github.com/andreyvit/diff
