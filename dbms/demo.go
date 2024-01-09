package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"time"
)

var genres = []string{"Action", "Comedy", "Drama", "Thriller", "Sci-Fi", "Documentary"}
var types = []string{"Show", "Movie"}

type Media struct {
	Title           string
	Type            string
	Rating          float64
	Length          float64
	Year            float64
	Popularity      float64
	CriticalAcclaim float64
	AudienceTarget  float64
	Genre           string
}

func roundToOneDecimal(num float64) float64 {
	return math.Round(num*10) / 10
}

func generateRandomMedia(count int) []Media {
	rand.Seed(time.Now().UnixNano())
	medias := make([]Media, count)

	for i := 0; i < count; i++ {
		genre := genres[rand.Intn(len(genres))]
		typ := types[rand.Intn(len(types))]
		title := fmt.Sprintf("%s %s %d", genre, typ, i+1)
		medias[i] = Media{
			Title:           title,
			Type:            typ,
			Rating:          roundToOneDecimal(rand.Float64() * 5),
			Length:          roundToOneDecimal(rand.Float64()),
			Year:            roundToOneDecimal(0.5 + rand.Float64()*0.5),
			Popularity:      roundToOneDecimal(rand.Float64()),
			CriticalAcclaim: roundToOneDecimal(rand.Float64() * 5),
			AudienceTarget:  roundToOneDecimal(rand.Float64()),
			Genre:           genre,
		}
	}

	return medias
}

func createMediaVector(media Media) VectorEntry {
	attributes := []float64{media.Rating, media.Length, media.Year, media.Popularity, media.CriticalAcclaim, media.AudienceTarget}

	encodedGenre := oneHotEncode(media.Genre, genres)

	encodedType := oneHotEncode(media.Type, types)

	finalVector := append(attributes, encodedGenre...)
	finalVector = append(finalVector, encodedType...)

	mediaData, _ := json.Marshal(media)

	return VectorEntry{Vector: finalVector, Key: media.Title, Data: json.RawMessage(mediaData)}
}
