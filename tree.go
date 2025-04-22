package main

import (
    "fmt"
    "os"
    "path/filepath"
    "sort"
)

func walk(path string, prefix string) {
    entries, err := os.ReadDir(path)
    if err != nil {
        fmt.Printf("%s[Permission Denied]\n", prefix)
        return
    }

    // 정렬: 디렉토리 먼저 나오게 하려면 여기 커스텀 정렬 넣을 수도 있음
    sort.Slice(entries, func(i, j int) bool {
        return entries[i].Name() < entries[j].Name()
    })

    for i, entry := range entries {
        connector := "├── "
        if i == len(entries)-1 {
            connector = "└── "
        }

        name := entry.Name()
        fmt.Printf("%s%s%s\n", prefix, connector, name)

        if entry.IsDir() {
            newPrefix := prefix
            if i == len(entries)-1 {
                newPrefix += "    "
            } else {
                newPrefix += "│   "
            }
            walk(filepath.Join(path, name), newPrefix)
        }
    }
}

func main() {
    root := "."
    if len(os.Args) > 1 {
        root = os.Args[1]
    }

    fmt.Printf("📁 Directory Tree for: %s\n\n", root)
    walk(root, "")
}

