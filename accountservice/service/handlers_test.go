package service

import (
	"encoding/json"
	"fmt"
	"github.com/phillipahereza/go_microservices/accountservice/dbclient"
	"github.com/phillipahereza/go_microservices/model"
	"github.com/smartystreets/goconvey/convey"
	"gopkg.in/h2non/gock.v1"
	"net/http/httptest"
	"testing"
)

func init() {
	gock.InterceptClient(client)
}


func TestGetAccountWithWrongPath(t *testing.T) {
	defer gock.Off()
	gock.New("http://quotes-service:8080").
		Get("/api/quote").
		MatchParam("strength", "4").
		Reply(200).
		BodyString(`{"quote":"May the source be with you. Always.","ipAddress":"10.0.0.5:8080","language":"en"}`)

	mockRepo := &dbclient.MockBoltClient{}

	mockRepo.On("QueryAccount", "123").Return(model.Account{Id:"123", Name:"Person_123"}, nil)
	mockRepo.On("QueryAccount", "456").Return(model.Account{}, fmt.Errorf("Some error"))

	DBClient = mockRepo

	convey.Convey("Given an HTTP Request for /invalid/123", t, func() {
		req := httptest.NewRequest("GET", "/invalid/123", nil)
		resp := httptest.NewRecorder()

		convey.Convey("When request is handled by the router", func() {
			NewRouter().ServeHTTP(resp, req)
			convey.Convey("Then the response should be 404", func() {
				convey.So(resp.Code, convey.ShouldEqual, 404)
			})
		})

	})

	convey.Convey("Given a HTTP request for /accounts/123", t, func() {
		req := httptest.NewRequest("GET", "/accounts/123", nil)
		resp := httptest.NewRecorder()

		convey.Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			convey.Convey("Then the response should be a 200", func() {
				convey.So(resp.Code, convey.ShouldEqual, 200)

				account := model.Account{}
				json.Unmarshal(resp.Body.Bytes(), &account)
				convey.So(account.Id, convey.ShouldEqual, "123")
				convey.So(account.Name, convey.ShouldEqual, "Person_123")
				convey.So(account.Quote.Text, convey.ShouldEqual, "May the source be with you. Always.")
			})
		})
	})

	convey.Convey("Given a HTTP request for /accounts/456", t, func() {
		req := httptest.NewRequest("GET", "/accounts/456", nil)
		resp := httptest.NewRecorder()

		convey.Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			convey.Convey("Then the response should be a 404", func() {
				convey.So(resp.Code, convey.ShouldEqual, 404)
			})
		})
	})
}
