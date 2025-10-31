package movie

type Movie struct{
	ID string `json:"id"`
	IsBn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	FirstName string `json:"firstname"` 
	LastName string `json:"lastname"`
}