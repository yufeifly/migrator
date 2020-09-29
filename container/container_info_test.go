package container

import (
	"fmt"
	"testing"
)

func TestInspect(t *testing.T) {
	got, _ := Inspect("bb19")
	fmt.Println("containerJson: ", got.ID)
}

func TestGetImageByImageID(t *testing.T) {
	got, err := GetImageByImageID("b5166e9de47d")
	if err != nil {
		fmt.Printf("TestGetImageByImageID err: %v\n", err)
	} else {
		fmt.Printf("TestGetImageByImageID: %v\n", got)
	}
}

func TestGetImageByContainer(t *testing.T) {
	got, err := GetImageByContainer("b5166e9de47d")
	if err != nil {
		fmt.Printf("TestGetImageByContainer err: %v\n", err)
	} else {
		fmt.Printf("TestGetImageByContainer: %v\n", got)
	}
}

func TestGetPortBinding(t *testing.T) {
	err := GetPortBinding("c8dc2bbf7c9a")
	if err != nil {
		fmt.Printf("TestGetPortBinding err: %v\n", err)
	} else {
		fmt.Printf("TestGetPortBinding: pass")
	}
}
