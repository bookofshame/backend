package sms

import (
	"context"
	"fmt"
	"net/url"

	"github.com/bookofshame/bookofshame/pkg/config"
	"github.com/bookofshame/bookofshame/pkg/fetch"
	"github.com/bookofshame/bookofshame/pkg/logging"
	"go.uber.org/zap"
)

type Client struct {
	cfg    config.Config
	fetch  *fetch.Fetch
	logger *zap.SugaredLogger
}

func NewClient(ctx context.Context, cfg config.Config) *Client {
	return &Client{
		cfg:    cfg,
		fetch:  fetch.NewFetch(ctx, nil),
		logger: logging.FromContext(ctx),
	}
}

func (s Client) Send(payload Payload) error {
	if err := payload.Validate(); err != nil {
		return err
	}

	if s.cfg.Env == "development" {
		s.logger.Debugln("sms not sent in development mode")
		return nil
	}

	requestURL := fmt.Sprintf("%s?api_key=%s&type=text&number=%s&senderid=%s&message=%s", s.cfg.SmsHost, s.cfg.SmsApiKey, payload.Number, s.cfg.SmsSenderId, url.QueryEscape(payload.Message))
	err := s.fetch.Get(requestURL, nil)

	if err != nil {
		s.logger.Errorf("failed to send sms. response: %w", err)
		return fmt.Errorf("failed to send sms")
	}

	return nil
}

func (s Payload) Validate() error {
	if s.Number == "" {
		return fmt.Errorf("number cannot be empty")
	}

	if s.Message == "" {
		return fmt.Errorf("message cannot be empty")
	}

	return nil
}
