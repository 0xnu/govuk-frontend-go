package error_summary

type Item struct {
	Href string
	Text string
}

type Model struct {
	Title string
	Items []Item
}
