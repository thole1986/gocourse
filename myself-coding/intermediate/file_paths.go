package intermediate

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	relativePath := "./data/file.txt"
	absolutePath := "D:\\my-learning\\go-learning\\gocourse"

	// Join paths using filepath.join
	joinedPath := filepath.Join("downloads", "file.zip")

	fmt.Println("Joined Path: ", joinedPath)

	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println("Normalized Path: ", normalizedPath)

	dir, file := filepath.Split("D:\\my-learning\\go-learning\\gocourse\\file.txt")
	fmt.Println("File:", file)
	fmt.Println("Path:", dir)
	// fmt.Println(filepath.Base("D:\\my-learning\\go-learning\\gocourse\\file.txt"))
	fmt.Println(filepath.Base("D:\\my-learning\\go-learning\\gocourse"))

	fmt.Println("Is relativePath variable: ", filepath.IsAbs(relativePath))
	fmt.Println("Is absolutePath: ", filepath.IsAbs(absolutePath))

	// Get the extension of the file.
	fmt.Println(filepath.Ext(file))

	fmt.Println(filepath.Ext(file))
	fmt.Println(strings.TrimSuffix(file, filepath.Ext(file)))

	rel, err := filepath.Rel("a/c", "a/b/t/file")

	if err != nil {
		panic(err)
	}

	fmt.Println(rel)

	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Absolute path", absPath)
	}
}
