package models

import (
	"context"
	"io"
	"px.dev/pixie/src/api/go/pxapi"
	"px.dev/pixie/src/api/go/pxapi/types"
)

type ItemType interface {
	string | HTTP_raw_data
}

func New[itemType ItemType]() *ItemMapMux[itemType] {
	return &ItemMapMux[itemType]{}
}

type ItemMapMux[itemType ItemType] struct {
	Table TablePrinterItemMap[itemType]
}

type TablePrinterItemMap[itemType ItemType] struct {
	Values []itemType
}

func (t *TablePrinterItemMap[itemType]) HandleInit(ctx context.Context, metadata types.TableMetadata) error {
	return nil
}

func (t *TablePrinterItemMap[itemType]) HandleRecord(ctx context.Context, r *types.Record) error {
	itemMap := ConvertPixieDataToItemStore[itemType](r)
	t.Values = append(t.Values, itemMap)
	return nil
}

func (t *TablePrinterItemMap[itemType]) HandleDone(ctx context.Context) error {
	return nil
}

func (s *ItemMapMux[itemType]) AcceptTable(ctx context.Context, metadata types.TableMetadata) (pxapi.TableRecordHandler, error) {
	return &s.Table, nil
}

func (s *ItemMapMux[itemType]) ExecutePxlScript(ctx context.Context, vz *pxapi.VizierClient, pxl string) (*pxapi.ScriptResults, error) {
	resultSet, err := vz.ExecuteScript(ctx, pxl, s)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return resultSet, nil
}

type PixieHTTPResponse[T ItemType] struct {
	ResultStats *pxapi.ResultsStats `json:"stats"`
	Results     []T                 `json:"results"`
}

func PixieResponseToHTTPResponse[T ItemType](results *pxapi.ScriptResults, mux *ItemMapMux[T]) *PixieHTTPResponse[T] {
	resp := PixieHTTPResponse[T]{}
	if results != nil {
		resp.ResultStats = results.Stats()
	}
	if mux != nil && mux.Table.Values != nil {
		resp.Results = mux.Table.Values
	}
	return &resp
}
