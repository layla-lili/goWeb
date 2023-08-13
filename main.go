package main

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			response := r.Form.Get("response")
			selectedOption := r.Form.Get("option")
			if selectedOption == "" { // if no option is selected, set a default value
				selectedOption = "option1"
			}
			file, err := os.Open(selectedOption + ".txt")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer file.Close()

			var fileContent []string

			scanner := bufio.NewScanner(file)
			///////////////////////////////
			// var ascii_art []string
			// var ascii_codes []int

			// ascii_code := 32
			// var lines []string

			// for scanner.Scan() {
			// 	lines = append(lines, scanner.Text())
			// 	if len(lines) == 8 {
			// 		ascii_codes = append(ascii_codes, ascii_code)
			// 		ascii_art = append(ascii_art, strings.Join(lines, "\n"))
			// 		ascii_code++
			// 		lines = nil
			// 	}
			// }

			// if err := scanner.Err(); err != nil {
			// 	panic(err)
			// }

			// mystr := "hello"
			// for i, code := range ascii_codes {
			// 	// fmt.Println("code: ", code)
			// 	// fmt.Printf("%d: %c\n", code, code)
			// 	// fmt.Println("ascii i", i)

			// 	for _, m := range mystr {
			// 		//if m == rune(code) {
			// 		fmt.Println("m ", m, "string m: ", string(m))
			// 		fmt.Println("code", code)
			// 		fmt.Println("ascii_code", ascii_code)
			// 		fmt.Printf("%s\n", ascii_art[i])
			// 		//}
			// 	}

			// }

			///////////////////////
			for scanner.Scan() {

				fileContent = append(fileContent, string(scanner.Bytes()))
				fileContent = append(fileContent, string('\n'))

			}
			
			if err := scanner.Err(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			tmpl := template.Must(template.ParseFiles("form.html"))
			data := struct {
				Response       string
				SelectedOption string
				Option1        bool
				Option2        bool
				Option3        bool
				FileContent    string
			}{
				Response:       response,
				SelectedOption: selectedOption,
				Option1:        selectedOption == "standard",
				Option2:        selectedOption == "shadow",
				Option3:        selectedOption == "ThinkerToy",
				FileContent: strings.Trim(fmt.Sprintf("%v", fileContent), "[]"),
			}
			tmpl.Execute(w, data)
			//fmt.Print("fileContent", fileContent)
		} else {
			tmpl := template.Must(template.ParseFiles("form.html"))
			tmpl.Execute(w, nil)
		}
	})

	http.ListenAndServe(":8080", nil)
}
