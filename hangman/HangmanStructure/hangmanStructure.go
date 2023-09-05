package HangmanStructure

// Creation of structure 'HangmanData' with all its fields
type HangmanData struct {
	GameFinished        bool
	Word                []string
	WordToFind          string
	Attempts            int
	SavePositionHangman int
	PositionHangman     []string
	RandomWordsFile     []string
}

// Method GET (getters)
// GET returns the value stored in the field specified
func (h HangmanData) GetGameFinished() bool {
	return h.GameFinished
}

func (h HangmanData) GetWord() []string {
	return h.Word
}

func (h HangmanData) GetWordToFind() string {
	return h.WordToFind
}

func (h HangmanData) GetAttempts() int {
	return h.Attempts
}

func (h HangmanData) GetSavePositionHangman() int {
	return h.SavePositionHangman
}

func (h HangmanData) GetPositionHangman() []string {
	return h.PositionHangman
}

func (h HangmanData) GetRandomWordsFile() []string {
	return h.RandomWordsFile
}

// Method SET (setters)
// SET changes the value of the field
func (h *HangmanData) SetGameFinished(GameFinished bool) {
	h.GameFinished = GameFinished
}

func (h *HangmanData) SetWord(Word []string) {
	h.Word = Word
}

func (h *HangmanData) SetWordToFind(WordToFind string) {
	h.WordToFind = WordToFind
}

func (h *HangmanData) SetAttempts(Attempts int) {
	h.Attempts = Attempts
}

func (h *HangmanData) SetSavePositionHangman(SavePositionHangman int) {
	h.SavePositionHangman = SavePositionHangman
}

func (h *HangmanData) SetPositionHangman(PositionHangman []string) {
	h.PositionHangman = PositionHangman
}

func (h *HangmanData) SetRandomWordsFile(RandomWordsFile []string) {
	h.RandomWordsFile = RandomWordsFile
}
