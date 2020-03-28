package music

import (
	"log"
	"musicMaestro/internal/domain"
	"musicMaestro/internal/network"
	"musicMaestro/internal/persistence"
	"musicMaestro/internal/token"
	"musicMaestro/internal/utils"
	"strings"
)

type TrackService struct {
	tokenService *token.Service
}

const url = "https://api.spotify.com/v1/me/tracks?offset=0&limit=50"

func (service *TrackService) FetchAllUserTracks() {
	persistence.DeleteAllTracks()

	var tracks []domain.Track
	var tracksUrl = url
	for {
		var userTracks []UserTrack

		trackResponse := service.getUserTracks(tracksUrl)

		userTracks = append(userTracks, trackResponse.Items...)

		tracks = append(tracks, service.mapUserTracks(userTracks)...)

		if trackResponse.Next == "" {
			break
		} else {
			tracksUrl = trackResponse.Next
		}
	}
	persistence.SaveTracks(tracks)
}

func (service *TrackService) GetAllUserTracks() []domain.Track {
	tracks := persistence.RetrieveTracks()
	return tracks
}

func (service *TrackService) getUserTracks(url string) *GetUserTracksResponse {
	mapper := NewTrackResponseMapper()
	response, err := network.Get(url, strings.NewReader(""), mapper, service.tokenService.GetAuthorizationToken())

	if err != nil {
		log.Fatal(err)
	}

	trackResponse := response.(*GetUserTracksResponse)
	return trackResponse
}

func (service *TrackService) mapUserTracks(userTracks []UserTrack) []domain.Track {
	var tracks []domain.Track

	for i := 0; i < len(userTracks); i++ {
		userTrackInfo := userTracks[i]
		tracks = append(tracks, service.createTrack(userTrackInfo))
	}

	return tracks
}

func (service *TrackService) createTrack(info UserTrack) domain.Track {
	userTrack := info.Track
	id := userTrack.Id
	name := userTrack.Name

	track := domain.NewTrack(id, name)
	track.AddedAt = utils.ParseDateTime(info.AddedAt)
	track.Popularity = userTrack.Popularity
	track.Preview = userTrack.PreviewUrl
	track.Urls = domain.NewUrls(userTrack.Href, userTrack.ExternalUrls.Spotify, userTrack.Uri)

	track.Artists = addTrackArtists(info)
	track.Album = addTrackAlbum(info)

	return *track
}

func addTrackArtists(info UserTrack) []domain.Artist {
	var artists []domain.Artist
	userTrack := info.Track

	for i := 0; i < len(info.Track.Artists); i++ {
		currentArtist := userTrack.Artists[i]
		id := currentArtist.Id
		name := currentArtist.Name
		urls := domain.NewUrls(currentArtist.Href, currentArtist.ExternalUrls.Spotify, currentArtist.Uri)
		artist := domain.NewArtist(id, name, urls)
		artists = append(artists, *artist)
	}

	return artists
}

func addTrackAlbum(info UserTrack) *domain.Album {
	userTrack := info.Track
	albumInfo := userTrack.Album

	id := albumInfo.Id
	name := albumInfo.Name
	numberOfTracks := albumInfo.TotalTracks
	album := domain.NewAlbum(id, name, numberOfTracks)
	album.ReleaseDate = utils.ParseDate(albumInfo.ReleaseDate, albumInfo.ReleaseDatePrecision)
	album.Urls = domain.NewUrls(albumInfo.Href, albumInfo.ExternalUrls.Spotify, albumInfo.Uri)
	album.AlbumArt = addAlbumArt(info)

	return album
}

func addAlbumArt(info UserTrack) []domain.AlbumArt {
	var albumArts []domain.AlbumArt
	images := info.Track.Album.Images

	for i := 0; i < len(images); i++ {
		albumArtInfo := images[i]
		albumArt := domain.NewAlbumArt(albumArtInfo.Width, albumArtInfo.Height, albumArtInfo.Url)
		albumArts = append(albumArts, *albumArt)
	}

	return albumArts
}

func NewTrackService(tokenService *token.Service) *TrackService {
	return &TrackService{tokenService: tokenService}
}
