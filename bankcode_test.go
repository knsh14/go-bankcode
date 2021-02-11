package bankcode

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestListBanks(t *testing.T) {
	t.Parallel()
	client := NewClient()
	banks, err := client.ListBanks(context.Background(), &ListBankRequest{
		Limit: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(banks, &Banks{
		Data: []*Bank{
			{
				Code:          "0000",
				Name:          "日本銀行",
				HalfWidthKana: "ﾆﾂﾎﾟﾝｷﾞﾝｺｳ",
				FullWidthKana: "ニツポンギンコウ",
				Hiragana:      "につぽんぎんこう",
			},
		},
		Size:       1,
		Limit:      1,
		HasNext:    true,
		NextCursor: "",
	}, cmpopts.IgnoreFields(Banks{}, "NextCursor", "Version")); diff != "" {
		t.Fatal(diff)
	}
}

func TestGetBanks(t *testing.T) {
	t.Parallel()
	client := NewClient()
	bank, err := client.GetBank(context.Background(), "0001", &GetBankRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(bank, &Bank{
		Code:          "0001",
		Name:          "みずほ銀行",
		HalfWidthKana: "ﾐｽﾞﾎｷﾞﾝｺｳ",
		FullWidthKana: "ミズホギンコウ",
		Hiragana:      "みずほぎんこう",
	}); diff != "" {
		t.Fatal(diff)
	}
}
