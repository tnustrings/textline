package main

import "fmt"
import "embed"
import "os"
import "encoding/json"
import "io"
import "regexp"
import "strings"

//go:embed dict/*
var dictfolder embed.FS

func Hello() string {
  message := fmt.Sprintf("hi")
  return message
}
// main runs Hello.
func main() {
  langa := os.Args[1]
  dict := make(map[string]interface{})
  content, _ := dictfolder.ReadFile("dict/lingo-dict-" + langa + ".json")
  json.Unmarshal(content, &dict)
  fmt.Println(Tagid(string(CatStdin()), langa, dict))
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
  re := regexp.MustCompile(`[\p{L}\p{M}]+'?[\p{L}\p{M}]+`)
  return SplitKeepSep(re, s)
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

// Cat returns the contents of a file as byte array
func Cat(path string) []byte {
  file, err := os.Open(path)
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  bytes, _ := io.ReadAll(file)

  return bytes
}

// CatStdin cats from stdin
func CatStdin() []byte {
  bytes, _ := io.ReadAll(os.Stdin)
  return bytes
}
// DictCopy puts values from the src dict into dest dict, overwriting
func DictCopy(dst map[string]interface{}, src map[string]interface{}) {
  for lang, _ := range src {
    _, ok := dst[lang]
    if !ok {
      if lang == "en" {
        dst[lang] = make(map[string]interface{})
      } else {
        dst[lang] = map[string]string{}
      }
    }
    if lang == "en" {
      DictCopy(dst["en"].(map[string]interface{}), src["en"].(map[string]interface{}))
    } else { 
      // maps.Copy(dst[lang], src[lang].(map[string]string))
      for word, tlate := range src[lang].(map[string]interface{}) {
        dst[lang].(map[string]string)[word] = tlate.(string)
      }
    }
  }
}
