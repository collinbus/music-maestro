package music

type GetUserTracksResponse struct {
	Items    []UserTrack `json:"items"`
	Limit    int         `json:"limit"`
	Next     string      `json:"next"`
	Offset   int         `json:"offset"`
	Previous string      `json:"previous"`
	Total    int         `json:"total"`
}

type UserTrack struct {
	AddedAt string `json:"added_at"`
	Track   struct {
		Id           string             `json:"id"`
		Name         string             `json:"name"`
		ExternalUrls ExternalSpotifyUrl `json:"external_urls"`
		Href         string             `json:"href"`
		Popularity   int                `json:"popularity"`
		PreviewUrl   string             `json:"preview_url"`
		TrackNumber  int                `json:"track_number"`
		Uri          string             `json:"uri"`
		Artists      []struct {
			Id           string             `json:"id"`
			Name         string             `json:"name"`
			Uri          string             `json:"uri"`
			ExternalUrls ExternalSpotifyUrl `json:"external_urls"`
		} `json:"artists"`
		Album struct {
			Id                   string             `json:"id"`
			Name                 string             `json:"name"`
			Uri                  string             `json:"uri"`
			ExternalUrls         ExternalSpotifyUrl `json:"external_urls"`
			ReleaseDate          string             `json:"release_date"`
			ReleaseDatePrecision string             `json:"release_date_precision"`
			TotalTracks          int                `json:"total_tracks"`
			Images               []struct {
				Height int    `json:"height"`
				Width  int    `json:"width"`
				Url    string `json:"url"`
			} `json:"images"`
		} `json:"album"`
	} `json:"track"`
}

type ExternalSpotifyUrl struct {
	Spotify string `json:"spotify"`
}
