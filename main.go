package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/getlantern/systray"
)

//go:embed icon.png
var icon []byte

var (
	durationFlag = flag.String("duration", "15m", "Countdown duration (e.g., 90s, 2m, 1h10m)")
	modeFlag     = flag.String("mode", "timer", "Mode: timer or stopwatch")

	countdown time.Duration
	remaining time.Duration
	ticker    *time.Ticker
	paused    bool = false
	isRunning bool = true
)

func main() {
	flag.Parse()
	// Show help if no arguments are passed
	if len(os.Args) == 1 {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	var err error
	if *modeFlag == "stopwatch" {
		countdown = 0
	} else {
		countdown, err = time.ParseDuration(*durationFlag)
	}

	if err != nil {
		fmt.Println("Invalid duration:", err)
		os.Exit(1)
	}
	remaining = countdown

	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTooltip("Countdown Timer")
	systray.SetTemplateIcon(icon, icon)
	systray.SetTitle(formatDuration(remaining))

	pauseResume := systray.AddMenuItem("Pause", "Pause or resume the countdown")
	reset := systray.AddMenuItem("Restart", "Reset the countdown")
	m5 := systray.AddMenuItem("Start for 5min", "Restart for 5min")
	m15 := systray.AddMenuItem("Start for 15min", "Restart for 15min")
	m30 := systray.AddMenuItem("Start for 30min", "Restart for 5min")
	m45 := systray.AddMenuItem("Start for 45min", "Restart for 45min")
	m60 := systray.AddMenuItem("Start for 60min", "Restart for 60min")
	m90 := systray.AddMenuItem("Start for 90min", "Restart for 90min")
	m120 := systray.AddMenuItem("Start for 120min", "Restart for 120min")
	quit := systray.AddMenuItem("Quit", "Exit the app")

	ticker = time.NewTicker(1 * time.Second)

	go func() {
		for range ticker.C {
			if paused || !isRunning {
				continue
			}

			if *modeFlag == "timer" {
				if remaining > 0 {
					remaining -= time.Second
					systray.SetTitle(formatDuration(remaining))
				} else {
					ticker.Stop()
					isRunning = false
					systray.SetTitle("Done!")
					sendNotification("Countdown Finished", "Time's up!")
				}
			} else if *modeFlag == "stopwatch" {
				remaining += time.Second
				systray.SetTitle(formatDuration(remaining))
			}
		}
	}()

	go func() {
		for {
			select {
			case <-pauseResume.ClickedCh:
				paused = !paused
				if paused {
					pauseResume.SetTitle("Resume")
					systray.SetTooltip("Resume")
				} else {
					pauseResume.SetTitle("Pause")
					systray.SetTooltip("Pause")
				}

			case <-reset.ClickedCh:
				remaining = countdown
				ticker = time.NewTicker(1 * time.Second)
				ticker.Stop()
			case <-m5.ClickedCh:
				remaining, _ = time.ParseDuration("5m")
			case <-m15.ClickedCh:
				remaining, _ = time.ParseDuration("15m")
			case <-m30.ClickedCh:
				remaining, _ = time.ParseDuration("30m")
			case <-m45.ClickedCh:
				remaining, _ = time.ParseDuration("45m")
			case <-m60.ClickedCh:
				remaining, _ = time.ParseDuration("60m")
			case <-m90.ClickedCh:
				remaining, _ = time.ParseDuration("90m")
			case <-m120.ClickedCh:
				remaining, _ = time.ParseDuration("120m")
			case <-quit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func onExit() {
	if ticker != nil {
		ticker.Stop()
	}
}

func formatDuration(d time.Duration) string {
	m := int(d.Minutes())
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d", m, s)
}

func sendNotification(title, message string) {
	de := getLinuxDesktopEnv()
	switch {
	case strings.Contains(de, "kde"):
		// Use kdialog
		exec.Command("kdialog", "--title", title, "--passivepopup", message, "5").Run()
	case strings.Contains(de, "xfce"), strings.Contains(de, "gnome"), strings.Contains(de, "unity"), strings.Contains(de, "mate"):
		exec.Command("notify-send", title, message).Run()
	default:
		// fallback
		exec.Command("notify-send", title, message).Run()
	}
}

func getLinuxDesktopEnv() string {
	envs := []string{
		"XDG_CURRENT_DESKTOP",
		"DESKTOP_SESSION",
		"GNOME_DESKTOP_SESSION_ID",
		"KDE_FULL_SESSION",
	}
	for _, env := range envs {
		if val := os.Getenv(env); val != "" {
			return strings.ToLower(val)
		}
	}
	return "unknown"
}
