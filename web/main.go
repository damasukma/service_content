package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	uuid "github.com/satori/go.uuid"
	"github.com/service_content/model"
	"github.com/service_content/repository"
)

type handle struct{}

func (h *handle) CreateNews(client model.EventStoreClient) gin.HandlerFunc {
	handler := func(ctx *gin.Context) {
		var (
			fields    repository.Article
			newsParam model.ArticleParam
		)

		if err := ctx.ShouldBindJSON(&fields); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		newsParam.Author = fields.Author
		newsParam.Title = fields.Title
		newsParam.Body = fields.Body

		data, _ := json.Marshal(newsParam)

		evId := uuid.NewV4().String()
		agId := uuid.NewV4().String()

		eventMap := &model.EventParam{
			EventId:       evId,
			EventType:     "news-created",
			AggregateId:   agId,
			AggregateType: "news",
			EventData:     string(data),
		}

		result, err := client.CreateEvent(ctx, eventMap)

		if err != nil {
			log.Fatal(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": result})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success"})
	}
	return handler
}
func main() {
	handle := &handle{}

	conn, err := grpc.Dial("127.0.0.1:4040", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	client := model.NewEventStoreClient(conn)

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version": 1,
		})
	})

	router.POST("/news", handle.CreateNews(client))

	srv := &http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to initialize server: %v \n", err)
		}
	}()

	log.Printf("Listening on port %v\n", srv.Addr)
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	log.Println("Shutting down server ...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}
