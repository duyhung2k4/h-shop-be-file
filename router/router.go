package router

import (
	"app/controller"
	"net/http"
	"time"

	middlewares "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func Router() http.Handler {
	app := chi.NewRouter()

	app.Use(middlewares.RequestID)
	app.Use(middlewares.RealIP)
	app.Use(middlewares.Logger)
	app.Use(middlewares.Recoverer)
	app.Use(middlewares.Timeout(60 * time.Second))

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	app.Use(cors.Handler)

	// middlewares := middlewares.NewMiddlewares()
	imageController := controller.NewImageProductController()
	app.Route("/file/api/v1", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			res := map[string]interface{}{
				"mess": "done",
			}
			render.JSON(w, r, res)
		})

		r.Route("/public", func(public chi.Router) {
			public.Route("/image-product", func(imageProduct chi.Router) {
				imageProduct.Get("/productId", imageController.GetImagesbyProductId)
				imageProduct.Get("/avatar", imageController.GetAvatarByProductId)
			})
		})
	})

	return app
}
