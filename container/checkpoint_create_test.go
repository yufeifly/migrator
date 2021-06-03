package container

import (
	"github.com/docker/docker/api/types"
	"testing"
)

func TestCreateCheckpoint(t *testing.T) {
	chOpts := types.CheckpointCreateOptions{
		CheckpointID:  "cp-redis",
		CheckpointDir: "/tmp",
		Exit:          false,
	}
	err := dockerCli.CheckpointCreate(ctx, "s1.c1", chOpts)
	if err != nil {
		t.Error("CheckpointCreate err")
	}
}
