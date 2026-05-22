package schema

type ApplicationParameters struct {
	OpeningStatement              string           `json:"opening_statement,omitempty"`
	SuggestedQuestions            []string         `json:"suggested_questions,omitempty"`
	SuggestedQuestionsAfterAnswer map[string]any   `json:"suggested_questions_after_answer,omitempty"`
	SpeechToText                  map[string]any   `json:"speech_to_text,omitempty"`
	RetrieverResource             map[string]any   `json:"retriever_resource,omitempty"`
	AnnotationReply               map[string]any   `json:"annotation_reply,omitempty"`
	UserInputForm                 []map[string]any `json:"user_input_form,omitempty"`
	FileUpload                    map[string]any   `json:"file_upload,omitempty"`
	SystemParameters              map[string]any   `json:"system_parameters,omitempty"`
}
