package entities

type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

type Location struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type PersonDTO struct {
	Gender   string   `json:"gender"`
	Name     Name     `json:"name"`
	Email    string   `json:"email"`
	Location Location `json:"location"`
}
