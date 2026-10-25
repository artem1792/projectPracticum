package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const USD = 83.92
const EUR = 94.04

func rubToUsd(rub float64) float64 {
	return rub / USD
}

func rubToEur(rub float64) float64 {
	return rub / EUR
}

func usdToEur(USD2 float64) float64 {
	rub := USD2 * USD
	return rub / EUR
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<html><body><a href="/convert"><button>Перейти к конвертеру</button></a></body></html>`)
	})

	http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "страница2")
	})

	http.HandleFunc("/page2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "страница 333")
	})

	http.HandleFunc("/convert", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<html><body><form action="/convert" method="GET">`)
		fmt.Fprintf(w, `<input type="text" name="rub" placeholder="Сумма"><br><br>`)
		fmt.Fprintf(w, `<input type="checkbox" name="to_usd" value="1"> В доллары<br>`)
		fmt.Fprintf(w, `<input type="checkbox" name="to_eur" value="1"> В евро<br><br>`)
		fmt.Fprintf(w, `<button type="submit">Конвертировать</button></form><hr>`)

		rubStr := r.URL.Query().Get("rub")
		if rubStr != "" {
			rub, err := strconv.ParseFloat(rubStr, 64)
			if err != nil {
				fmt.Fprint(w, "<p>Error</p>")
			} else {
				toUsd := r.URL.Query().Get("to_usd")
				toEur := r.URL.Query().Get("to_eur")

				if toUsd == "1" {
					fmt.Fprintf(w, "<p>%.2f руб. = %.2f долларов</p>", rub, rubToUsd(rub))
				}
				if toEur == "1" {
					fmt.Fprintf(w, "<p>%.2f руб. = %.2f евро</p>", rub, rubToEur(rub))
				}
				if toUsd == "1" && toEur == "1" {
					fmt.Fprintf(w, "<p>%.2f долларов = %.2f евро</p>", rubToUsd(rub), usdToEur(rubToUsd(rub)))
				}
			}
		}
		fmt.Fprintf(w, `</body></html>`)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}