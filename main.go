package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jacolate/go-pks/sequential"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: gopks [seq|threaded] <filepath>")
        os.Exit(1)
    }

    filePath := os.Args[2]
    
    // Verify that the path exists
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        fmt.Printf("Error: Path '%s' does not exist\n", filePath)
        os.Exit(1)
    }

    startTime := time.Now()

    switch os.Args[1] {
    case "seq":
        fmt.Println("Starting sequential processing...")
        hist, err := sequential.Start(filePath)
        if err != nil {
            fmt.Printf("Error during processing: %v\n", err)
            os.Exit(1)
        }
        
        duration := time.Since(startTime)
        fmt.Println("\n\nResults:")
        fmt.Printf("Time taken: %v\n", duration)
        fmt.Printf("Directories processed: %d\n", hist.Directories)
        fmt.Printf("Total files found: %d\n", hist.Files)
        fmt.Printf("Text files processed: %d\n", hist.ProcessedFiles)
        fmt.Printf("Total lines processed: %d\n", hist.Lines)
        
        // Print character distribution (optional)
        fmt.Println("\nCharacter distribution (non-zero counts):")

        fmt.Println(hist)

    case "threaded":
        fmt.Println("Threaded processing not implemented yet")
        fmt.Println("File Path:", filePath)

    default:
        fmt.Println("Unknown processing mode. Use 'seq' or 'threaded'")
        os.Exit(1)
    }
}
