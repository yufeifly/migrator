package container

import (
	"fmt"
	"testing"
)

func TestGetImageRepoTags(t *testing.T) {
	got, _ := GetImageRepoTags("bb22")
	fmt.Println("image name: ", got)
}

func TestInspect(t *testing.T) {
	got, _ := Inspect("bb19")
	fmt.Println("containerJson: ", got.ID)
}

func TestGetImageByImageID(t *testing.T) {
	got, err := GetImageByImageID("7e4d58f0e5f3")
	if err != nil {
		fmt.Printf("TestGetImageByImageID err: %v\n", err)
	} else {
		fmt.Printf("TestGetImageByImageID: %v\n", got)
	}
}
