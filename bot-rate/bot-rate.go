package main

import (
    "fmt"
    "regexp"
    "encoding/json"
    "github.com/PuerkitoBio/goquery"
)

func GetPage(url string) {
    data := make(map[string][]string)
    doc, _ := goquery.NewDocument(url)
    doc.Find("div#slice1 table").Eq(2).Find("tr").Each(func(i int, tr *goquery.Selection) {
        if i>1 {
            country := ""
            tr.Find("td").Each(func(j int, td *goquery.Selection) {
                txt := td.Text()
                if j==0 {
                    country = regexp.MustCompile(`\([A-Z]+\)`).FindString(txt)
                    country = country[1:4]
                    data[country] = []string{}
                } else if j<5 {
                    data[country] = append(data[country], txt)
                }
            })
        }
    })

    jsonString, err := json.Marshal(data)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(string(jsonString))
    }
}

func main() {
    GetPage("http://rate.bot.com.tw/Pages/Static/UIP003.zh-TW.htm")
}
