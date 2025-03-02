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
rtt .

# Convert directory to custom output file
rtt /path/to/dir -o output.md
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

By default, queries use the `Mixtral-8x7b` model with a 32k context window.

For more details:
```bash
rtt help
rtt query --help
```

## Author
[@shammianand](https://www.github.com/shammianand)

