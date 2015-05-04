package loader

import (
  "io/ioutil"
  "net/http"
  "encoding/json"
  "github.com/kyeah/gohunt/gohunt"
  "github.com/jzelinskie/geddit"
)

type Feed interface {
  Display()
  Setter(name string, title []string, url []string)
}

type JsonData struct {
  Title   string `json:"title"`
  Url     string `json:"url"`
}

type ResultData struct {
  Name string
  Title []string
  Url []string
}

func (r *ResultData) Display() {
  ppred("[" + r.Name + "]\n")
  for key, _ := range r.Title {
    pp(" - " + r.Title[key] + "\n")
    pp("   - " + r.Url[key] + "\n")
  }
  pp("\n")
}

func (r *ResultData) Setter(name string, title []string, url []string) {
  r.Name = name
  r.Title = title
  r.Url = url
}

func getIdsFromUrl(url string) (ids []int) {
  resp, err := http.Get(url)
  perror(err)
  jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)
  perror(err)
  err = json.Unmarshal([]byte(jsonDataFromHttp), &ids)
  perror(err)
  return ids
}

func getJsonDataFromUrl(url string) (jsonRes JsonData) {
  resp, err := http.Get(url)
  perror(err)
  jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)
  perror(err)
  err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonRes)
  perror(err)
  return jsonRes
}

func getRedditSession() (session *geddit.LoginSession) {
  session, _ = geddit.NewLoginSession(
    "techstk",
    "techstack",
    "gedditAgent v1",
  )
  return session
}

func GetPHFeed(ph chan ResultData) {
  var result ResultData
  var PHTitle, PHUrl []string

  client := gohunt.NewUserClient("a0ad779df4746d96d4d87a5a589f786277c4c78699445fc70abdc08d4be77b45")

  posts, err := client.GetPosts()
  perror(err)

  for _, post := range posts[0:5] {
    PHTitle = append(PHTitle, post.Name + " : " + post.Tagline)
    PHUrl = append(PHUrl, post.RedirectUrl)
  }

  result.Setter("ProductHunt", PHTitle, PHUrl)
  ph <- result
}

func GetRedditFeed(re chan ResultData) {
  var result ResultData
  var RETitle, REUrl []string

  session := getRedditSession()
  subOpts := geddit.ListingOptions{Limit: 5}
  submissions, _ := session.SubredditSubmissions("programming", geddit.TopSubmissions, subOpts)
  for _, s := range submissions {
    RETitle = append(RETitle, s.Title)
    REUrl = append(REUrl, s.URL)
  }

  result.Setter("Reddit", RETitle, REUrl)
  re <- result
}
