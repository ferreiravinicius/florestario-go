package disorder

import (
	"pesthub/contracts"
	"pesthub/failures"
	"strconv"
)

var OptStore contracts.DisorderStore
var OptMsgs contracts.Messages

var emptyOutput RegisterDisorderOutput

func RegisterOptimal(disorderInput RegisterDisorderInput) (RegisterDisorderOutput, error) {
	exists, err := OptStore.ExistsName(disorderInput.Name)
	if err != nil {
		return emptyOutput, failures.Internal(err)
	}
	if exists {
		message := OptMsgs.GetText(MsgNameAlreadyExists)
		return emptyOutput, failures.UseCase(message)
	}

	disorder := disorderInput.ToEntity()
	disorder, err = OptStore.Save(disorder)
	if err != nil {
		return emptyOutput, failures.Internal(err)
	}

	return RegisterDisorderOutput{
		Code: strconv.FormatUint(disorder.Code, 10),
	}, nil
}
