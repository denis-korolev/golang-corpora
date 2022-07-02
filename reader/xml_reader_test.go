package reader

import (
	"strings"
	"sync"
	"testing"
)

func TestXmlReader(t *testing.T) {
	inXml := `
<?xml version="1.0" encoding="utf-8" standalone="yes"?>
<dictionary version="0.92" revision="417150">
    <grammemes></grammemes>
    <restrictions>
        <restr type="obligatory" auto="0">
            <left type="lemma"></left>
            <right type="lemma">POST</right>
        </restr>
    </restrictions>
    <lemmata>
        <lemma id="1" rev="1">
            <l t="ёж">
                <g v="NOUN"/>
                <g v="anim"/>
                <g v="masc"/>
            </l>
            <f t="ёж">
                <g v="sing"/>
                <g v="nomn"/>
            </f>
        </lemma>
        <lemma id="2" rev="2">
            <l t="слон">
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
    </lemmata>
    <link_types>
        <type id="1">ADJF-ADJS</type>
    </link_types>
    <links>
    </links>
</dictionary>
`
	provider := NewXmlLemmaProviderFormStream(strings.NewReader(inXml))

	wg := new(sync.WaitGroup)

	wg.Add(1)
	ch := provider.GetLemmasChan(3, wg)

	expectId := []string{"1", "2"}
	expectTerm := []string{"ёж", "слон"}
	idx := 0

	for lemma := range ch {

		if lemma.ID != expectId[idx] {
			t.Errorf("expect lemma id=%s, actual=%s", expectId[idx], lemma.ID)
		}
		if lemma.L.T != expectTerm[idx] {
			t.Errorf("expect lemma term=%s, actual=%s", expectTerm[idx], lemma.L.T)
		}

		idx++
	}

	if idx != 2 {
		t.Errorf("expect lemma count 2, actual=%d", idx)
	}

	wg.Wait()

}
