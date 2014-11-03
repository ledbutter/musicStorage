package musicStorage

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//based in part on https://coderwall.com/p/unklzq

const (
	ADDRESS = "127.0.0.1:6379"
	LISTKEY = "albums"
)

var (
	c, e = redis.Dial("tcp", ADDRESS)
)

// Represents a stored wishlist album.
type SavedAlbum struct {
	Key    string
	Title  string
	Artist string
}

// Lists all the wishlist albums currently being tracked.
func ListAlbums() []SavedAlbum {
	/*
		let's use a redis list so we get ordered items
	*/
	vals, err := redis.Values(c.Do("LRANGE", LISTKEY, 0, -1))
	if err != nil {
		fmt.Println("Uh oh, Do")
		return nil
	} else {
		//i assume there is a more elegant way to do this, but that way is unknown
		//to this feeble mind
		albums := make([]SavedAlbum, len(vals))
		for i, v := range vals {
			albums[i] = v.(SavedAlbum)
		}
		return albums
	}
}

// Adds a new album to the list of wishlist albums.
func AddAlbum(album SavedAlbum) (err error) {
	c.Send("MULTI")

	c.Send("rpush", LISTKEY, album)
	_, err = c.Do("EXEC")
	return
}

// Removes an album from the list of wishlist albums.
func RemoveAlbum(album SavedAlbum) (err error) {
	c.Send("MULTI")
	c.Send("lrem", LISTKEY, 0, album)
	_, err = c.Do("EXEC")
	return
}
