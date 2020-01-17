package address

import (
	"testing"
)

func TestMakeAddress(t *testing.T) {
	tpub := "tpubD6NzVbkrYhZ4YNL2wZjwhx4VdV81ve5tbwdvB1gdLp3xMGmFiYm226ggo24D22UawGB7co9wmaEGPgq1vyDJVzL3SbjCE2hNXRNEnRxPCSB"
	vpub := "vpub5VopB7T65QAVURGyGhGxaNPChoPmLg6DwpvDfSib5V74LKcXaJFrWqSkDARcHQiUEJ9TKu7dK1BC15c1XXGzhuPWQCjPAeexueGSrmoCrh5"
	xpub := "xpub661MyMwAqRbcFtXgS5sYJABqqG9YLmC4Q1Rdap9gSE8NqtwybGhePY2gZ29ESFjqJoCu1Rupje8YtGqsefD265TMg7usUDFdp6W1EGMcet8"
	zpub := "zpub6rFR7y4Q2AijBEqTUquhVz398htDFrtymD9xYYfG1m4wAcvPhXNfE3EfH1r1ADqtfSdVCToUG868RvUUkgDKf31mGDtKsAYz2oz2AGutZYs"

	cases := []struct {
		masterPub string
		change    bool
		index     int
		wantAddr  string
	}{
		{xpub, false, 0, "12CL4K2eVqj7hQTix7dM7CVHCkpP17Pry3"},
		{xpub, false, 1, "13Q3u97PKtyERBpXg31MLoJbQsECgJiMMw"},
		{xpub, true, 0, "1NwEtFZ6Td7cpKaJtYoeryS6avP2TUkSMh"},
		{xpub, true, 1, "18FcseQ86zCaXzLbgDsH86292xb2EuKtFW"},
		{xpub, true, 9, "1eWNxfQVi6wRti9qqsDPQ9pqZqebpXxwF"},

		{zpub, false, 0, "bc1qcr8te4kr609gcawutmrza0j4xv80jy8z306fyu"},
		{zpub, false, 1, "bc1qnjg0jd8228aq7egyzacy8cys3knf9xvrerkf9g"},
		{zpub, true, 0, "bc1q8c6fshw2dlwun7ekn9qwf37cu2rn755upcp6el"},
		{zpub, true, 1, "bc1qggnasd834t54yulsep6fta8lpjekv4zj6gv5rf"},

		{tpub, false, 0, "mk92jDjQnRmLaTixco74cs6q7PMTUoQ6Hy"},
		{tpub, false, 1, "mke8TJ7M5yyNwoKu77YbJPuXcSWbX6tqkH"},
		{tpub, true, 0, "n43EMvNHLDizjZKtMmuqk81P8DsufWdV9u"},
		{tpub, true, 1, "mh3MDWji77ff2ptEWU4KRQWuSaXQjvkSxb"},

		{vpub, false, 0, "tb1qx70f4lyjslul8vase0nzmey40wcqgrrfk2rm3m"},
		{vpub, false, 1, "tb1qmuzfq06ryvw2wz5agcp3hqyhrl066n8gze4azs"},
		{vpub, true, 0, "tb1q6xcu59s8na4lp3uzdzly26459sdwy3sxnwj6ss"},
		{vpub, true, 1, "tb1q0w7h8r08hm0c7k50p4c2347kuectspx2rs7z7y"},
	}

	for _, tc := range cases {
		got, err := MakeAddress(tc.masterPub, tc.change, tc.index)
		if err != nil {
			t.Errorf("MakeAddress(%q, %v, %d) failed: %v.", tc.masterPub, tc.change, tc.index, err)
			continue
		}
		if got != tc.wantAddr {
			t.Errorf("MakeAddress(%q, %v, %d) returned %s, want %s.", tc.masterPub, tc.change, tc.index, got, tc.wantAddr)
		}
	}
}
