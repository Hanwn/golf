package golf

import "testing"

func TestShow(t *testing.T) {
    ret := Show()
    if ret != "OK" {
        t.Fatal("error")
    }
}
