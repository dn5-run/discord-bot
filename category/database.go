package category

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

const DATA_PATH = ".data/category.json"

func initDatabase() {
	os.MkdirAll(path.Dir(DATA_PATH), os.ModePerm)
	if _, err := os.Stat(DATA_PATH); err != nil {
		f, err := os.Create(DATA_PATH)
		if err != nil {
			panic(err)
		}

		f.Write([]byte("[]"))
		f.Close()
	}
}

func getCategories() []*Category {
	raw, err := ioutil.ReadFile(DATA_PATH)
	if err != nil {
		panic(err)
	}

	var categories []*Category
	err = json.Unmarshal(raw, &categories)
	if err != nil {
		panic(err)
	}

	return categories
}
func saveCategories(categories []*Category) {
	f, err := os.Create(DATA_PATH)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := json.Marshal(categories)
	if err != nil {
		panic(err)
	}

	f.Write(data)
}

func addCategory(c *Category) {
	categories := getCategories()

	categories = append(categories, c)
	saveCategories(categories)
}

func getCategory(id string) *Category {
	categories := getCategories()

	for _, category := range categories {
		if category.ID == id {
			return category
		}
	}

	return nil
}
