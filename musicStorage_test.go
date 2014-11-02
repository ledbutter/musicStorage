package musicStorage

import (
	"math/rand"
	"testing"
)

func TestListAlbums(t *testing.T) {
	items := ListAlbums()
	if items == nil {
		t.Errorf("No items")
	}
}

//this random string code courtesy of http://stackoverflow.com/a/22892986/1346943
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func TestAddAlbum(t *testing.T) {
	album := SavedAlbum{}
	album.Key = randSeq(15)
	album.Title = randSeq(10)
	album.Artist = randSeq(13)
	err := AddAlbum(album)
	if err != nil {
		t.Errorf("Error adding album: %v", err)
	}
	err = RemoveAlbum(album)
	if err != nil {
		t.Errorf("Error deleting album: %v", err)
	}
}
