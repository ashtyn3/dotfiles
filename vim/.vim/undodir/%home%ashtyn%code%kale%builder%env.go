Vim�UnDo� �\�e�kLT�(����&5g�J�y��S'����,  �                  	       	   	   	    `C�    _�                             ����                                                                                                                                                                                                                                                                                                                                                             `C��     �                 �              5�_�                          ����                                                                                                                                                                                                                                                                                                                                                  V        `C�     �  �              	// Unless the progr�  �            �                // Note that s�              �   }              "	err := runAndLogCommand(cmd, e.ve�   ~            �                  �               �                  package builder5�_�      	             �        ����                                                                                                                                                                                                                                                                                                                                                  V        `C�     �        �       5�_�                  	  �        ����                                                                                                                                                                                                                                                                                                                                                  V        `C�    �              �   package builder       import (   	"bytes"   		"errors"   	"flag"   	"fmt"   	"io"   	"io/ioutil"   	"log"   	"os"   
	"os/exec"   	"path/filepath"   
	"runtime"   
	"strconv"   
	"strings"   )       var (   :	// cgoEnvVars is the list of all cgo environment variable   S	cgoEnvVars = []string{"CGO_CFLAGS", "CGO_CXXFLAGS", "CGO_CPPFLAGS", "CGO_LDFLAGS"}   J	// cgoAbsEnvFlags are all the flags that need absolute path in cgoEnvVars   u	cgoAbsEnvFlags = []string{"-I", "-L", "-isysroot", "-isystem", "-iquote", "-include", "-gcc-toolchain", "--sysroot"}   )       G// env holds a small amount of Go environment and toolchain information   N// which is common to multiple builders. Most Bazel-agnostic build information   +// is collected in go/build.Default though.   //   E// See ./README.rst for more information about handling arguments and   // environment variables.   type env struct {   D	// sdk is the path to the Go SDK, which contains tools for the host   0	// platform. This may be different than GOROOT.   	sdk string       M	// installSuffix is the name of the directory below GOROOT/pkg that contains   B	// the .a files for the standard library we should build against.   "	// For example, linux_amd64_race.   	installSuffix string       I	// verbose indicates whether subprocess command lines should be printed.   	verbose bool       D	// workDirPath is a temporary work directory. It is created lazily.   	workDirPath string       	shouldPreserveWorkDir bool   }       J// envFlags registers flags common to multiple builders and returns an env   // configured with those flags.   )func envFlags(flags *flag.FlagSet) *env {   	env := &env{}   <	flags.StringVar(&env.sdk, "sdk", "", "Path to the Go SDK.")   E	flags.Var(&tagFlag{}, "tags", "List of build tags considered true.")   ^	flags.StringVar(&env.installSuffix, "installsuffix", "", "Standard library under GOROOT/pkg")   ^	flags.BoolVar(&env.verbose, "v", false, "Whether subprocess command lines should be printed")   t	flags.BoolVar(&env.shouldPreserveWorkDir, "work", false, "if true, the temporary work directory will be preserved")   	return env   }       K// checkFlags checks whether env flags were set to valid values. checkFlags   (// should be called after parsing flags.   "func (e *env) checkFlags() error {   	if e.sdk == "" {   '		return errors.New("-sdk was not set")   	}   	return nil   }       K// workDir returns a path to a temporary work directory. The same directory   H// is returned on multiple calls. The caller is responsible for cleaning   ,// up the work directory by calling cleanup.   Bfunc (e *env) workDir() (path string, cleanup func(), err error) {   	if e.workDirPath != "" {   &		return e.workDirPath, func() {}, nil   	}   K	// Keep the stem "rules_go_work" in sync with reproducible_binary_test.go.   :	e.workDirPath, err = ioutil.TempDir("", "rules_go_work-")   	if err != nil {   		return "", func() {}, err   	}   	if e.verbose {   (		log.Printf("WORK=%s\n", e.workDirPath)   	}   	if e.shouldPreserveWorkDir {   		cleanup = func() {}   		} else {   2		cleanup = func() { os.RemoveAll(e.workDirPath) }   	}   #	return e.workDirPath, cleanup, nil   }       A// goTool returns a slice containing the path to an executable at   <// $GOROOT/pkg/$GOOS_$GOARCH/$tool and additional arguments.   <func (e *env) goTool(tool string, args ...string) []string {   ?	platform := fmt.Sprintf("%s_%s", runtime.GOOS, runtime.GOARCH)   @	toolPath := filepath.Join(e.sdk, "pkg", "tool", platform, tool)   	if runtime.GOOS == "windows" {   		toolPath += ".exe"   	}   +	return append([]string{toolPath}, args...)   }       A// goCmd returns a slice containing the path to the go executable   // and additional arguments.   :func (e *env) goCmd(cmd string, args ...string) []string {   )	exe := filepath.Join(e.sdk, "bin", "go")   	if runtime.GOOS == "windows" {   		exe += ".exe"   	}   +	return append([]string{exe, cmd}, args...)   }       I// runCommand executes a subprocess that inherits stdout, stderr, and the   !// environment from this process.   /func (e *env) runCommand(args []string) error {   *	cmd := exec.Command(args[0], args[1:]...)   J	// Redirecting stdout to stderr. This mirrors behavior in the go command:   ]	// https://go.googlesource.com/go/+/refs/tags/go1.15.2/src/cmd/go/internal/work/exec.go#1958   	buf := &bytes.Buffer{}   	cmd.Stdout = buf   	cmd.Stderr = buf   (	err := runAndLogCommand(cmd, e.verbose)   .	os.Stderr.Write(relativizePaths(buf.Bytes()))   	return err   }       L// runCommandToFile executes a subprocess and writes the output to the given   
// writer.   Bfunc (e *env) runCommandToFile(w io.Writer, args []string) error {   *	cmd := exec.Command(args[0], args[1:]...)   	cmd.Stdout = w   	cmd.Stderr = os.Stderr   (	return runAndLogCommand(cmd, e.verbose)   }       ;func absEnv(envNameList []string, argList []string) error {   &	for _, envName := range envNameList {   2		splitedEnv := strings.Fields(os.Getenv(envName))   		absArgs(splitedEnv, argList)   K		if err := os.Setenv(envName, strings.Join(splitedEnv, " ")); err != nil {   			return err   		}   	}   	return nil   }       :func runAndLogCommand(cmd *exec.Cmd, verbose bool) error {   	if verbose {   -		fmt.Fprintln(os.Stderr, formatCommand(cmd))   	}   ,	cleanup := passLongArgsInResponseFiles(cmd)   	defer cleanup()   "	if err := cmd.Run(); err != nil {   E		return fmt.Errorf("error running subcommand %s: %v", cmd.Path, err)   	}   	return nil   }       <// expandParamsFiles looks for arguments in args of the form   P// "-param=filename". When it finds these arguments it reads the file "filename"   .// and replaces the argument with its content.   9func expandParamsFiles(args []string) ([]string, error) {   	var paramsIndices []int   	for i, arg := range args {   (		if strings.HasPrefix(arg, "-param=") {   +			paramsIndices = append(paramsIndices, i)   		}   	}   	if len(paramsIndices) == 0 {   		return args, nil   	}   	var expandedArgs []string   
	last := 0   #	for _, pi := range paramsIndices {   7		expandedArgs = append(expandedArgs, args[last:pi]...)   		last = pi + 1       '		fileName := args[pi][len("-param="):]   +		fileArgs, err := readParamsFile(fileName)   		if err != nil {   			return nil, err   		}   2		expandedArgs = append(expandedArgs, fileArgs...)   	}   4	expandedArgs = append(expandedArgs, args[last:]...)   	return expandedArgs, nil   }       I// readParamsFiles parses a Bazel params file in "shell" format. The file   L// should contain one argument per line. Arguments may be quoted with single   I// quotes. All characters within quoted strings are interpreted literally   L// including newlines and excepting single quotes. Characters outside quoted   +// strings may be escaped with a backslash.   4func readParamsFile(name string) ([]string, error) {   #	data, err := ioutil.ReadFile(name)   	if err != nil {   		return nil, err   	}       	var args []string   	var arg []byte   	quote := false   	escape := false   !	for p := 0; p < len(data); p++ {   		b := data[p]   
		switch {   		case escape:   			arg = append(arg, b)   			escape = false       		case b == '\'':   			quote = !quote       		case !quote && b == '\\':   			escape = true       		case !quote && b == '\n':   #			args = append(args, string(arg))   			arg = arg[:0]       
		default:   			arg = append(arg, b)   		}   	}   	if quote {   .		return nil, fmt.Errorf("unterminated quote")   	}   	if escape {   /		return nil, fmt.Errorf("unterminated escape")   	}   	if len(arg) > 0 {   "		args = append(args, string(arg))   	}   	return args, nil   }       S// writeParamsFile formats a list of arguments in Bazel's "shell" format and writes   // it to a file.   8func writeParamsFile(path string, args []string) error {   	buf := new(bytes.Buffer)   	for _, arg := range args {   )		if !strings.ContainsAny(arg, "'\n\\") {   			fmt.Fprintln(buf, arg)   			continue   		}   		buf.WriteByte('\'')   		for _, r := range arg {   			if r == '\'' {   				buf.WriteString(`'\''`)   			} else {   				buf.WriteRune(r)   			}   		}   		buf.WriteString("'\n")   	}   1	return ioutil.WriteFile(path, buf.Bytes(), 0666)   }       N// splitArgs splits a list of command line arguments into two parts: arguments   I// that should be interpreted by the builder (before "--"), and arguments   E// that should be passed through to the underlying tool (after "--").   Ifunc splitArgs(args []string) (builderArgs []string, toolArgs []string) {   	for i, arg := range args {   		if arg == "--" {   			return args[:i], args[i+1:]   		}   	}   	return args, nil   }       K// abs returns the absolute representation of path. Some tools/APIs require   K// absolute paths to work correctly. Most notably, golang on Windows cannot   L// handle relative paths to files whose absolute path is > ~250 chars, while   :// it can handle absolute paths. See http://goo.gl/eqeWjm.   //   N// Note that strings that begin with "__BAZEL_" are not absolutized. These are   G// used on macOS for paths that the compiler wrapper (wrapped_clang) is   // supposed to know about.   func abs(path string) string {   )	if strings.HasPrefix(path, "__BAZEL_") {   		return path   	}       0	if abs, err := filepath.Abs(path); err != nil {   		return path   		} else {   		return abs   	}   }       J// absArgs applies abs to strings that appear in args. Only paths that are   /// part of options named by flags are modified.   -func absArgs(args []string, flags []string) {   	absNext := false   	for i := range args {   		if absNext {   			args[i] = abs(args[i])   			absNext = false   			continue   		}   		for _, f := range flags {   &			if !strings.HasPrefix(args[i], f) {   				continue   			}   $			possibleValue := args[i][len(f):]   			if len(possibleValue) == 0 {   				absNext = true   					break   			}   			separator := ""   			if possibleValue[0] == '=' {   %				possibleValue = possibleValue[1:]   				separator = "="   			}   D			args[i] = fmt.Sprintf("%s%s%s", f, separator, abs(possibleValue))   			break   		}   	}   }       N// relativizePaths converts absolute paths found in the given output string to   6// relative, if they are within the working directory.   ,func relativizePaths(output []byte) []byte {   	dir, err := os.Getwd()   	if dir == "" || err != nil {   		return output   	}   /	dirBytes := make([]byte, len(dir), len(dir)+1)   	copy(dirBytes, dir)   ;	if bytes.HasSuffix(dirBytes, []byte{filepath.Separator}) {   0		return bytes.ReplaceAll(output, dirBytes, nil)   	}       	// This is the common case.   /	// Replace "$CWD/" with "" and "$CWD" with "."   0	dirBytes = append(dirBytes, filepath.Separator)   1	output = bytes.ReplaceAll(output, dirBytes, nil)   &	dirBytes = dirBytes[:len(dirBytes)-1]   7	return bytes.ReplaceAll(output, dirBytes, []byte{'.'})   }       I// formatCommand formats cmd as a string that can be pasted into a shell.   G// Spaces in environment variables and arguments are escaped as needed.   *func formatCommand(cmd *exec.Cmd) string {   )	quoteIfNeeded := func(s string) string {   $		if strings.IndexByte(s, ' ') < 0 {   			return s   		}   		return strconv.Quote(s)   	}   ,	quoteEnvIfNeeded := func(s string) string {   !		eq := strings.IndexByte(s, '=')   		if eq < 0 {   			return s   		}    		key, value := s[:eq], s[eq+1:]   (		if strings.IndexByte(value, ' ') < 0 {   			return s   		}   8		return fmt.Sprintf("%s=%s", key, strconv.Quote(value))   	}   	var w bytes.Buffer   	environ := cmd.Env   	if environ == nil {   		environ = os.Environ()   	}   	for _, e := range environ {   1		fmt.Fprintf(&w, "%s \\\n", quoteEnvIfNeeded(e))   	}       
	sep := ""   	for _, arg := range cmd.Args {   2		fmt.Fprintf(&w, "%s%s", sep, quoteIfNeeded(arg))   		sep = " "   	}   	return w.String()   }       :// passLongArgsInResponseFiles modifies cmd such that, for   E// certain programs, long arguments are passed in "response files", a   D// file on disk with the arguments, with one arg per line. An actual   D// argument starting with '@' means that the rest of the argument is   %// a filename of arguments to expand.   //   >// See https://github.com/golang/go/issues/18468 (Windows) and   6// https://github.com/golang/go/issues/37768 (Darwin).   Bfunc passLongArgsInResponseFiles(cmd *exec.Cmd) (cleanup func()) {   -	cleanup = func() {} // no cleanup by default   	var argLen int   	for _, arg := range cmd.Args {   		argLen += len(arg)   	}   C	// If we're not approaching 32KB of args, just pass args normally.   J	// (use 30KB instead to be conservative; not sure how accounting is done)   (	if !useResponseFile(cmd.Path, argLen) {   		return   	}   '	tf, err := ioutil.TempFile("", "args")   	if err != nil {   F		log.Fatalf("error writing long arguments to response file: %v", err)   	}   *	cleanup = func() { os.Remove(tf.Name()) }   	var buf bytes.Buffer   #	for _, arg := range cmd.Args[1:] {    		fmt.Fprintf(&buf, "%s\n", arg)   	}   1	if _, err := tf.Write(buf.Bytes()); err != nil {   		tf.Close()   		cleanup()   F		log.Fatalf("error writing long arguments to response file: %v", err)   	}   #	if err := tf.Close(); err != nil {   		cleanup()   F		log.Fatalf("error writing long arguments to response file: %v", err)   	}   2	cmd.Args = []string{cmd.Args[0], "@" + tf.Name()}   	return cleanup   }       4func useResponseFile(path string, argLen int) bool {   ?	// Unless the program uses objabi.Flagparse, which understands   -	// response files, don't use response files.   :	// TODO: do we need more commands? asm? cgo? For now, no.   8	prog := strings.TrimSuffix(filepath.Base(path), ".exe")   	switch prog {   	case "compile", "link":   		default:   		return false   	}   F	// Windows has a limit of 32 KB arguments. To be conservative and not   D	// worry about whether that includes spaces or not, just use 30 KB.   E	// Darwin's limit is less clear. The OS claims 256KB, but we've seen   *	// failures with arglen as small as 50KB.   	if argLen > (30 << 10) {   		return true   	}   	return false   }5�_�                          ����                                                                                                                                                                                                                                                                                                                                                             `C��     �                 package builder    �              �             �   import (   	"bufio"   	"bytes"   		"errors"   	"fmt"   	"io"   	"io/ioutil"   	"os"   	"path/filepath"   	"sort"   
	"strings"   )       type archive struct {   ,	label, importPath, packagePath, file string   .	importPathAliases                    []string   }       >// checkImports verifies that each import in files refers to a   D// direct dependendency in archives or to a standard library package   A// listed in the file at stdPackageListPath. checkImports returns   C// a map from source import paths to elements of archives or to nil   !// for standard library packages.   qfunc checkImports(files []fileInfo, archives []archive, stdPackageListPath string) (map[string]*archive, error) {   #	// Read the standard package list.   8	packagesTxt, err := ioutil.ReadFile(stdPackageListPath)   	if err != nil {   		return nil, err   	}   !	stdPkgs := make(map[string]bool)   	for len(packagesTxt) > 0 {   )		n := bytes.IndexByte(packagesTxt, '\n')   		var line string   		if n < 0 {   			line = string(packagesTxt)   			packagesTxt = nil   
		} else {   !			line = string(packagesTxt[:n])   "			packagesTxt = packagesTxt[n+1:]   		}    		line = strings.TrimSpace(line)   		if line == "" {   			continue   		}   		stdPkgs[line] = true   	}       	// Index the archives.   -	importToArchive := make(map[string]*archive)   2	importAliasToArchive := make(map[string]*archive)   	for i := range archives {   		arc := &archives[i]   '		importToArchive[arc.importPath] = arc   -		for _, imp := range arc.importPathAliases {   "			importAliasToArchive[imp] = arc   		}   	}       	// Build the import map.   %	imports := make(map[string]*archive)   	var derr depsError   	for _, f := range files {   !		for _, imp := range f.imports {   			path := imp.path   E			if _, ok := imports[path]; ok || path == "C" || isRelative(path) {   H				// TODO(#1645): Support local (relative) import paths. We don't emit   I				// errors for them here, but they will probably break something else.   				continue   			}   			if stdPkgs[path] {   				imports[path] = nil   7			} else if arc := importToArchive[path]; arc != nil {   				imports[path] = arc   <			} else if arc := importAliasToArchive[path]; arc != nil {   				imports[path] = arc   			} else {   E				derr.missing = append(derr.missing, missingDep{f.filename, path})   			}   		}   	}   	if len(derr.missing) > 0 {   		return nil, derr   	}   	return imports, nil   }       N// buildImportcfgFileForCompile writes an importcfg file to be consumed by the   N// compiler. The file is constructed from direct dependencies and std imports.   =// The caller is responsible for deleting the importcfg file.   kfunc buildImportcfgFileForCompile(imports map[string]*archive, installSuffix, dir string) (string, error) {   	buf := &bytes.Buffer{}   %	goroot, ok := os.LookupEnv("GOROOT")   		if !ok {   )		return "", errors.New("GOROOT not set")   	}   	goroot = abs(goroot)       1	sortedImports := make([]string, 0, len(imports))   	for imp := range imports {   ,		sortedImports = append(sortedImports, imp)   	}   	sort.Strings(sortedImports)       $	for _, imp := range sortedImports {   &		if arc := imports[imp]; arc == nil {   			// std package   O			path := filepath.Join(goroot, "pkg", installSuffix, filepath.FromSlash(imp))   7			fmt.Fprintf(buf, "packagefile %s=%s.a\n", imp, path)   
		} else {   			if imp != arc.packagePath {   ?				fmt.Fprintf(buf, "importmap %s=%s\n", imp, arc.packagePath)   			}   E			fmt.Fprintf(buf, "packagefile %s=%s\n", arc.packagePath, arc.file)   		}   	}       ,	f, err := ioutil.TempFile(dir, "importcfg")   	if err != nil {   		return "", err   	}   	filename := f.Name()   +	if _, err := io.Copy(f, buf); err != nil {   		f.Close()   		os.Remove(filename)   		return "", err   	}   "	if err := f.Close(); err != nil {   		os.Remove(filename)   		return "", err   	}   	return filename, nil   }       sfunc buildImportcfgFileForLink(archives []archive, stdPackageListPath, installSuffix, dir string) (string, error) {   	buf := &bytes.Buffer{}   %	goroot, ok := os.LookupEnv("GOROOT")   		if !ok {   )		return "", errors.New("GOROOT not set")   	}   ;	prefix := abs(filepath.Join(goroot, "pkg", installSuffix))   7	stdPackageListFile, err := os.Open(stdPackageListPath)   	if err != nil {   		return "", err   	}   !	defer stdPackageListFile.Close()   0	scanner := bufio.NewScanner(stdPackageListFile)   	for scanner.Scan() {   '		line := strings.TrimSpace(scanner.Tex�   �   �        �   �   �       [   +		line := strings.TrimSpace(scanner.Text())   		if line == "" {   			continue   		}   b		fmt.Fprintf(buf, "packagefile %s=%s.a\n", line, filepath.Join(prefix, filepath.FromSlash(line)))   	}   &	if err := scanner.Err(); err != nil {   		return "", err   	}    	depsSeen := map[string]string{}   	for _, arc := range archives {   -		if _, ok := depsSeen[arc.packagePath]; ok {   �			return "", fmt.Errorf("internal error: package %s provided multiple times. This should have been detected during analysis.", arc.packagePath)   		}   '		depsSeen[arc.packagePath] = arc.label   D		fmt.Fprintf(buf, "packagefile %s=%s\n", arc.packagePath, arc.file)   	}   ,	f, err := ioutil.TempFile(dir, "importcfg")   	if err != nil {   		return "", err   	}   	filename := f.Name()   +	if _, err := io.Copy(f, buf); err != nil {   		f.Close()   		os.Remove(filename)   		return "", err   	}   "	if err := f.Close(); err != nil {   		os.Remove(filename)   		return "", err   	}   	return filename, nil   }       type depsError struct {   	missing []missingDep   	known   []string   }       type missingDep struct {   	filename, imp string   }       var _ error = depsError{}       #func (e depsError) Error() string {   	buf := bytes.NewBuffer(nil)   3	fmt.Fprintf(buf, "missing strict dependencies:\n")    	for _, dep := range e.missing {   A		fmt.Fprintf(buf, "\t%s: import of %q\n", dep.filename, dep.imp)   	}   	if len(e.known) == 0 {   5		fmt.Fprintln(buf, "No dependencies were provided.")   		} else {   .		fmt.Fprintln(buf, "Known dependencies are:")   		for _, imp := range e.known {   "			fmt.Fprintf(buf, "\t%s\n", imp)   		}   	}   Y	fmt.Fprint(buf, "Check that imports in Go sources match importpath attributes in deps.")   	return buf.String()   }       #func isRelative(path string) bool {   G	return strings.HasPrefix(path, "./") || strings.HasPrefix(path, "../")   }       type archiveMultiFlag []archive       ,func (m *archiveMultiFlag) String() string {   	if m == nil || len(*m) == 0 {   		return ""   	}   	return fmt.Sprint(*m)   }       0func (m *archiveMultiFlag) Set(v string) error {   	parts := strings.Split(v, "=")   	if len(parts) != 3 {   4		return fmt.Errorf("badly formed -arc flag: %s", v)   	}   ,	importPaths := strings.Split(parts[0], ":")   	a := archive{   $		importPath:        importPaths[0],   %		importPathAliases: importPaths[1:],   		packagePath:       parts[1],   #		file:              abs(parts[2]),   	}   	*m = append(*m, a)   	return nil   }5�_�                    �        ����                                                                                                                                                                                                                                                                                                                                                             `C��     �         �       5�_�                    �        ����                                                                                                                                                                                                                                                                                                                                                             `C��    �       �       �   package builder       import (   	"bufio"   	"bytes"   		"errors"   	"fmt"   	"io"   	"io/ioutil"   	"os"   	"path/filepath"   	"sort"   
	"strings"   )       type archive struct {   ,	label, importPath, packagePath, file string   .	importPathAliases                    []string   }       >// checkImports verifies that each import in files refers to a   D// direct dependendency in archives or to a standard library package   A// listed in the file at stdPackageListPath. checkImports returns   C// a map from source import paths to elements of archives or to nil   !// for standard library packages.   qfunc checkImports(files []fileInfo, archives []archive, stdPackageListPath string) (map[string]*archive, error) {   #	// Read the standard package list.   8	packagesTxt, err := ioutil.ReadFile(stdPackageListPath)   	if err != nil {   		return nil, err   	}   !	stdPkgs := make(map[string]bool)   	for len(packagesTxt) > 0 {   )		n := bytes.IndexByte(packagesTxt, '\n')   		var line string   		if n < 0 {   			line = string(packagesTxt)   			packagesTxt = nil   
		} else {   !			line = string(packagesTxt[:n])   "			packagesTxt = packagesTxt[n+1:]   		}    		line = strings.TrimSpace(line)   		if line == "" {   			continue   		}   		stdPkgs[line] = true   	}       	// Index the archives.   -	importToArchive := make(map[string]*archive)   2	importAliasToArchive := make(map[string]*archive)   	for i := range archives {   		arc := &archives[i]   '		importToArchive[arc.importPath] = arc   -		for _, imp := range arc.importPathAliases {   "			importAliasToArchive[imp] = arc   		}   	}       	// Build the import map.   %	imports := make(map[string]*archive)   	var derr depsError   	for _, f := range files {   !		for _, imp := range f.imports {   			path := imp.path   E			if _, ok := imports[path]; ok || path == "C" || isRelative(path) {   H				// TODO(#1645): Support local (relative) import paths. We don't emit   I				// errors for them here, but they will probably break something else.   				continue   			}   			if stdPkgs[path] {   				imports[path] = nil   7			} else if arc := importToArchive[path]; arc != nil {   				imports[path] = arc   <			} else if arc := importAliasToArchive[path]; arc != nil {   				imports[path] = arc   			} else {   E				derr.missing = append(derr.missing, missingDep{f.filename, path})   			}   		}   	}   	if len(derr.missing) > 0 {   		return nil, derr   	}   	return imports, nil   }       N// buildImportcfgFileForCompile writes an importcfg file to be consumed by the   N// compiler. The file is constructed from direct dependencies and std imports.   =// The caller is responsible for deleting the importcfg file.   kfunc buildImportcfgFileForCompile(imports map[string]*archive, installSuffix, dir string) (string, error) {   	buf := &bytes.Buffer{}   %	goroot, ok := os.LookupEnv("GOROOT")   		if !ok {   )		return "", errors.New("GOROOT not set")   	}   	goroot = abs(goroot)       1	sortedImports := make([]string, 0, len(imports))   	for imp := range imports {   ,		sortedImports = append(sortedImports, imp)   	}   	sort.Strings(sortedImports)       $	for _, imp := range sortedImports {   &		if arc := imports[imp]; arc == nil {   			// std package   O			path := filepath.Join(goroot, "pkg", installSuffix, filepath.FromSlash(imp))   7			fmt.Fprintf(buf, "packagefile %s=%s.a\n", imp, path)   
		} else {   			if imp != arc.packagePath {   ?				fmt.Fprintf(buf, "importmap %s=%s\n", imp, arc.packagePath)   			}   E			fmt.Fprintf(buf, "packagefile %s=%s\n", arc.packagePath, arc.file)   		}   	}       ,	f, err := ioutil.TempFile(dir, "importcfg")   	if err != nil {   		return "", err   	}   	filename := f.Name()   +	if _, err := io.Copy(f, buf); err != nil {   		f.Close()   		os.Remove(filename)   		return "", err   	}   "	if err := f.Close(); err != nil {   		os.Remove(filename)   		return "", err   	}   	return filename, nil   }       sfunc buildImportcfgFileForLink(archives []archive, stdPackageListPath, installSuffix, dir string) (string, error) {   	buf := &bytes.Buffer{}   %	goroot, ok := os.LookupEnv("GOROOT")   		if !ok {   )		return "", errors.New("GOROOT not set")   	}   ;	prefix := abs(filepath.Join(goroot, "pkg", installSuffix))   7	stdPackageListFile, err := os.Open(stdPackageListPath)   	if err != nil {   		return "", err   	}   !	defer stdPackageListFile.Close()   0	scanner := bufio.NewScanner(stdPackageListFile)   	for scanner.Scan() {   +		line := strings.TrimSpace(scanner.Text())   		if line == "" {   			continue   		}   b		fmt.Fprintf(buf, "packagefile %s=%s.a\n", line, filepath.Join(prefix, filepath.FromSlash(line)))   	}   &	if err := scanner.Err(); err != nil {   		return "", err   	}    	depsSeen := map[string]string{}   	for _, arc := range archives {   -		if _, ok := depsSeen[arc.packagePath]; ok {   �			return "", fmt.Errorf("internal error: package %s provided multiple times. This should have been detected during analysis.", arc.packagePath)   		}   '		depsSeen[arc.packagePath] = arc.label   D		fmt.Fprintf(buf, "packagefile %s=%s\n", arc.packagePath, arc.file)   	}   ,	f, err := ioutil.TempFile(dir, "importcfg")   	if err != nil {   		return "", err   	}   	filename := f.Name()   +	if _, err := io.Copy(f, buf); err != nil {   		f.Close()   		os.Remove(filename)   		return "", err   	}   "	if err := f.Close(); err != nil {   		os.Remove(filename)   		return "", err   	}   	return filename, nil   }       type depsError struct {   	missing []missingDep   	known   []string   }       type missingDep struct {   	filename, imp string   }       var _ error = depsError{}       #func (e depsError) Error() string {   	buf := bytes.NewBuffer(nil)   3	fmt.Fprintf(buf, "missing strict dependencies:\n")    	for _, dep := range e.missing {   A		fmt.Fprintf(buf, "\t%s: import of %q\n", dep.filename, dep.imp)   	}   	if len(e.known) == 0 {   5		fmt.Fprintln(buf, "No dependencies were provided.")   		} else {   .		fmt.Fprintln(buf, "Known dependencies are:")   		for _, imp := range e.known {   "			fmt.Fprintf(buf, "\t%s\n", imp)   		}   	}   Y	fmt.Fprint(buf, "Check that imports in Go sources match importpath attributes in deps.")   	return buf.String()   }       #func isRelative(path string) bool {   G	return strings.HasPrefix(path, "./") || strings.HasPrefix(path, "../")   }       type archiveMultiFlag []archive       ,func (m *archiveMultiFlag) String() string {   	if m == nil || len(*m) == 0 {   		return ""   	}   	return fmt.Sprint(*m)   }       0func (m *archiveMultiFlag) Set(v string) error {   	parts := strings.Split(v, "=")   	if len(parts) != 3 {   4		return fmt.Errorf("badly formed -arc flag: %s", v)   	}   ,	importPaths := strings.Split(parts[0], ":")   	a := archive{   $		importPath:        importPaths[0],   %		importPathAliases: importPaths[1:],   		packagePath:       parts[1],   #		file:              abs(parts[2]),   	}   	*m = append(*m, a)   	return nil   }5�_�                            ����                                                                                                                                                                                                                                                                                                                                                V       `C��     �         �      Atype archive struct { label, importPath, packagePath, file string5�_�                             ����                                                                                                                                                                                                                                                                                                                                                V       `C��     �              5��