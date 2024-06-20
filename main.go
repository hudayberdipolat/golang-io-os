package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if err := createNewFile(); err != nil {
		log.Fatal(err)
	}
	if err := createNewDirectory(); err != nil {
		log.Fatal(err)
	}
	if err := createNewDirectoryIfNotExists(); err != nil {
		log.Fatal(err)
	}
	if err := copyTextOtherFile(); err != nil {
		log.Println()
	}
}

func createNewFile() error {
	newFile, err := os.Create("file1.txt")
	if err != nil {
		log.Println(err)
		return err
	}
	defer newFile.Close()
	return nil
}

func createNewDirectory() error {
	if err := os.Mkdir("./images", 0777); err != nil {
		return err
	}
	return nil
}

func createNewDirectoryIfNotExists() error {
	if _, err := os.Stat("./images"); os.IsNotExist(err) {
		if errCreateFolder := os.Mkdir("./images", 0777); errCreateFolder != nil {
			return errCreateFolder
		}
		return err
	}
	return nil
}

func copyTextOtherFile() error {
	file, err := os.Open("text.txt")
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()
	writeFile, err := os.Create("file.txt")
	if err != nil {
		log.Println(err)
		return err
	}
	defer writeFile.Close()
	_, err = io.Copy(writeFile, file)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("test")
	return nil
}
