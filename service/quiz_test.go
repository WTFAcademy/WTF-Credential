package service

import (
	"context"
	"fmt"
	"testing"
	"wtf-credential/request"

	jsoniter "github.com/json-iterator/go"
)

func TestGetChapterQuizzes(t *testing.T) {
	got, err := QuizService.GetChapterQuizzes(context.Background(), &request.GetChapterQuizzes{
		Path:      "Solidity101",
		RoutePath: "HelloWeb3",
	})
	if err != nil {
		t.Errorf("GetChapterQuizzes() error = %v", err)
		return
	}
	gotStr, _ := jsoniter.MarshalToString(got)
	fmt.Println(gotStr)
}

func TestGradeSubmit(t *testing.T) {
	reqStr := `{"chapter_id":95,"course_id":128,"answers":[{"id":366,"answers":["A"]},{"id":367,"answers":["A"]},{"id":368,"answers":["B"]},{"id":369,"answers":["B"]},{"id":370,"answers":["A"]}]}`
	var req request.QuizGradeSubmitReq
	jsoniter.UnmarshalFromString(reqStr, &req)
	got, err := QuizService.GradeSubmit(context.Background(), &req)
	if err != nil {
		t.Errorf("GradeSubmit() error = %v", err)
		return
	}
	gotStr, _ := jsoniter.MarshalToString(got)
	fmt.Println(gotStr)

}
