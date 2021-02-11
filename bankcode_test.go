package bankcode

import (
	"context"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func testClient(t *testing.T) *Client {
	t.Helper()
	apiKey := os.Getenv("BANKCODE_API_KEY")
	c, err := NewClient(WithAPIKey(apiKey))
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func TestListBanks(t *testing.T) {
	t.Parallel()
	client := testClient(t)
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
	client := testClient(t)
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

func TestListBranch(t *testing.T) {
	t.Parallel()
	client := testClient(t)
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

func TestGetBranch(t *testing.T) {
	t.Parallel()
	client := testClient(t)
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

func TestVersion(t *testing.T) {
	t.Parallel()
	client := testClient(t)
	version, err := client.GetVersion(context.Background(), "")
	if err != nil {
		t.Fatal(err)
	}
	if version == "" {
		t.Fatal("version response is empty")
	}
}

func TestGetParameter_Bank(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		title  string
		input  *GetParameter
		expect *Bank
	}{
		{
			title: "apiKey",
			input: &GetParameter{
				APIKey: os.Getenv("BANKCODE_API_KEY"),
			},
			expect: &Bank{
				Code:          "0001",
				Name:          "みずほ銀行",
				HalfWidthKana: "ﾐｽﾞﾎｷﾞﾝｺｳ",
				FullWidthKana: "ミズホギンコウ",
				Hiragana:      "みずほぎんこう",
			},
		},
		{
			title: "fields",
			input: &GetParameter{
				Fields: []string{"code", "name"},
			},
			expect: &Bank{
				Code: "0001",
				Name: "みずほ銀行",
			},
		},
	}
	for _, tc := range testcases {
		tc := tc
		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()
			client := testClient(t)
			bank, err := client.GetBank(context.Background(), "0001", tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(bank, tc.expect); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestGetParameter_Branch(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		title  string
		input  *GetParameter
		expect []*Branch
	}{
		{
			title: "apiKey",
			input: &GetParameter{
				APIKey: os.Getenv("BANKCODE_API_KEY"),
			},
			expect: []*Branch{
				{
					Code:          "093",
					Name:          "本店",
					HalfWidthKana: "ﾎﾝﾃﾝ",
					FullWidthKana: "ホンテン",
					Hiragana:      "ほんてん",
				},
			},
		},
		{
			title: "fields",
			input: &GetParameter{
				Fields: []string{"code", "name"},
			},
			expect: []*Branch{
				{
					Code: "093",
					Name: "本店",
				},
			},
		},
	}
	for _, tc := range testcases {
		tc := tc
		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()
			client := testClient(t)
			branches, err := client.GetBranch(context.Background(), "0000", "093", tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(branches, tc.expect); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestListParameter_Banks(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		title  string
		input  *ListParameter
		expect *Banks
	}{
		{
			title: "apiKey",
			input: &ListParameter{
				APIKey: os.Getenv("BANKCODE_API_KEY"),
				Limit:  1,
			},
			expect: &Banks{
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
				NextCursor: "nob4gqVAoYiC4vLAuyWZuA70I-m_Gn96eQd4N6hveLc",
			},
		},
		{
			title: "limit",
			input: &ListParameter{
				Limit: 1,
			},
			expect: &Banks{
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
				NextCursor: "nob4gqVAoYiC4vLAuyWZuA70I-m_Gn96eQd4N6hveLc",
			},
		},
		{
			title: "fields",
			input: &ListParameter{
				Limit:  1,
				Fields: []string{"code", "name"},
			},
			expect: &Banks{
				Data: []*Bank{
					{
						Code: "0000",
						Name: "日本銀行",
					},
				},
				Size:       1,
				Limit:      1,
				HasNext:    true,
				NextCursor: "nob4gqVAoYiC4vLAuyWZuA70I-m_Gn96eQd4N6hveLc",
			},
		},
		{
			title: "cursor",
			input: &ListParameter{
				Limit:  1,
				Cursor: "nob4gqVAoYiC4vLAuyWZuA70I-m_Gn96eQd4N6hveLc",
			},
			expect: &Banks{
				Data: []*Bank{
					{
						Code:          "0001",
						Name:          "みずほ銀行",
						HalfWidthKana: "ﾐｽﾞﾎｷﾞﾝｺｳ",
						FullWidthKana: "ミズホギンコウ",
						Hiragana:      "みずほぎんこう",
					},
				},
				Size:       1,
				Limit:      1,
				HasNext:    true,
				NextCursor: "GqahEPgLJcmJvi0n-ogqc8WmBgJaWA-WiZ7aDkl5MuQgFmdp2aO_ak6yQhTQ_kjg",
				HasPrev:    true,
				PrevCursor: "w_JTggxW8eHCK9fMFbcwyA",
			},
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.title, func(t *testing.T) {
			client := testClient(t)
			banks, err := client.ListBanks(context.Background(), tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(banks, tc.expect, cmpopts.IgnoreFields(Banks{}, "Version")); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
