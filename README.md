# gowttr

A simple command line app written in go, that shows you the current weather. The app uses the data from [wttr.in](https://wttr.in/) and is just a simple app for me to learn go.

![gowttr usage](https://f.j4n.net/gowttr.gif)

## 📦 Installation

### Build it yourself

You can install gowttr in two ways: using an install script or building it manually.

### 🔧 Method 1: Install via Script

Run this to download the latest release and install it:

```bash
curl -sL https://gowttr.j4n.net/install | bash
```

This will:

- [x] Download the latest compiled binary from GitHub releases
- [x] Make it executable
- [x] Move it to /usr/local/bin/ so you can run gowttr from anywhere

After installation, you can check if it works by running:

```bash
gowttr -h
```

### 🛠️ Method 2: Build from Source

If you prefer to build gowttr yourself, follow these steps:

#### 1. Install Go (if not already installed)

```bash
sudo apt install golang  # Debian/Ubuntu  
brew install go          # macOS  
```

#### 2. Clone the repository

```bash
git clone https://github.com/KartoffelChipss/gowttr.git
cd gowttr
```

#### 3. Build the binary

```bash
go build -o gowttr .
```

#### 4. Move the binary to your PATH

```bash
chmod +x gowttr
sudo mv gowttr /usr/local/bin/
```

#### 5. Verify the installation

```bash
gowttr --help
```