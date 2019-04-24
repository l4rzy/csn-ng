package main

import (
	"fmt"
)

const (
	OP_MUSIC  = 1 << 1
	OP_VIDEO  = 1 << 2
	OP_ARTIST = 1 << 3
	OP_ALBUM  = 1 << 4
)

const (
	CSN_HOME = "https://chiasenhac.com"
)

type CSNMusic struct {
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
	MusicDownloads   int    `json:"music_downloads"`
}

type CSNVideo struct {
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

type CSNArtist struct {
	ArtistID       int    `json:"artist_id"`
	ArtistNickname string `json:"artist_nickname"`
	ArtistLink     string `json:"artist_link"`
	ArtistCover    string `json:"artist_cover"`
	ArtistAvatar   string `json:"artist_avatar"`
}

type CSNAlbum struct {
	CoverID         int         `json:"cover_id"`
	MusicAlbum      string      `json:"music_album"`
	AlbumLink       string      `json:"album_link"`
	AlbumID         string      `json:"album_id"`
	AlbumBitrate    interface{} `json:"album_bitrate"`
	AlbumArtist     string      `json:"album_artist"`
	AlbumArtistHTML string      `json:"album_artist_html"`
	AlbumCover      string      `json:"album_cover"`
}

type RespJson []struct {
	Q     string `json:"q"`
	Music struct {
		Data     []CSNMusic `json:"data"`
		Rows     string     `json:"rows"`
		Page     int        `json:"page"`
		RowTotal int        `json:"row_total"`
	} `json:"music"`
	MusicPlayback struct {
		Data     []interface{} `json:"data"`
		Rows     int           `json:"rows"`
		Page     int           `json:"page"`
		RowTotal int           `json:"row_total"`
	} `json:"music_playback"`
	Video struct {
		Data     []CSNVideo `json:"data"`
		Rows     string     `json:"rows"`
		Page     int        `json:"page"`
		RowTotal int        `json:"row_total"`
	} `json:"video"`
	Artist struct {
		Data     []CSNArtist `json:"data"`
		Rows     string      `json:"rows"`
		Page     int         `json:"page"`
		RowTotal int         `json:"row_total"`
	} `json:"artist"`
	Album struct {
		Data     []CSNAlbum `json:"data"`
		Rows     string     `json:"rows"`
		Page     int        `json:"page"`
		RowTotal int        `json:"row_total"`
	} `json:"album"`
}

type DlAble struct {
	link map[string]string
}

// interface
type CSNObject interface {
	print()
}

func (m CSNMusic) print() {
	fmt.Printf("%v\n", m.MusicLink)
}

func (v CSNVideo) print() {
	fmt.Printf("%v\n", v.VideoLink)
}

func (a CSNArtist) print() {
	fmt.Printf("%s%v\n", CSN_HOME, a.ArtistLink)
}

func (a CSNAlbum) print() {
	fmt.Printf("%s%v\n", CSN_HOME, a.AlbumLink)
}
