package models

type Joke struct {
	JokeNumber int    `json:"jokeNumber"`
	Joke       string `json:"joke"`
}

// In-memory database
var JokesDB = []Joke{
	{JokeNumber: 1, Joke: "Why don't scientists trust atoms? Because they make up everything!"},
	{JokeNumber: 2, Joke: "Did you hear about the mathematician who's afraid of negative numbers? He'll stop at nothing to avoid them."},
	{JokeNumber: 3, Joke: "Why don't skeletons fight each other? They don't have the guts."},
	{JokeNumber: 4, Joke: "Why did the scarecrow win an award? Because he was outstanding in his field!"},
}
