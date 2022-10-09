package lemma

type ListResponse struct {
	Data []ListItem `json:"data"`
}

type ListItem struct {
	Text string `json:"text" example:"Бугульма"`
}
