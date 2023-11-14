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

func DirListing(wordlist, target string, delay int) {
	if !strings.HasSuffix(target, "/"){
		target+= "/"
	}
	initLogger()
	l.Println("[PARAMS] --target", target ,"--wordlist", wordlist, "--delay", delay)
	l.Println("[INFO] Directory listing attack started for", target)
	fmt.Println(Yellow("[INFO]"), "Directory listing attack started for", Green(target))

	file, err := os.Open(wordlist)
	if err != nil {
		fmt.Println(Red("[ERROR] ", err))
	}
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
				fmt.Println(Red("[ERROR] " ,err))
				break
			}
		
		defer response.Body.Close()

		if response.StatusCode == 200 {
			fmt.Println("")
			fmt.Print("\033[F")
			fmt.Print("\033[K")
			doc, err := goquery.NewDocumentFromReader(response.Body)
			if err != nil {
				fmt.Println(Red("[ERROR] " ,err))
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
