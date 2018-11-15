package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	//"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

//var inFlag = flag.String("in", "", "Input file")
//var outFlag = flag.String("out", "", "Output file")

func main() {
	flag.Parse()
	//	fmt.Println(*inFlag, *outFlag)
	//Open file
	f, _ := os.Open("/home/ivan/bill_temp/91.214.205.34_2018_08.flw")
	//f, _ := os.Open("/home/ivan/ttt/test.flw")
	scanner := bufio.NewScanner(f)
	file, _ := os.Create("result.csv")
	writer := csv.NewWriter(file)

	//Flags

	for scanner.Scan() {
		line := scanner.Text()

		//Split line
		parts := strings.Split(line, "|")
		//fmt.Println(parts)
		//fmt.Println(stringToTime(parts[1]))
		r, _ := stringToTime(parts[1])
		date := r.Format("02.01.2006 15:04:05")
		//fmt.Println(date)

		outcsv := []string{date, parts[3], parts[4], protoName(parts[9]), parts[10], parts[11], parts[12]}

		writer.Write(outcsv)
	}

	writer.Flush()
	file.Close()
}

func stringToTime(s string) (time.Time, error) {
	sec, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(sec, 0), nil
}

func protoName(p string) string {
	switch p {
	case "1":
		return "ICMP"
	case "6":
		return "TCP"
	case "17":
		return "UDP"
	case "47":
		return "PPTP data over GRE"
	case "51":
		return "AH"
	case "50":
		return "ESP"
	case "8":
		return "EGP"
	case "3":
		return "GGP"
	case "20":
		return "HMP"
	case "88":
		return "IGMP"
	case "66":
		return "RVD"
	case "89":
		return "OSPF Open Shortest Path First"
	case "12":
		return "PUP"
	case "27":
		return "RDP"
	case "46":
		return "RSVP"
	default:
		return p
	}

}
