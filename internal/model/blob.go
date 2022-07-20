package model

type BlobAtr struct {
	Title string `json:"title"`
}

type BlobLinks struct {
	Self    string `json:"self"`
	Related string `json:"related"`
}

type BlobData struct {
	Type string `json:"type"`
	ID   int    `json:"id"`
}

type BlobAuthor struct {
	Links *BlobLinks `json:"links"`
	Data  *BlobData  `json:"data"`
}

type BlobRelationships struct {
	Author *BlobAuthor `json:"author"`
}

type BlobMainData struct {
	Type          string             `json:"type"`
	ID            int                `json:"id"`
	Attributes    *BlobAtr           `json:"attributes"`
	Relationships *BlobRelationships `json:"relationships"`
}
