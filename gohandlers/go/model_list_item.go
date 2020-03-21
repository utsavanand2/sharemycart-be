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

// ListItem is the shopping list item
type ListItem struct {
	ID string `json:"id,omitempty"`

	Name string `json:"name"`

	Amount float32 `json:"amount"`

	Unit string `json:"unit,omitempty"`
}
