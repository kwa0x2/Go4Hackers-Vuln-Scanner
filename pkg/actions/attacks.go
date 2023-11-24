package actions

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/cheggaaa/pb/v3"
	"github.com/fatih/color"
)

var (
	l *log.Logger
)

func checkError(err error) {
	if err != nil {
		fmt.Println(Red("[ERROR] ", err))
	}
}

func initLogger() {
	logsFolderPath := "logs/"

	if err := os.Mkdir(logsFolderPath, os.ModePerm); err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	dateFormat := time.Now().Format("2006-01-02-15-04-05")

	logFileName := fmt.Sprintf("%s%s", logsFolderPath, dateFormat)

	outfile, err := os.Create(logFileName)
	if err != nil {
		log.Fatal(err)
	}
	l = log.New(outfile, "", log.LstdFlags|log.Lshortfile)
}

var (
	Red    = color.New(color.FgRed).SprintFunc()
	Yellow = color.New(color.FgYellow).SprintFunc()
	Green  = color.New(color.FgGreen).SprintFunc()
	Blue   = color.New(color.FgBlue).SprintFunc()
)

func CheckXFrameOptions(target string) {
	initLogger()
	l.Println("[INFO] X-Frame-Options header passive attack started for", target)

	resp, err := http.Get(target)
	checkError(err)
	if resp.StatusCode > 400 {
		l.Println("[NOT VULN] Cannot connect to the", target)
		fmt.Println(Red("[WARN]"), "Cannot connect to the", Green(target))
	} else if resp.StatusCode == http.StatusOK {
		x_frame_options := resp.Header.Get("X-Frame-Options")

		if x_frame_options != "" {
			l.Println("[VULN] X-Frame-Options header implemented for", target)
			fmt.Println(Blue("[VULN]"), "[VULN] X-Frame-Options header implemented for", Green(target))
		} else {
			l.Println("[VULN] X-Frame-Options header not implemented for", target)
			fmt.Println(Red("[NOT VULN]"), "[VULN] X-Frame-Options header not implemented for", Green(target))
		}

	}

}

func CheckTrace(target string) {
	initLogger()
	l.Println("[INFO] HTTP Trace attack started for", target)

	client := &http.Client{}

	req, err := http.NewRequest("TRACE", target, nil)
	checkError(err)

	resp, err := client.Do(req)
	checkError(err)

	defer resp.Body.Close()

	if resp.StatusCode > 400 {
		l.Println("[NOT VULN] HTTP TRACE method is not allowed for", target)
		fmt.Println(Red("[NOT VULN]"), "HTTP TRACE method is not allowed for", Green(target))
	} else if resp.StatusCode == http.StatusOK {
		l.Println("[VULN] HTTP TRACE method is allowed for for", target)
		fmt.Println(Blue("[VULN]"), "HTTP TRACE method is allowed for", Green(target))
	} else {
		fmt.Println(Red("Unexpected response status:", resp.Status))
	}
}

func CheckDirListing(wordlist, target string, delay int) {
	if !strings.HasSuffix(target, "/") {
		target += "/"
	}
	initLogger()
	l.Println("[INFO] Directory listing attack started for", target)
	fmt.Println(Yellow("[INFO]"), "Directory listing attack started for", Green(target))

	file, err := os.Open(wordlist)
	checkError(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	tmpl := `{{ red "Scanning:" }} {{ bar . "<" "-" (cycle . "↖" "↗" "↘" "↙" ) "." ">"}} {{speed . | rndcolor }} {{percent .}}`

	bar := pb.ProgressBarTemplate(tmpl).Start64(int64(len(lines)))
	bar.SetWidth(100)

	var totalVuln int

	for _, line := range lines {
		bar.Increment()
		time.Sleep(time.Duration(delay) * time.Millisecond)
		url := target + line
		response, err := http.Get(url)
		if err != nil {
			fmt.Println(Red("[ERROR] ", err))
			break
		}

		defer response.Body.Close()

		if response.StatusCode == 200 {
			fmt.Println("")
			fmt.Print("\033[F")
			fmt.Print("\033[K")
			doc, err := goquery.NewDocumentFromReader(response.Body)
			if err != nil {
				fmt.Println(Red("[ERROR] ", err))
				continue
			}

			title := doc.Find("title").Text()
			titleLower := strings.ToLower(title)
			if strings.Contains(titleLower, "index of") || strings.Contains(titleLower, "directory list") {
				totalVuln++

				l.Println("[VULN] Directory listing vulnerability detected for", url)
				fmt.Println(Blue("[VULN]"), "Directory listing vulnerability detected for", Green(url))
			}
		}
	}
	bar.Finish()
	fmt.Print("\033[F")
	fmt.Print("\033[K")
	if err == nil {
		if totalVuln == 0 {
			l.Println("[NOT VULN] No directory listing vulnerability found for", target)
			fmt.Println(Red("[NOT VULN]"), "No directory listing vulnerability found for", Green(target))
		} else {
			l.Println("[FINISH] Total number of directories with listings:", totalVuln)
			fmt.Println(Green("[FINISH]"), "Total number of directories with listings:", totalVuln)
		}
	}

}
