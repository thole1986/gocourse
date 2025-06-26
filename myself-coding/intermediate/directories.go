package intemediate

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		panic(err)

	}
}

func main() {
	// The folder name need to create
	// and the permission for folder.
	// err := os.Mkdir("subdir", 0755)
	// checkError(err)

	// checkError(os.Mkdir("subdir", 0755))

	// defer os.RemoveAll("subdir")

	// os.WriteFile("subdir/file", []byte(""), 0755)

	// checkError(os.MkdirAll("subdir/parent/child", 0755))
	// checkError(os.MkdirAll("subdir/parent/child1", 0755))
	// checkError(os.MkdirAll("subdir/parent/child2", 0755))
	// checkError(os.MkdirAll("subdir/parent/child3", 0755))
	// os.WriteFile("subdir/parent/file", []byte(""), 0755)
	// os.WriteFile("subdir/parent/child/file", []byte(""), 0755)

	// Read Dir
	result, err := os.ReadDir("subdir/parent")
	checkError(err)

	for _, entry := range result {
		fmt.Println(entry)
	}

	checkError(os.Chdir("subdir/parent/child"))

	result, err = os.ReadDir(".")
	checkError(err)

	fmt.Println("Reading subdir/parent/child")
	for _, entry := range result {
		fmt.Println(entry)
	}

	checkError(os.Chdir("../../.."))
	dir, err := os.Getwd()

	checkError(err)
	fmt.Println(dir)

	// filepath.Walk and filepath.WalkDir
	// filepath.WalkDir is better than .Walk
	// for performance.

	pathfile := "subdir"

	fmt.Println("Walking Directory")
	err = filepath.WalkDir(pathfile, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error: ", err)
			return err
		}
		fmt.Println(path)
		return nil
	})

	checkError(err)

	// checkError(os.RemoveAll("./subdir")) // The name of directory to delete.
	checkError(os.Remove("subdir"))
}
