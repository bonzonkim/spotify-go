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
	Href         string            `json:"href,omitempty"`
	Id           string            `json:"id,omitempty"`
	IdType       string            `json:"type,omitempty"`
	Uri          string            `json:"uri,omitempty"`
	Followers    SpotifyFollowers  `json:"followers"`
	Images       []SpotifyImage    `json:"images,omitempty"`
}

//	{
//	 "href": "https://api.spotify.com/v1/me/shows?offset=0&limit=20",
//	 "limit": 20,
//	 "next": "https://api.spotify.com/v1/me/shows?offset=1&limit=1",
//	 "offset": 0,
//	 "previous": "https://api.spotify.com/v1/me/shows?offset=1&limit=1",
//	 "total": 4,
//	 "items": [
//	   {
//	     "external_urls": {
//	       "spotify": "string"
//	     },
//	     "followers": {
//	       "href": "string",
//	       "total": 0
//	     },
//	     "genres": ["Prog rock", "Grunge"],
//	     "href": "string",
//	     "id": "string",
//	     "images": [
//	       {
//	         "url": "https://i.scdn.co/image/ab67616d00001e02ff9ca10b55ce82ae553c8228",
//	         "height": 300,
//	         "width": 300
//	       }
//	     ],
//	     "name": "string",
//	     "popularity": 0,
//	     "type": "artist",
//	     "uri": "string"
//	   }
//	 ]
//	}
type SpotifyTopTracksItem struct {
	ExternalUrls map[string]string `json:"external_urls,omitempty"`
	Followers    SpotifyFollowers  `json:"followers,omitempty"`
	Genres       []string          `json:"genres,omitempty"`
	Href         string            `json:"href,omitempty"`
	Id           string            `json:"id,omitempty"`
	Images       SpotifyImage      `json:"images,omitempty"`
	Name         string            `json:"name,omitempty"`
	Popularity   int               `json:"popularity,omitempty"`
	ItemType     string            `json:"item_type,omitempty"`
	Uri          string            `json:"uri,omitempty"`
}
type SpotifyTopTracks struct {
	Href     string                 `json:"href,omitempty"`
	Limit    int                    `json:"limit,omitempty"`
	Next     string                 `json:"next,omitempty"`
	Offset   int                    `json:"offset,omitempty"`
	Previous string                 `json:"previous,omitempty"`
	Total    int                    `json:"total,omitempty"`
	Items    []SpotifyTopTracksItem `json:"items,omitempty"`
}
