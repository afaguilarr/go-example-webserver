package business

import (
	"github.com/afaguilarr/go-example-webserver/app/src/dao"
	"github.com/afaguilarr/go-example-webserver/proto"
)

func ProtoLocationToDaoLocation(protoLocation *proto.Location) *dao.Location {
	var sp, cm, n, zp *string

	if protoLocation.StateOrProvince != nil {
		sp = &protoLocation.StateOrProvince.Value
	}
	if protoLocation.CityOrMunicipality != nil {
		sp = &protoLocation.CityOrMunicipality.Value
	}
	if protoLocation.Neighborhood != nil {
		sp = &protoLocation.Neighborhood.Value
	}
	if protoLocation.ZipCode != nil {
		sp = &protoLocation.ZipCode.Value
	}

	return &dao.Location{
		Country:            protoLocation.Country,
		StateOrProvince:    sp,
		CityOrMunicipality: cm,
		Neighborhood:       n,
		ZipCode:            zp,
	}
}
