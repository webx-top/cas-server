package storage

import (
	"github.com/webx-top/cas-server/types"
)

// IStorage interface for all Storages
type IStorage interface {
	SaveTicket(*types.Ticket)
	DoesTicketExist(ticket string) *types.Ticket
	DeleteTicket(ticket string)
}
