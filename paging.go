package spotify

// Paging represents a PagingObject in the Spotify API
// https://developer.spotify.com/documentation/web-api/reference/#object-pagingobject
type Paging struct {
	Tracks TrackPage `json:"tracks"`
}

type TrackPage struct {
	PagingMeta
	Items []*Track `json:"items"`
}

type PlaylistPage struct {
	PagingMeta
	Items []*Playlist `json:"items"`
}

type PlaylistTrackPage struct {
	PagingMeta
	Items []*PlaylistTrack `json:"items"`
}
