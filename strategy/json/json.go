package json

import (
	"fmt"
	"reflect"

	"github.com/goccy/go-json"
	"github.com/sg3t41/gocnf/strategy"
)

type JSONStrategy struct {
	strategy.Strategy
}

func (js *JSONStrategy) Unmarshal(out any) error {
	if reflect.TypeOf(out).Kind() != reflect.Ptr {
		return fmt.Errorf("gocnf.Unmarshalの引数は構造体のポインタ型である必要があります。要求された型: %s", reflect.TypeOf(out).Kind())
	}

	kind := reflect.TypeOf(out).Elem().Kind()
	if kind != reflect.Struct {
		return fmt.Errorf("gocnf.Unmarshalの引数は構造体のポインタ型である必要があります。要求された型: %s", kind)
	}

	in := js.Data
	if err := json.Unmarshal(in, out); err != nil {
		return fmt.Errorf("Unmarshalに失敗しました エラー: %v", err)
	}

	return nil
}