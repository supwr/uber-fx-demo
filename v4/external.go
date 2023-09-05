package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/supwr/uber-fx-demo/entities"
)

type External struct {
	HttpClient *resty.Client
}

func NewExternal(httpClient *resty.Client) *External {
	return &External{
		HttpClient: httpClient,
	}
}

func (e *External) getRandomPersons(ctx context.Context) ([]entities.PersonDTO, error) {
	var result struct {
		Persons []entities.PersonDTO `json:"results"`
	}

	response, err := e.HttpClient.
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
