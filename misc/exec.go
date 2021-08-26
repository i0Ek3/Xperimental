package main

import (
        "fmt"
        "os"
        "os/exec"
        "log"
)

func lsl1() {
        env := os.Environ();
        procAttr := &os.ProcAttr{
                Env: env,
                Files:  []*os.File{
                        os.Stdin,
                        os.Stdout,
                        os.Stderr,
                },
        }
        pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
        if err != nil {
                fmt.Printf("Error %v starting process!", err)
                os.Exit(1)
        }
        fmt.Printf("The process id is %v.\n", pid)

        cmd := exec.Command("pwd")
        err = cmd.Run()
        if err != nil {
                fmt.Printf("Error %v executing command!", err)
                os.Exit(1)
        }
        fmt.Printf("The command is %v.", cmd)
}

func lsl2() {
    cmd := exec.Command("ls", "-l", "/usr/local/bin/")
    if err := cmd.Run(); err != nil {
        log.Fatalf("cmd.Run() failed with %s\n.", err)
    }
}

func main() {
    lsl1()
    //lsl2()
}
