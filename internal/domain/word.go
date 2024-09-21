package domain

type Word struct {
	ID                          int64  `json:"id"`
	Word                        string `json:"word"`
	TranslationWord             string `json:"translation_word"`
	DictionaryID                int64  `json:"dictionary_id"`
	WordLanguageCode            int64  `json:"word_language_code"`
	TranslationWordLanguageCode int64  `json:"translation_word_language_code"`
}

type UpdateWordInput struct {
	ID                          *int64  `json:"id"`
	Word                        *string `json:"word"`
	TranslationWord             *string `json:"translation_word"`
	DictionaryID                *int64  `json:"dictionary_id"`
	WordLanguageCode            *int64  `json:"word_language_code"`
	TranslationWordLanguageCode *int64  `json:"translation_word_language_code"`
}
