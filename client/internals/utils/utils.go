package utils

import (
	"fmt"
	"io"
)

func Copy(dst io.Writer, src io.Reader) error {
	var input = make([]byte, 1024)
	n, err := src.Read(input)
	if err != nil {
		return fmt.Errorf("utils.Copy: Error during read: %w", err)
	}
	_, err = dst.Write(input[:n])
	if err != nil {
		return fmt.Errorf("utils.Copy: Error during write: %w", err)
	}
	return nil
}
