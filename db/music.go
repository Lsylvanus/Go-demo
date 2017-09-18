package db

type SearchMusic struct {
	Data SongData `json:"data"`
	Pro  []Pro    `json:"-"`
}

type SongData struct {
	Songs  []Song   `json:"song"`
	Album  []Album  `json:"album"`
	Artist []Artist `json:"artist"`
}

type Pro struct {
	Artist string `json:"-"`
	Song   string `json:"-"`
	Album  string `json:"-"`
}

type Song struct {
	Id               int64
	BitrateFee       string `json:"bitrate_fee"`
	Weight           string `json:"weight"`
	Songname         string `json:"songname"`
	Songid           string `json:"songid"`
	HasMv            string `json:"has_mv"`
	YyrArtist        string `json:"yyr_artist"`
	ResourceTypeExt  string `json:"resource_type_ext"`
	Artistname       string `json:"artistname"`
	Info             string `json:"info"`
	ResourceProvider string `json:"resource_provider"`
	EncryptedSongid  string `json:"encrypted_songid"`
}

type Album struct {
	Id              int64
	Albumname       string `json:"albumname"`
	Weight          string `json:"weight"`
	Artistname      string `json:"artistname"`
	ResourceTypeExt string `json:"resource_type_ext"`
	Artistpic       string `json:"artistpic"`
	Albumid         string `json:"albumid"`
}

type Artist struct {
	Id         int64
	YyrArtist  string `json:"yyr_artist"`
	Artistname string `json:"artistname"`
	Artistid   string `json:"artistid"`
	Artistpic  string `json:"artistpic"`
	Weight     string `json:"weight"`
}

type DownLink struct {
	ErrorCode int      `json:"errorCode"`
	Data      LinkData `json:"data"`
}

type LinkData struct {
	XCode    string     `json:"xcode"`
	SongList []SongList `json:"songList"`
}

type SongList struct {
	Id         int64
	QueryId    string `json:"queryId"`
	SongId     string `json:"songId"`
	SongName   string `json:"songName"`
	ArtistId   string `json:"artistId"`
	ArtistName string `json:"artistName"`

	AlbumId      string `json:"albumId"`
	AlbumName    string `json:"albumName"`
	SongPicSmall string `json:"songPicSmall"`
	SongPicBig   string `json:"songPicBig"`
	SongPicRadio string `json:"songPicRadio"`

	LrcLink  string  `json:"lrcLink"`
	Version  string  `json:"version"`
	CopyType float64 `json:"copyType"`
	Time     float64 `json:"time"`
	LinkCode float64 `json:"linkCode"`

	SongLink string `json:"songLink"`
	ShowLink string `json:"showLink"`
	Format   string `json:"format"`
	Rate     string `json:"rate"`
	Size     string `json:"size"`

	RelateStatus string `json:"relateStatus"`
	ResourceType string `json:"resourceType"`
	Source       string `json:"source"`
}
