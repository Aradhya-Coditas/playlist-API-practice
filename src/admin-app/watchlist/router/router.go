package router

import (
	"admin-app/watchlist/business"
	serviceConstant "admin-app/watchlist/commons/constants"
	"admin-app/watchlist/docs"
	"admin-app/watchlist/handlers"
	"admin-app/watchlist/repositories"
	"fmt"
	"net/http"
	genericConstants "omnenest-backend/src/constants"

	"omnenest-backend/src/utils/configs"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

// GetRouter is used to get the router configured with the middlewares and the routes
func GetRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.New()
	router.Use(middlewares...)
	router.Use(gin.Recovery())

	applicationConfig := configs.GetApplicationConfig()

	// Swagger
	docs.SwaggerInfo.Host = fmt.Sprintf("%v", applicationConfig.SwaggerConfig.SwaggerHost)
	docs.SwaggerInfo.Schemes = []string{genericConstants.SwaggerInfoHttpSchemeConfig, genericConstants.SwaggerInfoHttpsSchemeConfig}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{genericConstants.PostMethod, genericConstants.GetMethod, genericConstants.PutMethod, genericConstants.DeleteMethod, genericConstants.PatchMethod},
		AllowHeaders:     []string{genericConstants.AllowHeaderOriginConfig, genericConstants.Authorization},
		ExposeHeaders:    []string{genericConstants.ExposeHeaderContentLengthConfig},
		AllowCredentials: true,
		MaxAge:           300 * time.Second,
	}))

	// Swagger Configuration
	router.GET(serviceConstant.SwaggerRoute, ginSwagger.WrapHandler(swaggerFiles.Handler))

	// useMocks := applicationConfig.AppConfig.UseMocks
	useDBMocks := applicationConfig.AppConfig.UseDBMocks

	// Initialize controller and services
	// nestApiWrapper := nestIntegration.GetNestWrapper(useMocks)

	// metrics.Init()
	// router.GET(serviceConstant.Metrics, gin.WrapH(prometheusHandler.Handler()))

	// enableUIBFFEncDec := applicationConfig.AppConfig.EnableUIBFFEncDec
	// enableRateLimit := applicationConfig.AppConfig.EnableRateLimit

	// jwtUtils := jwtUtils.NewJwtTokenUtils()
	// jwtMiddleware := authorization.AuthorizeJWtToken(jwtUtils, genericConstants.MiddlewareEncryptionKey)

	deleteWatchlistRepository := repositories.GetDeleteWatchlistRepository(useDBMocks)
	deleteWatchlistService := business.NewDeleteWatchlistService(deleteWatchlistRepository)
	deleteWatchlistController := handlers.NewDeleteWatchlistController(deleteWatchlistService)

	getWatchListRepository := repositories.NewGetWatchiistScripsRepository(useDBMocks)
	getWatchListService := business.NewGetWatchlistScripsService(getWatchListRepository)
	getWatchListController := handlers.NewGetWatchlistScripsController(getWatchListService)

	getWatchlistRepository := repositories.GetGetWatchlistRepository(useDBMocks)
	getWatchlistService := business.NewGetWatchlistService(getWatchlistRepository)
	getWatchlistController := handlers.NewGetWatchlistController(getWatchlistService)

	createUserPlaylistRepository := repositories.GetCreateUserPlaylistRepository(useDBMocks)
	createUserService := business.NewCreateUserPlaylistService(createUserPlaylistRepository)
	createUserController := handlers.NewPlaylistController(createUserService)

	modifyPlaylistsRepository := repositories.NewAdPlaylistRepository(useDBMocks)
	modifyPlaylistsService := business.NewAdPlaylistService(modifyPlaylistsRepository)
	modifyPlaylistsController := handlers.NewAdPlaylistHandler(modifyPlaylistsService)

	v1Routes := router.Group(genericConstants.RouterV1Config)
	{
		v1Routes.GET(serviceConstant.WatchlistServiceHealthCheck, func(c *gin.Context) {
			response := map[string]string{
				genericConstants.ResponseMessageKey: genericConstants.BFFResponseSuccessMessage,
			}
			c.JSON(http.StatusOK, response)
		})
		// if enableRateLimit {
		// 	v1Routes.Use(rateLimitMiddleware.RateLimitMiddleware(limiter))
		// }
		// v1Routes.Use(headerMiddleware.HeaderCheck(serviceConstant.ServiceName), jwtMiddleware)
		v1Routes.DELETE(serviceConstant.DeleteWatchlist, deleteWatchlistController.HandleDeleteWatchlist)
		// v1Routes.Use(headerMiddleware.HeaderCheck(serviceConstant.ServiceName), metricsMiddleware.Metric(), jwtMiddleware)
		// if enableUIBFFEncDec {
		// 	v1Routes.Use(encryptMiddleware, decryptMiddleware)
		// }
		v1Routes.POST(serviceConstant.GetWatchListScrips, getWatchListController.HandleGetWatchlistScrips)
		v1Routes.POST(serviceConstant.GetWatchlist, getWatchlistController.HandleGetWatchlist)
		v1Routes.POST(serviceConstant.CreatePlaylist, createUserController.HandleCreatePlaylist)
		v1Routes.POST(serviceConstant.ModifyPlaylist, modifyPlaylistsController.HandleModifyPlaylistSongs())
	}
	return router
}
