package main

import (
	"fmt"
	"github.com/nulpatrol/wordslearn/controllers/home"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	//"encoding/json"
)

type Subtitle struct {
	Captions []Caption
}

func main() {
	file, _ := ioutil.ReadFile("sub.srt")
	content := string(file)
	res, _ := NewFromSRT(content)

	db := dbConn()
	defer db.Close()

	m := map[string]int{}

	for _, caption := range res.Captions {
		for _, p := range caption.Text {
			words := strings.Split(p, " ")
			for _, word := range words {
				word = strings.Trim(word, "\".?!- ,'():")
				if word == "" {
					continue
				}
				word = strings.ToLower(word)

				matched, _ := regexp.Match(`[a-z]`, []byte(word))
				if !matched {
					continue
				}

				m[word]++
			}
		}
	}

	for _, query := range getQueryForWordsForms(m) {
		stmt, _ := db.Prepare(query.sql)
		defer stmt.Close()
		s := make([]interface{}, len(query.bindings))
		for i, v := range query.bindings {
    		s[i] = v
		}

		rows, err := stmt.Query(s...)
		if err != nil {
        	panic(err.Error())
    	}

		wordForm := WordForm{}
	    for rows.Next() {
	        err = rows.Scan(&wordForm.Id, &wordForm.WordId, &wordForm.Form)
	        if err != nil {
	            panic(err.Error())
	        }
	        
	        delete(m, wordForm.Form)
	    }
	}

	//jsonString, _ := json.MarshalIndent(m, "", "    ")
	//fmt.Println(string(jsonString))


    http.Handle("/", new(HomeHandler))
    http.ListenAndServe(":80", nil)
}

func NewFromSRT(s string) (res Subtitle, err error) {
	r1 := regexp.MustCompile("([0-9:.,]*) --> ([0-9:.,]*)")
	lines := strings.Split(s, "\n")
	outSeq := 1

	for i := 0; i < len(lines); i++ {
		seq := strings.Trim(lines[i], "\r ")
		if seq == "" {
			continue
		}

		_, err = strconv.Atoi(seq)
		if err != nil {
			err = fmt.Errorf("srt: atoi error at line %d: %v", i, err)
			break
		}

		var o Caption
		o.Seq = outSeq

		i++
		if i >= len(lines) {
			break
		}

		matches := r1.FindStringSubmatch(lines[i])
		if len(matches) < 3 {
			err = fmt.Errorf("srt: parse error at line %d (idx out of range) for input '%s'", i, lines[i])
			break
		}

		o.Start, err = parseSrtTime(matches[1])
		if err != nil {
			err = fmt.Errorf("srt: start error at line %d: %v", i, err)
			break
		}

		o.End, err = parseSrtTime(matches[2])
		if err != nil {
			err = fmt.Errorf("srt: end error at line %d: %v", i, err)
			break
		}

		i++
		if i >= len(lines) {
			break
		}

		textLine := 1
		for {
			line := strings.Trim(lines[i], "\r ")
			if line == "" && textLine > 1 {
				break
			}
			if line != "" {
				o.Text = append(o.Text, line)
			}

			i++
			if i >= len(lines) {
				break
			}

			textLine++
		}

		if len(o.Text) > 0 {
			res.Captions = append(res.Captions, o)
			outSeq++
		}
	}
	return
}
