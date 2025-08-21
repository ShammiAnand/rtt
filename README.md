# rtt

`rtt` is a cli application which allows you to convert a repository of code/files
and webpages into flat files (both `txt` and `md` formats are supported) and 
interact with LLMs all using one CLI

## Installation

currently only supports macOS & Linux distros

```bash
curl -sfL https://raw.githubusercontent.com/shammianand/rtt/main/install.sh | sh
```

```bash
# Add to your shell config (~/.zshrc or ~/.bashrc):
export GROQ_API_KEY=sk_xxxx 
```

## Usage

### Convert Local Directory
```bash
# Convert current directory to rtt.md
rtt

# Convert current directory to rtt.md (explicit)
rtt .

# Convert directory to custom output file
rtt /path/to/dir -o output.md

# Skip cache files and non-code files
rtt /path/to/dir --skip-cache -o clean.md

# Create optimized text file (minimal formatting, saves tokens)
rtt /path/to/dir --optimized-text -o optimized.txt

# Combine both options for maximum efficiency
rtt /path/to/dir --skip-cache --optimized-text -o efficient.txt
```

### Download Web Pages
```bash
# Save webpage as markdown
rtt url https://example.com -o page.md
```

### Query Content
```bash
# Query current directory
rtt query . "Explain this codebase"

# Query webpage
rtt query url https://example.com "Summarize this article"
```

## Features

- **Current Directory Support**: Run `rtt` without arguments to process the current directory
- **Smart File Filtering**: Use `--skip-cache` to automatically skip cache files, build artifacts, and non-code files
- **Optimized Text Output**: Use `--optimized-text` to create minimal-formatting text files that save tokens for LLM processing
- **Comprehensive File Support**: Processes code files, text files, markdown, configuration files, and more

## File Filtering

When using `--skip-cache`, the tool automatically skips:
- Cache files (`.cache`, `.tmp`, `.log`, etc.)
- Build artifacts (`node_modules`, `target`, `build`, `dist`)
- System files (`.DS_Store`, `Thumbs.db`)
- Binary and non-text files

## Output Formats

- **Markdown** (default): Rich formatting with syntax highlighting and file separators
- **Optimized Text**: Minimal formatting, no extra spaces, ideal for token-efficient LLM processing

By default, queries use the `Mixtral-8x7b` model with a 32k context window.

For more details:
```bash
rtt help
rtt query --help
```

## Author
[@shammianand](https://www.github.com/shammianand)

