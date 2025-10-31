package movie

var movies []Movie
func SeedData(){
	movies = []Movie{
		{ID: "1", IsBn: "438227", Title: "Movie One", Director: &Director{FirstName: "Shahriar", LastName: "Stranger"}},
		{ID: "2", IsBn: "739012", Title: "Echoes of Time", Director: &Director{FirstName: "Hakim", LastName: "Akash"}},
		{ID: "3", IsBn: "562934", Title: "Quantum Mirage", Director: &Director{FirstName: "Liam", LastName: "Noir"}},
		{ID: "4", IsBn: "823675", Title: "Fragments of Tomorrow", Director: &Director{FirstName: "Ava", LastName: "Lumen"}},
		{ID: "5", IsBn: "982341", Title: "Silent Horizon", Director: &Director{FirstName: "Noah", LastName: "Vale"}},
		{ID: "6", IsBn: "114589", Title: "The Last Algorithm", Director: &Director{FirstName: "Elena", LastName: "Cross"}},
		{ID: "7", IsBn: "7781178", Title: "Quats Lots", Director: &Director{FirstName: "Jhon", LastName: "Deh"}},
	}
}