package url

import (
	"encoding/json"
	"encurtUrl/db"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/asaskevich/govalidator"
	hashids "github.com/speps/go-hashids"
)

func CreateEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(db.DB)
	var err error
	var url MyUrl
	var body, _ = ioutil.ReadAll((r.Body))
	_ = json.Unmarshal(body, &url)

	if govalidator.IsURL(url.LongUrl) {
		if err = psql.Select("id,longurl,shorturl").
			From("url").
			Where(sq.Eq{
				"longurl": url.LongUrl,
			}).
			Scan(&url.ID, &url.LongUrl, &url.ShortUrl); err != nil {
			hd := hashids.NewData()
			h, _ := hashids.NewWithData(hd)
			now := time.Now()
			url.ID, _ = h.Encode([]int{int(now.Unix())})
			url.ShortUrl = "http://localhost:8080/" + url.ID

			if _, err := psql.Insert("url").
				Columns("id", "longurl", "shorturl").
				Values(url.ID, url.LongUrl, url.ShortUrl).
				Exec(); err != nil {
				log.Println(err)
			}
		}
	}
	json.NewEncoder(w).Encode(url)
}
