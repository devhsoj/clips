package data

// Login is a struct representing the login/signup data from the forms Login.svelte & Signup.svelte.
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
