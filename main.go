package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type CompileRequest struct {
	Code     string `json:"code"`
	Compiler string `json:"compiler"` // "gcc", "clang"
	Mode     string `json:"mode"`     // "preprocess", "ir", "asm", "obj", "bin", "hex"
}

type CompileResponse struct {
	Output string `json:"output"`
	Error  string `json:"error,omitempty"`
}

func compileHandler(w http.ResponseWriter, r *http.Request) {
	var req CompileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	tmpFile, err := ioutil.TempFile("", "*.c")
	if err != nil {
		http.Error(w, "could not create temp file", 500)
		return
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.WriteString(req.Code)
	tmpFile.Close()

	var cmd *exec.Cmd
	switch req.Compiler {
	case "gcc":
		switch req.Mode {
		case "preprocess":
			cmd = exec.Command("gcc", "-E", tmpFile.Name(), "-o", "-")
		case "asm":
			cmd = exec.Command("gcc", "-S", tmpFile.Name(), "-o", "-")
		case "bin":
			cmd = exec.Command("gcc", "-c", tmpFile.Name(), "-o", "-")
		default:
			http.Error(w, "unsupported mode", http.StatusBadRequest)
			return
		}
	case "clang":
		switch req.Mode {
		case "preprocess":
			cmd = exec.Command("clang", "-E", tmpFile.Name(), "-o", "-")
		case "asm":
			cmd = exec.Command("clang", "-S", tmpFile.Name(), "-o", "-")
		case "llvm ir":
			cmd = exec.Command("clang", "-S", "-emit-llvm", tmpFile.Name(), "-o", "-")
		case "bitcode":
			cmd = exec.Command("clang", "-c", "-emit-llvm", tmpFile.Name(), "-o", "-")
		case "bin":
			cmd = exec.Command("clang", "-c", tmpFile.Name(), "-o", "-")
		default:
			http.Error(w, "unsupported mode", http.StatusBadRequest)
			return
		}
	}

	output, err := cmd.CombinedOutput()
	res := CompileResponse{Output: string(output)}
	if err != nil {
		res.Error = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/compile", compileHandler)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
