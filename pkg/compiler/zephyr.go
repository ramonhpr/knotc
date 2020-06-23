package compiler

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/ramonhpr/knotc/pkg/listener"
	tmplt "github.com/ramonhpr/knotc/pkg/template"
)

// ZephyrCompiler compile it to target zephyr
type ZephyrCompiler struct {
	*BaseKnotCompiler
	listener listener.ListenerImpl
	outPath string
}

// NewZephyrCompiler creates a new instance of zephyr compiler
func NewZephyrCompiler(outPath string) *ZephyrCompiler {
	return &ZephyrCompiler{listener: listener.ListenerImpl{}, outPath: outPath}
}

// Compile the input
func (zc *ZephyrCompiler) Compile(input string) string {
	zc.Walk(input, &zc.listener)
	tmp := template.New("compiler")
	funcs := template.FuncMap{
		"Up": strings.ToUpper,
		"Lower": strings.ToLower,
	}
	tmp.Funcs(funcs)
	sourceCode := new(bytes.Buffer)
	template.Must(tmp.Parse(tmplt.ZephyrSourceCode)).Execute(sourceCode, zc.listener)

	x := struct {
		Name string
		Len int
	}{Name: zc.listener.Name, Len: len(zc.listener.Sensors)}
	prjConf := new(bytes.Buffer)
	template.Must(tmp.Parse(tmplt.ZephyrPrjConf)).Execute(prjConf, x)

	setupConf := new(bytes.Buffer)
	template.Must(tmp.Parse(tmplt.ZephyrSetupConf)).Execute(setupConf, zc.listener)

	cmake := new(bytes.Buffer)
	template.Must(tmp.Parse(tmplt.ZephyrCMakeLists)).Execute(cmake, zc.listener)

	err := zc.createTreeStructureFolder()
	if err != nil {
		panic(err)
	}
	err = zc.outputTopSrcDir("proj.conf", prjConf)
	if err != nil {
		panic(err)
	}
	err = zc.outputTopSrcDir("setup.conf", setupConf)
	if err != nil {
		panic(err)
	}
	err = zc.outputTopSrcDir("CMakeLists.txt", cmake)
	if err != nil {
		panic(err)
	}
	err = zc.outputSrc(sourceCode)
	if err != nil {
		panic(err)
	}

	return sourceCode.String()
}

func (zc *ZephyrCompiler) createTreeStructureFolder() error {
	return os.MkdirAll(filepath.Join(zc.outPath, "zephyr", zc.listener.Name, "src"), os.ModePerm)
}

func (zc *ZephyrCompiler) outputSrc(buf *bytes.Buffer) error {
	return ioutil.WriteFile(filepath.Join(filepath.Join(zc.outPath, "zephyr", zc.listener.Name, "src", zc.listener.Name + ".c")), buf.Bytes(), 0644)
}

func (zc *ZephyrCompiler) outputTopSrcDir(filename string, buf *bytes.Buffer) error {
	return ioutil.WriteFile(filepath.Join(filepath.Join(zc.outPath, "zephyr", zc.listener.Name, filename)), buf.Bytes(), 0644)
}