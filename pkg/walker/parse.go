package walker

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/shammianand/rtt/utils/logger"
)

// WalkAndExtract walks the given path and extracts the contents to the output file
func WalkAndExtract(walkPath string, outputPath string) error {

	var buf []byte
	buf = append(buf, []byte(fmt.Sprintf("# %s\n\n", walkPath))...)

	err := filepath.WalkDir(walkPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			logger.Log.Warnf("Error accessing path %q: %v", path, err)
			return nil
		}

		if d.Name()[0] == '.' {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if !d.IsDir() {

			if d.Type() == fs.ModeSymlink {
				logger.Log.Warnf("Skipping symlink: %s", path)
				return nil
			}

			content, err := os.ReadFile(path)
			if err != nil {
				logger.Log.Warnf("reading file %q: %v", path, err)
				return nil
			}
			logger.Log.Debugf("File: %s | Size: %d bytes ", path, len(content))

			buf = append(buf, []byte(fmt.Sprintf("### %s\n```%s\n", path, getFileExtension(path)))...)
			buf = append(buf, content...)
			buf = append(buf, []byte("\n```\n\n---\n")...)
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("Error walking directory: %v", err)
	}

	if err := os.WriteFile(outputPath, buf, 0644); err != nil {
		return fmt.Errorf("Error writing to file: %v", err)
	}

	logger.Log.Infof("successfully written to file: %s | size: %d bytes", outputPath, len(buf))

	return nil
}

func getFileExtension(path string) (ext string) {
	defer func() {
		if r := recover(); r != nil {
			logger.Log.Debugf("no file extension for %s", path)
			ext = "bash"
		}
	}()
	ext = filepath.Ext(path)[1:]
	return ext
}
