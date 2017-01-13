package main
import (
    "strings"
    "strconv"
)
// Analyze summarizes a stream of input without modifying user files.
func Analyze(input string) map[string]interface{} {

methods:=map[string]int{};statuses:=map[string]int{};valid:=0
for _,line:=range strings.Split(input,"\n") {fields:=strings.Fields(line);if len(fields)<3{continue}
method:=fields[0];status,err:=strconv.Atoi(fields[len(fields)-1]);if err!=nil{continue}
methods[method]++;statuses[strconv.Itoa(status/100)+"xx"]++;valid++}
return map[string]interface{}{"requests":valid,"methods":methods,"statusClasses":statuses}
}
