package chatbot

type FormData struct {
	data map[string]string
}

func (f *FormData) GetValue(key string) string {
	return f.data[key]
}

func (f *FormData) Text() string {
	return f.data[TextKey]
}

func (f *FormData) GetMessengerID() string {
	return f.data[MessagenerIDKey]
}

func (f *FormData) GetReplyToken() string {
	return f.data[ReplyTokenKey]
}

func (f *FormData) load(data map[string]string) {
	for k, v := range data {
		f.data[k] = v
	}
}

func (f *FormData) store(key, value string) {
	f.data[key] = value
}

func (f *FormData) Bind(obj interface{}, tag string) error {
	return bind(f.data, obj, tag)
}
