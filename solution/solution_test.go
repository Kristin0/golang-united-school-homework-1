package solution

import (
    "fmt"

    "testing"
)


func TestEmoji(t *testing.T) {
    message := GetMessage()
    if message == "Hello 🗺️!" {
        fmt.Print("Test passed")
    }
}