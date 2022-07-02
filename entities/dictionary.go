package entities

import "fmt"

type Dictionary struct {
	Version   string `xml:"version,attr"`
	Revision  string `xml:"revision,attr"`
	Grammemes struct {
		Grammeme []struct {
			Parent      string `xml:"parent,attr"`
			Name        string `xml:"name"`
			Alias       string `xml:"alias"`
			Description string `xml:"description"`
		} `xml:"grammeme"`
	} `xml:"grammemes"`
	Restrictions struct {
		Restr []struct {
			Type string `xml:"type,attr"`
			Auto string `xml:"auto,attr"`
			Left struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"left"`
			Right struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"right"`
		} `xml:"restr"`
	} `xml:"restrictions"`
	Lemmata struct {
		Lemma []Lemma `xml:"lemma"`
	} `xml:"lemmata"`
	LinkTypes struct {
		Type []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
		} `xml:"type"`
	} `xml:"link_types"`
	Links struct {
		Link []struct {
			ID   string `xml:"id,attr"`
			From string `xml:"from,attr"`
			To   string `xml:"to,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
	} `xml:"links"`
}

type Lemma struct {
	ID  string           `xml:"id,attr" json:"ID"`
	Rev string           `xml:"rev,attr" json:"Rev"`
	L   LemmaAttribute   `xml:"l" json:"L"`
	F   []LemmaAttribute `xml:"f" json:"F"`
}

func (l Lemma) ShortString() string {
	return fmt.Sprintf("Lemma id=%s rev=%s term=%s", l.ID, l.Rev, l.L.T)
}

type LemmaAttribute struct {
	T string `xml:"t,attr" json:"T"`
	G []struct {
		V string `xml:"v,attr" json:"V"`
	} `xml:"g" json:"G"`
}
