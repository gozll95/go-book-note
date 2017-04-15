package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// func NewCommandJob(id int, name string, command string) *Job {
// 	job := &Job{
// 		id:   id,
// 		name: name,
// 	}
// 	job.runFunc = func(timeout time.Duration) (string, string, error, bool) {
// 		bufOut := new(bytes.Buffer)
// 		bufErr := new(bytes.Buffer)
// 		cmd := exec.Command("/bin/bash", "-c", command)
// 		cmd.Stdout = bufOut
// 		cmd.Stderr = bufErr
// 		cmd.Start()
// 		err, isTimeout := runCmdWithTimeout(cmd, timeout)

// 		return bufOut.String(), bufErr.String(), err, isTimeout
// 	}
// 	return job
// }

func main() {
	var pid []string
	var ppid []string

	result := make(map[string][]string)

	out, err := exec.Command("/bin/bash", "-c", "ps -opid,ppid,command").Output()
	if err != nil {
		log.Fatal(err)
	}

	for _, value := range strings.Split(string(out), "\n")[1:] {
		if len(value) != 0 {
			//fmt.Println(strings.Fields(value))
			realOut := strings.Fields(value)
			pid = append(pid, realOut[0])
			ppid = append(ppid, realOut[1])
		}
	}
	//fmt.Println(pid)
	//fmt.Println(ppid)

	for iPpid, vPpid := range ppid {
		result[vPpid] = append(result[vPpid], pid[iPpid])
	}
	//fmt.Println(result)

	for i, v := range result {
		fmt.Printf("Pid %s has %d children: %s\n", i, len(v), v)
	}
}
