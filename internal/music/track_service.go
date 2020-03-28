package music

import (
	"log"
	"musicMaestro/internal/network"
	"musicMaestro/internal/token"
	"strings"
)

type TrackService struct {
	tokenService *token.Service
}

const url = "https://api.spotify.com/v1/me/tracks?offset=0&limit=50"

func (service *TrackService) FetchUserTracks() {
	var tracks []UserTrack

	mapper := NewTrackResponseMapper()
	response, err := network.Get(url, strings.NewReader(""), mapper, service.tokenService.GetAuthorizationToken())

	if err != nil {
		log.Fatal(err)
	}

	trackResponse := response.(*GetUserTracksResponse)

	tracks = append(tracks, trackResponse.Items...)
}

func NewTrackService(tokenService *token.Service) *TrackService {
	return &TrackService{tokenService: tokenService}
}
