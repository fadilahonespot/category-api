package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Category struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CategoryStore struct {
	categories []Category
	nextId     int
}

func NewCategoryStore() *CategoryStore {
	now := time.Now()
	dummyCategories := []Category{
		{
			Id:          1,
			Name:        "Elektronik",
			Description: "Perangkat elektronik dan peralatan",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			Id:          2,
			Name:        "Pakaian",
			Description: "Pakaian dan aksesori pakaian",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			Id:          3,
			Name:        "Makanan & Minuman",
			Description: "Makanan dan minuman",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			Id:          4,
			Name:        "Buku",
			Description: "Buku dan bahan bacaan",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			Id:          5,
			Name:        "Olahraga & Outdoor",
			Description: "Peralatan olahraga dan peralatan outdoor",
			CreatedAt:   now,
			UpdatedAt:   now,
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
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()
	cs.categories = append(cs.categories, category)
	return category
}

func (cs *CategoryStore) Update(id int, updated Category) (*Category, error) {
	for i := range cs.categories {
		if cs.categories[i].Id == id {
			updated.Id = id
			updated.CreatedAt = cs.categories[i].CreatedAt
			updated.UpdatedAt = time.Now()
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

var store = NewCategoryStore()

func Handler(w http.ResponseWriter, r *http.Request) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	path := r.URL.Path
	path = strings.TrimPrefix(path, "/.netlify/functions/category-api")

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

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

	// Update request path and serve
	r.URL.Path = path
	router.ServeHTTP(w, r)
}

func main() {
	// Untuk development lokal
	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)

	http.HandleFunc("/", Handler)
	http.ListenAndServe(port, nil)
}
