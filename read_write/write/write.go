package w

import "os"

func WriteFileAll(filename, content string) error {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func WriteFile(filename, content string) error {
	// os.O_WRONLY	write only
	// os.O_CREATE	create if not existed
	// os.O_RDONLY	read only
	// os.O_RDWR	read and write only
	// os.O_TRUNC	clear
	// os.O_APPEND	append
	// os.Create(filename)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644) // cover
	if err != nil {
		return err
	}
	defer file.Close()
	// file.Write([]byte(content))
	if _, err := file.WriteString(content); err != nil {
		return err
	}
	return nil
}
