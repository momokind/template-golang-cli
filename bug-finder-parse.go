package main

import (
	"fmt"
	"strings"

	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Check if filename is provided
	if len(os.Args) < 2 {
		fmt.Println("Error! 请提供HTML文件名作为参数。")
		return
	}

	filename := os.Args[1]

	// 从文件中读取 HTML 内容
	html, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error! 无法读取文件。")
		return
	}

	// 创建 GoQuery 文档对象
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(html)))
	if err != nil {
		fmt.Println("Error! 无法解析 HTML:")
		return
	}

	// 解析 "Project Summary" 表格数据
	doc.Find("p.rgTableTitle").Each(func(i int, s *goquery.Selection) {
		if s.Find("span.rgTableTitleText").Text() == "Project Summary" {
			table := s.Next()
			table.Find("tbody tr").Each(func(j int, row *goquery.Selection) {
				rowData := []string{}
				row.Find("td p span").Each(func(k int, cell *goquery.Selection) {
					rowData = append(rowData, cell.Text())
				})
				fmt.Printf("Total Count:%s, Reviewed:%s, Unreviewed:%s, Pass/Fail:%s\n", rowData[1], rowData[2], rowData[3], rowData[4])
			})
		}
	})
}
