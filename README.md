# cloudseek
This is a simple tool that reads a list of IP addresses from stdin or an input file and prints which IP addresses belong to Microsoft Azure, Amazon AWS, Google Cloud, or DigitalOcean.

Can be used together with DNSX (https://github.com/projectdiscovery/dnsx), resolve domain names to IP and check if it belongs to a Cloud Provider e.g.

```bash
cat domains.txt | dnsx -ro | cloudseek
```

## Installation

```bash
go install github.com/ninposec/cloudseek@latest
```

## Usage
To use this tool, you can either provide a list of IP addresses via stdin or specify an input file with the -ips flag.

#### Reading from stdin
To provide a list of IP addresses via stdin, simply pipe them to cloudseek like this:

```bash
cat ips.txt | cloudseek
```

#### Reading from an input file
To specify an input file containing a list of IP addresses, use the -ips flag followed by the path to the input file, like this:

```bash
cloudseek -ips ips.txt
```

#### Help
To display the usage information, use the -h flag, like this:

```bash
./cloudseek -h

██████╗██╗      ██████╗ ██╗   ██╗██████╗
██╔════╝██║     ██╔═══██╗██║   ██║██╔══██╗
██║     ██║     ██║   ██║██║   ██║██║  ██║
██║     ██║     ██║   ██║██║   ██║██║  ██║
╚██████╗███████╗╚██████╔╝╚██████╔╝██████╔╝
 ╚═════╝╚══════╝ ╚═════╝  ╚═════╝ ╚═════╝
										
███████╗███████╗███████╗██╗  ██╗
██╔════╝██╔════╝██╔════╝██║ ██╔╝
███████╗█████╗  █████╗  █████╔╝
╚════██║██╔══╝  ██╔══╝  ██╔═██╗
███████║███████╗███████╗██║  ██╗
╚══════╝╚══════╝╚══════╝╚═╝  ╚═╝
															
			
		
cloudseek v.0.1
Author: ninposec

Checks if IPs is hosted on Microsoft Azure, Amazon AWS, Google Cloud, or DigitalOcean.
Provide IPs as a list in a file or through stdin.

  -h	Display usage
  -ips string
    	Path to file ro read IPs from
```