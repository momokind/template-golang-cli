package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Project struct {
	XMLName      xml.Name `xml:"Project"`
	CheckedFiles int      `xml:"checkedFiles,attr"`
	TotalFiles   int      `xml:"totFiles,attr"`
	TotalLines   int      `xml:"totLns,attr"`
	TotalErrors  int      `xml:"totErrs,attr"`
}

type CodingStandards struct {
	XMLName xml.Name `xml:"CodingStandards"`
	Project Project  `xml:"Projects>Project"`
}

type ResultsSession struct {
	XMLName         xml.Name        `xml:"ResultsSession"`
	CodingStandards CodingStandards `xml:"CodingStandards"`
}

func main() {
	// Check if filename is provided
	if len(os.Args) < 2 {
		fmt.Println("请提供XML文件名作为参数")
		return
	}

	filename := os.Args[1]

	// 从文件中读取 XML 内容
	xmlFile, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error! 无法读取文件 %s: %v\n", filename, err)
		return
	}

	// 解析 XML 文件
	var resultsSession ResultsSession
	err = xml.Unmarshal(xmlFile, &resultsSession)
	if err != nil {
		fmt.Println("Error! 无法解析 XML:", err)
		return
	}

	// 打印 STATIC ANALYSIS 的汇总信息
	fmt.Printf("Checked Files: %d, Total Files: %d, Total Lines: %d, Total Errors: %d\n", resultsSession.CodingStandards.Project.CheckedFiles, resultsSession.CodingStandards.Project.TotalFiles, resultsSession.CodingStandards.Project.TotalLines, resultsSession.CodingStandards.Project.TotalErrors)
}
