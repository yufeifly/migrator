package migration

import (
	"github.com/gin-gonic/gin"
	"github.com/yufeifly/proxyd/model"
)

var DefaultChkPDirPrefix = "/var/lib/docker/containers/"

// PushCheckpoint push checkpoint to destination and deliver restore request
func PushCheckpoint(migOpts model.MigrationOpts) error {
	// 这里应该与目的端有一个交互，以便知道目的端是否真实接收到检查点
	// 这里应该作为文件传输的客户端

	return nil
}

// ReceiveCheckpointAndRestore get checkpoint from source node and restore from it
func ReceiveCheckpointAndRestore(c *gin.Context) {
	// 文件传输的服务器端

	// 启动容器
	c.JSON(200, gin.H{
		"result": "success",
	})
}
