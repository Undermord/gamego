package main

import (
	"fmt"
	"sync"
	"net/http"
	"io"

	
)

/* 
Напишите функцию FetchURLs(urls []string) map[string]string, которая:

Принимает слайс URL-адресов.
Конкурентно делает HTTP-запросы к каждому URL.
Собирает результаты (код ответа и часть тела) в map[string]string, где:
ключ — URL
значение — содержимое ответа (ограниченное, например, 100 символами)
Использует sync.WaitGroup и sync.Mutex для защиты записи в map.
В случае ошибки записывает "error" как значени 
*/


func FetchURLs(urls []string) map[string]string {
	var wg sync.WaitGroup
	var mu sync.Mutex
	result := make(map[string]string)

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				mu.Lock()
				result[url] = "error"
				mu.Unlock()
				return
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				mu.Lock()
				result[url] = "error"
				mu.Unlock()
				return
			}

			content := string(body)
			if len(content) > 100 {
				content = content[:100]
			}
			mu.Lock()
			result[url] = content
			mu.Unlock()
		}(url)
	}
	wg.Wait()
	return result
}



func main() {
	urls := []string{
		"https://google.com", 
		"https://golang.org",
		"https://www.nonexistenturl.xyz",
		"https://stepik.org",
		"https://github.com",
		}
	results := FetchURLs (urls)
	for url, content := range results {
		fmt.Printf("URL: %s, Content: %s\n", url, content)
	}
}