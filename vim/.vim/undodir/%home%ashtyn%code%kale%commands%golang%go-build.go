Vim�UnDo� �P�n��$�	R�� �����H�������J�   >                 	       	   	   	    `��F    _�                            ����                                                                                                                                                                                                                                                                                                                                                             `��(     �          =      package command5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             `��+     �          =      package 5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             `��/    �               =   package goBuild       import (   	"fmt"   	"kale/utils"   
	"strings"       	"github.com/muesli/termenv"   )       1func indexOf(element string, data []string) int {   	for k, v := range data {   		if element == v {   			return k   		}   	}   	return -1 //not found.   }       type GO struct {   	Platforms []string   	Params    []string   	Target    []string   	Out       string   }       func (g *GO) Build() {   	c := utils.InitColors()   	if len(g.Platforms) != 0 {   #		for _, str := range g.Platforms {   			outName := g.Out   			g.Out = outName       (			platDouble := strings.Split(str, " ")   -			os := strings.Split(platDouble[0], "=")[1]   /			arch := strings.Split(platDouble[0], "=")[1]       K			//filePather := regexp.MustCompile(`^(.*/)?(?:$|(.+?)(?:(\.[^.]*$)|$))`)   6			//name := (filePather.FindStringSubmatch(g.Out))[2]   2			g.Out = g.Out + "/" + os + "/" + "kale-" + arch       T			build := Builder{ProcName: "Building", Output: g.Out, Cmd: "go", Env: platDouble}   			build.AddArgs("build")   			build.AddArgs(g.Params...)   			build.AddTarget(g.Target...)   			build.Construct()       +			utils.FPrint(c.Green, "Compiled", g.Out)   J			fmt.Println(termenv.String("\t- OS:").Foreground(c.Cyan).Bold(), g.Out)   T			fmt.Println(termenv.String("\t- ARCHITECTURE:").Foreground(c.Cyan).Bold(), g.Out)   d			fmt.Println(termenv.String("\t- PARAMS:").Foreground(c.Cyan).Bold(), strings.Join(g.Params, " "))   			g.Out = outName   		}   		} else {   B		build := Builder{ProcName: "Building", Output: g.Out, Cmd: "go"}   		build.AddArgs("build")   		build.AddTarget(g.Target...)   		build.Construct()   *		utils.FPrint(c.Green, "Compiled", g.Out)   	}   }5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             `��/     �          =      package goBuild5�_�                    *       ����                                                                                                                                                                                                                                                                                                                                                             `��5     �   )   +   =      T			build := Builder{ProcName: "Building", Output: g.Out, Cmd: "go", Env: platDouble}5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             `��8    �               =   package goBuild       import (   	"fmt"   	"kale/utils"   
	"strings"       	"github.com/muesli/termenv"   )       1func indexOf(element string, data []string) int {   	for k, v := range data {   		if element == v {   			return k   		}   	}   	return -1 //not found.   }       type GO struct {   	Platforms []string   	Params    []string   	Target    []string   	Out       string   }       func (g *GO) Build() {   	c := utils.InitColors()   	if len(g.Platforms) != 0 {   #		for _, str := range g.Platforms {   			outName := g.Out   			g.Out = outName       (			platDouble := strings.Split(str, " ")   -			os := strings.Split(platDouble[0], "=")[1]   /			arch := strings.Split(platDouble[0], "=")[1]       K			//filePather := regexp.MustCompile(`^(.*/)?(?:$|(.+?)(?:(\.[^.]*$)|$))`)   6			//name := (filePather.FindStringSubmatch(g.Out))[2]   2			g.Out = g.Out + "/" + os + "/" + "kale-" + arch       \			build := command.Builder{ProcName: "Building", Output: g.Out, Cmd: "go", Env: platDouble}   			build.AddArgs("build")   			build.AddArgs(g.Params...)   			build.AddTarget(g.Target...)   			build.Construct()       +			utils.FPrint(c.Green, "Compiled", g.Out)   J			fmt.Println(termenv.String("\t- OS:").Foreground(c.Cyan).Bold(), g.Out)   T			fmt.Println(termenv.String("\t- ARCHITECTURE:").Foreground(c.Cyan).Bold(), g.Out)   d			fmt.Println(termenv.String("\t- PARAMS:").Foreground(c.Cyan).Bold(), strings.Join(g.Params, " "))   			g.Out = outName   		}   		} else {   B		build := Builder{ProcName: "Building", Output: g.Out, Cmd: "go"}   		build.AddArgs("build")   		build.AddTarget(g.Target...)   		build.Construct()   *		utils.FPrint(c.Green, "Compiled", g.Out)   	}   }5�_�                    +       ����                                                                                                                                                                                                                                                                                                                                                             `��9     �   *   ,   >      \			build := command.Builder{ProcName: "Building", Output: g.Out, Cmd: "go", Env: platDouble}5�_�      	              8       ����                                                                                                                                                                                                                                                                                                                                                             `��D     �   7   9   >      B		build := Builder{ProcName: "Building", Output: g.Out, Cmd: "go"}5�_�                  	   8       ����                                                                                                                                                                                                                                                                                                                                                             `��E    �               >   package goBuild       import (   	"fmt"   	command "kale/commands"   	"kale/utils"   
	"strings"       	"github.com/muesli/termenv"   )       1func indexOf(element string, data []string) int {   	for k, v := range data {   		if element == v {   			return k   		}   	}   	return -1 //not found.   }       type GO struct {   	Platforms []string   	Params    []string   	Target    []string   	Out       string   }       func (g *GO) Build() {   	c := utils.InitColors()   	if len(g.Platforms) != 0 {   #		for _, str := range g.Platforms {   			outName := g.Out   			g.Out = outName       (			platDouble := strings.Split(str, " ")   -			os := strings.Split(platDouble[0], "=")[1]   /			arch := strings.Split(platDouble[0], "=")[1]       K			//filePather := regexp.MustCompile(`^(.*/)?(?:$|(.+?)(?:(\.[^.]*$)|$))`)   6			//name := (filePather.FindStringSubmatch(g.Out))[2]   2			g.Out = g.Out + "/" + os + "/" + "kale-" + arch       \			build := command.Builder{ProcName: "Building", Output: g.Out, Cmd: "go", Env: platDouble}   			build.AddArgs("build")   			build.AddArgs(g.Params...)   			build.AddTarget(g.Target...)   			build.Construct()       +			utils.FPrint(c.Green, "Compiled", g.Out)   J			fmt.Println(termenv.String("\t- OS:").Foreground(c.Cyan).Bold(), g.Out)   T			fmt.Println(termenv.String("\t- ARCHITECTURE:").Foreground(c.Cyan).Bold(), g.Out)   d			fmt.Println(termenv.String("\t- PARAMS:").Foreground(c.Cyan).Bold(), strings.Join(g.Params, " "))   			g.Out = outName   		}   		} else {   J		build := command.Builder{ProcName: "Building", Output: g.Out, Cmd: "go"}   		build.AddArgs("build")   		build.AddTarget(g.Target...)   		build.Construct()   *		utils.FPrint(c.Green, "Compiled", g.Out)   	}   }5��