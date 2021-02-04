package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type NginxBlock struct {
	StartLine   string
	EndLine     string
	AllContents string
	// split lines by \n on AllContents,
	// use make to create *[],
	// first create make([]*Type..)
	// then use &var to make it *
	AllLines          *[]*string
	NestedBlocks      []*NginxBlock
	TotalBlocksInside int
}

func (ngBlock *NginxBlock) IsBlock(line string) bool {
	// TODO Solve it using regex
	startParenthesis,  err := regexp.Match("^\w*|\s*(\{)\w*|\s*$", []byte(ngBlock.StartLine))
	endParenthesis,  err := regexp.Match("^\w*|\s*(\})\w*|\s*$", []byte(ngBlock.EndLine))
}
//
//func (ngBlock *NginxBlock) IsLine(line string) bool {
//	// TODO Solve it using regex
//}
//
//func (ngBlock *NginxBlock) HasComment(line string) bool {
//	// TODO Solve it using regex
//}

type NginxBlocks struct {
	blocks      *[]*NginxBlock
	AllContents string
	// split lines by \n on AllContents
	AllLines *[]*string
}

//func GetNginxBlock(
//	lines *[]*string,
//	startIndex,
//	endIndex,
//	recursionMax int,
//) *NginxBlock {
//
//}
//
//func GetNginxBlocks(configContent string) *NginxBlocks {
//
//}


func main() {
	data, _ := ioutil.ReadFile("nginx.config")
	file := string(data)
	line := 1
	var lines []*string
	temp := strings.Split(file, "\n")
	for _, item := range temp {
		//fmt.Println("[", line, "]\t", item)
		line++
		lines = append(lines, &item)
	}
	for _, item := range lines{
		fmt.Println(item)

	}
}
