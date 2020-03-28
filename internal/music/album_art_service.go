package music

import (
	"musicMaestro/internal/domain"
	"musicMaestro/internal/image"
	"musicMaestro/internal/persistence"
	"sync"
	"time"
)

func UpdateAlbumArtOf(tracks []domain.Track) {
	var wg sync.WaitGroup
	for i := 0; i < len(tracks); i++ {
		wg.Add(1)
		go func(track domain.Track) {
			time.Sleep(5 * time.Millisecond)
			defer wg.Done()
			updateAlbumArtOf(track)
		}(tracks[i])
	}
	wg.Wait()

	persistence.DeleteAllTracks()
	persistence.SaveTracks(tracks)
}

func updateAlbumArtOf(track domain.Track) {
	albumArt := track.Album.AlbumArt
	imageUrl := albumArt[0].ImageUrl
	base64Image := image.DownloadImage(imageUrl)
	albumArt[0].ImageData = base64Image
}
