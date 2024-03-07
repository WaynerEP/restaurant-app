package menu

import "github.com/WaynerEP/restaurant-app/server/models/common"

// NutritionalValue .
type NutritionalValue struct {
	common.ModelId
	ItemID             uint    `json:"dishId" gorm:"not null"` // ID del plato asociado
	Calories           float64 `json:"calories"`               // Calorías del plato
	Carbohydrates      float64 `json:"carbohydrates"`          // Carbohidratos del plato
	Proteins           float64 `json:"proteins"`               // Proteínas del plato
	Fats               float64 `json:"fats"`                   // Grasas del plato
	SaturatedFat       float64 `json:"saturatedFat"`           // Grasas saturadas del plato
	TransFat           float64 `json:"transFat"`               // Grasas trans del plato
	MonounsaturatedFat float64 `json:"monounsaturatedFat"`     // Grasas monoinsaturadas del plato
	PolyunsaturatedFat float64 `json:"polyunsaturatedFat"`     // Grasas poliinsaturadas del plato
	Cholesterol        float64 `json:"cholesterol"`            // Colesterol del plato
	Sodium             float64 `json:"sodium"`                 // Sodio del plato
	Potassium          float64 `json:"potassium"`              // Potasio del plato
	Fiber              float64 `json:"fiber"`                  // Fibra del plato
	Sugar              float64 `json:"sugar"`                  // Azúcar del plato
	VitaminA           float64 `json:"vitaminA"`               // Vitamina A del plato
	VitaminC           float64 `json:"vitaminC"`               // Vitamina C del plato
	Calcium            float64 `json:"calcium"`                // Calcio del plato
	Iron               float64 `json:"iron"`                   // Hierro del plato
	CommonServing      string  `json:"commonServing"`          // Porción común del plato (ej. "100g")
	Allergens          string  `json:"allergens"`              // Lista de alérgenos del plato
	IsVegetarian       *bool   `json:"isVegetarian"`           // Indica si el plato es vegetariano
	IsVegan            *bool   `json:"isVegan"`                // Indica si el plato es vegano
	IsGlutenFree       *bool   `json:"isGlutenFree"`           // Indica si el plato es libre de gluten
	IsOrganic          *bool   `json:"isOrganic"`              // Indica si el plato es orgánico
}
