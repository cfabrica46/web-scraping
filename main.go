package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	tag := "h1"

	resp, err := http.Get("https://www.hectormainar.com/h1.php")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	//regex method
	/* exampleText := `<h2>holasss ssg sdfa</h2>

	<p><h1 class='hola'>hola q talk</h1></p>`
		reader := strings.NewReader(exampleText)
		scanner := bufio.NewScanner(reader) */

	expression := fmt.Sprintf("(?s)<%s.*>.+</%s>", tag, tag)

	r, err := regexp.Compile(expression)
	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		text := r.FindString(scanner.Text())
		if len(text) != 0 {
			r := regexp.MustCompile(`(<\/?[a-zA-A]+?[^>]*\/?>)*`)
			matchs := r.FindAllString(text, -1)

			for i := range matchs {
				if strings.TrimSpace(matchs[i]) != "" {
					text = strings.ReplaceAll(text, matchs[i], "")
				}
			}
			fmt.Println("First h1 is:")
			fmt.Println(text)
			break
		}
	}
}
