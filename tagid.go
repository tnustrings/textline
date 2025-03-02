package main

import "fmt"
import "os"
import "encoding/json"
import "io"
import "regexp"
import "strings"


func Hello() string {
  message := fmt.Sprintf("hi")
  return message
}


// main runs Hello.
func main() {
  gtdictfile := "gt-scrape-dict.json"

  file, err := os.Open(gtdictfile)
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()
  bytes, _ := io.ReadAll(file)

  var dict map[string]interface{}
  json.Unmarshal([]byte(bytes), &dict)
  bytes, _ = io.ReadAll(os.Stdin)
  langa := os.Args[1]
  fmt.Print(Tagid(string(bytes), langa, dict))
  //fmt.Println(dict["de"].(map[string]interface{})["hallo"])


}


// Tagid puts word-by-word translations into text
func Tagid(text string, langa string, dict map[string]interface{}) string {
  wab := WordsAndBetween(text)
  out := ""
  hasLetter := regexp.MustCompile(`\p{L}`)
  for _, w := range wab {
    if !hasLetter.MatchString(w) {
      out += w
    } else {
      insert := ""
      insert = Lookup(dict, langa, w)
      out += w
      out += " [[" + insert + "]]"
      

    }
  }
  return out


}


// WordsAndBetween splits the text into words and non-words
func WordsAndBetween(s string)[]string {
  re := regexp.MustCompile(`[^\p{L}\p{M}]+`)
  return SplitKeepSep(re, s)
  // return strings.Split(s, " ")
}


// Lookup looks up a word in dict
func Lookup(dict map[string]interface{}, lang string, word string) string {
  word = strings.ToLower(word)
  val, ok := dict[lang]
  // no dict for lang
  if !ok { return "" }
  val, ok = dict[lang].(map[string]interface{})[word]
  // word not in dict
  if !ok { return "" }

  return val.(string)


}



func SplitKeepSep(re *regexp.Regexp, s string) []string {
  matches := re.FindAllStringIndex(s, -1)
  var out []string
  start := 0
  for _, match := range matches {
    out = append(out, s[start:match[0]])
    out = append(out, s[match[0]:match[1]])
    // the next starts after match
    start = match[1]
  }
  if start < len(s) {
    out = append(out, s[start:len(s)-1])
  }
  return out
  

}



