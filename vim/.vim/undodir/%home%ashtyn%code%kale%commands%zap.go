Vim�UnDo� ��7��e�u�/=h���%^yiB���G   N           !      $       $   $   $    `��J    _�                             ����                                                                                                                                                                                                                                                                                                                                                             `���     �         O      func Err(err string) {5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             `���     �         O      	c := utils.InitColors()5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             `���     �         O      E	fmt.Println(termenv.String("Error: ").Foreground(c.Red).Bold(), err)5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             `���     �         O      	os.Exit(0)5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             `���     �         O      }5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             `���     �         O      ?}5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             `���     �         O      }5�_�      	                     ����                                                                                                                                                                                                                                                                                                                                                             `���     �         O      	//c := utils.InitColors()   G	//fmt.Println(termenv.String("Error: ").Foreground(c.Red).Bold(), err)   	//os.Exit(0)5�_�      
           	           ����                                                                                                                                                                                                                                                                                                                                                             `���    �               O   package command       import (   	"fmt"   	"io/ioutil"   	"kale/utils"   	"os"   
	"os/exec"   		"regexp"   
	"strings"       	"github.com/muesli/termenv"   )       var path string       type valid struct {   	Name  string   	Value string   }       //func Err(err string) {   //c := utils.InitColors()   F//fmt.Println(termenv.String("Error: ").Foreground(c.Red).Bold(), err)   //os.Exit(0)   //}       func Transfer(path string) {   	rm := exec.Command("rm", path)   	rm.Stdout = os.Stdout   	rm.Stderr = os.Stdin   		rm.Run()   -	mv := exec.Command("mv", path+".copy", path)   	mv.Stdout = os.Stdout   	mv.Stderr = os.Stderr   		mv.Run()   }       %func Zap(evs []string, name string) {   	var aval []valid   	for _, v := range evs {   		tv := strings.Split(v, "=")   7		aval = append(aval, valid{Name: tv[0], Value: tv[1]})   	}   	// fmt.Println(flag.Param)   N	// f, _ := os.OpenFile(flag.Param, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)   '	contentB, err := ioutil.ReadFile(name)   	if err != nil {   		Err(err.Error())   	}   =	er := ioutil.WriteFile(name+".copy", []byte(contentB), 0644)   	if er != nil {   		Err(er.Error())   	}   	path = name   	content := string(contentB)   #	sN := strings.Split(content, "\n")   	for i, line := range sN {   F		if m, _ := regexp.MatchString(`//\s*@zap\s+var:`, line); m == true {   M			if m, _ := regexp.MatchString(`var\s+\S+\s+\w+\s*=`, sN[i+1]); m == true {   %				params := strings.Fields(sN[i+1])   				name := params[1]   				tp := params[2]   /				r := regexp.MustCompile(`//\s*@zap\s+var:`)   				s := r.Split(line, -1)       				for _, a := range aval {   I					if a.Name == strings.Replace(strings.Join(s[1:], ""), " ", "", -1) {   B						newVar := fmt.Sprintln("var", name, tp, "= \""+a.Value+"\"")   						sN[i+1] = newVar   					}   				}   			}   		}   	}       	file := strings.Join(sN, "\n")   +	ioutil.WriteFile(name, []byte(file), 0644)   }5�_�   	              
          ����                                                                                                                                                                                                                                                                                                                                                             `���     �         L      //}5�_�   
                 .       ����                                                                                                                                                                                                                                                                                                                                                             `���     �   -   /   L      		Err(err.Error())5�_�                    .       ����                                                                                                                                                                                                                                                                                                                                                             `���     �   -   /   L      		(err.Error())5�_�                    .       ����                                                                                                                                                                                                                                                                                                                                                             `���    �               L   package command       import (   	"fmt"   	"io/ioutil"   	"os"   
	"os/exec"   		"regexp"   
	"strings"   )       var path string       type valid struct {   	Name  string   	Value string   }       //func Err(err string) {   //c := utils.InitColors()   F//fmt.Println(termenv.String("Error: ").Foreground(c.Red).Bold(), err)   //os.Exit(0)   //}       func Transfer(path string) {   	rm := exec.Command("rm", path)   	rm.Stdout = os.Stdout   	rm.Stderr = os.Stdin   		rm.Run()   -	mv := exec.Command("mv", path+".copy", path)   	mv.Stdout = os.Stdout   	mv.Stderr = os.Stderr   		mv.Run()   }       %func Zap(evs []string, name string) {   	var aval []valid   	for _, v := range evs {   		tv := strings.Split(v, "=")   7		aval = append(aval, valid{Name: tv[0], Value: tv[1]})   	}   	// fmt.Println(flag.Param)   N	// f, _ := os.OpenFile(flag.Param, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)   '	contentB, err := ioutil.ReadFile(name)   	if err != nil {    		utils.FormatPrint(err.Error())   	}   =	er := ioutil.WriteFile(name+".copy", []byte(contentB), 0644)   	if er != nil {   		Err(er.Error())   	}   	path = name   	content := string(contentB)   #	sN := strings.Split(content, "\n")   	for i, line := range sN {   F		if m, _ := regexp.MatchString(`//\s*@zap\s+var:`, line); m == true {   M			if m, _ := regexp.MatchString(`var\s+\S+\s+\w+\s*=`, sN[i+1]); m == true {   %				params := strings.Fields(sN[i+1])   				name := params[1]   				tp := params[2]   /				r := regexp.MustCompile(`//\s*@zap\s+var:`)   				s := r.Split(line, -1)       				for _, a := range aval {   I					if a.Name == strings.Replace(strings.Join(s[1:], ""), " ", "", -1) {   B						newVar := fmt.Sprintln("var", name, tp, "= \""+a.Value+"\"")   						sN[i+1] = newVar   					}   				}   			}   		}   	}       	file := strings.Join(sN, "\n")   +	ioutil.WriteFile(name, []byte(file), 0644)   }5�_�                    .   	    ����                                                                                                                                                                                                                                                                                                                                                             `��     �   -   /   L       		utils.FormatPrint(err.Error())5�_�                    .   	    ����                                                                                                                                                                                                                                                                                                                                                             `��     �   -   /   L      		utils.FrmatPrint(err.Error())5�_�                    .   	    ����                                                                                                                                                                                                                                                                                                                                                             `��     �   -   /   L      		utils.FmatPrint(err.Error())5�_�                    .   	    ����                                                                                                                                                                                                                                                                                                                                                             `��     �   -   /   L      		utils.FatPrint(err.Error())5�_�                    .   	    ����                                                                                                                                                                                                                                                                                                                                                             `��     �   -   /   L      		utils.FtPrint(err.Error())5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             `��    �               L   package command       import (   	"fmt"   	"io/ioutil"   	"os"   
	"os/exec"   		"regexp"   
	"strings"   )       var path string       type valid struct {   	Name  string   	Value string   }       //func Err(err string) {   //c := utils.InitColors()   F//fmt.Println(termenv.String("Error: ").Foreground(c.Red).Bold(), err)   //os.Exit(0)   //}       func Transfer(path string) {   	rm := exec.Command("rm", path)   	rm.Stdout = os.Stdout   	rm.Stderr = os.Stdin   		rm.Run()   -	mv := exec.Command("mv", path+".copy", path)   	mv.Stdout = os.Stdout   	mv.Stderr = os.Stderr   		mv.Run()   }       %func Zap(evs []string, name string) {   	var aval []valid   	for _, v := range evs {   		tv := strings.Split(v, "=")   7		aval = append(aval, valid{Name: tv[0], Value: tv[1]})   	}   	// fmt.Println(flag.Param)   N	// f, _ := os.OpenFile(flag.Param, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)   '	contentB, err := ioutil.ReadFile(name)   	if err != nil {   		utils.FPrint(err.Error())   	}   =	er := ioutil.WriteFile(name+".copy", []byte(contentB), 0644)   	if er != nil {   		Err(er.Error())   	}   	path = name   	content := string(contentB)   #	sN := strings.Split(content, "\n")   	for i, line := range sN {   F		if m, _ := regexp.MatchString(`//\s*@zap\s+var:`, line); m == true {   M			if m, _ := regexp.MatchString(`var\s+\S+\s+\w+\s*=`, sN[i+1]); m == true {   %				params := strings.Fields(sN[i+1])   				name := params[1]   				tp := params[2]   /				r := regexp.MustCompile(`//\s*@zap\s+var:`)   				s := r.Split(line, -1)       				for _, a := range aval {   I					if a.Name == strings.Replace(strings.Join(s[1:], ""), " ", "", -1) {   B						newVar := fmt.Sprintln("var", name, tp, "= \""+a.Value+"\"")   						sN[i+1] = newVar   					}   				}   			}   		}   	}       	file := strings.Join(sN, "\n")   +	ioutil.WriteFile(name, []byte(file), 0644)   }5�_�                    /   	    ����                                                                                                                                                                                                                                                                                                                                                             `��     �   .   0   M      		utils.FPrint(err.Error())5�_�                    /       ����                                                                                                                                                                                                                                                                                                                                                             `��     �   .   0   M      		utils.FPrint(err.Error())5�_�                    /       ����                                                                                                                                                                                                                                                                                                                                                             `��     �   .   0   M      		utils.FPrint(,err.Error())5�_�                    /       ����                                                                                                                                                                                                                                                                                                                                                             `��     �   .   0   M      		utils.FPrint(,,err.Error())5�_�                    /       ����                                                                                                                                                                                                                                                                                                                                                             `��     �   .   0   M      		utils.FPrint(,"",err.Error())5�_�                    %       ����                                                                                                                                                                                                                                                                                                                                                             `��+     �   %   '   M    �   %   &   M    5�_�                    &        ����                                                                                                                                                                                                                                                                                                                                                             `��,     �   %   '   N      //c := utils.InitColors()5�_�                    &        ����                                                                                                                                                                                                                                                                                                                                                             `��-    �   %   '   N      /c := utils.InitColors()5�_�                    0       ����                                                                                                                                                                                                                                                                                                                                                             `��-     �   /   1   N      $		utils.FPrint(,"Error",err.Error())5�_�                    0       ����                                                                                                                                                                                                                                                                                                                                                             `��4     �   /   1   N      $		utils.FPrint(,"Error",err.Error())5�_�                    0       ����                                                                                                                                                                                                                                                                                                                                                             `��:     �   %   '   N      c := utils.InitColors()�   /   1   N      )		utils.FPrint(c.Red,"Error",err.Error())5�_�                     0       ����                                                                                                                                                                                                                                                                                                                                                             `��:    �               N   package command       import (   	"fmt"   	"io/ioutil"   	"kale/utils"   	"os"   
	"os/exec"   		"regexp"   
	"strings"   )       var path string       type valid struct {   	Name  string   	Value string   }       //func Err(err string) {   //c := utils.InitColors()   F//fmt.Println(termenv.String("Error: ").Foreground(c.Red).Bold(), err)   //os.Exit(0)   //}       func Transfer(path string) {   	rm := exec.Command("rm", path)   	rm.Stdout = os.Stdout   	rm.Stderr = os.Stdin   		rm.Run()   -	mv := exec.Command("mv", path+".copy", path)   	mv.Stdout = os.Stdout   	mv.Stderr = os.Stderr   		mv.Run()   }       %func Zap(evs []string, name string) {   	c := utils.InitColors()   	var aval []valid   	for _, v := range evs {   		tv := strings.Split(v, "=")   7		aval = append(aval, valid{Name: tv[0], Value: tv[1]})   	}   	// fmt.Println(flag.Param)   N	// f, _ := os.OpenFile(flag.Param, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)   '	contentB, err := ioutil.ReadFile(name)   	if err != nil {   +		utils.FPrint(c.Red, "Error", err.Error())   	}   =	er := ioutil.WriteFile(name+".copy", []byte(contentB), 0644)   	if er != nil {   		Err(er.Error())   	}   	path = name   	content := string(contentB)   #	sN := strings.Split(content, "\n")   	for i, line := range sN {   F		if m, _ := regexp.MatchString(`//\s*@zap\s+var:`, line); m == true {   M			if m, _ := regexp.MatchString(`var\s+\S+\s+\w+\s*=`, sN[i+1]); m == true {   %				params := strings.Fields(sN[i+1])   				name := params[1]   				tp := params[2]   /				r := regexp.MustCompile(`//\s*@zap\s+var:`)   				s := r.Split(line, -1)       				for _, a := range aval {   I					if a.Name == strings.Replace(strings.Join(s[1:], ""), " ", "", -1) {   B						newVar := fmt.Sprintln("var", name, tp, "= \""+a.Value+"\"")   						sN[i+1] = newVar   					}   				}   			}   		}   	}       	file := strings.Join(sN, "\n")   +	ioutil.WriteFile(name, []byte(file), 0644)   }5�_�      !               0       ����                                                                                                                                                                                                                                                                                                                                                             `��:     �   /   1   N      +		utils.FPrint(c.Red, "Error", err.Error())5�_�       "           !   3       ����                                                                                                                                                                                                                                                                                                                                                             `��?     �   3   5   N    �   3   4   N    5�_�   !   #           "   5       ����                                                                                                                                                                                                                                                                                                                                                             `��A     �   4   5          		Err(er.Error())5�_�   "   $           #   4   !    ����                                                                                                                                                                                                                                                                                                                                                             `��D     �   3   5   N      +		utils.FPrint(c.Red, "Error", err.Error())5�_�   #               $   J       ����                                                                                                                                                                                                                                                                                                                                                             `��I    �               N   package command       import (   	"fmt"   	"io/ioutil"   	"kale/utils"   	"os"   
	"os/exec"   		"regexp"   
	"strings"   )       var path string       type valid struct {   	Name  string   	Value string   }       //func Err(err string) {   //c := utils.InitColors()   F//fmt.Println(termenv.String("Error: ").Foreground(c.Red).Bold(), err)   //os.Exit(0)   //}       func Transfer(path string) {   	rm := exec.Command("rm", path)   	rm.Stdout = os.Stdout   	rm.Stderr = os.Stdin   		rm.Run()   -	mv := exec.Command("mv", path+".copy", path)   	mv.Stdout = os.Stdout   	mv.Stderr = os.Stderr   		mv.Run()   }       %func Zap(evs []string, name string) {   	c := utils.InitColors()   	var aval []valid   	for _, v := range evs {   		tv := strings.Split(v, "=")   7		aval = append(aval, valid{Name: tv[0], Value: tv[1]})   	}   	// fmt.Println(flag.Param)   N	// f, _ := os.OpenFile(flag.Param, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)   '	contentB, err := ioutil.ReadFile(name)   	if err != nil {   +		utils.FPrint(c.Red, "Error", err.Error())   	}   =	er := ioutil.WriteFile(name+".copy", []byte(contentB), 0644)   	if er != nil {   *		utils.FPrint(c.Red, "Error", er.Error())   	}   	path = name   	content := string(contentB)   #	sN := strings.Split(content, "\n")   	for i, line := range sN {   F		if m, _ := regexp.MatchString(`//\s*@zap\s+var:`, line); m == true {   M			if m, _ := regexp.MatchString(`var\s+\S+\s+\w+\s*=`, sN[i+1]); m == true {   %				params := strings.Fields(sN[i+1])   				name := params[1]   				tp := params[2]   /				r := regexp.MustCompile(`//\s*@zap\s+var:`)   				s := r.Split(line, -1)       				for _, a := range aval {   I					if a.Name == strings.Replace(strings.Join(s[1:], ""), " ", "", -1) {   B						newVar := fmt.Sprintln("var", name, tp, "= \""+a.Value+"\"")   						sN[i+1] = newVar   					}   				}   			}   		}   	}       	file := strings.Join(sN, "\n")   +	ioutil.WriteFile(name, []byte(file), 0644)   }5��