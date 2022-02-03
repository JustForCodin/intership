package treejsonified

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func TreeToJSON(root string) error {
	tree := make(map[string][]string)

	file, err := os.Stat(root)
	if err != nil {
		return fmt.Errorf("Failed to read directory %s: %v", root, err)
	}

	tree[file.Name()] = append(tree[file.Name()], file.Name())
	// fmt.Println(tree)

	if !file.IsDir() {
		return nil
	}

	files, err := ioutil.ReadDir(root)
	if err != nil {
		return fmt.Errorf("Failed to read directory %s: %v", root, err)
	}

	for _, file := range files {
		if string(file.Name()[0]) == string(".") {
			continue
		}

		if !file.IsDir() {
			tree[file.Name()] = []string{}
		} else {
			subfiles, _ := ioutil.ReadDir(file.Name())

			for _, sf := range subfiles {
				tree[file.Name()] = append(tree[file.Name()], sf.Name())
			}
		}

		if err := TreeToJSON(filepath.Join(root, file.Name())); err != nil {
			return err
		}
	}

	j, err := json.MarshalIndent(tree, "", "\t")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(j))
	}

	return nil
}
