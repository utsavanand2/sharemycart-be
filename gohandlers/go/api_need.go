// Package swagger ...
/*
 * Share my Cart
 *
 * Stay home. Stay safe. I'll bring your groceries along as I get mine.
 *
 * API version: 0.1.0
 * Contact: sharemycart@beimir.net
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"net/http"
)

// AddNeedList adds a needs list
func (f *Firebase) AddNeedList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// GetNeedLists gets the need list
func (f *Firebase) GetNeedLists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// GetNeeds retrieves all items which are needed to be supplied
func (f *Firebase) GetNeeds(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// DeleteNeeds removes all the needed items
func (f *Firebase) DeleteNeeds(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// DeleteNeedList deletes the need list
func (f *Firebase) DeleteNeedList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// AddNeededItem adds a needed item
func (f *Firebase) AddNeededItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// AddNeededItemToShoppingListItem adds a needed item similar like one which already exists on another shopping list
func (f *Firebase) AddNeededItemToShoppingListItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// DeleteNeededItem deletes an existing needed item
func (f *Firebase) DeleteNeededItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// UpdateNeededItem updates an existing needed item
func (f *Firebase) UpdateNeededItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
