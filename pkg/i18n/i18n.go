package i18n

import (
	"path/filepath"
	"runtime"
	"sync"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

var (
	bundle = i18n.NewBundle(language.English)
	once   sync.Once
)

func Initi18n() {
	once.Do(func() {
		_, f, _, _ := runtime.Caller(0)
		dir := filepath.Dir(f)
		bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
		bundle.LoadMessageFile(filepath.Join(dir, "../../assets/i18n/zh-TW.yaml"))
	})
}

func ZhTW(msgID string, data ...interface{}) (string, error) {
	localizer := i18n.NewLocalizer(bundle, "zh-TW")
	cfg := &i18n.LocalizeConfig{
		MessageID: msgID,
	}
	if len(data) > 0 {
		cfg.TemplateData = data[0]
	}

	return localizer.Localize(cfg)
}
