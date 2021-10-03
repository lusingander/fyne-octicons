package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

const (
	srcFile = "./bundle.go"
	dstFile = "./icons.go"
)

// Go reserved keywords
var keywords = map[string]bool{
	"break":       true,
	"default":     true,
	"func":        true,
	"interface":   true,
	"select":      true,
	"case":        true,
	"defer":       true,
	"go":          true,
	"map":         true,
	"struct":      true,
	"chan":        true,
	"else":        true,
	"goto":        true,
	"package":     true,
	"switch":      true,
	"const":       true,
	"fallthrough": true,
	"if":          true,
	"range":       true,
	"type":        true,
	"continue":    true,
	"for":         true,
	"import":      true,
	"return":      true,
	"var":         true,
}

func cutResourceName(s string) string {
	s = strings.TrimLeft(s, "resource")
	s = strings.TrimRight(s, "Svg")
	s = strings.ToLower(s)
	if _, exist := keywords[s]; exist {
		s += "_"
	}
	return s
}

func createGetterFuncName(s string) (string, string) {
	if _, exist := keywords[s[:len(s)-1]]; exist {
		s = s[:len(s)-1]
	}
	s = strings.Title(s)
	return s + "Icon", s
}

func generateIconsFile(vars []string) error {
	names := make([]string, 0, len(vars))
	for _, v := range vars {
		names = append(names, cutResourceName(v))
	}
	buf := newBufferWrapper()
	buf.writeln("// AUTO-GENERATED: DO NOT EDIT")
	buf.writeln("")
	buf.writeln("package octicons")
	buf.writeln("")
	buf.writeln("import (")
	buf.writeln("\"fyne.io/fyne/v2\"")
	buf.writeln("\"fyne.io/fyne/v2/theme\"")
	buf.writeln(")")
	buf.writeln("var (")
	for _, n := range names {
		buf.writeln("%s *theme.ThemedResource", n)
	}
	buf.writeln(")")
	buf.writeln("func init() {")
	for i, n := range names {
		buf.writeln("%s = theme.NewThemedResource(%s)", n, vars[i])
	}
	buf.writeln("}")
	for _, n := range names {
		f, s := createGetterFuncName(n)
		buf.writeln("// %s returns %s icon resource", f, s)
		buf.writeln("func %s() fyne.Resource {", f)
		buf.writeln("return %s", n)
		buf.writeln("}")
	}
	source, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	dst, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer dst.Close()
	dst.Write(source)
	return nil
}

func listVariables() ([]string, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, srcFile, nil, 0)
	if err != nil {
		return nil, err
	}
	variables := make([]string, 0)
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.ValueSpec:
			name := x.Names[0].Name
			variables = append(variables, name)
		}
		return true
	})
	return variables, nil
}

type bufferWrapper struct {
	*bytes.Buffer
}

func newBufferWrapper() *bufferWrapper {
	return &bufferWrapper{&bytes.Buffer{}}
}

func (b *bufferWrapper) writeln(s string, a ...interface{}) {
	b.WriteString(fmt.Sprintf(s+"\n", a...))
}

func run(args []string) error {
	vars, err := listVariables()
	if err != nil {
		return err
	}
	err = generateIconsFile(vars)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run(os.Args); err != nil {
		panic(err)
	}
}
