package twitterdownloader

import (
	"os"
	"testing"
)

const uri = "/ext_tw_video/1035056301657583617/pu/pl/540x960/Vyiq7upZ9HjHDruX.m3u8"

func TestPlayList(t *testing.T) {

	file, err := os.Open("./testdata/playlist.m3u8")
	if err != nil {
		t.Errorf("Open test file failed")
	}
	str, err := playList(file)
	if err != nil {
		t.Errorf(err.Error())
	}
	if str != uri {
		t.Errorf("Expected uri is %s, but return[%s]", uri, str)
	}
}

func TestVideoList(t *testing.T) {
	ts0 := "/ext_tw_video/1035056301657583617/pu/vid/0/3000/540x960/DJ0xXa4TpDDejYlV.ts"
	ts1 := "/ext_tw_video/1035056301657583617/pu/vid/3000/6000/540x960/Ij-0cqTubNDJbd1E.ts"
	ts2 := "/ext_tw_video/1035056301657583617/pu/vid/6000/10043/540x960/kljEHBu_sl5nP90X.ts"

	file, err := os.Open("./testdata/video.m3u8")
	if err != nil {
		t.Errorf("Open test file failed")
	}
	str, err := videoList(file)
	if err != nil {
		t.Errorf(err.Error())
	}
	if ts0 != str[0] {
		t.Errorf("parse video m3u8 0 failed")
	}
	if ts1 != str[1] {
		t.Errorf("parse video m3u8 1 failed")
	}
	if ts2 != str[2] {
		t.Errorf("parse video m3u8 2 failed")
	}
}
