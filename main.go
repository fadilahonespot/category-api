package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Category struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type CategoryStore struct {
	categories []Category
	nextId     int
}

func NewCategoryStore() *CategoryStore {
	dummyCategories := []Category{
		{
			Id:          1,
			Name:        "Elektronik",
			Description: "Perangkat elektronik dan peralatan",
		},
		{
			Id:          2,
			Name:        "Pakaian",
			Description: "Pakaian dan aksesori pakaian",
		},
		{
			Id:          3,
			Name:        "Makanan & Minuman",
			Description: "Makanan dan minuman",
		},
		{
			Id:          4,
			Name:        "Buku",
			Description: "Buku dan bahan bacaan",
		},
		{
			Id:          5,
			Name:        "Olahraga & Outdoor",
			Description: "Peralatan olahraga dan peralatan outdoor",
		},
	}

	return &CategoryStore{
		categories: dummyCategories,
		nextId:     6,
	}
}

func (cs *CategoryStore) GetAll() []Category {
	return cs.categories
}

func (cs *CategoryStore) GetByID(id int) (*Category, error) {
	for i := range cs.categories {
		if cs.categories[i].Id == id {
			return &cs.categories[i], nil
		}
	}
	return nil, fmt.Errorf("category not found")
}

func (cs *CategoryStore) Create(category Category) Category {
	category.Id = cs.nextId
	cs.nextId++
	cs.categories = append(cs.categories, category)
	return category
}

func (cs *CategoryStore) Update(id int, updated Category) (*Category, error) {
	for i := range cs.categories {
		if cs.categories[i].Id == id {
			updated.Id = id
			cs.categories[i] = updated
			return &cs.categories[i], nil
		}
	}
	return nil, fmt.Errorf("category not found")
}

func (cs *CategoryStore) Delete(id int) error {
	for i := range cs.categories {
		if cs.categories[i].Id == id {
			cs.categories = append(cs.categories[:i], cs.categories[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("category not found")
}

func main() {
	store := NewCategoryStore()

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Server is running",
			"time":    time.Now().Format(time.RFC3339),
		})
	})

	router.GET("/categories", func(c *gin.Context) {
		categories := store.GetAll()
		c.JSON(http.StatusOK, categories)
	})

	router.POST("/categories", func(c *gin.Context) {
		var category Category
		if err := c.ShouldBindJSON(&category); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		if category.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
			return
		}

		created := store.Create(category)
		c.JSON(http.StatusCreated, created)
	})

	router.PUT("/categories/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var category Category
		if err := c.ShouldBindJSON(&category); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		updated, err := store.Update(id, category)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, updated)
	})

	router.GET("/categories/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		category, err := store.GetByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, category)
	})

	router.DELETE("/categories/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		if err := store.Delete(id); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	})

	// Gunakan PORT dari environment variable (untuk Render) atau default ke 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	port = ":" + port

	router.Run(port)
}
