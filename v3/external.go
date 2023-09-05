package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/supwr/uber-fx-demo/entities"
	"go.uber.org/fx"
)

type ExternalParams struct {
	fx.In

	HttpClient *resty.Client
}

type External struct {
	httpClient *resty.Client
}

func NewExternal(params ExternalParams) *External {
	return &External{
		httpClient: params.HttpClient,
	}
}

func (e *External) getRandomPersons(ctx context.Context) ([]entities.PersonDTO, error) {
	var result struct {
		Persons []entities.PersonDTO `json:"results"`
	}

	response, err := e.httpClient.
		SetTimeout(30*time.Second).
		SetBaseURL("http://randomuser.me").R().
		SetContext(ctx).
		SetResult(&result).
		Execute(resty.MethodGet, "/api/")

	if err != nil {
		return nil, err
	}

	if response.IsError() {
		return nil, errors.New(fmt.Sprintf("Error code: %s", strconv.Itoa(response.StatusCode())))
	}

	return result.Persons, nil
}
