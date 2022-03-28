package main

import (
	"Dict/utils"
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
)

var (
	parentId int
	// 单词正则
	wordRe = regexp.MustCompile(`\n[a-zA-Z'\s\-]+[^<\s@]+`)
	// 基础词正则
	phraseRe = regexp.MustCompile(`</>\n([^\n]+)\n@@@LINK=([^\n]+)\n</>`)
	// 英式发音音标正则
	pronEnRe = regexp.MustCompile(
		`pronunciation[&#10;\s]+English"[^>]+>\s</a><span class="phon">([^<]+)</span>`)
	// 美式发音音标正则
	pronAmRe = regexp.MustCompile(
		`pronunciation[&#10;\s]+American[^>]+>\s</a><span class="phon">([^<]+)</span>`)
	// 单词释义正则
	transCnFromWordRe = regexp.MustCompile(`<defT><chn>([^<]+)</chn></defT>`)
	// 词性正则
	grammarRe = regexp.MustCompile(
		`<span class="grammar" [^>]+>([^<]+)</span>`)
)

// main is Program entry.
func main() {
	file := utils.OpenFile("./dict_file/牛津高阶（第10版 英汉双解） V12_1.txt")
	defer file.Close()
	b := bufio.NewReaderSize(bufio.NewReader(file), 2048)
	for {
		// 从文件中读取整个单词信息块
		bytes, err := utils.GetWordBlock(b)
		if err == io.EOF {
			break
		}
		wordMatch := wordRe.FindSubmatch(bytes)
		if len(wordMatch) > 0 { // 匹配到单词
			word := strings.TrimSpace(string(wordMatch[0]))
			log.Println("word:", word)
		}

		//fmt.Printf("\n\n+++++++++++++++++++++++++++\n%s\n+++++++++++++++++++++++++++\n\n",
		//	string(bytes))

		phraseMatch := phraseRe.FindSubmatch(bytes)
		if len(phraseMatch) > 0 { // 匹配到短语
			phrase := strings.TrimSpace(string(phraseMatch[1]))
			phraseSource := strings.TrimSpace(string(phraseMatch[2]))
			log.Println("phrase:", phrase)
			log.Println("phraseSource:", phraseSource)
		}

		pronEnRe := pronEnRe.FindSubmatch(bytes)
		if len(pronEnRe) > 0 { // 匹配到英式发音
			pronEn := strings.TrimSpace(string(pronEnRe[1]))
			log.Println("pronEn:", pronEn)
		}

		pronAmRe := pronAmRe.FindSubmatch(bytes)
		if len(pronAmRe) > 0 { // 匹配到美式发音
			pronAm := strings.TrimSpace(string(pronAmRe[1]))
			log.Println("pronAm:", pronAm)
		}

		transCnFromWordMatch := transCnFromWordRe.FindSubmatch(bytes)
		if len(transCnFromWordMatch) > 0 { // 匹配到单词释义
			transCnFromWord := strings.TrimSpace(string(transCnFromWordMatch[1]))
			log.Println("transCnFromWord:", transCnFromWord)
		}

		grammarMatch := grammarRe.FindSubmatch(bytes)
		if len(grammarMatch) > 0 { // 匹配到词性
			grammar := strings.TrimSpace(string(grammarMatch[1]))
			log.Println("grammar:", grammar)
		}

		fmt.Println("\n")
	}
}
