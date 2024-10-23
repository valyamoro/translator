package domain

type Word struct {
	ID                          int64  `json:"id"`
	Word                        string `json:"word" validate:"correct"`
	TranslationWord             string `json:"translation_word" validate:"correct"`
	DictionaryID                int64  `json:"dictionary_id" validate:"dictionary_exists"`
	WordLanguageCode            string `json:"word_language_code" validate:"language"`
	TranslationWordLanguageCode string `json:"translation_word_language_code" validate:"language"`
}

type UpdateWordInput struct {
	ID                          *int64  `json:"id"`
	Word                        *string `json:"word"`
	TranslationWord             *string `json:"translation_word"`
	DictionaryID                *int64  `json:"dictionary_id"`
	WordLanguageCode            *int64  `json:"word_language_code"`
	TranslationWordLanguageCode *int64  `json:"translation_word_language_code"`
}
