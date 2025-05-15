package routers

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	r.GET("/playlists", controllers.GetPlaylists)
	r.GET("/playlists/:id", controllers.GetPlaylist)
	r.POST("/playlists", controllers.CreatePlaylist)
	r.PUT("/playlists/:id", controllers.UpdatePlaylist)
	r.DELETE("/playlists/:id", controllers.DeletePlaylist)
	r.POST("/playlists/:id/tracks", controllers.AddTrackToPlaylist)

	r.GET("/tracks", controllers.GetTracks)
	r.GET("/tracks/:id", controllers.GetTrack)
	r.POST("/tracks", controllers.CreateTrack)
	r.PUT("/tracks/:id", controllers.UpdateTrack)
	r.DELETE("/tracks/:id", controllers.DeleteTrack)
}
