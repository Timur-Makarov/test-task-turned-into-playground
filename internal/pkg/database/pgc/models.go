// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package pgc

type Banner struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type CounterStatistic struct {
	BannerID      int32 `json:"bannerId"`
	TimestampFrom int64 `json:"timestampFrom"`
	TimestampTo   int64 `json:"timestampTo"`
	Count         int64 `json:"count"`
}