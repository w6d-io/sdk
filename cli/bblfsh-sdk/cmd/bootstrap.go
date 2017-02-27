package cmd

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"fmt"

	"github.com/bblfsh/sdk/assets/skeleton"
	"github.com/bblfsh/sdk/manifest"
)

const (
	tplExtension = ".tpl"
	manifestTpl  = "manifest.toml.tpl"
)

// managedFiles are files that always are overwritten
var managedFiles = map[string]bool{
	".travis.yml":        true,
	"README.md.tpl":      true,
	"LICENSE":            true,
	"driver/main.go.tpl": true,
}

type BootstrapCommand struct {
	Init   bool `long:"init" description:"creates an initial manifest.toml, based on the given lang"`
	DryRun bool `long:"dry-run" description:"don't writes nothing just checks if something should be written"`
	Args   struct {
		Language string `positional-arg-name:"language"  description:"target langunge of the driver"`
		OS       string `positional-arg-name:"os" description:"distribution used to run the runtime. (Values: alpine or debian)"`
	} `positional-args:"yes"`

	changes int
	context map[string]interface{}
	command
}

func (c *BootstrapCommand) Execute(args []string) error {
	if c.Init && c.DryRun {
		return fmt.Errorf("can't use --init and --dry-run at the same time")
	}

	if err := c.processManifest(); err != nil {
		return err
	}

	m, err := c.readManifest()
	if err != nil {
		return err
	}

	c.context = map[string]interface{}{
		"Manifest": m,
	}

	for _, file := range skeleton.AssetNames() {
		if file == manifestTpl {
			continue
		}

		if err := c.processAsset(file); err != nil {
			return err
		}
	}

	if c.DryRun && c.changes > 0 {
		return fmt.Errorf("changes are required")
	}

	return nil
}

func (c *BootstrapCommand) processManifest() error {
	if !c.Init {
		return nil
	}

	if c.Args.Language == "" || c.Args.OS == "" {
		return fmt.Errorf("`language` and `os` arguments are mandatory in combination with --init")
	}

	notice.Printf("initializing driver %q, creating new manifest\n", c.Args.Language)
	if _, err := c.readManifest(); err == nil {
		warning.Printf("driver already initialized. %q detected\n", manifest.Filename)
	}

	return c.processTemplateAsset(manifestTpl, c.Args, false)
}

func (c *BootstrapCommand) processAsset(name string) error {
	overwrite := managedFiles[name]

	if strings.HasSuffix(name, tplExtension) {
		return c.processTemplateAsset(name, c.context, overwrite)
	}

	return c.processFileAsset(name, overwrite)
}

func (c *BootstrapCommand) processFileAsset(name string, overwrite bool) error {
	content := skeleton.MustAsset(name)
	info, _ := skeleton.AssetInfo(name)

	name = fixGitFolder(name)
	return c.writeFile(filepath.Join(c.Root, name), content, info.Mode(), overwrite)
}

//	rel, _ := filepath.Rel(c.Root, file)

var funcs = map[string]interface{}{
	"escape_shield": escapeShield,
}

func (c *BootstrapCommand) processTemplateAsset(name string, v interface{}, overwrite bool) error {
	tpl := string(skeleton.MustAsset(name))

	t, err := template.New(name).Funcs(funcs).Parse(tpl)
	if err != nil {
		return err
	}

	name = fixGitFolder(name)
	file := filepath.Join(c.Root, name[:len(name)-len(tplExtension)])

	buf := bytes.NewBuffer(nil)
	if err := t.Execute(buf, v); err != nil {
		return err
	}

	info, _ := skeleton.AssetInfo(name)
	return c.writeFile(file, buf.Bytes(), info.Mode(), overwrite)
}

func (c *BootstrapCommand) writeFile(file string, content []byte, m os.FileMode, overwrite bool) error {
	f, err := os.Open(file)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if f == nil {
		c.notifyMissingFile(file)
		return c.doWriteFile(file, content, m)
	}

	if !overwrite {
		return nil
	}

	original, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	if bytes.Compare(original, content) == 0 {
		return nil
	}

	c.notifyChangedFile(file)
	return c.doWriteFile(file, content, m)
}

func (c *BootstrapCommand) doWriteFile(file string, content []byte, m os.FileMode) error {
	if c.DryRun {
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(file), 0755); err != nil {
		return err
	}

	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, m)
	if err != nil {
		return err
	}

	defer f.Close()

	if c.Verbose {
		debug.Printf("file %q has be written\n", file)
	}

	_, err = f.Write(content)
	return err
}

func (c *BootstrapCommand) readManifest() (*manifest.Manifest, error) {
	return manifest.Load(filepath.Join(c.Root, manifest.Filename))
}

func (c *BootstrapCommand) notifyMissingFile(file string) {
	c.changes++

	if !c.DryRun {
		notice.Printf("creating file %q\n", file)
		return
	}

	warning.Printf("missing file %q\n", file)
}

func (c *BootstrapCommand) notifyChangedFile(file string) {
	c.changes++

	if !c.DryRun {
		warning.Printf("managed file %q has changed, discarding changes\n", file)
		return
	}

	warning.Printf("managed file changed %q\n", file)
}

func escapeShield(text interface{}) string {
	return strings.Replace(fmt.Sprintf("%s", text), "-", "--", -1)
}

func fixGitFolder(path string) string {
	return strings.Replace(path, "git/", ".git/", 1)
}
