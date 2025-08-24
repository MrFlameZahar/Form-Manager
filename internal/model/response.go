package model

import "time"

type ResponseID uint

func NewResponseID(id uint) (ResponseID, error) {
	return ResponseID(id), nil
}

type AnswerValue string

func NewAnswerValue(value string) (AnswerValue, error) {
	return AnswerValue(value), nil
}

type Response struct {
	ResponseID  ResponseID
	FormID      uint
	SubmittedAt time.Time
	Answers     []Answer
}

type Answer struct {
	AnswerID    uint
	FormFieldID uint
	Value       AnswerValue
}
