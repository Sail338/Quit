package tree

import (
    "strings"
)

func Base64FSCompat(b64str string) string {
    return strings.Replace(strings.Replace(b64str, "/", "-", -1), "=", "_", -1)
}
