// Цель упражнения: Написать утилиту для управления закладками, не сохраняющую данные локально, а сохраняющую их только в памяти.
// Функциональность приложения:
// Просмотр всех закладок
// Добавление новой закладки
// Удаление закладки
// Выход из приложения
// Шаги Создания Утилиты:
// Подготовка интерфейса пользователя:
//     При запуске приложения выводить меню с возможностями.
//     Использовать бесконечный цикл для возврата к меню после каждого действия.
//     Реализация функции получения выбора пользователя.
// Работа с закладками:
//     Для хранения закладок используем мапу (map[string]string), ключ - название закладки, значение - адрес.
//     Функция для вывода всех закладок с учетом пустого списка.
//     Добавление новой закладки с вводом пользователя (название и адрес).
//     Функция удаления закладки по названию.
//     Вывод сообщения при пустом списке закладок.
// Цикл обработки действий пользователя:
//     Использование switch для обработки выбора пользователя.
//     Определение и реализация функций для каждого действия (просмотр, добавление, удаление).
//     Завершение работы приложения при выборе соответствующего пункта.
// Тестирование приложения:
//     Прогон различных сценариев использования приложения (просмотр, добавление, удаление, проверка на пустоту).
//     Выход из приложения.

// еще нужно чтобы список закладок выводился и можно было прилепить комментарий к закладке
// сохранение закладок после перезапуска проекта (не через json)
// сделать поиск по закладкам
// исправить адреса на русском языке (punycode)


package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"os"
	"bufio"
	"net/url"
)

type Bookmark struct {
	URL     string
	Comment string
}

var bookmarks = make(map[string]Bookmark)

const bookmarksFile = "bookmarks.txt"

func saveBookmarks() {
	file, err := os.Create(bookmarksFile)
	if err != nil {
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for name, info := range bookmarks {
		fmt.Fprintf(writer, "%s|%s|%s\n", name, info.URL, info.Comment)
	}
	writer.Flush()
}

func loadBookmarks() {
	file, err := os.Open(bookmarksFile)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "|", 3)
		if len(parts) == 3 {
			bookmarks[parts[0]] = Bookmark{
				URL:     parts[1],
				Comment: parts[2],
			}
		}
	}
}

func main() {
	loadBookmarks()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		searchQuery := r.URL.Query().Get("search")

		fmt.Fprintf(w, "<html><body>")
		fmt.Fprintf(w, `<form action="/add" method="GET">`)
		fmt.Fprintf(w, `название<br><input type="text" name="name" required><br><br>`)
		fmt.Fprintf(w, `адрес<br><input type="text" name="url" required><br><br>`)
		fmt.Fprintf(w, `коммент<br><input type="text" name="comment"><br><br>`)
		fmt.Fprintf(w, `<button type="submit">добавить</button>`)
		fmt.Fprintf(w, `</form>`)

		fmt.Fprintf(w, "<h2>список закладок ")
		fmt.Fprintf(w, `<form action="/" method="GET" style="display:inline;">`)
		fmt.Fprintf(w, `<input type="text" name="search" placeholder="поиск" value="%s">`, searchQuery)
		fmt.Fprintf(w, `<button type="submit">найти</button>`)
		fmt.Fprintf(w, `</form></h2>`)

		displayedBookmarks := make(map[string]Bookmark)
		if searchQuery != "" {
			for name, info := range bookmarks {
				if strings.Contains(strings.ToLower(name), strings.ToLower(searchQuery)) {
					displayedBookmarks[name] = info
				}
			}
		} else {
			displayedBookmarks = bookmarks
		}

		if len(displayedBookmarks) == 0 {
			fmt.Fprintf(w, "<p>список пуст</p>")
		} else {
			fmt.Fprintf(w, "<ul>")
			for name, info := range displayedBookmarks {
				href := info.URL
				if !strings.HasPrefix(href, "http://") && !strings.HasPrefix(href, "https://") {
					href = "http://" + href
				}
				
				displayURL := info.URL
				if decoded, err := url.PathUnescape(info.URL); err == nil {
					displayURL = decoded
				}
				
				fmt.Fprintf(w, "<li><b>%s</b>: <a href=\"%s\" target=\"_blank\">%s</a> ", name, href, displayURL)
				
				if info.Comment != "" {
					fmt.Fprintf(w, "(%s)", info.Comment)
				}
				
				fmt.Fprintf(w, ` <form action="/delete" method="GET" style="display:inline;"><input type="hidden" name="name" value="%s"><button type="submit">удалить</button></form></li>`, name)
			}
			fmt.Fprintf(w, "</ul>")
		}
		
		fmt.Fprintf(w, "</body></html>")
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		url := r.URL.Query().Get("url")
		comment := r.URL.Query().Get("comment")

		if name != "" && url != "" {
			bookmarks[name] = Bookmark{
				URL:     url,
				Comment: comment,
			}
			saveBookmarks()
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name != "" {
			delete(bookmarks, name)
			saveBookmarks()
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}