package redis

import (
	"fmt"
	"github.com/yufeifly/proxyd/model"

	"github.com/gin-gonic/gin"
	"github.com/yufeifly/proxyd/container"
	"github.com/yufeifly/proxyd/dal"
	"github.com/yufeifly/proxyd/migration"
)

func Get(c *gin.Context) {
	key := c.Request.URL.Query().Get("key")

	val := dal.GetKV(key)
	c.JSON(200, gin.H{
		"value: ": val,
	})
}

func Set(c *gin.Context) {
	key := c.Request.URL.Query().Get("key")
	val := c.Request.URL.Query().Get("val")
	fmt.Printf("about to set key: %v, val: %v", key, val)

	dal.SetKV(key, val)

	c.JSON(200, gin.H{
		"result": "ok",
	})
}

func MigrateRedis(c *gin.Context) {
	// 获取请求参数
	containerName := c.Request.URL.Query().Get("container")
	checkpointID := c.Request.URL.Query().Get("checkpointID")
	destIP := c.Request.URL.Query().Get("destIP")
	checkpointDir := c.Request.URL.Query().Get("checkpointDir")
	if checkpointDir == "" {
		checkpointDir = migration.DefaultChkPDirPrefix + migration.GetContainerFullID(containerName) + "/" + checkpointID
	}

	chOpts := model.CheckpointOpts{
		CheckPointID:  checkpointID,
		CheckPointDir: checkpointDir,
	}

	// 1 创建检查点
	err := container.CreateCheckpoint(c, chOpts)
	if err != nil {
		fmt.Printf("CreateCheckpoint err: %v\n", err)
		container.ReportErr(c, err)
		panic(err)
	}

	// 2 推送检查点到目标节点
	err = migration.PushCheckpoint(checkpointDir, destIP)
	if err != nil {
		fmt.Printf("Push Checkpoint err: %v\n", err)
		container.ReportErr(c, err)
		panic(err)
	}

	c.JSON(200, gin.H{
		"result": "success",
	})
}
