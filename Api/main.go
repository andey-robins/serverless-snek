package api

// Game is a top level game object
type Game struct {
	ID      string `json:"id"`
	Timeout int32  `json:"timeout"`
}

// Coord is a single coordinate object
type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Battlesnake is a struct for holding a single battlesnake
type Battlesnake struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Health int32   `json:"health"`
	Body   []Coord `json:"body"`
	Head   Coord   `json:"head"`
	Length int32   `json:"length"`
	Shout  string  `json:"shout"`
}

// Board is a struct that holds an entire board
type Board struct {
	Height int           `json:"height"`
	Width  int           `json:"width"`
	Food   []Coord       `json:"food"`
	Snakes []Battlesnake `json:"snakes"`
}

// BattlesnakeInfoResponse is what you respond with to start the game
type BattlesnakeInfoResponse struct {
	APIVersion string `json:"apiversion"`
	Author     string `json:"author"`
	Color      string `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
}

// GameRequest is a struct that does something
type GameRequest struct {
	Game  Game        `json:"game"`
	Turn  int         `json:"turn"`
	Board Board       `json:"board"`
	You   Battlesnake `json:"you"`
}

// MoveResponse is what you respond with to do something
type MoveResponse struct {
	Move  string `json:"move"`
	Shout string `json:"shout,omitempty"`
}
