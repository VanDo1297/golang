package scripts

//import (
//	"demo/internal/instances/servinstances"
//	"demo/internal/integrations/stripeapi"
//	"demo/internal/models/dbmodels"
//	"github.com/go-errors/errors"
//	"github.com/stripe/stripe-go/v71"
//)

func AddBaseData(){
	//user := &dbmodels.User{
	//	Email:        "test",
	//	Password:     "test",
	//}
	//err := sharedinstances.Get().GetDAO().GetUserDAO().CreateUser(user)
	//if err!=nil {
	//	panic(err)
	//}
	//
	//card := dbmodels.PaymentCard{
	//	StripeID:            "234",
	//	SetupIntentStripeID: "234",
	//	Brand:               "1234",
	//	Last4:               "234",
	//	ExpMonth:            1,
	//	ExpYear:             1,
	//	DefaultPayment:      false,
	//	Name:                "234",
	//	BillingInfo:         dbmodels.PaymentCardBillingInfo{
	//		AddressLine1:    "123",
	//		AddressLine2:    "123",
	//		AddressCity:     "123",
	//		AddressState:    "123",
	//		AddressZip:      "123",
	//		AddressCountry:  "123",
	//	},
	//}
	//
	//e := sharedinstances.Get().GetDB().GetConnection().Model(user).Association("PaymentCards").Append(&card)
	//if e!=nil {
	//	panic(e)
	//}
}


//func AddBaseData(){
//	sProd, price, err := servinstances.Get().GetStripeAPI().CreateOneTimePaymentProduct("OneTimeBPv2", stripeapi.ProductPriceParams{
//		Name:            "One Time Purchase Price",
//		Price:           9.99,
//	})
//	if err !=nil{
//		panic(err)
//	}
//	err = saveProdAndPrice(sProd,[]*stripe.Price{price},false)
//	if err !=nil{
//		panic(err)
//	}
//
//	monthInterval := stripe.PlanIntervalMonth
//	monthIntervalCount := int64(1)
//	yearlyInterval := stripe.PlanIntervalYear
//	yearIntervalCount := int64(1)
//	sProd, prices, err := servinstances.Get().GetStripeAPI().CreateSubscriptionProduct("SubscriptionBPv2",[]*stripeapi.ProductPriceParams{{
//		Name:            "Monthly Subscription",
//		Price:           9.99,
//		Interval: &monthInterval,
//		IntervalCount: 	&monthIntervalCount,
//	},{
//		Name:            "Annual Subscription",
//		Price:           99.99,
//		Interval: &yearlyInterval,
//		IntervalCount: 	&yearIntervalCount,
//	}})
//	if err !=nil{
//		panic(err)
//	}
//
//	err = saveProdAndPrice(sProd,prices,true)
//	if err !=nil{
//		panic(err)
//	}
//}
//
//func saveProdAndPrice(prod *stripe.Product, prices []*stripe.Price, subscription bool) error{
//	bpProd := &dbmodels.BPProduct{
//		StripeID:     prod.ID,
//		Name:         prod.Name,
//		Subscription: subscription,
//	}
//	err := servinstances.Get().GetDAO().GetProductDAO().CreateProduct(bpProd)
//	if err !=nil{
//		return err
//	}
//
//	for _, price := range prices{
//		bpPrice := &dbmodels.BPProductPrice{
//			ProductID:    bpProd.ID,
//			StripeID:     price.ID,
//			Name:         price.Nickname,
//			Price:        float64(price.UnitAmount)/float64(100),
//		}
//		if price.Recurring!=nil{
//			i := string(price.Recurring.Interval)
//			bpPrice.IntervalType = &i
//			ic := uint64(price.Recurring.IntervalCount)
//			bpPrice.Interval = &ic
//			td := uint64(price.Recurring.TrialPeriodDays)
//			bpPrice.TrialDays = &td
//		}
//		err := servinstances.Get().GetDAO().GetProductDAO().CreateProductPrice(bpPrice)
//		if err !=nil{
//			return err
//		}
//	}
//
//	return nil
//}
