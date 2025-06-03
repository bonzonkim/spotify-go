package types

//{
//  "display_name": "B9",
//  "email": "kbg8609@gmail.com",
//  "external_urls": {
//    "spotify": "https://open.spotify.com/user/3k75udaz705fa2fh6evjykafn"
//  },
//  "followers": {
//    "href": null,
//    "total": 2
//  },
//  "href": "https://api.spotify.com/v1/users/3k75udaz705fa2fh6evjykafn",
//  "id": "3k75udaz705fa2fh6evjykafn",
//  "images": [ ],
//  "type": "user",
//  "uri": "spotify:user:3k75udaz705fa2fh6evjykafn"
//}

type SpotifyImage struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type SpotifyFollowers struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type SpotifyUser struct {
	DisplayName  string            `json:"display_name,omitempty"`
	Email        string            `json:"email,omitempty"`
	ExternalUrls map[string]string `json:"external_urls,omitempty"`
	//Followers    SpotifyFollowers  `json:"followers,omitempty"`
	Href string `json:"href,omitempty"`
	Id   string `json:"id,omitempty"`
	//Images       SpotifyImage      `json:"images,omitempty"`
	IdType string `json:"type,omitempty"`
	Uri    string `json:"uri,omitempty"`
}
