package main

import (
	"common"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
	"service"
	"syscall"
)

func daemon(nochdir, noclose int) int {
	var ret, ret2 uintptr
	var err syscall.Errno

	darwin := runtime.GOOS == "darwin"

	// already a daemon
	if syscall.Getppid() == 1 {
		return 0
	}
	// fork off the parent process
	ret, ret2, err = syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)
	if err != 0 {
		return -1
	}

	// failure
	if ret2 < 0 {
		os.Exit(-1)
	}

	// handle exception for darwin
	if darwin && ret2 == 1 {
		ret = 0
	}

	// if we got a good PID, then we call exit the parent process.
	if ret > 0 {
		os.Exit(0)
	}

	/* Change the file mode mask */
	_ = syscall.Umask(0)

	// create a new SID for the child process
	s_ret, s_errno := syscall.Setsid()
	if s_errno != nil {
		log.Printf("Error: syscall.Setsid errno: %d", s_errno)
	}
	if s_ret < 0 {
		return -1
	}

	if nochdir == 0 {
		os.Chdir("/")
	}

	if noclose == 0 {
		f, e := os.OpenFile("/dev/null", os.O_RDWR, 0)
		if e == nil {
			fd := f.Fd()
			syscall.Dup2(int(fd), int(os.Stdin.Fd()))
			syscall.Dup2(int(fd), int(os.Stdout.Fd()))
			syscall.Dup2(int(fd), int(os.Stderr.Fd()))
		}
	}

	return 0
}

func initWebHandlers() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		service.Dispatch(w, r)
	})
}

func main() {
	port := flag.Int64("port", 8080, "Specify web server port")
	host := flag.String("host", "0.0.0.0", "Specify web server host")
	flag.Parse()
	daemon(1, 1)
	common.InitLogger()

	initWebHandlers()
	if l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port)); nil != err {
		fmt.Printf("Failed to launch app http server for reason:%v\n", err)
	} else {
		pid := os.Getpid()
		ioutil.WriteFile("server.pid", []byte(fmt.Sprintf("%d", pid)), 0666)
		http.Serve(l, nil)
	}
}
