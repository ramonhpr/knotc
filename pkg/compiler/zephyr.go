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
func (zc *ZephyrCompiler) Compile(input string) {
	zc.Walk(input, &zc.listener)
	tmp := template.New("compiler")
	funcs := template.FuncMap{
		"Up": strings.ToUpper,
		"Lower": strings.ToLower,
		"UpRaw": func (x string) string {
			if x == "char[]" {
				return "RAW"

			} else {
				return strings.ToUpper(x)
			}
		},
	}
	tmp.Funcs(funcs)
	
	for _, thing := range zc.listener.Things {
		sourceCode := new(bytes.Buffer)
		template.Must(tmp.Parse(tmplt.ZephyrSourceCode)).Execute(sourceCode, thing)
		
		x := struct {
			Name string
			Len int
		}{Name: thing.Name, Len: len(thing.Sensors)}
		prjConf := new(bytes.Buffer)
		template.Must(tmp.Parse(tmplt.ZephyrPrjConf)).Execute(prjConf, x)
	
		setupConf := new(bytes.Buffer)
		template.Must(tmp.Parse(tmplt.ZephyrSetupConf)).Execute(setupConf, thing)
	
		cmake := new(bytes.Buffer)
		template.Must(tmp.Parse(tmplt.ZephyrCMakeLists)).Execute(cmake, thing)
	
		err := zc.createTreeStructureFolder(thing.Name)
		if err != nil {
			panic(err)
		}
		err = zc.outputTopSrcDir(thing.Name, "prj.conf", prjConf)
		if err != nil {
			panic(err)
		}
		err = zc.outputTopSrcDir(thing.Name, "setup.conf", setupConf)
		if err != nil {
			panic(err)
		}
		err = zc.outputTopSrcDir(thing.Name, "CMakeLists.txt", cmake)
		if err != nil {
			panic(err)
		}
		err = zc.outputSrc(thing.Name, sourceCode)
		if err != nil {
			panic(err)
		}
	}
}

func (zc *ZephyrCompiler) createTreeStructureFolder(folderName string) error {
	return os.MkdirAll(filepath.Join(zc.outPath, "zephyr", folderName, "src"), os.ModePerm)
}

func (zc *ZephyrCompiler) outputSrc(folderName string, buf *bytes.Buffer) error {
	return ioutil.WriteFile(filepath.Join(filepath.Join(zc.outPath, "zephyr", folderName, "src", folderName + ".c")), buf.Bytes(), 0644)
}

func (zc *ZephyrCompiler) outputTopSrcDir(folderName string, filename string, buf *bytes.Buffer) error {
	return ioutil.WriteFile(filepath.Join(filepath.Join(zc.outPath, "zephyr", folderName, filename)), buf.Bytes(), 0644)
}