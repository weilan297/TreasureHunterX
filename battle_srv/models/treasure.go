package models

import (
	"github.com/ByteArena/box2d"
)

type Treasure struct {
	Id              int32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	LocalIdInBattle int32         `protobuf:"varint,2,opt,name=localIdInBattle,proto3" json:"localIdInBattle,omitempty"`
	Score           int32         `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty"`
	X               float64       `protobuf:"fixed64,4,opt,name=x,proto3" json:"x,omitempty"`
	Y               float64       `protobuf:"fixed64,5,opt,name=y,proto3" json:"y,omitempty"`
	Removed         bool          `protobuf:"varint,6,opt,name=removed,proto3" json:"removed,omitempty"`
  Type            int32         `protobuf:"varint,7,opt,name=Type,proto3" json:"type,omitempty"`
 
	PickupBoundary  *Polygon2D    `json:"-"`
	CollidableBody  *box2d.B2Body `json:"-"`
}
