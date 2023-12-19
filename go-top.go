package main

import (
    "fmt"
    "os"
    "os/exec"
    "sort"
    "strconv"
    "strings"
    "time"
)

// processInfo stores process information
type ProcessInfo struct {
    pid int
    cpu float64
    mem float64 
    command string
}

// ByCPU implements sort.Interface for []ProcessInfo based on the cpu field
type ByCPU []ProcessInfo 

func (a ByCPU) Len() int { return len(a) }
func (a ByCPU) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByCPU) Less(i, j int) bool { return a[i].cpu > a[j].cpu }

func getProcessInfo() ([]ProcessInfo, error) {
    // run ps command to get process info
    out, err := exec.Command("ps", "--no-headers", "-eo", "pid,pcpu,pmem,command").Output()
    if err != nil {
        return nil, err
    }

    lines := strings.Split(string(out), "\n")

    var processes []ProcessInfo
    for _, line := range lines {
        if line == "" {
            continue 
        }
        parts := strings.Fields(line)

        pid, _ := strconv.Atoi(parts[0])
        cpu, _ := strconv.ParseFloat(parts[1], 64)
        mem, _ := strconv.ParseFloat(parts[2], 64)

        processes = append(processes, ProcessInfo {
            pid: pid,
            cpu: cpu,
            mem: mem,
            command: strings.Join(parts[3:], " "),
        })
    }   

    return processes, nil
}

func printProcessTable(processes []ProcessInfo) {
    fmt.Printf(" PID\tCPU%%\tMEM%%\tCOMMAND\n")
    for _, proc := range processes {
       fmt.Printf("%d\t%.2f\t%.2f\t%s\n", proc.pid, proc.cpu, proc.mem, proc.command) 
    }
}

func main() {
    for { 
        processes, err := getProcessInfo()
        if err != nil {
            fmt.Printf("Error getting process info: %v\n", err)
            os.Exit(1)
        }

        // clear screen for each refresh
        print("\033[H\033[2J")

        // sort by CPU usage
        sort.Sort(ByCPU(processes))

        printProcessTable(processes)

        time.Sleep(time.Second)
    }
}

