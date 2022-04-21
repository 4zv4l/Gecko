package main

import (
	"fmt"
	hash "gecko/hash"
	"os"

	"github.com/charmbracelet/lipgloss"
)

// show usage
func printUsage() {
	println("Usage:")
	fmt.Print(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("86")).Render(" gecko"))
	println("  [options] [file 1 file2 ...]")
	println("Options:")
	println("  -o <file>  write output to file (not append)")
	println("  -stdin     read from stdin")
	println("  -quiet     doesn't print the filename and the hash type")
	println("  -md5       show md5 checksum")
	println("  -sha1      show sha1 checksum")
	println("  -sha256    show sha256 checksum")
	println("  -b64       show base64 encoding")
	println("  -d64       show base64 decoding")
}

// parse flags
func parse() (out, quiet, showMD5, showSHA1, showSHA256, base64, dec64, stdin bool) {
	for _, arg := range os.Args[1:] {
		switch arg {
		case "-o":
			out = true
		case "-quiet":
			quiet = true
		case "-md5":
			showMD5 = true
		case "-sha1":
			showSHA1 = true
		case "-sha256":
			showSHA256 = true
		case "-b64":
			base64 = true
		case "-d64":
			dec64 = true
		case "-stdin":
			stdin = true
		default:
			// do nothing
		}
	}
	return out, quiet, showMD5, showSHA1, showSHA256, base64, dec64, stdin
}

// get files from arguments without the output file
func getFiles(args []string) []string {
	files := []string{}
	outputBefore := false
	for _, arg := range args {
		if arg == "-o" && len(files) != 0 {
			break
		}
		if arg == "-o" && len(files) == 0 {
			outputBefore = true
		}
		if arg[0] != '-' {
			files = append(files, arg)
		}
	}
	if outputBefore {
		files = files[1:]
	}
	return files
}

// check output file
func checkOutput(output bool) *os.File {
	var out string
	if output {
		// check if there is a filename after the -o flag
		for i, arg := range os.Args[1:] {
			if arg == "-o" {
				if i+3 > len(os.Args) {
					printUsage()
					os.Exit(1)
				}
				if os.Args[i+2][0] == '-' {
					printUsage()
					os.Exit(1)
				}
				out = os.Args[i+2]
				break
			}
		}
		f, err := os.OpenFile(out, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0700)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		return f
	}
	return os.Stdout
}

// show the asked hashes of a file
func showHash(stdin bool, files []string, out *os.File, quiet, showMD5, showSHA1, showSHA256, showB64, showD64 bool) {
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Println("Error opening file", file, ":", err)
			os.Exit(1)
		}
		defer f.Close()
		if !quiet {
			if out == os.Stdout {
				fmt.Fprintf(out, "%s\n", lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("86")).Render(file))
			} else {
				fmt.Fprintf(out, "%s\n", file)
			}
			if showMD5 {
				fmt.Fprintf(out, "MD5: %s\n", hash.GetMD5(f))
			}
			if showSHA1 {
				fmt.Fprintf(out, "SHA1: %s\n", hash.GetSHA1(f))
			}
			if showSHA256 {
				fmt.Fprintf(out, "SHA256: %s\n", hash.GetSHA256(f))
			}
			if showB64 {
				fmt.Fprintf(out, "B64: %s\n", hash.GetB64(f))
			}
			if showD64 {
				fmt.Fprintf(out, "D64: %s\n", hash.GetD64(f))
			}
		} else {
			if showMD5 {
				fmt.Fprintf(out, "%s\n", hash.GetMD5(f))
			}
			if showSHA1 {
				fmt.Fprintf(out, "%s\n", hash.GetSHA1(f))
			}
			if showSHA256 {
				fmt.Fprintf(out, "%s\n", hash.GetSHA256(f))
			}
			if showB64 {
				fmt.Fprintf(out, "%s", hash.GetB64(f))
			}
			if showD64 {
				fmt.Fprintf(out, "%s", hash.GetD64(f))
			}
		}
	}
	// if stdin is true, read from stdin
	if stdin {
		if !quiet {
			if showMD5 {
				fmt.Fprintf(out, "MD5: %s\n", hash.GetMD5(os.Stdin))
			}
			if showSHA1 {
				fmt.Fprintf(out, "SHA1: %s\n", hash.GetSHA1(os.Stdin))
			}
			if showSHA256 {
				fmt.Fprintf(out, "SHA256: %s\n", hash.GetSHA256(os.Stdin))
			}
			if showB64 {
				fmt.Fprintf(out, "B64: %s\n", hash.GetB64(os.Stdin))
			}
			if showD64 {
				fmt.Fprintf(out, "D64: %s\n", hash.GetD64(os.Stdin))
			}
		} else {
			if showMD5 {
				fmt.Fprintf(out, "%s\n", hash.GetMD5(os.Stdin))
			}
			if showSHA1 {
				fmt.Fprintf(out, "%s\n", hash.GetSHA1(os.Stdin))
			}
			if showSHA256 {
				fmt.Fprintf(out, "%s\n", hash.GetSHA256(os.Stdin))
			}
			if showB64 {
				fmt.Fprintf(out, "%s", hash.GetB64(os.Stdin))
			}
			if showD64 {
				fmt.Fprintf(out, "%s", hash.GetD64(os.Stdin))
			}
		}
	}
}

func main() {
	output, quiet, showMD5, showSHA1, showSHA256, showB64, showD64, stdin := parse()
	if !showMD5 && !showSHA1 && !showSHA256 && !showB64 && !showD64 {
		printUsage()
		return
	}
	f := checkOutput(output)
	defer f.Close()
	files := getFiles(os.Args[1:])
	if len(files) == 0 && !stdin {
		printUsage()
		return
	}
	showHash(stdin, files, f, quiet, showMD5, showSHA1, showSHA256, showB64, showD64)
}
