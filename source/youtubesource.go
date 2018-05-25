package source

type YoutubeSource struct {
	*GenericSource
}

func (y *YoutubeSource) Regex() string {
	return `http(?:s?):\/\/(?:www\.)?youtu(?:be\.com\/watch\?v=|\.be\/)([\w\-\_]*)(&(amp;)?‌​[\w\?‌​=]*)?`
}
