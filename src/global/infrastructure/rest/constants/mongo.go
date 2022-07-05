package restconst

import (
	evtlog "github.com/jdpadillaac/canditickets_api/src/packages/event/log/domain"
	mdbevtlog "github.com/jdpadillaac/canditickets_api/src/packages/event/log/infra/mongo"
	"github.com/jdpadillaac/canditickets_api/src/packages/event/tag/domain"
	"github.com/jdpadillaac/canditickets_api/src/packages/event/tag/infra/mongo"
	otp "github.com/jdpadillaac/canditickets_api/src/packages/otp/domain"
	mdbotp "github.com/jdpadillaac/canditickets_api/src/packages/otp/infra/mongo"
	"github.com/jdpadillaac/canditickets_api/src/packages/product/domain"
	"github.com/jdpadillaac/canditickets_api/src/packages/product/infra/mongo"
	"github.com/jdpadillaac/canditickets_api/src/packages/promotional_code/domain"
	"github.com/jdpadillaac/canditickets_api/src/packages/promotional_code/infra/mongo"
	"github.com/jdpadillaac/canditickets_api/src/packages/return_policy/domain"
	"github.com/jdpadillaac/canditickets_api/src/packages/return_policy/infra/mongo"
	"github.com/jdpadillaac/canditickets_api/src/packages/tickets/payment/domain"
	"github.com/jdpadillaac/canditickets_api/src/packages/tickets/payment/infra/mongo"
	"github.com/lea55/BACKENDCANDITICKET/src/packages/error/domain"
	"github.com/lea55/BACKENDCANDITICKET/src/packages/error/infra/mongo"
)

type MongoRepo struct {
	EventTag            eventtag.Repository
	Tag                 eventtag.TagRepo
	Ticket              ticket.Repository
	PromCodeApp         promcode.CodeAppRepo
	ReturnPl            returnpl.Repository
	Product             pdt.Repository
	PurchaseOrder       tktpo.Repository
	PurchaseOrderStatus tktpo.StatusRepo
	TktPayment          tktpayment.Repository
	ErrLog              errlog.Repository
	EventLog            evtlog.Repository
	Otp                 otp.Repository
}

func GetMongoRepo() *MongoRepo {
	return &MongoRepo{
		EventTag:            mdbeventcat.NewEventCat(),
		Tag:                 mdbeventcat.NewTagRepo(),
		Ticket:              mdbtkt.NewRepository(),
		PromCodeApp:         mdbpromcode.NewPromCodeApplication(),
		ReturnPl:            mdbreturnpl.NewRepo(),
		Product:             mdbpdt.NewRepo(),
		PurchaseOrder:       mdbtktpo.NewRepo(),
		PurchaseOrderStatus: mdbtktpo.NewStatusRepo(),
		TktPayment:          mdbtktpmt.NewRepo(),
		ErrLog:              mdberrlog.NewRepo(),
		EventLog:            mdbevtlog.NewRepo(),
		Otp:                 mdbotp.New(),
	}
}
