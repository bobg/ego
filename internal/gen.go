// +build ignore

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path"
	"runtime"
	"sort"
	"strings"
	"unicode"
)

var pkgnames = []string{
	"archive/tar",
	"archive/zip",
	"bufio",
	"builtin",
	"bytes",
	"compress/bzip2",
	"compress/flate",
	"compress/gzip",
	"compress/lzw",
	"compress/zlib",
	"container/heap",
	"container/list",
	"container/ring",
	"context",
	"crypto",
	"crypto/aes",
	"crypto/cipher",
	"crypto/des",
	"crypto/dsa",
	"crypto/ecdsa",
	"crypto/elliptic",
	"crypto/hmac",
	"crypto/md5",
	"crypto/rand",
	"crypto/rc4",
	"crypto/rsa",
	"crypto/sha1",
	"crypto/sha256",
	"crypto/sha512",
	"crypto/subtle",
	"crypto/tls",
	"crypto/x509",
	"crypto/x509/pkix",
	"database/sql",
	"database/sql/driver",
	"debug/dwarf",
	"debug/elf",
	"debug/gosym",
	"debug/macho",
	"debug/pe",
	"debug/plan9obj",
	"encoding",
	"encoding/ascii85",
	"encoding/asn1",
	"encoding/base32",
	"encoding/base64",
	"encoding/binary",
	"encoding/csv",
	"encoding/gob",
	"encoding/hex",
	"encoding/json",
	"encoding/pem",
	"encoding/xml",
	"errors",
	"expvar",
	"flag",
	"fmt",
	"go/ast",
	"go/build",
	"go/build/testdata/multi",
	"go/build/testdata/other",
	"go/build/testdata/other/file",
	"go/constant",
	"go/doc",
	"go/doc/testdata",
	"go/format",
	"go/importer",
	"go/parser",
	"go/printer",
	"go/printer/testdata",
	"go/scanner",
	"go/token",
	"go/types",
	"hash",
	"hash/adler32",
	"hash/crc32",
	"hash/crc64",
	"hash/fnv",
	"html",
	"html/template",
	"image",
	"image/color",
	"image/color/palette",
	"image/draw",
	"image/gif",
	"image/jpeg",
	"image/png",
	"index/suffixarray",
	"io",
	"io/ioutil",
	"log",
	"log/syslog",
	"math",
	"math/big",
	"math/bits",
	"math/cmplx",
	"math/rand",
	"mime",
	"mime/multipart",
	"mime/quotedprintable",
	"net",
	"net/http",
	"net/http/cgi",
	"net/http/cookiejar",
	"net/http/fcgi",
	"net/http/httptest",
	"net/http/httptrace",
	"net/http/httputil",
	"net/http/pprof",
	"net/mail",
	"net/rpc",
	"net/rpc/jsonrpc",
	"net/smtp",
	"net/textproto",
	"net/url",
	"os",
	"os/exec",
	"os/signal",
	"os/user",
	"path",
	"path/filepath",
	"plugin",
	"reflect",
	"regexp",
	"regexp/syntax",
	"runtime",
	"runtime/cgo",
	"runtime/debug",
	"runtime/msan",
	"runtime/pprof",
	"runtime/race",
	"runtime/race/testdata",
	"runtime/testdata/testprog",
	"runtime/testdata/testprogcgo",
	"runtime/testdata/testprogcgo/windows",
	"runtime/testdata/testprognet",
	"runtime/trace",
	"sort",
	"strconv",
	"strings",
	"sync",
	"sync/atomic",
	"syscall",
	"testing",
	"testing/iotest",
	"testing/quick",
	"text/scanner",
	"text/tabwriter",
	"text/template",
	"text/template/parse",
	"time",
	"unicode",
	"unicode/utf16",
	"unicode/utf8",
	"unsafe",
}

func notests(info os.FileInfo) bool {
	return !strings.HasSuffix(info.Name(), "_test.go")
}

func main() {
	for _, pkgname := range pkgnames {
		log.Printf("Parsing %s", pkgname)

		inPkgFullDir := path.Join(runtime.GOROOT(), "src", pkgname)
		fset := token.NewFileSet()

		pkgs, err := parser.ParseDir(fset, inPkgFullDir, notests, 0)
		must(err)

		pkgBaseName := path.Base(pkgname)

		for parsedPkgname, pkg := range pkgs {
			if parsedPkgname != pkgBaseName {
				continue
			}

			outPkgFullDir := path.Join("standard", pkgname) // xxx
			err = os.MkdirAll(outPkgFullDir, 0755)
			must(err)

			outFileName := path.Join(outPkgFullDir, "stubs.go")
			out, err := os.Create(outFileName)
			must(err)
			func() {
				defer out.Close()

				fmt.Fprintf(out, "package %s\n\n", pkgBaseName)
				fmt.Fprintf(out, "import (\n")
				imports := []string{"reflect", pkgname}
				sort.Strings(imports)
				for _, imp := range imports {
					fmt.Fprintf(out, "\t\"%s\"\n", imp)
				}
				fmt.Fprintf(out, ")\n\n")

				for _, f := range pkg.Files {

					for objname, obj := range f.Scope.Objects {
						if !unicode.IsUpper(rune(objname[0])) {
							continue
						}
						if obj.Kind != ast.Fun {
							continue
						}
						decl, ok := obj.Decl.(*ast.FuncDecl)
						if !ok {
							fmt.Printf("* %s is a %T, want *ast.FuncDecl\n", objname, obj.Decl)
							continue // xxx error?
						}
						if decl.Recv != nil {
							fmt.Printf("* skipping %s, which is a method\n", objname)
							continue // xxx not implemented
						}
						fmt.Fprintf(out, "var %s = reflect.ValueOf(%s.%s)\n", objname, pkgBaseName, objname)
					}
				}
			}()
		}
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
