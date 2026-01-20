package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"GoWebapitest/internal/adapter/handler"
	"GoWebapitest/internal/adapter/repository"
	oracle "GoWebapitest/internal/adapter/repository/Oracle"
	postgres "GoWebapitest/internal/adapter/repository/Postgres"
	jwt "GoWebapitest/internal/adapter/token"
	"GoWebapitest/internal/core/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var NumDeclarationPM int = 0
var NumDeclarationENG int = 0
var NumDeclarationENC int = 0
var NumDeclarationDEB int = 0

func main() {
	fmt.Println("Initialisation du serveur")

	err := godotenv.Load()
	if err != nil {
		log.Println("Aucun fichier .env trouvé")
	}

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN non défini")
	}

	db, err := oracle.ConnectDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	pgdb := postgres.ConnectDB()

	// Initialiser le service JWT
	jwtSecret := os.Getenv("JWT_SECRET")

	tokenService := jwt.NewJWTService(jwtSecret, 1*time.Hour)

	repoPP := repository.NewPersonnePhysiqueRepo(db)
	servicePP := service.NewPersonnePhysiqueService(repoPP)
	HandlerPersonnePhysique := handler.HandlerPersonnePhysique(servicePP)

	repoPM := repository.NewPersonneMoralRepo(db)
	servicePM := service.NewPersonneMoralService(repoPM)
	HandlerPersonneMorale := handler.HandlerPersonneMoral(servicePM)

	repoENG := repository.NewEngagementRepo(db)
	serviceENG := service.NewEngagementService(repoENG)
	HandlerEngagement := handler.HandlerEngagement(serviceENG)

	repoENC := repository.NewEncoursRepo(db)
	serviceENC := service.NewEncoursService(repoENC)
	HandlerEncours := handler.HandlerEncours(serviceENC)

	repoDEB := repository.NewCompteDebiteursRepo(db)
	serviceDEB := service.NewCompteDebiteursService(repoDEB)
	HandlerCompteDebiteurs := handler.HandlerCompteDebiteurs(serviceDEB)

	AuthRepo := repository.NewUserRepo(pgdb)
	AuthService := service.NewAuthService(AuthRepo, tokenService)
	HandlerAuth := handler.NewAuthHandler(AuthService)

	UserRepo := repository.NewUserRepo(pgdb)
	UserService := service.NewUserService(UserRepo)
	HandlerUser := handler.HandlerUser(UserService)

	AddUserRepo := repository.NewUserRepo(pgdb)
	AddUserService := service.NewUserService(AddUserRepo)
	HandlerAddUser := handler.HandlerAddUser(AddUserService)

	ModifyStatusRepo := repository.NewUserRepo(pgdb)
	ModifyStatusService := service.NewUserService(ModifyStatusRepo)
	HandlerModifyStatus := handler.ModifyStatus(ModifyStatusService)

	HistoriqueRepo := repository.NewHistoriqueRepository(pgdb)
	HistoriqueService := service.NewHistoriqueService(HistoriqueRepo)


	//API avec Gin
	router := gin.Default()

	// Routes publiques (sans JWT)
	router.POST("/login", HandlerAuth.Login)

	// Routes protégées (avec JWT)
	protected := router.Group("/declaration")
	protected.Use(handler.JWTMiddleware(tokenService),handler.HistoriqueMiddleware(HistoriqueService))
	{
		protected.GET("/personnephysique", gin.WrapH(HandlerPersonnePhysique))
		protected.GET("/personnemorale", gin.WrapH(HandlerPersonneMorale))
		protected.GET("/engagements", gin.WrapH(HandlerEngagement))
		protected.GET("/encours", gin.WrapH(HandlerEncours))
		protected.GET("/comptedebiteurs", gin.WrapH(HandlerCompteDebiteurs))
	}

	protectedUser := router.Group("/users")
	protectedUser.Use(handler.JWTMiddleware(tokenService),handler.HistoriqueMiddleware(HistoriqueService))
	{
		protectedUser.GET("/", gin.WrapH(HandlerUser))
		protectedUser.POST("/register", gin.WrapH(HandlerAddUser))
		protectedUser.POST("/updatestate", gin.WrapH(HandlerModifyStatus))
		// protectedUser.POST("/historique", gin.WrapH(HandlerHistorique))
	}

	log.Fatal(router.Run("10.0.20.32:8181"))
}
