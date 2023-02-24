package api

import (
	"github.com/Shemistan/english_bot/internal/domain"
	v1 "github.com/Shemistan/english_bot/pkg/admin/v1"
)

type Bot struct {
	v1.UnimplementedBotServiceServer

	BotService domain.Service
}
