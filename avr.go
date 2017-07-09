package main

import(
"os"
"fmt"
"strings"
"bytes"
"os/exec"
"io/ioutil"
"log"
)

func main(){
fmt.Print("New Project!!\n");
fmt.Print("Start[Y/N] := ");
var va string;
fmt.Scanln(&va)
if va == "N"{
log.Fatal("Bummer Dude");
}

if len(va) == 0 || va != "Y" && va != "N"{
log.Fatal("Not An Appropriate Answer")
}
if va == "Y"{
fmt.Print("Project Name := ");
os.Chdir("C:\\AVRSketchbook")
var ar string;
fmt.Scan(&ar)
if len(ar) == 0{
log.Fatal("Not An Appropriate Answer")
}
os.Mkdir(ar ,os.ModeDir | 0755);
fmt.Print("Programmer := ")
var po string
fmt.Scan(&po)
fmt.Print("MCU := ")
var pa string
fmt.Scan(&pa)
fmt.Print("Target := ")
var pl string
fmt.Scan(&pl)
fmt.Print("Port := ")
var pq string
fmt.Scan(&pq)
fmt.Print("F_CPU := ")
var qe string
fmt.Scan(&qe)
ae := "AVRDUDE_PROGRAMMER = " + po
ao := "MCU = " + pa 
lo := "TARGET = " + pl
lp := "AVRDUDE_PORT = " + pq
pm := "F_CPU = " + qe + "000000"
op := "C:\\AVRSketchbook\\Makefile"
oe := "C:\\AVRSketchbook\\"+ ar
  input, errr := ioutil.ReadFile(op)
  inputs, errrr := ioutil.ReadFile(op)
  inputss, errrrr := ioutil.ReadFile(op)
  inputsss, errrrrr := ioutil.ReadFile(op)
  inputssss, errrrrrr := ioutil.ReadFile(op)
        if errr != nil ||errrr != nil || errrrr != nil || errrrrr != nil || errrrrrr != nil {
                log.Fatalln(errr)
                log.Fatalln(errrr)
                log.Fatalln(errrrr)
                log.Fatalln(errrrrr)
                log.Fatalln(errrrrrr)
        }

        lines := strings.Split(string(input), "\n")
        liness := strings.Split(string(inputs), "\n")
        linesss:= strings.Split(string(inputss), "\n")
        linessss:= strings.Split(string(inputsss), "\n")
        linesssss := strings.Split(string(inputssss), "\n")

        for i, line := range lines {
                if strings.Contains(line, "AVRDUDE_PROGRAMMER = ") {
                        lines[i] = ae
                }
        }
        output := strings.Join(lines, "\n")
        errr = ioutil.WriteFile("Makefile", []byte(output), 0644)
        if errr != nil {
                log.Fatalln(errr)
        }
		
		for a, lins := range liness {
                if strings.Contains(lins, "MCU = ") {
                        liness[a] = ao
                }
        }
        outputs := strings.Join(liness, "\n")
        errrr = ioutil.WriteFile("Makefile", []byte(outputs), 0644)
        if errrr != nil {
                log.Fatalln(errrr)
        }
		
		for b, linss := range linesss {
                if strings.Contains(linss, "AVRDUDE_PORT =") {
                        linesss[b] = lp
                }
        }
        outputss := strings.Join(linesss, "\n")
        errrrr = ioutil.WriteFile("Makefile", []byte(outputss), 0644)
        if errrrr != nil {
                log.Fatalln(errrrr)
        }
		
		for c, linsss := range linessss {
                if strings.Contains(linsss, "TARGET = ") {
                        linessss[c] = lo
                }
        }
        outputsss := strings.Join(linessss, "\n")
        errrrrr = ioutil.WriteFile("Makefile", []byte(outputsss), 0644)
        if errrrrr != nil {
                log.Fatalln(errrrrr)
        }
		
		for d, linssss := range linesssss {
                if strings.Contains(linssss, "F_CPU = ") {
                        linesssss[d] = pm
                }
        }
        outputssss := strings.Join(linesssss, "\n")
        errrrrrr = ioutil.WriteFile("Makefile", []byte(outputssss), 0644)
        if errrrrrr != nil {
                log.Fatalln(errrrrrr)
        }
		
cmd := exec.Command("xcopy",op,oe)

// Stdout buffer
cmdOutput := &bytes.Buffer{}
// Attach buffer to command
cmd.Stdout = cmdOutput
 
// Execute command
printCommand(cmd)
err := cmd.Run() // will wait for command to return
printError(err)
// Only output the commands stdout
printOutput(cmdOutput.Bytes())
fmt.Print("Have Fun")

}
}

func printCommand(cmd *exec.Cmd) {
  fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
  if err != nil {
    os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
  }
}

func printOutput(outs []byte) {
  if len(outs) > 0 {
    fmt.Printf("==> Output: %s\n", string(outs))
  }
}

