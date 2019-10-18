package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	seccomp "github.com/elastic/go-seccomp-bpf"
)

var (
	policyFile string
	noNewPrivs bool
)

func main() {
	flag.BoolVar(&noNewPrivs, "no-new-privs", true, "set no new privs bit")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "You must specify a command and args to execute.\n")
		os.Exit(1)
	}

	var policy = &seccomp.Policy{
		DefaultAction: seccomp.ActionAllow,
		Syscalls: []seccomp.SyscallGroup{
			{
				Action: seccomp.ActionErrno,
				Names: []string{
					"connect",
					"accept",
					"sendto",
					"recvfrom",
					"sendmsg",
					"recvmsg",
					"bind",
					"listen",
					"getpid",
					"kill",
				},
			},
		},
	}

	// Create a filter based on config.
	filter := seccomp.Filter{
		NoNewPrivs: noNewPrivs,
		Flag:       seccomp.FilterFlagTSync,
		Policy:     *policy,
	}

	// Load the BPF filter using the seccomp system call.
	if err := seccomp.LoadFilter(filter); err != nil {
		fmt.Fprintf(os.Stderr, "error loading filter: %v\n", err)
		os.Exit(1)
	}

	// Execute the specified command (requires execve).
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
