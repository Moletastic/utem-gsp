package decoder

import (
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
)

type GSPDecoder struct {
	Decoder *mapstructure.Decoder
}

func NewDecoder(result interface{}) (*GSPDecoder, error) {
	gsp := new(GSPDecoder)
	cfg := &mapstructure.DecoderConfig{
		Metadata:   nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(ToTimeHookFunc()),
		Result:     result,
	}
	decoder, err := mapstructure.NewDecoder(cfg)
	if err != nil {
		return nil, err
	}
	gsp.Decoder = decoder
	return gsp, nil
}

func (gsp *GSPDecoder) Decode(input map[string]interface{}) error {
	err := gsp.Decoder.Decode(&input)
	if err != nil {
		return err
	}
	return nil
}

func ToTimeHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		switch f.Kind() {
		case reflect.String:
			return time.Parse(time.RFC3339, data.(string))
		case reflect.Float64:
			return time.Unix(0, int64(data.(float64))*int64(time.Millisecond)), nil
		case reflect.Int64:
			return time.Unix(0, data.(int64)*int64(time.Millisecond)), nil
		default:
			return data, nil
		}
	}
}
