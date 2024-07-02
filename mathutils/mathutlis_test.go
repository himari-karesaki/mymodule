package mathutils

import (
	"testing"
)

// Add関数に対するテストケースを定義
func TestAdd(t *testing.T) {
	// テストケースのスライスを定義
	// 各ケースは加算の入力値 (a, b) と期待する結果 (expected)
	testCases := []struct {
		a, b, expected int
	}{
		{2, 3, 5},
		{-5, -3, -8},
		{7, -2, 5},
		{0, 5, 5},
		{3, 0, 3},
	}

	//各テストケースに対して繰り返し処理
	for _, tc := range testCases {
		// Add関数を呼び出し、結果をresultに代入
		result := Add(tc.a, tc.b)
		// 結果が期待する値と異なる場合、テスト失敗としてエラーメッセージを表示
		if result != tc.expected {
			t.Errorf("Add(%d, %d) = %d, want %d", tc.a, tc.b, result, tc.expected)
		}
	}
}

func TestSubstract(t *testing.T) {
	// テストケースのスライスを定義
	// 各ケースは加算の入力値 (a, b) と期待する結果 (expected)
	testCases := []struct {
		a, b, expected int
	}{
		{2, 3, -1},
		{-5, -3, -2},
		{7, -2, 9},
		{0, 5, -5},
		{3, 0, 3},
	}

	//各テストケースに対して繰り返し処理
	for _, tc := range testCases {
		// Add関数を呼び出し、結果をresultに代入
		result := Subtract(tc.a, tc.b)
		// 結果が期待する値と異なる場合、テスト失敗としてエラーメッセージを表示
		if result != tc.expected {
			t.Errorf("Add(%d, %d) = %d, want %d", tc.a, tc.b, result, tc.expected)
		}
	}
}
