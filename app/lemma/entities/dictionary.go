package entities

type Dictionary struct {
	Version   string `xml:"version,attr"`
	Revision  string `xml:"revision,attr"`
	Grammemes struct {
		Grammeme []Grammeme `xml:"grammeme"`
	} `xml:"grammemes"`
	Restrictions struct {
		Restr []Restriction `xml:"restr"`
	} `xml:"restrictions"`
	Lemmata   `xml:"lemmata"`
	LinkTypes struct {
		Type []LinkType `xml:"type"`
	} `xml:"link_types"`
	Links struct {
		Link []Link `xml:"link"`
	} `xml:"links"`
}

type Lemmata struct {
	Lemma []Lemma `xml:"lemma" json:"Lemma"`
}

type Lemma struct {
	ID  string `xml:"id,attr" json:"ID" example:"ID леммы"`
	Rev string `xml:"rev,attr" json:"Rev" example:"номер последней ревизии этой леммы"`
	L   struct {
		T string `xml:"t,attr" json:"T" example:"текст леммы"`
		G []struct {
			V string `xml:"v,attr" json:"V" example:"значение"`
		} `xml:"g" json:"G"`
	} `xml:"l" json:"L"`
	F []struct {
		T string `xml:"t,attr" json:"T" example:"текст"`
		G []struct {
			V string `xml:"v,attr" json:"V" example:"значение"`
		} `xml:"g" json:"G"`
	} `xml:"f" json:"F"`
}

type Restriction struct {
	Type  string `xml:"type,attr"`
	Auto  string `xml:"auto,attr"`
	Left  Left   `xml:"left"`
	Right Right  `xml:"right"`
}

type LinkType struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr"`
}

type Link struct {
	ID   string `xml:"id,attr"`
	From string `xml:"from,attr"`
	To   string `xml:"to,attr"`
	Type string `xml:"type,attr"`
}

type Left struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

type Grammeme struct {
	Parent      string `xml:"parent,attr"`
	Name        string `xml:"name"`
	Alias       string `xml:"alias"`
	Description string `xml:"description"`
}
type Right struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
}
