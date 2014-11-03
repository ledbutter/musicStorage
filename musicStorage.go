package musicStorage

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	//	"strings"
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
	Key    string `redis:"key"`
	Title  string `redis:"title"`
	Artist string `redis:"artist"`
}

// Lists all the wishlist albums currently being tracked.
func ListAlbums() []SavedAlbum {
	/*
		let's use a redis list so we get ordered items
	*/
	vals, err := redis.Values(c.Do("SORT", LISTKEY,
		"BY", "nosort",
		"GET", "*->key",
		"GET", "*->title",
		"GET", "*->artist"))
	if err != nil {
		fmt.Println("Uh oh, Do")
		return nil
	} else {

		var albums []SavedAlbum

		if err := redis.ScanSlice(vals, &albums); err != nil {
			panic(err)
		}

		return albums
	}
}

// Adds a new album to the list of wishlist albums.
func AddAlbum(album SavedAlbum) (err error) {
	c.Send("MULTI")

	c.Send("HMSET", "album:"+album.Key, "key", album.Key, "title", album.Title, "artist", album.Artist)
	c.Send("RPUSH", LISTKEY, "album:"+album.Key)
	//c.Send("rpush", LISTKEY, album.Key, album.Title, album.Artist)
	_, err = c.Do("EXEC")
	return
}

// Removes an album from the list of wishlist albums.
func RemoveAlbum(album SavedAlbum) (err error) {
	c.Send("MULTI")
	c.Send("HDEL", "album:"+album.Key)
	c.Send("lrem", LISTKEY, 0, album.Key)
	_, err = c.Do("EXEC")
	return
}
