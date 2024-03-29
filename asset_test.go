package asset

import (
	"fmt"
	"testing"
)

//TestGetFileVersion  check for versioning up filepath
func TestGetFileVersionWithFolder(t *testing.T) {
	res := UpFileName("/Users/buck/Desktop/test/v02/Screen_v02.mp4", "%03d")

	if res != "/Users/buck/Desktop/test/v003/Screen_v003.mp4" {
		t.Errorf("Should be '/Users/buck/Desktop/test/v003/Screen_v003.mp4' -> got : %s", res)
	}
	fmt.Println(res)
}

func TestGetFileVersionWithoutFolder(t *testing.T) {
	res := UpFileName("/Users/buck/Desktop/test/Screen_v02.mp4", "%03d")

	if res != "/Users/buck/Desktop/test/Screen_v003.mp4" {
		t.Errorf("Should be '/Users/buck/Desktop/test/Screen_v003.mp4' -> got : %s", res)
	}
	fmt.Println(res)
}

func TestGetFileVersionWithTwoDigits(t *testing.T) {
	res := UpFileName("/Users/buck/Desktop/test/Screen_v02.mp4", "%02d")

	if res != "/Users/buck/Desktop/test/Screen_v03.mp4" {
		t.Errorf("Should be '/Users/buck/Desktop/test/Screen_v03.mp4' -> got : %s", res)
	}
	fmt.Println(res)
}

func TestGetFilesFromPath(t *testing.T) {
	expected := []string{"/Users/buck/Desktop/test/Screen_v01.mp4",
		"/Users/buck/Desktop/test/Screen_v02.mp4",
		"/Users/buck/Desktop/test/Screen_v03.mp4",
		"/Users/buck/Desktop/test/Screen_v04.mp4",
		"/Users/buck/Desktop/test/Screen_v05.mp4"}
	res := getFilesFromPath("/Users/buck/Desktop/test/Screen_v01.mp4", ".mp4", []string{}, false)
	if len(res) != len(expected) {
		t.Errorf("Should match some files : %s \n but got : %s  ", expected, res)
	}
	fmt.Println(res)
}

func TestGetAepFilesFromPath(t *testing.T) {
	expected := []string{"/Users/buck/Desktop/current/mock/work/current/TestDesign/Production/TEST010/Work/AE/TEST010_v003.aep",
		"/Users/buck/Desktop/current/mock/work/current/TestDesign/Production/TEST010/Work/AE/_archive/TEST010_v002.aep",
		"/Users/buck/Desktop/current/mock/work/current/TestDesign/Production/TEST010/Work/AE/_archive/TEST010_v001.aep"}

	res := getFilesFromPath("/Users/buck/Desktop/current/mock/work/current/TestDesign/Production/TEST010/Work/AE/TEST010_v003.aep", AfterEffect.Extension, AfterEffect.Ignore, true)
	if len(res) != len(expected) {
		t.Errorf("Should contains some files got: %s", res)
	}
	fmt.Println(res)
}

func TestGetAepFilesFromPathWithFilter(t *testing.T) {
	expected := []string{"/Users/buck/Desktop/current/mock/work/current/TestDesign/Production/TEST010/Work/AE/TEST010_v003.aep"}
	res := getFilesFromPath("/Users/buck/Desktop/current/mock/work/current/TestDesign/Production/TEST010/Work/AE/TEST010_v003.aep", AfterEffect.Extension, AfterEffect.Ignore, false)
	if len(res) != len(expected) {
		t.Errorf("Should contains some files got: %s", res)
	}
	fmt.Println(res)
}
