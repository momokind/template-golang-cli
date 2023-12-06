package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Total struct {
	XMLName xml.Name `xml:"Total"`
	Pass    int      `xml:"pass,attr"`
	Fail    int      `xml:"fail,attr"`
	Total   int      `xml:"total,attr"`
}

type ExecutedTestsDetails struct {
	XMLName xml.Name `xml:"ExecutedTestsDetails"`
	Total   Total    `xml:"Total"`
}

type CvgInfo struct {
	XMLName xml.Name `xml:"CvgInfo"`
	Elem    string   `xml:"elem,attr"`
	Num     int      `xml:"num,attr"`
	Total   int      `xml:"total,attr"`
	Val     int      `xml:"val,attr"`
}

type CvgStats struct {
	XMLName xml.Name  `xml:"CvgStats"`
	CvgInfo []CvgInfo `xml:"CvgInfo"`
}

type Coverage struct {
	XMLName  xml.Name `xml:"Coverage"`
	CvgStats CvgStats `xml:"CvgStats"`
}

type Exec struct {
	XMLName  xml.Name `xml:"Exec"`
	Coverage Coverage `xml:"Coverage"`
}

type ResultsSession struct {
	XMLName              xml.Name             `xml:"ResultsSession"`
	ExecutedTestsDetails ExecutedTestsDetails `xml:"ExecutedTestsDetails"`
	Exec                 Exec                 `xml:"Exec"`
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
		fmt.Printf("无法读取文件 %s: %v\n", filename, err)
		return
	}

	// 解析 XML 文件
	var resultsSession ResultsSession
	err = xml.Unmarshal(xmlFile, &resultsSession)
	if err != nil {
		fmt.Println("无法解析 XML:", err)
		return
	}

	// 打印单元测试的汇总信息
	fmt.Printf("TEST EXECUTION:\n")
	fmt.Printf("Passed Tests: %d\n", resultsSession.ExecutedTestsDetails.Total.Pass)
	fmt.Printf("Failed Tests: %d\n", resultsSession.ExecutedTestsDetails.Total.Fail)
	fmt.Printf("Total Tests: %d\n", resultsSession.ExecutedTestsDetails.Total.Total)

	// 打印代码覆盖率的汇总信息
	fmt.Printf("\nCoverage Summary:\n")
	if len(resultsSession.Exec.Coverage.CvgStats.CvgInfo) > 0 {
		cvgInfo := resultsSession.Exec.Coverage.CvgStats.CvgInfo[0]
		coveragePercentage := 0.0
		if cvgInfo.Total > 0 {
			coveragePercentage = float64(cvgInfo.Num) / float64(cvgInfo.Total) * 100
		}
		fmt.Printf("Covered Elements: %d\n", cvgInfo.Num)
		fmt.Printf("Total Elements: %d\n", cvgInfo.Total)
		fmt.Printf("Coverage Percentage: %.2f%%\n", coveragePercentage)
	} else {
		fmt.Println("No coverage information available.")
	}

}
