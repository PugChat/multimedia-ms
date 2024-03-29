package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pugchat_multimedia/infrastructure/api/handler"
	"github.com/pugchat_multimedia/infrastructure/persistance"
)

func main() {
	conn := persistance.NewSQLDB()
	defer conn.Close()
	entryRepo := persistance.NewEntryRepository(conn)
	tagRepo := persistance.NewTagRepository(conn)

	eH := handler.NewEntryHandler(entryRepo, tagRepo)
	tH := handler.NewTagHandler(tagRepo)

	r := gin.Default()

	entriesGroup := r.Group("/entries")
	entriesGroup.GET("", eH.GetEntries)
	entriesGroup.POST("", eH.CreateEntry)
	entriesGroup.GET(":id", eH.GetEntry)
	entriesGroup.PUT(":id", eH.UpdateEntry)
	entriesGroup.DELETE(":id", eH.DeleteEntry)

	tagsGroup := r.Group("/tags")
	tagsGroup.GET("", tH.GetTags)
	tagsGroup.POST("", tH.CreateTag)
	tagsGroup.GET(":id", tH.GetTag)
	tagsGroup.PUT(":id", tH.UpdateTag)
	tagsGroup.DELETE(":id", tH.DeleteTag)

	r.Run(":8081")
}
