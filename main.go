package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// TODO 填入你的 client_id
	clientID = "XXXXXXXXX"
	// TODO 填入你的 client_secret
	clientSecret = "XXXXXXXXX"

	githubToken = "https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/oauth/redirect", oauthDemo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}

}
func oauthDemo(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	code := req.URL.Query().Get("code")
	fmt.Println(req.URL.String(), code)

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
	for k, v := range req.Header {
		fmt.Println(k, v)
	}

	c := &http.Client{}
	r, err := http.NewRequest(http.MethodPost, fmt.Sprintf(githubToken, clientID, clientSecret, code), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	r.Header.Add("accept", "application/json")

	res, err := c.Do(r)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("--------->", string(body))

	http.Redirect(w, req, fmt.Sprintf("/welcome.html?name=%s", getGitHubUserName(body)), 301)
}
func getGitHubUserName(jsonBody []byte) string {
	data := map[string]interface{}{}
	err := json.Unmarshal(jsonBody, &data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for k, v := range data {
		fmt.Println(k, v)
	}

	c := &http.Client{}
	r, err := http.NewRequest(http.MethodGet, "https://api.github.com/user", nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	r.Header.Add("accept", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("token %s", data["access_token"]))

	res, err := c.Do(r)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println("--------->", string(body))

	for k, v := range res.Header {
		fmt.Println(k, v)
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for k, v := range data {
		fmt.Println(k, v)
	}
	return data["name"].(string)
}
