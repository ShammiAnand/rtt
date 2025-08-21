package walker

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/shammianand/rtt/utils/logger"
)

// FileFilter determines which files should be processed
type FileFilter struct {
	SkipCacheFiles bool
	OptimizedText  bool
}

// WalkAndExtract walks the given path and extracts the contents to the output file
func WalkAndExtract(walkPath string, outputPath string, filter FileFilter) error {

	var buf []byte
	
	// Use directory name for the title
	dirName := filepath.Base(walkPath)
	if dirName == "." {
		dirName = filepath.Base(GetCurrentDir())
	}
	
	if filter.OptimizedText {
		buf = append(buf, []byte(fmt.Sprintf("%s\n\n", dirName))...)
	} else {
		buf = append(buf, []byte(fmt.Sprintf("# %s\n\n", dirName))...)
	}

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

			// Skip cache files and non-code/text files
			if filter.SkipCacheFiles && shouldSkipFile(path) {
				logger.Log.Debugf("Skipping cache/non-code file: %s", path)
				return nil
			}

			content, err := os.ReadFile(path)
			if err != nil {
				logger.Log.Warnf("reading file %q: %v", path, err)
				return nil
			}
			logger.Log.Debugf("File: %s | Size: %d bytes ", path, len(content))

			if filter.OptimizedText {
				// Optimized text format: no extra spaces, minimal formatting
				buf = append(buf, []byte(fmt.Sprintf("FILE: %s\n", path))...)
				buf = append(buf, content...)
				buf = append(buf, []byte("\n\n")...)
			} else {
				// Markdown format
				buf = append(buf, []byte(fmt.Sprintf("### %s\n```%s\n", path, getFileExtension(path)))...)
				buf = append(buf, content...)
				buf = append(buf, []byte("\n```\n\n---\n")...)
			}
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

// shouldSkipFile determines if a file should be skipped based on its extension or name
func shouldSkipFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	name := strings.ToLower(filepath.Base(path))
	
	// Skip common cache and build artifacts
	cacheExtensions := []string{
		".cache", ".tmp", ".temp", ".log", ".lock", ".pid",
		".swp", ".swo", ".bak", ".backup", ".orig",
	}
	
	cacheNames := []string{
		"node_modules", ".git", ".svn", ".hg", ".bzr",
		"__pycache__", ".pytest_cache", ".mypy_cache",
		"target", "build", "dist", "out", "bin", "obj",
		".DS_Store", "Thumbs.db", "desktop.ini",
	}
	
	// Check extensions
	for _, cacheExt := range cacheExtensions {
		if ext == cacheExt {
			return true
		}
	}
	
	// Check names
	for _, cacheName := range cacheNames {
		if strings.Contains(name, cacheName) {
			return true
		}
	}
	
	// Only process code, text, and markdown files
	allowedExtensions := []string{
		".go", ".py", ".js", ".ts", ".jsx", ".tsx", ".java", ".c", ".cpp", ".h", ".hpp",
		".cs", ".php", ".rb", ".rs", ".swift", ".kt", ".scala", ".clj", ".hs", ".ml",
		".sql", ".sh", ".bash", ".zsh", ".fish", ".ps1", ".bat", ".cmd",
		".html", ".htm", ".css", ".scss", ".sass", ".less", ".xml", ".json", ".yaml", ".yml",
		".toml", ".ini", ".cfg", ".conf", ".config", ".env", ".properties",
		".md", ".markdown", ".txt", ".text", ".rst", ".adoc", ".tex",
		".dockerfile", ".dockerignore", ".gitignore", ".gitattributes",
		".makefile", ".cmake", ".cmakelists", ".gradle", ".maven", ".pom",
		".gemfile", ".package", ".composer", ".cargo", ".mix", ".rebar",
		".vim", ".emacs", ".vscode", ".idea", ".xcode", ".sublime",
	}
	
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			return false
		}
	}
	
	// Skip files without extensions unless they're common text files
	if ext == "" {
		// Allow common text files without extensions
		textFiles := []string{"makefile", "dockerfile", "readme", "license", "changelog", "contributing"}
		for _, textFile := range textFiles {
			if strings.Contains(name, textFile) {
				return false
			}
		}
		return true
	}
	
	return true
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

// GetCurrentDir returns the current working directory
func GetCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		logger.Log.Error(err)
		return "."
	}
	return dir
}
