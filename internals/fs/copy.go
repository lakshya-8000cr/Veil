package fs

import (
	"io"
	"os"
	"path/filepath"
)

//in this feature we are doing likw , finding the source file of changed then replacing that from the upperdir
// also we ae not using the os.readFile or write file bcz in future if we have the so much file then all file will get load in the ram eventually which is not good
//insted we are using the io.Copy , it reads the data in chunks at a time like 32kb etc   

func CopyFile(src string, dst string) error {

	                                                               // Open source file
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	                                                               // Ensure destination directory exists
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	                                                                // Create destination file
	dest, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dest.Close()

	                                                                // Copy contents
	_, err = io.Copy(dest, source)
	if err != nil {
		return err
	}

	return nil
}