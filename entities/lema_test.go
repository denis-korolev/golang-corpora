package entities

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"testing"
)

func TestLemmaEncoding(t *testing.T) {
	// todo add asserts package.
	xmlData := []byte(`
<lemma id="2" rev="2">
	<l t="ёж">
		<g v="NOUN"/>
		<g v="inan"/>
		<g v="masc"/>
	</l>
	<f t="ёж">
		<g v="sing"/>
		<g v="nomn"/>
	</f>
	<f t="ежа">
		<g v="sing"/>
		<g v="gent"/>
	</f>
</lemma>
`)

	lemma := new(Lemma)

	err := xml.Unmarshal(xmlData, lemma)
	if err != nil {
		t.Error(err)
	}

	if lemma.L.T != "ёж" {
		t.Error(fmt.Errorf("ожидается что lema.L.T == ёж"))
	}

	if lemma.F[0].T != "ёж" {
		t.Error(fmt.Errorf("ожидается что lema.F[0].T == ёж"))
	}

	if lemma.F[0].G[0].V != "sing" {
		t.Error(fmt.Errorf("ожидается что lemma.F[0].G[0].V == sing"))
	}

	dataBytes, err := json.MarshalIndent(&lemma, "", "  ")
	data := string(dataBytes)

	expect := `{
  "ID": "2",
  "Rev": "2",
  "L": {
    "T": "ёж",
    "G": [
      {
        "V": "NOUN"
      },
      {
        "V": "inan"
      },
      {
        "V": "masc"
      }
    ]
  },
  "F": [
    {
      "T": "ёж",
      "G": [
        {
          "V": "sing"
        },
        {
          "V": "nomn"
        }
      ]
    },
    {
      "T": "ежа",
      "G": [
        {
          "V": "sing"
        },
        {
          "V": "gent"
        }
      ]
    }
  ]
}`

	if data != expect {
		t.Errorf("expect: \n%s\n actual: \n%s\n", expect, data)
	}
}
