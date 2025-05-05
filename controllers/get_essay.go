package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"wenxinhexuan/api/response"
	GenerateGrpc "wenxinhexuan/grpc/go"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// GetGeneratedEssay godoc
// @Summary 获取生成的文章
// @Description 根据传入的 topic 参数，通过 gRPC 流式调用生成文章，并使用 SSE 向前端实时推送生成结果
// @Tags 文章, SSE, gRPC
// @Accept json
// @Produce text/event-stream
// @Param topic query string true "生成文章的主题"
// @Success 200 {object} response.Response "通信成功（通过code来判断具体情况）"
// @Router /api/essay [get]
func (uc *UserController) GetGeneratedEssay(c *gin.Context) {
	topic := c.Query("topic")

	// c.Writer 实现了 http.ResponseWriter，可以用于 Write()、WriteHeader() 这些方法
	// c.Writer.(http.Flusher) 试图将 c.Writer 转换为 http.Flusher 类型
	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		log.Println("Streaming unsupported")
		response.FailWithMessage("Streaming unsupported", c)
		return
	}

	// 连接Python gRPC服务器
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("无法连接 gRPC 服务器: %v", err)
		response.FailWithMessage("无法连接 gRPC 服务器", c)
		return
	}
	defer conn.Close()

	client := GenerateGrpc.NewGenerateClient(conn)

	// 设置 gRPC 调用的超时时间，防止请求长时间无响应
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()

	// 发送 gRPC 请求
	stream, err := client.GenerateStream(ctx, &GenerateGrpc.GenerateRequest{Topic: topic})
	if err != nil {
		log.Fatalf("请求失败: %v", err)
	}

	// 逐步接收流式响应并推送 SSE
	for {
		res, err := stream.Recv()
		if err != nil {
			break
		}
		_, _ = c.Writer.Write([]byte("data= " + res.Response + "\n\n"))
		flusher.Flush() // 立刻推送数据到前端
	}
}
