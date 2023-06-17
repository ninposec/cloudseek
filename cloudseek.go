package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	flag.Usage = func() {
		fmt.Println(`
		
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
																																																   
			
		`)
		fmt.Printf("cloudseek v.0.1\n")
		fmt.Printf("Author: ninposec\n")
		fmt.Println("")
		fmt.Println("Checks if IPs is hosted on Microsoft Azure, Amazon AWS, Google Cloud, or DigitalOcean.")
		fmt.Println("Provide IPs as a list in a file or through stdin.")
		fmt.Println("")
		flag.PrintDefaults()
	}

	var inputFile string
	flag.StringVar(&inputFile, "ips", "", "Path to file ro read IPs from")
	var inputHelp bool
	flag.BoolVar(&inputHelp, "h", false, "Display usage")
	flag.Parse()

	if inputHelp == true || flag.Arg(0) == "-h" {
		flag.Usage()
		os.Exit(0)
	}

	var ips []string

	if inputFile != "" {
		// Read from input file
		file, err := os.Open(inputFile)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ips = append(ips, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			os.Exit(1)
		}
	} else {
		// Read from stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			ips = append(ips, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			os.Exit(1)
		}
	}

	if len(ips) == 0 {
		fmt.Println("Please provide at least one IP address.")
		os.Exit(1)
	}

	for _, ip := range ips {
		ip = strings.TrimSpace(ip)
		if isAzure(ip) {
			fmt.Printf("%s belongs to Microsoft Azure\n", ip)
		} else if isAWS(ip) {
			fmt.Printf("%s belongs to Amazon AWS\n", ip)
		} else if isGCP(ip) {
			fmt.Printf("%s belongs to Google Cloud\n", ip)
		} else if isDigitalOcean(ip) {
			fmt.Printf("%s belongs to DigitalOcean\n", ip)
		} else {
			//fmt.Printf("%s does not belong to any of the supported providers\n", ip)
		}
	}
}

func isAzure(ip string) bool {
	_, azure, _ := net.ParseCIDR("13.64.0.0/11")
	return azure.Contains(net.ParseIP(ip))
}

func isAWS(ip string) bool {
	_, aws1, _ := net.ParseCIDR("52.95.245.0/24")
	_, aws2, _ := net.ParseCIDR("18.208.0.0/13")
	_, aws3, _ := net.ParseCIDR("52.47.139.64/26")
	_, aws4, _ := net.ParseCIDR("54.144.0.0/14")
	_, aws5, _ := net.ParseCIDR("52.46.0.0/18")
	_, aws6, _ := net.ParseCIDR("35.154.0.0/16")
	return aws1.Contains(net.ParseIP(ip)) || aws2.Contains(net.ParseIP(ip)) || aws3.Contains(net.ParseIP(ip)) || aws4.Contains(net.ParseIP(ip)) || aws5.Contains(net.ParseIP(ip)) || aws6.Contains(net.ParseIP(ip))
}

func isGCP(ip string) bool {
	_, gcp1, _ := net.ParseCIDR("34.64.0.0/11")
	_, gcp2, _ := net.ParseCIDR("35.184.0.0/14")
	_, gcp3, _ := net.ParseCIDR("35.192.0.0/12")
	_, gcp4, _ := net.ParseCIDR("35.208.0.0/13")
	_, gcp5, _ := net.ParseCIDR("35.216.0.0/14")
	_, gcp6, _ := net.ParseCIDR("35.220.0.0/16")
	return gcp1.Contains(net.ParseIP(ip)) || gcp2.Contains(net.ParseIP(ip)) || gcp3.Contains(net.ParseIP(ip)) || gcp4.Contains(net.ParseIP(ip)) || gcp5.Contains(net.ParseIP(ip)) || gcp6.Contains(net.ParseIP(ip))
}

func isDigitalOcean(ip string) bool {
	_, do1, _ := net.ParseCIDR("174.138.0.0/17")
	_, do2, _ := net.ParseCIDR("157.230.0.0/16")
	_, do3, _ := net.ParseCIDR("138.197.0.0/16")
	return do1.Contains(net.ParseIP(ip)) || do2.Contains(net.ParseIP(ip)) || do3.Contains(net.ParseIP(ip))
}

func usage() {
	fmt.Printf("Usage: %s [OPTIONS]\n\n", os.Args[0])
	fmt.Println("cloudseek v.0.1\n\nAuthor: ninposec\n\n")
	fmt.Println("  Reads a list of IP addresses from stdin or from an input file and prints")
	fmt.Println("  which IP addresses belong to Microsoft Azure, Amazon AWS, Google Cloud,")
	fmt.Println("  or DigitalOcean.")
	fmt.Println()
	fmt.Println("Options:")
	flag.PrintDefaults()
}
