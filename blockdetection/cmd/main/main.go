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
	startParenthesis,  _ := regexp.Match("^\w*|\s*(\{)\w*|\s*$", []byte(ngBlock.StartLine))
	endParenthesis,  _ := regexp.Match("^\w*|\s*(\})\w*|\s*$", []byte(ngBlock.EndLine))
	if startParenthesis && endParenthesis {
		return true
	} else {
		return false
	}
}
//
func (ngBlock *NginxBlock) IsLine(line string) bool {
	// TODO Solve it using regex
	chk,  _ := regexp.Match(".*\n$", []byte(ngBlock.StartLine))
	if chk {
		return true
	} else {
		return false
	}
}

func (ngBlock *NginxBlock) HasComment(line string) bool {
	chk,  _ := regexp.Match("^\#.*", []byte(ngBlock.StartLine))
	if chk {
		return true
	} else {
		return false
	}
	// TODO Solve it using regex
}

type NginxBlocks struct {
	blocks      *[]*NginxBlock
	AllContents string
	// split lines by \n on AllContents
	AllLines *[]*string
}
//
func GetNginxBlock(
	lines *[]*string,
	startIndex,
	endIndex,
	recursionMax int,
) *NginxBlock {
	return GetNginxBlock(lines, 0, 0,  recursionMax+1)
}

func GetNginxBlocks(configContent string) *NginxBlocks {
	//var lines *[]*string := &configContent
	temp := strings.Split(configContent, "\n")
	var temp2 []*string
	for _, item := range temp {
		temp2 = append(temp2, &item)
	}
	GetNginxBlock( &temp2, 0, 0, 0)
	var obj NginxBlocks
	return &obj
}


func main() {
	data, _ := ioutil.ReadFile("nginx.config")
	file := string(data)
	fmt.Println(file)
	Ngblock := GetNginxBlocks(file)
	fmt.Println(Ngblock)
}
