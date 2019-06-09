package framework

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Vue holds the end result
// code and styles in order of execution / interpretation
type Vue struct {
	Style template.CSS
	Code  template.JS
}

//StyleString is a helper function when the go template engine is not beign used
func (v *Vue) StyleString() string {
	return StyleStart + string(v.Style) + StyleEnd
}

//CodeString is a helper function when the go template engine is not beign used
func (v *Vue) CodeString() string {
	return ScriptStart + string(v.Code) + ScriptEnd
}

// component holds the filename as well as the styles, templates and code
// of all files we process.
// which are Single File Components, scripts like Data Store, Router
// and the actual App
// important convention app must be called app.js
// scripts must be called *.js and components must not be *.js files
// but can have every file extension you want.
type component struct {
	FileName string
	Style    string
	Script   string
	Template string
}

// Compile walks through all directories containing single file components and bring them
// into a form that doesn't need Vue Loader and can be used just with the In Browser Vue Template Compiler
// e.g template: `<p>{{msg}}</p>` instead of <template><p>{{msg}}</p></template>
// on top of that it bundles all plain javascript files for example router and store together
// with the transformed Components into a single Javascript file
func Compile(dirNames ...string) (v *Vue, err error) {
	var (
		components = make([]*component, 0)
		cs         = make([]*component, 0)
		bc         = make([]string, 0)
		bs         = make([]string, 0)
		pre        = make([]string, 0)
		app        = ""
	)
	//process all directories
	for _, dirName := range dirNames {
		if cs, err = parseDirectory(dirName); err != nil {
			return nil, err
		}
		components = append(cs)
	}
	// process all components
	for _, c := range components {
		// insert style
		bs = append(bs, c.Style)
		if strings.HasSuffix(c.FileName, ".js") {
			// app must come after everything is declared
			// so we put it into app variable
			if strings.HasSuffix(c.FileName, "app.js") {
				app = c.Script
				continue
			}
			// scripts must come after components
			// and before the actual app so we capture it
			// into pre
			pre = append(pre, c.Script)
			continue
		}
		// if there is no template we just take the code as is
		if strings.TrimSpace(c.Template) == "" {
			bc = append(bc, c.Script)
			continue
		}
		// if it has a template
		if strings.Contains(c.Script, "//>>") {
			// inject the template into the code
			tpl := fmt.Sprintf("template: `%s`,", c.Template)
			bc = append(bc, strings.Replace(c.Script, "//>>", tpl, 1))
			continue
		}
		// if it has template but no magic comment throw an error
		return nil, fmt.Errorf("No Magic Comment %s in %s", "//>>", c.FileName)
	}
	// bc hold all the components code, then add the scripts
	bc = append(bc, pre...)
	// and finally the app that bootstraps the whole process
	bc = append(bc, app)
	v = new(Vue)
	// wrap it into a typed javascript string
	v.Code = template.JS(strings.Join(bc, "\n"))
	// wrap it into a typed CSS string
	v.Style = template.CSS(strings.Join(bs, "\n"))
	return v, nil
}
func parseDirectory(dirName string) (cs []*component, err error) {
	err = filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
		var c *component
		// dont go into subdirectories
		if info.IsDir() {
			return nil
		}
		// don't process hidden files
		// TODO this looks wonky
		// i probably checked only the root directory of this path
		if strings.HasPrefix(info.Name(), ".") {
			return nil
		}
		// parse the file as a component or javascript file
		if c, err = parseComponent(path); err != nil {
			return err
		}
		// collect the components into an array of pointers
		cs = append(cs, c)
		return nil
	})
	return cs, err
}
func parseComponent(fileName string) (c *component, err error) {
	c = new(component)
	c.FileName = fileName
	var buf []byte
	if buf, err = ioutil.ReadFile(fileName); err != nil {
		return nil, fmt.Errorf("parseComponent-File: %s\n%s", c.FileName, err.Error())
	}
	// if its plain javascript we just take the sourcecode
	// and save it in the script part of the component
	if strings.HasSuffix(c.FileName, ".js") {
		c.Script = string(buf)
		return c, nil
	}
	s := string(buf)
	m, err := getIndexes(s)
	if err != nil {
		return nil, fmt.Errorf("parseComponent-getIndexes: %s\n%s", c.FileName, err.Error())
	}
	// get a substring between the style tags
	v1, ok := m[StyleStart]
	if ok {
		c.Style = s[v1+len(StyleStart) : m[StyleEnd]]
	}
	// get a substring between the template tags
	v1, ok = m[TemplateStart]
	if ok {
		c.Template = s[v1+len(TemplateStart) : m[TemplateEnd]]
	}
	// get a substring between the script tags
	v1, ok = m[ScriptStart]
	if ok {
		c.Script = s[v1+len(ScriptStart) : m[ScriptEnd]]
	}
	return c, err
}

// define the start and end of the tags we want to process
const (
	TemplateStart = "<template>"
	TemplateEnd   = "</template>"
	ScriptStart   = "<script>"
	ScriptEnd     = "</script>"
	StyleStart    = "<style>"
	StyleEnd      = "</style>"
)

func getIndexes(s string) (m map[string]int, err error) {
	m = make(map[string]int)
	// get the first and last index of the start and end tags
	// first and last so we dont accidentally capture other
	// occurences of our tag delimiters
	m[TemplateStart] = strings.Index(s, TemplateStart)
	m[TemplateEnd] = strings.LastIndex(s, TemplateEnd)
	m[ScriptStart] = strings.Index(s, ScriptStart)
	m[ScriptEnd] = strings.LastIndex(s, ScriptEnd)
	m[StyleStart] = strings.Index(s, StyleStart)
	m[StyleEnd] = strings.LastIndex(s, StyleEnd)
	// if the start and end of the template are the same
	// it means we got -1 aka we have none
	// so delete the indices from the map
	if m[TemplateStart] == m[TemplateEnd] {
		delete(m, TemplateStart)
		delete(m, TemplateEnd)
		goto checkScriptTags
	}
	// if one of both is -1 it means one of the tags is missing
	// or if the end comes before the start it means the order of
	// the tags is bollocks so we throw an error in any of these cases
	// same goes with all the other tags we process down below
	if (m[TemplateStart] == -1 || m[TemplateEnd] == -1) || (m[TemplateEnd] < m[TemplateStart]) {
		return nil, fmt.Errorf("Template Tag not well formed")
	}

checkScriptTags:
	if m[ScriptStart] == m[ScriptEnd] {
		delete(m, ScriptStart)
		delete(m, ScriptEnd)
		goto checkStyleTags
	}
	if (m[ScriptStart] == -1 || m[ScriptEnd] == -1) || (m[ScriptEnd] < m[ScriptStart]) {
		return nil, fmt.Errorf("Script Tag not well formed")
	}
checkStyleTags:
	if m[StyleStart] == m[StyleEnd] {
		delete(m, StyleStart)
		delete(m, StyleEnd)
		goto result
	}
	if (m[StyleStart] == -1 || m[StyleEnd] == -1) || (m[StyleEnd] < m[StyleStart]) {
		return nil, fmt.Errorf("Style Tag not well formed")
	}

result:
	return m, nil
}
