package digitalaccount

import (
	"fmt"
	"pesthub/contracts"
	"pesthub/entities"
)

type CreateOfferInput struct {
	Name          string
	CodeName      string
	SomethingElse int64
}

type CreateOfferOutput interface {
	OnSuccess(r string)
	OnInvalid(e string)
}

type CreateOffer struct {
	insert contracts.InsertOffer
	output CreateOfferOutput
}

func NewCreateOffer(implInsert contracts.InsertOffer, implOutput CreateOfferOutput) *CreateOffer {
	return &CreateOffer{
		insert: implInsert,
		output: implOutput,
	}
}

func (o *CreateOffer) Execute(offerInput *CreateOfferInput) {
	if err := validate(offerInput); err != nil {
		o.output.OnInvalid(err.Error())
	}
	offer := convert(offerInput)
	o.insert("data here")
	fmt.Println(offer)
	o.output.OnSuccess("everything is good")
}

func validate(offerInput *CreateOfferInput) error {
	return nil
}

func convert(offerInput *CreateOfferInput) *entities.Damage {
	return &entities.Damage{
		//
		//
		//
	}
}
