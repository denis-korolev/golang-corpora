package lemma

import "parser/app/lemma/entities"

type ListResponse struct {
	Data []entities.Lemma `json:"data"`
}
