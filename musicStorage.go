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

type SavedAlbum struct {
	Key    string
	Title  string
	Artist string
}

func ListAlbums() []SavedAlbum {
	/*
		let's use a redis list so we get ordered items
	*/
	var albums = make([]SavedAlbum, 10, 100)

	//var vals []interface{}
	vals, err := redis.Strings(c.Do("LRANGE", LISTKEY, 0, -1))
	if err != nil {
		fmt.Println("Uh oh, Do")
	}
	fmt.Printf("%#v", vals)

	return albums
}

func AddAlbum(album SavedAlbum) (err error) {
	c.Send("MULTI")

	c.Send("rpush", LISTKEY, album)
	_, err = c.Do("EXEC")
	return
}

func RemoveAlbum(album SavedAlbum) (err error) {
	c.Send("MULTI")
	c.Send("lrem", LISTKEY, 0, album)
	_, err = c.Do("EXEC")
	return
}
