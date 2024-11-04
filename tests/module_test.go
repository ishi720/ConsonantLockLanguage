package tests

import (
	"testing"
	"myapp/module"
)

func TestConsonantLockLanguage(t *testing.T) {

	result1 := module.ConsonantLockLanguage("こんにちは", "pa")
	expected1 := "ぽんぴぴぱ"

	if result1 != expected1 {
		t.Errorf("文字列が一致しません。Result=%s, Expected=%s", result1, expected1)
	}

	result2 := module.ConsonantLockLanguage("ありがとう", "sa")
	expected2 := "さしさそす"
	if result2 != expected2 {
		t.Errorf("文字列が一致しません。Result=%s, Expected=%s", result2, expected2)
	}

	result3 := module.ConsonantLockLanguage("おやすみなさい", "ga")
	expected3 := "ごがぐぎががぎ"
	if result3 != expected3 {
		t.Errorf("文字列が一致しません。Result=%s, Expected=%s", result3, expected3)
	}
}