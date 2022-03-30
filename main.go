package main

import (
	"Dict/bootstrap"
	"Dict/model"
	"Dict/utils"
	"bufio"
	"fmt"
	"gorm.io/gorm"
	"io"
	"regexp"
	"strings"
)

var (
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

type WordDetail struct {
	word            string
	phrase          string
	phraseSource    string
	pronEn          string
	pronAm          string
	transCnFromWord string
	grammar         string
}

// main is Program entry.
func main() {
	bootstrap.Setup()
	file := utils.OpenFile("./dict_file/牛津高阶（第10版 英汉双解） V12_1.txt")
	//file := utils.OpenFile("./dict_file/example.txt")
	defer file.Close()
	b := bufio.NewReaderSize(bufio.NewReader(file), 2048)
	for {
		// 从文件中读取整个单词信息块
		bytes, err := utils.GetWordBlock(b)
		if err == io.EOF {
			break
		}
		wordDetail := matchWord(bytes)
		db := bootstrap.DB
		// ------------------------插入单词数据 -------------------------
		//if wordDetail.word != "" {
		//	tx := db.Table(model.TableDictWord).Create(&model.DictWordModel{
		//		ParentID: 0,
		//		Word:     wordDetail.word,
		//		PronEN:   wordDetail.pronEn,
		//		PronAM:   wordDetail.pronAm,
		//		Type:     0,
		//	})
		//	if err = tx.Error; err != nil {
		//		panic(err)
		//	}
		//}

		// ------------------------设置单词的基础词和翻译 -------------------------
		wordParent := model.DictWordModel{}
		tx := db.Table(model.TableDictWord).
			Where("binary word = ?", wordDetail.phraseSource).
			First(&wordParent)
		if err = tx.Error; err != nil && err != gorm.ErrRecordNotFound {
			panic(err)
		}
		if wordParent.ID != 0 {
			tx := db.Table(model.TableDictWord).
				Where("binary word = ?", wordDetail.word).
				Update("parent_id", wordParent.ID)
			if err = tx.Error; err != nil {
				panic(err)
			}
		}

		// ------------------------设置单词的释义 -------------------------
		//if wordDetail.transCnFromWord == "" {
		//	continue
		//}
		//wordParent := model.DictWordModel{}
		//tx := db.Table(model.TableDictWord).
		//	Where("binary word = ?", wordDetail.word).
		//	First(&wordParent)
		//if err = tx.Error; err != nil && err != gorm.ErrRecordNotFound {
		//	panic(err)
		//}
		//tx = db.Table(model.TableDictWordTransMap).
		//	Create(&model.DictWordTransMapModel{
		//		Wid:     wordParent.ID,
		//		TransCN: wordDetail.transCnFromWord,
		//		Part:    wordDetail.grammar,
		//	})
		//if err = tx.Error; err != nil {
		//	panic(err)
		//}
	}
}

func matchWord(content []byte) WordDetail {
	var wordDetail WordDetail
	// 匹配单词
	wordMatch := wordRe.FindSubmatch(content)
	if len(wordMatch) > 0 {
		wordDetail.word = strings.TrimSpace(string(wordMatch[0]))
	}
	// 匹配短语
	phraseMatch := phraseRe.FindSubmatch(content)
	if len(phraseMatch) > 0 {
		wordDetail.phrase = strings.TrimSpace(string(phraseMatch[1]))
		wordDetail.phraseSource = strings.TrimSpace(string(phraseMatch[2]))
	}
	// 匹配英式发音
	pronEnReMatch := pronEnRe.FindSubmatch(content)
	if len(pronEnReMatch) > 0 {
		wordDetail.pronEn = strings.TrimSpace(string(pronEnReMatch[1]))
	}
	// 匹配美式发音
	pronAmMatch := pronAmRe.FindSubmatch(content)
	if len(pronAmMatch) > 0 {
		wordDetail.pronAm = strings.TrimSpace(string(pronAmMatch[1]))
	}
	// 匹配单词释义
	transCnFromWordMatch := transCnFromWordRe.FindSubmatch(content)
	if len(transCnFromWordMatch) > 0 {
		wordDetail.transCnFromWord = strings.TrimSpace(string(transCnFromWordMatch[1]))
	}
	// 匹配词性
	grammarMatch := grammarRe.FindSubmatch(content)
	if len(grammarMatch) > 0 {
		wordDetail.grammar = strings.TrimSpace(string(grammarMatch[1]))
	}
	fmt.Println(wordDetail)
	return wordDetail
}
