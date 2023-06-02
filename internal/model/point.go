package model

import (
	"context"
	"encoding/binary"
	"fmt"
	"math"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Point struct {
	Lat float64
	Lng float64
}

func (loc Point) GormDataType() string {
	return "point"
}

func (point Point) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_PointFromText(?)",
		Vars: []interface{}{fmt.Sprintf("POINT(%f %f)", point.Lat, point.Lng)},
	}
}

func (point *Point) Scan(v interface{}) error {
	u, ok := v.([]uint8)
	if !ok {
		return nil
	}

	point.Lat = float64frombytes(u[9:17])
	point.Lng = float64frombytes(u[17:25])
	return nil
}

func float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}
