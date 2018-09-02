package twitterdownloader

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/grafov/m3u8"
)

func saveFile(filename string, reader io.ReadCloser) error {
	body, err := ioutil.ReadAll(reader)

	if err = ioutil.WriteFile(filename, body, os.ModePerm); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func playList(reader io.ReadCloser) (string, error) {
	// saveFile("playlist.m3u8", reader)
	// f, err := os.Open("playlist.m3u8")
	// if err != nil {
	// 	panic(err)
	// }

	p, listType, err := m3u8.DecodeFrom(reader, true)
	if err != nil {
		panic(err)
	}

	switch listType {
	case m3u8.MEDIA:
		mediapl := p.(*m3u8.MediaPlaylist)
		fmt.Printf("%+v\n", mediapl)
	case m3u8.MASTER:
		masterpl := p.(*m3u8.MasterPlaylist)
		// fmt.Printf("%+v\n", masterpl)
		var target struct {
			index     int
			bandwidth uint32
		}
		target.index = -1
		target.bandwidth = 0

		for k, v := range masterpl.Variants {
			fmt.Println("URI:", v.URI)
			fmt.Println("Resolution:", v.Resolution)
			fmt.Println("Bandwidth:", v.Bandwidth)
			if target.bandwidth < v.Bandwidth {
				target.bandwidth = v.Bandwidth
				target.index = k
			}
		}
		if target.index > 0 {
			return masterpl.Variants[target.index].URI, nil
		}
		return masterpl.Variants[0].URI, nil
	}

	return "", errors.New("Not find playlist m3u8")
}

func videoList(reader io.ReadCloser) ([]string, error) {
	// f, err := os.Open("video.m3u8")
	// if err != nil {
	// 	panic(err)
	// }
	// p, listType, err := m3u8.DecodeFrom(bufio.NewReader(f), true)
	p, listType, err := m3u8.DecodeFrom(reader, true)
	if err != nil {
		panic(err)
	}
	var s []string
	switch listType {
	case m3u8.MEDIA:
		mediapl := p.(*m3u8.MediaPlaylist)
		// fmt.Printf("%+v\n", mediapl)
		for _, v := range mediapl.Segments {
			if v == nil {
				fmt.Println("Nil")
				break
			}
			fmt.Println("URI:", v.URI)
			fmt.Println("Resolution:", v.Duration)
			s = append(s, v.URI)
		}
		return s, nil
	case m3u8.MASTER:
		masterpl := p.(*m3u8.MasterPlaylist)
		fmt.Printf("%+v\n", masterpl)
	}
	return nil, errors.New("Not find video list")
}
