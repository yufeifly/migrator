package container

import (
	"fmt"
	"testing"
)

func TestGetImageRepoTags(t *testing.T) {
	got, _ := GetImageRepoTags("bb19")
	fmt.Println("image name: ", got)
}

func TestInspect(t *testing.T) {
	got, _ := Inspect("bb19")
	fmt.Println("containerJson: ", got.ID)
}
