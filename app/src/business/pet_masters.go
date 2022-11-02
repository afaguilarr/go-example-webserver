package business

import (
	"github.com/afaguilarr/go-example-webserver/app/src/dao"
	"github.com/afaguilarr/go-example-webserver/proto"
)

func ProtoPetMasterToDaoPetMaster(protoPetMaster *proto.PetMasterInfo) *dao.PetMaster {
	var cn *string
	location := protoPetMaster.Location

	l := ProtoLocationToDaoLocation(location)

	if protoPetMaster.ContactNumber != nil {
		cn = &protoPetMaster.ContactNumber.Value
	}
	return &dao.PetMaster{
		Name:          protoPetMaster.Name,
		ContactNumber: cn,
		Location:      l,
	}
}
