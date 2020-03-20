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

// AddUser creates a new user in Firebase Auth
func (f *Firebase) AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// EmailVerification verifies the email of the user
func (f *Firebase) EmailVerification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// LogoutUser logs the user out of Firebase Auth
func (f *Firebase) LogoutUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// ResetPassword resets the password of the user
func (f *Firebase) ResetPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// UserLogin logs the user in in Firebase Auth
func (f *Firebase) UserLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
