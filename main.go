package main

import (
	"html/template"
	"net/http"

	"github.com/Masterminds/sprig/v3"
	"github.com/gin-gonic/gin"
)

type Instruction struct {
	Title       string
	Description string
}

type Nutrition struct {
	Name  string
	Value string
}

func main() {
	r := gin.Default()
	templ := template.Must(
		template.New("base").Funcs(sprig.FuncMap()).ParseGlob("views/*.html"),
	)
	r.SetHTMLTemplate(templ)
	r.Static("/assets", "./assets")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"title": "Recipe Page",
			"preparation": []string{
				"Total: Approximately 10 minutes",
				"Preparation: 5 minutes",
				"Cooking: 5 minutes",
			},
			"ingredients": []string{
				"2-3 large eggs",
				"Salt, to taste",
				"Pepper, to taste",
				"1 tablespoon of butter or oil",
				"Optional fillings: cheese, diced vegetables, cooked meats, herbs",
			},
			"instructions": []Instruction{
				{
					Title: "Beat the eggs",
					Description: `In a bowl, beat the eggs with a pinch of salt and pepper until they are well mixed.
      You can add a tablespoon of water or milk for a fluffier texture.`,
				},
				{
					Title:       "Heat the pan",
					Description: `Place a non-stick frying pan over medium heat and add butter or oil.`,
				},
				{
					Title: "Cook the omelette",
					Description: `Once the butter is melted and bubbling, pour in the eggs. Tilt the pan to ensure
the eggs evenly coat the surface.`,
				},
				{
					Title: "Add fillings (optional)",
					Description: `When the eggs begin to set at the edges but are still slightly runny in the 
middle, sprinkle your chosen fillings over one half of the omelette.
`,
				},
				{
					Title: "Fold and serve",
					Description: `As the omelette continues to cook, carefully lift one edge and fold it over the
fillings. Let it cook for another minute, then slide it onto a plate.
`,
				},
				{
					Title:       "Enjoy",
					Description: `Serve hot, with additional salt and pepper if needed.`,
				},
			},
			"nutrition": []Nutrition{
				{
					Name:  "Calories",
					Value: "277kcal",
				},
				{
					Name:  "Carbs",
					Value: "0g",
				},
				{
					Name:  "Protein",
					Value: "20g",
				},
				{
					Name:  "Fat",
					Value: "22g",
				},
			},
		})
	})

	r.Run(":8080")
}
