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
var ar string;
fmt.Scan(&ar)
op := "C:\\AVRSketchbook\\Makefile"
oe := "C:\\AVRSketchbook\\"+ ar
oa := "C:\\AVRSketchbook"
os.Chdir(oa)

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
ae := "AVRDUDE_PROGRAMMER = " + po
ao := "MCU = " + pa 
lo := "TARGET = " + pl
lp := "AVRDUDE_PORT = " + pq
  input, errr := ioutil.ReadFile(op)
        if errr != nil {
                log.Fatalln(errr)
        }

        lines := strings.Split(string(input), "\n")

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
		
		
		 inputs, errrr := ioutil.ReadFile(op)
        if errrr != nil {
                log.Fatalln(errrr)
        }

        liness := strings.Split(string(inputs), "\n")

        for a, linesss := range liness {
                if strings.Contains(linesss, "MCU = ") {
                        liness[a] = ao
                }
        }
        outputs := strings.Join(liness, "\n")
        errrr = ioutil.WriteFile("Makefile", []byte(outputs), 0644)
        if errrr != nil {
                log.Fatalln(errrr)
        }
		
		
		
		
		 inputss, errrrr := ioutil.ReadFile(op)
        if errrrr != nil {
                log.Fatalln(errrrr)
        }

        linesss := strings.Split(string(inputss), "\n")

        for b, linessss := range linesss {
                if strings.Contains(linessss, "TARGET = ") {
                        linesss[b] = lo
                }
        }
        outputss := strings.Join(linesss, "\n")
        errrrr = ioutil.WriteFile("Makefile", []byte(outputss), 0644)
        if errrrr != nil {
                log.Fatalln(errrrr)
        }

		
		inputsss, errrrrr := ioutil.ReadFile(op)
        if errrrrr != nil {
                log.Fatalln(errrrrr)
        }

        linessss := strings.Split(string(inputsss), "\n")

        for c, linesssss := range linessss {
                if strings.Contains(linesssss, "AVRDUDE_PORT =") {
                        linessss[c] = lp
                }
        }
        outputsss := strings.Join(linessss, "\n")
        errrrrr = ioutil.WriteFile("Makefile", []byte(outputsss), 0644)
        if errrrrr != nil {
                log.Fatalln(errrrrr)
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

