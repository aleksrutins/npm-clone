package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/munchkinhalfling/npm-clone/depinstall"
	"github.com/munchkinhalfling/npm-clone/executil"
)

func main() {
	fmt.Printf("Fetching...\n")
	resp, err := http.Get("https://registry.npmjs.org/" + os.Args[1])
	if err != nil {
		defer resp.Body.Close()
		fmt.Println("There was an error.")
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	var decl interface{}
	err = json.Unmarshal(body, &decl)
	m := decl.(map[string]interface{})
	latest := m["dist-tags"].(map[string]interface{})["latest"].(string)
	latestData := m["versions"].(map[string]interface{})[latest]
	var tarballUrl string = latestData.(map[string]interface{})["dist"].(map[string]interface{})["tarball"].(string)
	fmt.Printf("Actually fetching...\n")
	resp, err = http.Get(tarballUrl)
	if err != nil {
		defer resp.Body.Close()
		fmt.Println("There was an error.")
		return
	}
	body, err = ioutil.ReadAll(resp.Body)
	f, err := os.Create(os.Args[1] + ".tar.gz")
	f.Write(body)
	f.Close()
	fmt.Println("Untarring...")
	executil.Run("tar", "xzf", os.Args[1]+".tar.gz")
	fmt.Println("Organizing...")
	executil.Run("rm", os.Args[1]+".tar.gz")
	executil.Run("mv", "package", os.Args[1])
	executil.Run("rmdir", "package")
	depinstall.InstallDeps(os.Args[1])
	fmt.Println("Done!")
}
