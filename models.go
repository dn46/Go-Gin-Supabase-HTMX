package main

type Books struct {
	Title  string  `form:"title"`
	Author string  `form:"author"`
	Price  float64 `form:"price"`
	ISBN   string  `form:"isbn"`
}
