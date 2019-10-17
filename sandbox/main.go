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
	// flag.StringVar(&policyFile, "policy", "seccomp.yml", "seccomp policy file")
	flag.BoolVar(&noNewPrivs, "no-new-privs", true, "set no new privs bit")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "You must specify a command and args to execute.\n")
		os.Exit(1)
	}

	// Load policy from file.
	// // policy, err := parsePolicy()
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	// 	os.Exit(1)
	// }
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
					"writev",
					"getpid",
					// RELRO
					// "accept",
					// "accept4",
					// "arch_prctl",
					// "bind",
					// "brk",
					// "chdir",
					// "chroot",
					// "clone",
					// "close",
					// "connect",
					// "dup",
					// "dup2",
					// "epoll_create",
					// "epoll_create1",
					// "epoll_ctl",
					// "epoll_wait",
					// "execve",
					// "exit",
					// "exit_group",
					// "fchdir",
					// "fchmod",
					// "fchown",
					// "fcntl",
					// "fstat",
					// "fsync",
					// "ftruncate",
					// "futex",
					// "getcwd",
					// "getdents64",
					// "geteuid",
					// "getgid",
					// "getpeername",
					// "getpid",
					// "getppid",
					// "getrandom",
					// "getrusage",
					// "getsockname",
					// "getsockopt",
					// "gettid",
					// "getuid",
					// "ioctl",
					// "kill",
					// "listen",
					// "lseek",
					// "lstat",
					// "madvise",
					// "mincore",
					// "mkdirat",
					// "mmap",
					// "mount",
					// "munmap",
					// "open",
					// "openat",
					// "pipe",
					// "pipe2",
					// "prctl",
					// "pread64",
					// "pselect6",
					// "ptrace",
					// "pwrite64",
					// "read",
					// "readlinkat",
					// "recvfrom",
					// "recvmsg",
					// "renameat",
					// "rt_sigaction",
					// "rt_sigprocmask",
					// "rt_sigreturn",
					// "sched_getaffinity",
					// "sched_yield",
					// "sendfile",
					// "sendmsg",
					// "sendto",
					// "setgid",
					// "setgroups",
					// "setitimer",
					// "setpgid",
					// "setsid",
					// "setsockopt",
					// "setuid",
					// "shutdown",
					// "sigaltstack",
					// "socket",
					// "stat",
					// "statfs",
					// "sysinfo",
					// "tkill",
					// "unlinkat",
					// "unshare",
					// "wait4",
					// "waitid",
					// "write",
					// "writev",
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
