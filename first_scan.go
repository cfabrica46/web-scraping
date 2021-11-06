package main

/* resp, err := http.Get("https://golang.org/dl/")
if err != nil {
	log.Fatal(err)
}
defer resp.Body.Close()

scanner := bufio.NewScanner(resp.Body)

r, err := regexp.Compile(`\/dl\/go([0-9\.]+)\.linux-amd64\.tar\.gz`)
if err != nil {
	log.Fatal(err)
} */

//print all coincidences
/* for scanner.Scan() {
	text := r.FindString(scanner.Text())
	if len(text) != 0 {
		fmt.Println(text)
	}
} */

//get last version
/* var value string
for scanner.Scan() {
	text := r.FindString(scanner.Text())
	if len(text) != 0 {
		value = text
		break
	}
}
url := fmt.Sprintf("golang.org%s", value)

fmt.Printf("The url of the last go version is: %s\n", url) */
