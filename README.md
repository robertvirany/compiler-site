# C Compiler Explorer (Flask Backend)

A minimal web-based compiler explorer for C code, built with Python (Flask). This app supports compilation of user-submitted C code using either `gcc` or `clang`, and provides output at different stages like preprocessing, assembly, LLVM IR, or binary object code.

## Features

- 🧠 Live compilation of C code via web UI
- 🛠️ Supports:
  - Preprocessing (`-E`)
  - Assembly output (`-S`)
  - LLVM IR (`clang`)
  - Binary object files (`-c`)
- 🔄 Output returned as plain text or hex-encoded binary
- 🔐 Server runs locally by default (localhost only)

## Technologies Used

- Python 3.12+
- Flask
- `gcc` and/or `clang` (must be installed on your system)

## Setup Instructions

### 1. Install Python dependencies

```bash
pip install -r requirements.txt
```

### 2. Run the server

```bash
python app.py
```
This starts the Flask server at `http://127.0.0.1:8080`


### 3. Visit in your browser
Open: `http://127.0.0.1:8080`. Use the interface to write C code, choose a compiler and mode, and view the results.

## Notes
- Object files are not executed or linked; binary output is returned as a hex string.
- Compilation errors are shown in the output pane.
- I wrote this on a Linux box, and it only compiles to ELF format. If you need PE check out `https://godbolt.org`. They also have a vast range of coding languages and other options for you to choose from.