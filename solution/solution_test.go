package solution

import (
    "testing"
)


func TestEmoji(t *testing.T) {
    msg := GetMessage()
    if msg != "Hello 🗺️ !"  {
        t.Fatalf(`GetMessage() = %q, want "Hello 🗺️ !"`, msg)
    }
}