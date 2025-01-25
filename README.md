# RTT (Repo to Text)

CLI tool to convert repository contents into a single text file.

## Installation

```bash
# Via Go
go install github.com/shammianand/rtt@latest

# Via Binary
curl -sfL https://raw.githubusercontent.com/shammianand/rtt/main/install.sh | sh

# Manual
wget https://github.com/shammianand/rtt/releases/latest/download/rtt_$(uname)_$(uname -m).tar.gz
tar xvf rtt_*.tar.gz
sudo mv rtt /usr/local/bin/
```

## Usage

```bash
# Convert current directory
rtt .

# Convert specific directory
rtt /path/to/directory

# Specify output file
rtt -o output.txt /path/to/directory

# Set author
rtt -a "Your Name" /path/to/directory
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -am 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

Apache License 2.0
