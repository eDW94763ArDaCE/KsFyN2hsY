// 代码生成时间: 2025-08-13 00:00:43
package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"
)

// FolderOrganizer is a struct that holds the source and target directories.
type FolderOrganizer struct {
    SourceDir string
    TargetDir string
}

// NewFolderOrganizer creates a new instance of FolderOrganizer.
func NewFolderOrganizer(sourceDir, targetDir string) *FolderOrganizer {
    return &FolderOrganizer{
        SourceDir: sourceDir,
        TargetDir: targetDir,
    }
}

// Organize moves files from the source directory to the target directory according to their extension.
func (f *FolderOrganizer) Organize(ctx context.Context) error {
    // Read the source directory and get all files.
    files, err := ioutil.ReadDir(f.SourceDir)
    if err != nil {
        return fmt.Errorf("failed to read source directory: %w", err)
    }

    for _, file := range files {
        if file.IsDir() {
            // Skip subdirectories.
            continue
        }

        // Extract the file extension.
        extension := strings.TrimPrefix(filepath.Ext(file.Name()), ".")

        // Define the directory for the file based on its extension.
        targetSubDir := filepath.Join(f.TargetDir, extension)

        // Create the target subdirectory if it does not exist.
        if _, err := os.Stat(targetSubDir); os.IsNotExist(err) {
            if err := os.MkdirAll(targetSubDir, os.ModePerm); err != nil {
                return fmt.Errorf("failed to create target directory: %w", err)
            }
        }

        // Construct the full paths for source and target files.
        sourcePath := filepath.Join(f.SourceDir, file.Name())
        targetPath := filepath.Join(targetSubDir, file.Name())

        // Move the file.
        if err := os.Rename(sourcePath, targetPath); err != nil {
            return fmt.Errorf("failed to move file: %w", err)
        }
    }

    return nil
}

func main() {
    // Example usage of FolderOrganizer.
    ctx := context.Background()
    sourceDir := "./source"
    targetDir := "./target"
    organizer := NewFolderOrganizer(sourceDir, targetDir)

    if err := organizer.Organize(ctx); err != nil {
        log.Fatalf("error organizing folders: %v", err)
    }
    fmt.Println("Folder organization completed successfully.")
}