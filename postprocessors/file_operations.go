package postprocessors

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func RenderFilePath(targetPath string) string {
	// prepend a "/" on linux
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		targetPath = "/" + targetPath
	}
	// Handle ":" drive on windows
	targetPath = strings.Replace(targetPath, "?", ":", 1)
	return targetPath
}

func FindTargetFileCaseInsensitive(targetPath string) string {
	targetPath = RenderFilePath(targetPath)
	targetFolder, targetFile := filepath.Split(targetPath)
	fileInfos, err := os.ReadDir(targetFolder)
	if err != nil {
		return ""
	}

	for _, info := range fileInfos {
		if strings.ToLower(info.Name()) == strings.ToLower(targetFile) {
			targetFile = info.Name()
			break
		}
	}
	targetPath = filepath.Join(targetFolder, targetFile)
	return targetPath
}

func AppendTargetFile(targetPath string, appendLines []string) error {
	//fmt.Fprintf(os.Stderr, "Target file: %s\n", targetPath)
	fileBytes, err := os.ReadFile(targetPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fileBytes = []byte{}
		} else {
			return err
		}
	}
	fileLines := strings.Split(string(fileBytes), "\n")
	newLines := append(fileLines, appendLines...)

	outputString := strings.Join(newLines, "\n")
	//fmt.Fprintln(os.Stderr, outputString)
	err = os.WriteFile(targetPath, []byte(outputString), 0777)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTargetFile(targetPath string, templateLines []string) error {
	//fmt.Fprintf(os.Stderr, "Target file: %s\n", targetPath)
	fileBytes, err := os.ReadFile(targetPath)
	if err != nil {
		return err
	}

	firstLine := strings.TrimSpace(templateLines[0])
	lastLine := strings.TrimSpace(templateLines[len(templateLines)-1])

	//fmt.Fprintln(os.Stderr, firstLine, lastLine)

	fileLines := strings.Split(string(fileBytes), "\n")
	newLines := make([]string, 0)

	for idx, line := range fileLines {
		testLine := strings.TrimSpace(line)
		//fmt.Fprintln(os.Stderr, testLine+" :: "+firstLine)
		if strings.HasPrefix(testLine, firstLine) {
			//fmt.Fprintln(os.Stderr, fileLines[0:idx])
			newLines = append(newLines, fileLines[0:idx]...)
			newLines = append(newLines, templateLines...)
		} else if strings.HasPrefix(testLine, lastLine) {
			//fmt.Fprintln(os.Stderr, fileLines[idx+1:])
			newLines = append(newLines, fileLines[idx+1:]...)
		}
	}

	// TODO - Refactor this behavior?
	if path.Ext(targetPath) == ".md" {
		//fmt.Fprintf(os.Stderr, "Correcting ``` escapement in file:%s\n", targetPath)
		// Handle the ``` issue
		for idx, line := range newLines {
			newLines[idx] = strings.ReplaceAll(line, "'''", "```")
		}
	}

	// If we didn't find anything, don't write anything
	if len(newLines) == 0 {
		fmt.Fprintln(os.Stderr, "No target comments found for file: "+targetPath)
		return nil
	}

	outputString := strings.Join(newLines, "\n")
	//fmt.Fprintln(os.Stderr, outputString)
	err = os.WriteFile(targetPath, []byte(outputString), 0777)
	if err != nil {
		return err
	}

	return nil
}
