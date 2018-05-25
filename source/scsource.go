package source

type SCSource struct {
	*GenericSource
}

/* Current Idea: Factory (Makes own special kind of videos htat contain the dl logic)
 */
func (s *SCSource) Regex() string {
	return `(https?:\/\/soundcloud\.com\/[a-zA-z\-\d]+\/[a-zA-z\-\d]+)`
}
