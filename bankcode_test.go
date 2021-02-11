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
	banks, err := client.ListBanks(context.Background(), &ListParameter{
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
	bank, err := client.GetBank(context.Background(), "0001", &GetParameter{})
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

func TestGetBranch(t *testing.T) {
	t.Parallel()
	client := NewClient()
	branches, err := client.GetBranch(context.Background(), "0000", "093", &GetParameter{})
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(branches, []*Branch{
		{
			Code:          "093",
			Name:          "本店",
			HalfWidthKana: "ﾎﾝﾃﾝ",
			FullWidthKana: "ホンテン",
			Hiragana:      "ほんてん",
		},
	}); diff != "" {
		t.Fatal(diff)
	}
}

func TestListBranch(t *testing.T) {
	t.Parallel()
	client := NewClient()
	branches, err := client.ListBranch(context.Background(), "0001", &ListParameter{
		Limit: 2,
	})
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(branches, &Branches{
		Data: []*Branch{
			{
				Code:          "001",
				Name:          "東京営業部",
				HalfWidthKana: "ﾄｳｷﾖｳｴｲｷﾞﾖｳﾌﾞ",
				FullWidthKana: "トウキヨウエイギヨウブ",
				Hiragana:      "とうきようえいぎようぶ",
			},
			{
				Code:          "001",
				Name:          "東京都庁公営企業出張所",
				HalfWidthKana: "ﾄｳｷﾖｳﾄﾁﾖｳｺｳｴｲｼﾕﾂﾁﾖｳｼﾞﾖ",
				FullWidthKana: "トウキヨウトチヨウコウエイシユツチヨウジヨ",
				Hiragana:      "とうきようとちようこうえいしゆつちようじよ",
			},
		},
		Size:       2,
		Limit:      2,
		HasNext:    true,
		NextCursor: "",
	}, cmpopts.IgnoreFields(Branches{}, "NextCursor", "Version")); diff != "" {
		t.Fatal(diff)
	}
}
func TestVersion(t *testing.T) {
	t.Parallel()
	client := NewClient()
	version, err := client.GetVersion(context.Background(), "")
	if err != nil {
		t.Fatal(err)
	}
	if version == "" {
		t.Fatal("version response is empty")
	}
}
