package unigreek

import (
	"fmt"
	"testing"
)

var TestStrings = [][2]string{
	[2]string{"a", "α"},
	[2]string{"abg", "αβγ"},
	[2]string{"*abg", "Αβγ"},
	[2]string{"a/bg", "άβγ"},
	[2]string{"plh\\n *milhsi/wn.", "πλὴν Μιλησίων."},
	[2]string{"mh=nin a)/eide qea\\ *phlhi+a/dew *)axilh=os", "μῆνιν ἄειδε θεὰ Πηληϊάδεω Ἀχιλῆος"},
	[2]string{"[*)axilh=os] ", "[Ἀχιλῆος] "},
	[2]string{"tw=|", "τῷ"},
	[2]string{"*t*w=|", "Τῼ͂"},
	[2]string{"ss", "σς"},
	[2]string{"ss ss", "σς σς"},
}

func TestConvert(t *testing.T) {
	for _, tcase := range TestStrings {
		bcode := tcase[0]
		ucode := tcase[1]

		converted, err := Convert(bcode)

		if converted != ucode || err != nil {
			fmt.Printf(`FAILURE: Conversion of "%s" should return: "%s"` + "\n",
				bcode, ucode)
			fmt.Printf(`         Instead returns: "%s"` + "\n", converted)

			if err != nil {
				fmt.Printf(`         Error value: "%s"` + "\n", err.Error())
			}
                        fmt.Printf("\n")
			t.Fail()
		}
	}
}
