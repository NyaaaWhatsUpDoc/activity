// This file is generated - do not edit.

package jen

// standardLibraryHints contains package name hints
var standardLibraryHints = map[string]string{
	"archive/tar":                          "tar",
	"archive/zip":                          "zip",
	"bufio":                                "bufio",
	"bytes":                                "bytes",
	"compress/bzip2":                       "bzip2",
	"compress/flate":                       "flate",
	"compress/gzip":                        "gzip",
	"compress/lzw":                         "lzw",
	"compress/zlib":                        "zlib",
	"constraints":                          "constraints",
	"container/heap":                       "heap",
	"container/list":                       "list",
	"container/ring":                       "ring",
	"context":                              "context",
	"crypto":                               "crypto",
	"crypto/aes":                           "aes",
	"crypto/cipher":                        "cipher",
	"crypto/des":                           "des",
	"crypto/dsa":                           "dsa",
	"crypto/ecdsa":                         "ecdsa",
	"crypto/ed25519":                       "ed25519",
	"crypto/ed25519/internal/edwards25519": "edwards25519",
	"crypto/ed25519/internal/edwards25519/field": "field",
	"crypto/elliptic":                 "elliptic",
	"crypto/elliptic/internal/fiat":   "fiat",
	"crypto/elliptic/internal/nistec": "nistec",
	"crypto/hmac":                     "hmac",
	"crypto/internal/randutil":        "randutil",
	"crypto/internal/subtle":          "subtle",
	"crypto/md5":                      "md5",
	"crypto/rand":                     "rand",
	"crypto/rc4":                      "rc4",
	"crypto/rsa":                      "rsa",
	"crypto/sha1":                     "sha1",
	"crypto/sha256":                   "sha256",
	"crypto/sha512":                   "sha512",
	"crypto/subtle":                   "subtle",
	"crypto/tls":                      "tls",
	"crypto/x509":                     "x509",
	"crypto/x509/internal/macos":      "macOS",
	"crypto/x509/pkix":                "pkix",
	"database/sql":                    "sql",
	"database/sql/driver":             "driver",
	"debug/buildinfo":                 "buildinfo",
	"debug/dwarf":                     "dwarf",
	"debug/elf":                       "elf",
	"debug/gosym":                     "gosym",
	"debug/macho":                     "macho",
	"debug/pe":                        "pe",
	"debug/plan9obj":                  "plan9obj",
	"embed":                           "embed",
	"embed/internal/embedtest":        "embedtest",
	"encoding":                        "encoding",
	"encoding/ascii85":                "ascii85",
	"encoding/asn1":                   "asn1",
	"encoding/base32":                 "base32",
	"encoding/base64":                 "base64",
	"encoding/binary":                 "binary",
	"encoding/csv":                    "csv",
	"encoding/gob":                    "gob",
	"encoding/hex":                    "hex",
	"encoding/json":                   "json",
	"encoding/pem":                    "pem",
	"encoding/xml":                    "xml",
	"errors":                          "errors",
	"expvar":                          "expvar",
	"flag":                            "flag",
	"fmt":                             "fmt",
	"go/ast":                          "ast",
	"go/build":                        "build",
	"go/build/constraint":             "constraint",
	"go/constant":                     "constant",
	"go/doc":                          "doc",
	"go/format":                       "format",
	"go/importer":                     "importer",
	"go/internal/gccgoimporter":       "gccgoimporter",
	"go/internal/gcimporter":          "gcimporter",
	"go/internal/srcimporter":         "srcimporter",
	"go/internal/typeparams":          "typeparams",
	"go/parser":                       "parser",
	"go/printer":                      "printer",
	"go/scanner":                      "scanner",
	"go/token":                        "token",
	"go/types":                        "types",
	"hash":                            "hash",
	"hash/adler32":                    "adler32",
	"hash/crc32":                      "crc32",
	"hash/crc64":                      "crc64",
	"hash/fnv":                        "fnv",
	"hash/maphash":                    "maphash",
	"html":                            "html",
	"html/template":                   "template",
	"image":                           "image",
	"image/color":                     "color",
	"image/color/palette":             "palette",
	"image/draw":                      "draw",
	"image/gif":                       "gif",
	"image/internal/imageutil":        "imageutil",
	"image/jpeg":                      "jpeg",
	"image/png":                       "png",
	"index/suffixarray":               "suffixarray",
	"internal/abi":                    "abi",
	"internal/buildcfg":               "buildcfg",
	"internal/bytealg":                "bytealg",
	"internal/cfg":                    "cfg",
	"internal/cpu":                    "cpu",
	"internal/execabs":                "execabs",
	"internal/fmtsort":                "fmtsort",
	"internal/fuzz":                   "fuzz",
	"internal/goarch":                 "goarch",
	"internal/godebug":                "godebug",
	"internal/goexperiment":           "goexperiment",
	"internal/goos":                   "goos",
	"internal/goroot":                 "goroot",
	"internal/goversion":              "goversion",
	"internal/intern":                 "intern",
	"internal/itoa":                   "itoa",
	"internal/lazyregexp":             "lazyregexp",
	"internal/lazytemplate":           "lazytemplate",
	"internal/nettrace":               "nettrace",
	"internal/obscuretestdata":        "obscuretestdata",
	"internal/oserror":                "oserror",
	"internal/poll":                   "poll",
	"internal/profile":                "profile",
	"internal/race":                   "race",
	"internal/reflectlite":            "reflectlite",
	"internal/singleflight":           "singleflight",
	"internal/syscall/execenv":        "execenv",
	"internal/syscall/unix":           "unix",
	"internal/sysinfo":                "sysinfo",
	"internal/testenv":                "testenv",
	"internal/testlog":                "testlog",
	"internal/trace":                  "trace",
	"internal/unsafeheader":           "unsafeheader",
	"internal/xcoff":                  "xcoff",
	"io":                              "io",
	"io/fs":                           "fs",
	"io/ioutil":                       "ioutil",
	"log":                             "log",
	"log/syslog":                      "syslog",
	"math":                            "math",
	"math/big":                        "big",
	"math/bits":                       "bits",
	"math/cmplx":                      "cmplx",
	"math/rand":                       "rand",
	"mime":                            "mime",
	"mime/multipart":                  "multipart",
	"mime/quotedprintable":            "quotedprintable",
	"net":                             "net",
	"net/http":                        "http",
	"net/http/cgi":                    "cgi",
	"net/http/cookiejar":              "cookiejar",
	"net/http/fcgi":                   "fcgi",
	"net/http/httptest":               "httptest",
	"net/http/httptrace":              "httptrace",
	"net/http/httputil":               "httputil",
	"net/http/internal":               "internal",
	"net/http/internal/ascii":         "ascii",
	"net/http/internal/testcert":      "testcert",
	"net/http/pprof":                  "pprof",
	"net/internal/socktest":           "socktest",
	"net/mail":                        "mail",
	"net/netip":                       "netip",
	"net/rpc":                         "rpc",
	"net/rpc/jsonrpc":                 "jsonrpc",
	"net/smtp":                        "smtp",
	"net/textproto":                   "textproto",
	"net/url":                         "url",
	"os":                              "os",
	"os/exec":                         "exec",
	"os/exec/internal/fdtest":         "fdtest",
	"os/signal":                       "signal",
	"os/signal/internal/pty":          "pty",
	"os/user":                         "user",
	"path":                            "path",
	"path/filepath":                   "filepath",
	"plugin":                          "plugin",
	"reflect":                         "reflect",
	"reflect/internal/example1":       "example1",
	"reflect/internal/example2":       "example2",
	"regexp":                          "regexp",
	"regexp/syntax":                   "syntax",
	"runtime":                         "runtime",
	"runtime/cgo":                     "cgo",
	"runtime/debug":                   "debug",
	"runtime/internal/atomic":         "atomic",
	"runtime/internal/math":           "math",
	"runtime/internal/sys":            "sys",
	"runtime/metrics":                 "metrics",
	"runtime/pprof":                   "pprof",
	"runtime/race":                    "race",
	"runtime/trace":                   "trace",
	"sort":                            "sort",
	"strconv":                         "strconv",
	"strings":                         "strings",
	"sync":                            "sync",
	"sync/atomic":                     "atomic",
	"syscall":                         "syscall",
	"testing":                         "testing",
	"testing/fstest":                  "fstest",
	"testing/internal/testdeps":       "testdeps",
	"testing/iotest":                  "iotest",
	"testing/quick":                   "quick",
	"text/scanner":                    "scanner",
	"text/tabwriter":                  "tabwriter",
	"text/template":                   "template",
	"text/template/parse":             "parse",
	"time":                            "time",
	"time/tzdata":                     "tzdata",
	"unicode":                         "unicode",
	"unicode/utf16":                   "utf16",
	"unicode/utf8":                    "utf8",
	"unsafe":                          "unsafe",
}
