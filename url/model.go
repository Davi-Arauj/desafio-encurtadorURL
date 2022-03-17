package url

type MyUrl struct {
	ID       string `json:"id,omitempty"       binding:"required" `
	LongUrl  string `json:"longUrl,omitempty" binding:"required" valid:"url" `
	ShortUrl string `json:"shortUrl,omitempty" binding:"required" `
}
