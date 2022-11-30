package schema

type User struct {
	Username   string   `json:"username"`
	ShowEns    bool     `json:"show_ens"`
	Title      string   `json:"title"`
	ImageUrl   string   `json:"image_url"`
	ShowNFT    bool     `json:"show_nft"`
	NFTAddress string   `json:"nft_address"`
	NFTId      string   `json:"nft_id"`
	Email      string   `json:"email"`
	Bio        string   `json:"bio"`
	Links      []string `json:"links"`
	Skills     []Skill  `json:"skills"`
}

type Skill struct {
	Address      string   `json:"address"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Links        []string `json:"links"`
	ImageUrls    []string `json:"image_urls"`
	MinimumPrice float64  `json:"minimum_price"`
}
