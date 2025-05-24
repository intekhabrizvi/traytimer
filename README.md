
![Frame 2](https://github.com/user-attachments/assets/25d682ab-f2ca-487c-bccb-3bf4176b2137)

# TrayTimer — Simple System Tray-based Countdown & Stopwatch App

> A tiny system tray utility that helps you track time effortlessly with countdowns ⏱️ and stopwatches ⏲️ — complete with native desktop notifications 🎉.


## 🚀 Features

- 🖥️ **Supported-platform**: Linux.
- 🔔 **Native notifications**: Uses OS-specific tools (like `notify-send`, `kdialog`) for alerts.
- 🔁 **Reset, pause, and resume**: Right-click tray menu to control the timer.
- 📌 **Modes**:
  - Countdown ⏳
  - Stopwatch ⏱️

## 🎯 Demo
[traytimer.webm](https://github.com/user-attachments/assets/5916b0eb-f676-4081-9931-9a644ee2cc6d)


## 📦 Installation

### 📥 1. Download Precompiled Binaries

👉 [Releases Page](https://github.com/intekhabrizvi/traytimer/releases)

Download the binary, extract it, and run:

```bash
chmod +x traytimer
./traytimer --duration 15m
```


### 🐧 2. Ubuntu `.deb` Installation

```bash
wget https://github.com/intekhabrizvi/traytimer/releases/download/v1.0.0/traytimer_1.0.0_amd64.deb
sudo apt install ./traytimer_1.0.0_amd64.deb
```
Then launch with:
```bash
traytimer --duration 15m
```



## 🧪 Usage Examples

### Basic countdown for 15 minutes:

```bash
traytimer --duration 15m
```

### Basic stopwatch:

```bash
traytimer --mode stopwatch
```

### Run in Background

Use `setsid` for CLI:

```bash
setsid traytimer --duration 15m &
```

Or launch it from the application menu if installed via `.deb`.

---

## 📋 Command-line Options

| Flag         | Description                                  | Default |
|--------------|----------------------------------------------|---------|
| `--duration` | Duration for countdown (e.g. `5m`, `90s`)     | `15m`   |
| `--mode`     | Clock mode: `timer`, `stopwatch`    | `timer`  |

---

## 🖼️ Tray Menu Options

- ▶️ **Pause / Resume**
- 🔁 **Reset**
- ⏲️ **5min / 15min / 30min / 45min / 60min / 90min / 120min shortcuts**
- 🚪 **Quit**

---

## 💻 Developer Setup

### 1. Install Go

If you don’t already have Go installed:

```bash
sudo apt install golang-go  # Ubuntu
```

Or visit [golang.org](https://golang.org/dl/)

### 2. Clone the repository

```bash
git clone https://github.com/intekhabrizvi/traytimer.git
cd traytimer
```

### 3. Install Dependency

```bash
sudo apt-get install gcc libgtk-3-dev libayatana-appindicator3-dev
```

### 4. Build

```bash
go build -o traytimer .
```

You should now see the `traytimer` binary in the current directory.

---

## 🤝 Contributing

We ❤️ contributions! Here's how to help:

### 1. Fork & Clone

```bash
git clone https://github.com/intekhabrizvi/traytimer.git
cd traytimer
```

### 2. Create a feature branch

```bash
git checkout -b feature/my-improvement
```

### 3. Make your changes & commit

```bash
git add .
git commit -m "Added awesome feature"
```

### 4. Push and open PR

```bash
git push origin feature/my-improvement
```

Then open a Pull Request on GitHub 🚀

---

## 🧰 Dev Tools Used

- 🛠️ [Go](https://golang.org)
- 📦 [systray](https://github.com/getlantern/systray)
- 💡 Desktop-specific notification tools
- 🖼️ [3d icons](https://www.figma.com/community/file/999547190555074906) - Embedded icons


---

## 🗃️ Packaging for Other Distros

### 📦 RPM (Fedora, CentOS)

Use `fpm` or `rpmbuild` to package:
```bash
fpm -s dir -t rpm -n traytimer -v 1.0.0 traytimer=/usr/local/bin/traytimer
```

---

## 📝 License

MIT License © [Intekhab rizvi](https://github.com/intekhabrizvi)

---

## 💬 Feedback

Found a bug? Want a new feature? [Open an issue](https://github.com/intekhabrizvi/traytimer/issues) or start a discussion!

---

✨ **Happy timing!**
