package main

import (
	"fmt"
	"os"
	"path/filepath"

	//install this 3rd party package: https://github.com/nicksnyder/go-i18n
	"github.com/nicksnyder/go-i18n/i18n"
)

func getCWD() (dir string, err error) {
	dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	return
}

func loadTranslationFiles(locale string) (err error) {
	cwd, err := getCWD()
	if err != nil {
		return
	}
	i18n.MustLoadTranslationFile(cwd + "/i18n/en-US.all.json")
	i18n.LoadTranslationFile(cwd + fmt.Sprintf("/i18n/%s.all.json", locale))
	return
}

func main() {

	err := loadTranslationFiles("zh-TW")
	if err != nil {
		panic(err)
	}

	T, err := i18n.Tfunc("zh-TW", "zh-TW", "en-US")
	if err != nil {
		panic(err)
	}
	fmt.Println("=== ZH-TW demo ===")
	fmt.Println(T("settings_title", map[string]interface{}{
		"SettingTarget": "一個設定目標",
	}))
	fmt.Println("")

	ST, err := i18n.Tfunc("zh-CN", "zh-CN", "en-US")
	if err != nil {
		panic(err)
	}

	fmt.Println("=== ZH CN fall back demo ===")
	fmt.Println(ST("settings_title", map[string]interface{}{
		"SettingTarget": "for a setting target",
	}))
	fmt.Println("")

	err = loadTranslationFiles("zh-CN")
	if err != nil {
		panic(err)
	}

	anotherT, err := i18n.Tfunc("zh", "zh-CN", "zh-TW", "en-US")
	if err != nil {
		panic(err)
	}

	fmt.Println("=== ZH demo for mapping lang precedence (zh-CN first) ===")
	fmt.Println(anotherT("settings_title", map[string]interface{}{
		"SettingTarget": "一个设定目标",
	}))
	fmt.Println("")

	thridT, err := i18n.Tfunc("zh", "zh-TW", "zh-CN", "en-US")
	if err != nil {
		panic(err)
	}

	fmt.Println("=== ZH demo for mapping lang precedence (zh-TW first) ===")
	fmt.Println(thridT("settings_title", map[string]interface{}{
		"SettingTarget": "一個設定目標",
	}))
	fmt.Println("")

	unSupportT, err := i18n.Tfunc("ar-JO", "ar-JO", "en-US")
	if err != nil {
		panic(err)
	}

	fmt.Println("=== ar-JO demo for unsupported lang fallback to en-US ===")
	fmt.Println(unSupportT("settings_title", map[string]interface{}{
		"SettingTarget": "for a setting target",
	}))
	fmt.Println("")
}
