/*
 * Copyright (C) 2019 l4rzy
 * MIT License
 */

package csn

type CSNMusicBase struct {
	MusicTitleURL string `json:"music_title_url"`
	MusicID       string `json:"music_id"`
	CatID         string `json:"cat_id"`
	MusicTitle    string `json:"music_title"`
	MusicArtist   string `json:"music_artist"`
}

type CSNMusicHot struct {
	CSNMusicBase
	CatLevel        string `json:"cat_level"`
	CoverID         string `json:"cover_id"`
	MusicDownloads  string `json:"music_downloads"`
	MusicBitrate    string `json:"music_bitrate"`
	MusicLength     string `json:"music_length"`
	MusicThumbsTime string `json:"music_thumbs_time"`
	CoverImg        string `json:"cover_img"`
}

// new search api that bases on the
// new site's xhr
type CSNMusicSearchNew struct {
	MusicDownloads   int    `json:"music_downloads"`
	MusicID          int    `json:"music_id"`
	MusicTitle       string `json:"music_title"`
	MusicArtist      string `json:"music_artist"`
	MusicBitrate     string `json:"music_bitrate"`
	MusicBitrateHTML string `json:"music_bitrate_html"`
	MusicLink        string `json:"music_link"`
	CatID            int    `json:"cat_id"`
	MusicListen      int    `json:"music_listen"`
	MusicFilename    string `json:"music_filename"`
	MusicCover       string `json:"music_cover"`
	MusicTitleURL    string `json:"music_title_url"`
}

type CSNMusicSearch struct {
	CSNMusicBase
	ID             int    `json:"id"`
	Thumbnail      string `json:"thumbnail"`
	Preview        string `json:"preview"`
	CatLevel       string `json:"cat_level"`
	MusicComposer  string `json:"music_composer"`
	MusicAlbum     string `json:"music_album"`
	MusicYear      string `json:"music_year"`
	MusicDownloads string `json:"music_downloads"`
	MusicBitrate   string `json:"music_bitrate"`
	MusicLength    string `json:"music_length"`
	MusicWidth     string `json:"music_width"`
	MusicHeight    string `json:"music_height"`
	CoverImg       string `json:"cover_img"`
}

type CSNMusicAlbum struct {
	CSNMusicBase
	CatLevel        string `json:"cat_level"`
	CatCustom       string `json:"cat_custom"`
	MusicTrackID    string `json:"music_track_id"`
	MusicShortlyric string `json:"music_shortlyric"`
	MusicLength     string `json:"music_length"`
}

type CSNMusicInfo struct {
	CSNMusicBase
	CatLevel              string `json:"cat_level"`
	CatSublevel           string `json:"cat_sublevel"`
	CoverID               string `json:"cover_id"`
	MusicComposer         string `json:"music_composer"`
	MusicAlbum            string `json:"music_album"`
	MusicProduction       string `json:"music_production"`
	MusicAlbumID          string `json:"music_album_id"`
	MusicYear             string `json:"music_year"`
	MusicListen           string `json:"music_listen"`
	MusicDownloads        string `json:"music_downloads"`
	MusicTime             string `json:"music_time"`
	MusicBitrate          string `json:"music_bitrate"`
	MusicLength           string `json:"music_length"`
	Music32Filesize       string `json:"music_32_filesize"`
	MusicFilesize         string `json:"music_filesize"`
	Music320Filesize      string `json:"music_320_filesize"`
	MusicM4AFilesize      string `json:"music_m4a_filesize"`
	MusicLosslessFilesize string `json:"music_lossless_filesize"`
	MusicWidth            string `json:"music_width"`
	MusicHeight           string `json:"music_height"`
	MusicUsername         string `json:"music_username"`
	MusicLyric            string `json:"music_lyric"`
	MusicImgHeight        string `json:"music_img_height"`
	MusicImgWidth         string `json:"music_img_width"`
	MusicImg              string `json:"music_img"`
	VideoThumbnail        string `json:"video_thumbnail"`
	VideoPreview          string `json:"video_preview"`
	FileURL               string `json:"file_url"`
	File32URL             string `json:"file_32_url"`
	File320URL            string `json:"file_320_url"`
	FileM4AURL            string `json:"file_m4a_url"`
	FileLosslessURL       string `json:"file_lossless_url"`
	FullURL               string `json:"full_url"`
	MusicGenre            string `json:"music_genre"`
}

type CSNVideoSearch struct {
	VideoID         int         `json:"video_id"`
	VideoTitle      string      `json:"video_title"`
	VideoArtist     string      `json:"video_artist"`
	VideoBitrate    string      `json:"video_bitrate"`
	VideoLink       string      `json:"video_link"`
	VideoCover      string      `json:"video_cover"`
	VideoListen     interface{} `json:"video_listen"`
	VideoLength     string      `json:"video_length"`
	VideoLengthHTML string      `json:"video_length_html"`
	VideoDownloads  int         `json:"video_downloads"`
	VideoTitleURL   string      `json:"video_title_url"`
}

type CSNArtistSearch struct {
	ArtistID       int    `json:"artist_id"`
	ArtistNickname string `json:"artist_nickname"`
	ArtistLink     string `json:"artist_link"`
	ArtistCover    string `json:"artist_cover"`
	ArtistAvatar   string `json:"artist_avatar"`
}

type CSNAlbumSearch struct {
	CoverID         int         `json:"cover_id"`
	MusicAlbum      string      `json:"music_album"`
	AlbumLink       string      `json:"album_link"`
	AlbumID         string      `json:"album_id"`
	AlbumBitrate    interface{} `json:"album_bitrate"`
	AlbumArtist     string      `json:"album_artist"`
	AlbumArtistHTML string      `json:"album_artist_html"`
	AlbumCover      string      `json:"album_cover"`
}

type CSNSearchResp struct {
	MusicList []CSNMusicSearch `json:"music_list"`
}

type CSNSearchNewResp []struct {
	Q     string `json:"q"`
	Music struct {
		Data     []CSNMusicSearchNew `json:"data"`
		Rows     string              `json:"rows"`
		Page     int                 `json:"page"`
		RowTotal int                 `json:"row_total"`
	} `json:"music"`
	MusicPlayback struct {
		Data     []interface{} `json:"data"`
		Rows     int           `json:"rows"`
		Page     int           `json:"page"`
		RowTotal int           `json:"row_total"`
	} `json:"music_playback"`
	Video struct {
		Data     []CSNVideoSearch `json:"data"`
		Rows     string           `json:"rows"`
		Page     int              `json:"page"`
		RowTotal int              `json:"row_total"`
	} `json:"video"`
	Artist struct {
		Data     []CSNArtistSearch `json:"data"`
		Rows     string            `json:"rows"`
		Page     int               `json:"page"`
		RowTotal int               `json:"row_total"`
	} `json:"artist"`
	Album struct {
		Data     []CSNAlbumSearch `json:"data"`
		Rows     string           `json:"rows"`
		Page     int              `json:"page"`
		RowTotal int              `json:"row_total"`
	} `json:"album"`
}

type CSNMusicInfoResp struct {
	MusicInfo CSNMusicInfo    `json:"music_info"`
	TrackList []CSNMusicAlbum `json:"track_list"`
	Related   struct {
		MusicTotal int `json:"music_total"`
		MusicList  []struct {
			MusicID         string `json:"music_id"`
			CatID           string `json:"cat_id"`
			CatLevel        string `json:"cat_level"`
			MusicTitleURL   string `json:"music_title_url"`
			MusicTitle      string `json:"music_title"`
			MusicArtist     string `json:"music_artist"`
			MusicBitrate    string `json:"music_bitrate"`
			MusicLength     string `json:"music_length"`
			MusicWidth      string `json:"music_width"`
			MusicHeight     string `json:"music_height"`
			MusicThumbsTime string `json:"music_thumbs_time"`
			MusicDownloads  string `json:"music_downloads"`
			ThumbnailURL    string `json:"thumbnail_url"`
		} `json:"music_list"`
	} `json:"related"`
	Recent struct {
		MusicTotal int `json:"music_total"`
		MusicList  []struct {
			MusicID             string `json:"music_id"`
			CatID               string `json:"cat_id"`
			CatLevel            string `json:"cat_level"`
			CoverID             string `json:"cover_id"`
			MusicTitleURL       string `json:"music_title_url"`
			MusicTitle          string `json:"music_title"`
			MusicArtist         string `json:"music_artist"`
			MusicDownloads      string `json:"music_downloads"`
			MusicListen         string `json:"music_listen"`
			MusicBitrate        string `json:"music_bitrate"`
			MusicWidth          string `json:"music_width"`
			MusicHeight         string `json:"music_height"`
			MusicLength         string `json:"music_length"`
			MusicDownloadsToday string `json:"music_downloads_today"`
			MusicThumbsTime     string `json:"music_thumbs_time"`
			MusicDeleted        string `json:"music_deleted"`
			ThumbnailURL        string `json:"thumbnail_url"`
		} `json:"music_list"`
	} `json:"recent"`
	Artist struct {
		MusicTotal int `json:"music_total"`
		MusicList  []struct {
			MusicID         string `json:"music_id"`
			CatID           string `json:"cat_id"`
			CatLevel        string `json:"cat_level"`
			MusicTitleURL   string `json:"music_title_url"`
			MusicTitle      string `json:"music_title"`
			MusicArtist     string `json:"music_artist"`
			MusicBitrate    string `json:"music_bitrate"`
			MusicLength     string `json:"music_length"`
			MusicWidth      string `json:"music_width"`
			MusicHeight     string `json:"music_height"`
			MusicThumbsTime string `json:"music_thumbs_time"`
			MusicDownloads  string `json:"music_downloads"`
			ThumbnailURL    string `json:"thumbnail_url"`
		} `json:"music_list"`
	} `json:"artist"`
}

type DlAble struct {
	Link map[string]string
}

type CSNUrlInfo struct {
	Url      string
	Kind     int
	Category string
	UrlName  string
}

// interface
type CSNObjectSearch interface {
	Print()
	GetLink() string
	GetID() interface{}
	GetInfo() (CSNMusicInfo, error)
}

type CSNObjectInfo interface {
	PrintLinks(prefix bool, opt int)
}
