package tree

import (
    "strings"
    b64 "encoding/base64"
)

func Base64FSCompat(b64str string) string {
    return strings.Replace(
                strings.Replace(
                    strings.Replace(b64str, "/", "-", -1), "=", "_", -1), "+", ".", -1)
}

func FSCompatToB64(FSstr string) string {
    return strings.Replace(
                strings.Replace(
                    strings.Replace(FSstr, "-", "/", -1), "_", "=", -1), ".", "+", -1)
}

func B64Enc(data []byte) string {
    return b64.StdEncoding.EncodeToString(data)
}
