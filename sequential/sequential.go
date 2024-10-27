package sequential

import (
    h "github.com/jacolate/go-pks/utils/histogram"
    "os"
    "path/filepath"
    "fmt"
    "io"
    "bufio"
    "bytes"
)

var counter = 0

func Start(path string) (h.Histogram, error) {
    hist := h.Histogram{
        Distribution: make([]int64, 26),
        Lines:        0,
        Files:        0,
        ProcessedFiles: 0,
        Directories:    0,
    }
    
    err := traverse(path, &hist)
    if err != nil {
        return hist, fmt.Errorf("error during traversal: %v", err)
    }
    return hist, nil
}

func traverse(filePath string, hist *h.Histogram) error {
    fileInfo, err := os.Stat(filePath)
    if err != nil {
        return fmt.Errorf("error getting file info: %v", err)
    }

    if fileInfo.IsDir() {
        hist.Directories++
        
        files, err := os.ReadDir(filePath)
        if err != nil {
            return fmt.Errorf("error reading directory: %v", err)
        }

        for _, file := range files {
            fullPath := filepath.Join(filePath, file.Name())
            err = traverse(fullPath, hist)
            if err != nil {
                return err
            }
        }
        fmt.Printf("%d - Directory %s finished !\n", counter, filePath)
        counter++
    } else {
        hist.Files++
        
        if filepath.Ext(filePath) == ".txt" {
            hist.ProcessedFiles++
            processFile(filePath, hist)
            if err != nil {
                return fmt.Errorf("error processing file %s: %v", filePath, err)
            }
        }
    }
    return nil
}

func processFile(filePath string, hist *h.Histogram) error {
    file, err := os.Open(filePath)
    if err != nil {
        return fmt.Errorf("error opening file: %v", err)
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    buffer := make([]byte, 1024)
    
    for {
        n, err := reader.Read(buffer)
        if err != nil {
            if err == io.EOF {
                break
            }
            return fmt.Errorf("error reading file: %v", err)
        }

        newlines := int64(bytes.Count(buffer[:n], []byte{'\n'}))
        hist.Lines += newlines

        for i := 0; i < n; i++ {
            char := buffer[i]
            
            if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
                if char >= 'A' && char <= 'Z' {
                    char = char + 32 
                }
                index := char - 'a'
                hist.Distribution[index]++
            }
        }
        fmt.Printf("%d - File %s finished !\n", counter, filePath)
        counter++
    }
    return nil
}
