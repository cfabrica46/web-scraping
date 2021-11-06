package main

/* resp, err := http.Get("https://www.hectormainar.com/h1.php")
if err != nil {
	log.Fatal(err)
}
defer resp.Body.Close()

tokenizer := html.NewTokenizer(resp.Body)
var enter bool

for {
	tt := tokenizer.Next()
	token := tokenizer.Token()

	err := tokenizer.Err()
	if err == io.EOF {
		break
	}

	switch tt {
	case html.ErrorToken:
		log.Fatal(err)

	case html.StartTagToken, html.SelfClosingTagToken:
		enter = false

		tag := token.Data
		if tag == "h1" {
			enter = true
		}
	case html.TextToken:
		if enter {
			data := strings.TrimSpace(token.Data)

			if len(data) > 0 {
				fmt.Println("First h1:")
				fmt.Println(data)
				return
			}
		}
	}
} */
