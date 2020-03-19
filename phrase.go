package pinyin

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/yanyiwu/gojieba"
)

var (
	dictDir = path.Join(filepath.Dir(os.Args[0]), "configs/api/dict")
	jiebaPath = path.Join(dictDir, "jieba.dict.utf8")
	hmmPath = path.Join(dictDir, "hmm_model.utf8")
	userPath = path.Join(dictDir, "user.dict.utf8")
	idfPath = path.Join(dictDir, "idf.utf8")
	stopPath = path.Join(dictDir, "stop_words.utf8")
	jieba = gojieba.NewJieba(jiebaPath, hmmPath, userPath, idfPath, stopPath)
)

func cutWords(s string) []string {
	return jieba.CutAll(s)
}

func pinyinPhrase(s string) string {
	words := cutWords(s)
	for _, word := range words {
		match := phraseDict[word]
		if match == "" {
			match = phraseDictAddition[word]
		}

		match = toFixed(match, paragraphOption)
		if match != "" {
			s = strings.Replace(s, word, " "+match+" ", 1)
		}
	}

	return s
}

