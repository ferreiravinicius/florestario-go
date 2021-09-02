package digitalaccount

import "pesthub/contracts"

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
	o.insert("data here")
	o.output.OnSuccess("everything is good")
}
