package vkapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetWallPost() {

	/*Все параметры хранить в базе.
	access_token хранить в зашифрованном виде, дешифровывать при первоначальном запуске программы и использовать повторно.
	Логи тоже хранить в базе.*/

	type WallPost struct {
		Response struct {
			Items []struct {
				CopyHistory []struct {
					ID          int    `json:"id"`
					TextPost    string `json:"text"`
					Attachments []struct {
						Link struct {
							UrlPost string `json:"url"`
							Photo   struct {
								OrigPhoto struct {
									UrlPhoto string `json:"url"`
								} `json:"orig_photo"`
							} `json:"photo"`
						} `json:"link"`
					} `json:"attachments"`
				} `json:"copy_history"`
			} `json:"items"`
		} `json:"response"`
	}

	access_token := "beedb8b3beedb8b3beedb8b3c6bdc98e54bbeedbeedb8b3d9d7fc987b6687e723b10668"
	versionAPI := "5.199"
	owner_id := "-74598681"
	domain := "elite_dangerous_ru"
	count := "2"

	queryString := fmt.Sprintf("https://api.vk.com/method/wall.get?&v=%s&access_token=%s&owner_id=%s&domain=%s&count=%s", versionAPI, access_token, owner_id, domain, count)

	resp, err := http.Get(queryString)

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Не удалось получить данные по запросу. Код состояния: ", resp.Status)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	wp := WallPost{}

	err = json.Unmarshal(body, &wp)

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range wp.Response.Items {
		for _, text := range v.CopyHistory {
			for _, att := range text.Attachments {
				fmt.Println(att.Link.Photo.OrigPhoto.UrlPhoto)
			}
		}
	}
}
