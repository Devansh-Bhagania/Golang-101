
# Installing Go (Golang)

This guide will walk you through the installation of Go (Golang) on Windows, Linux, and macOS.

## System Requirements
- **Windows**: 64-bit version
- **macOS**: macOS 10.13 or higher
- **Linux**: 64-bit distribution

## 1. Installing Go on Windows

### Step 1: Download the Go Installer
- Go to the official Go download page: [https://go.dev/dl/](https://go.dev/dl/)
- Download the Windows `.msi` installer.

### Step 2: Run the Installer
- Double-click the downloaded `.msi` file.
- Follow the installation wizard prompts, keeping the default installation path (usually `C:\Go`).

### Step 3: Verify Installation
1. Open **Command Prompt** or **PowerShell**.
2. Run the following command to verify the installation:
   ```bash
   go version
   ```

You should see the Go version installed.

### Step 4: Set Up the Go Workspace
- By default, Go will use `C:\Users\YourUserName\go` as the workspace.
- Create a new directory for your Go projects if desired:
  ```bash
  mkdir %USERPROFILE%\go
  ```

## 2. Installing Go on Linux

### Step 1: Download Go Tarball
1. Open your terminal.
2. Download the latest tarball from the official website:
   ```bash
   wget https://go.dev/dl/go1.x.x.linux-amd64.tar.gz
   ```
   Replace `1.x.x` with the latest version.

### Step 2: Extract and Install Go
1. Extract the tarball to `/usr/local`:
   ```bash
   sudo tar -C /usr/local -xzf go1.x.x.linux-amd64.tar.gz
   ```

### Step 3: Set Up Go Path and Environment Variables
1. Add Go to your `PATH` by adding this to your `.bashrc` (or `.zshrc`):
   ```bash
   export PATH=$PATH:/usr/local/go/bin
   ```
2. Apply the changes:
   ```bash
   source ~/.bashrc
   ```

### Step 4: Verify Installation
1. Run the following command to verify the installation:
   ```bash
   go version
   ```

You should see the Go version installed.

### Step 5: Set Up the Go Workspace
1. Create a Go workspace directory:
   ```bash
   mkdir -p $HOME/go
   ```

2. Add the workspace directory to the Go environment:
   ```bash
   export GOPATH=$HOME/go
   ```

## 3. Installing Go on macOS

### Step 1: Download Go Installer
- Go to the official Go download page: [https://go.dev/dl/](https://go.dev/dl/)
- Download the `.pkg` installer for macOS.

### Step 2: Run the Installer
- Open the downloaded `.pkg` file and follow the installation instructions.

### Step 3: Verify Installation
1. Open **Terminal**.
2. Run the following command to verify the installation:
   ```bash
   go version
   ```

You should see the Go version installed.

### Step 4: Set Up Go Path and Environment Variables
1. Add Go to your `PATH` by adding this to your `.bash_profile` (or `.zshrc`):
   ```bash
   export PATH=$PATH:/usr/local/go/bin
   ```

2. Apply the changes:
   ```bash
   source ~/.bash_profile
   ```

### Step 5: Set Up the Go Workspace
1. Create a Go workspace directory:
   ```bash
   mkdir -p $HOME/go
   ```

2. Add the workspace directory to the Go environment:
   ```bash
   export GOPATH=$HOME/go
   ```

---

## Common Post-Installation Steps

### Verifying Installation
After installation on any system, you can verify by running:
```bash
go version
```

### Setting Up a Simple Go Program
1. Create a directory for your Go project:
   ```bash
   mkdir -p $HOME/go/src/hello
   cd $HOME/go/src/hello
   ```

2. Create a `hello.go` file:
   ```go
   package main

   import "fmt"

   func main() {
       fmt.Println("Hello, Go!")
   }
   ```

3. Run the program:
   ```bash
   go run hello.go
   ```

You should see the output:
```bash
Hello, Go!
```

---

For more information, visit the official [Go documentation](https://go.dev/doc/).
