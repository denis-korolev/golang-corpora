package entities

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
		Lemma []struct {
			ID  string `xml:"id,attr"`
			Rev string `xml:"rev,attr"`
			L   struct {
				T string `xml:"t,attr"`
				G []struct {
					V string `xml:"v,attr"`
				} `xml:"g"`
			} `xml:"l"`
			F []struct {
				T string `xml:"t,attr"`
				G []struct {
					V string `xml:"v,attr"`
				} `xml:"g"`
			} `xml:"f"`
		} `xml:"lemma"`
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

type LemmaJson struct {
	ID  string `json:"ID"`
	Rev string `json:"Rev"`
	L   struct {
		T string `json:"T"`
		G []struct {
			V string `json:"V"`
		} `json:"G"`
	} `json:"L"`
	F []struct {
		T string `json:"T"`
		G []struct {
			V string `json:"V"`
		} `json:"G"`
	} `json:"F"`
}
