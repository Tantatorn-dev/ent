package integration

import (
	"context"
	"math"
	"testing"

	"fbc/ent/entc/integration/ent"

	"github.com/stretchr/testify/require"
)

func Types(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	require := require.New(t)

	ft := client.FieldType.Create().
		SetInt(1).
		SetInt8(8).
		SetInt16(16).
		SetInt32(32).
		SetInt64(64).
		SaveX(ctx)

	require.Equal(1, ft.Int)
	require.Equal(int8(8), ft.Int8)
	require.Equal(int16(16), ft.Int16)
	require.Equal(int32(32), ft.Int32)
	require.Equal(int64(64), ft.Int64)

	ft = client.FieldType.Create().
		SetInt(1).
		SetInt8(math.MaxInt8).
		SetInt16(math.MaxInt16).
		SetInt32(math.MaxInt16).
		SetInt64(math.MaxInt16).
		SetOptionalInt8(math.MaxInt8).
		SetOptionalInt16(math.MaxInt16).
		SetOptionalInt32(math.MaxInt32).
		SetOptionalInt64(math.MaxInt64).
		SetNullableInt8(math.MaxInt8).
		SetNullableInt16(math.MaxInt16).
		SetNullableInt32(math.MaxInt32).
		SetNullableInt64(math.MaxInt64).
		SaveX(ctx)

	require.Equal(int8(math.MaxInt8), ft.OptionalInt8)
	require.Equal(int16(math.MaxInt16), ft.OptionalInt16)
	require.Equal(int32(math.MaxInt32), ft.OptionalInt32)
	require.Equal(int64(math.MaxInt64), ft.OptionalInt64)
	require.Equal(int8(math.MaxInt8), *ft.NullableInt8)
	require.Equal(int16(math.MaxInt16), *ft.NullableInt16)
	require.Equal(int32(math.MaxInt32), *ft.NullableInt32)
	require.Equal(int64(math.MaxInt64), *ft.NullableInt64)

	ft = client.FieldType.UpdateOne(ft).
		SetInt(1).
		SetInt8(math.MaxInt8).
		SetInt16(math.MaxInt16).
		SetInt32(math.MaxInt16).
		SetInt64(math.MaxInt16).
		SetOptionalInt8(math.MaxInt8).
		SetOptionalInt16(math.MaxInt16).
		SetOptionalInt32(math.MaxInt32).
		SetOptionalInt64(math.MaxInt64).
		SetNullableInt8(math.MaxInt8).
		SetNullableInt16(math.MaxInt16).
		SetNullableInt32(math.MaxInt32).
		SetNullableInt64(math.MaxInt64).
		SaveX(ctx)

	require.Equal(int8(math.MaxInt8), ft.OptionalInt8)
	require.Equal(int16(math.MaxInt16), ft.OptionalInt16)
	require.Equal(int32(math.MaxInt32), ft.OptionalInt32)
	require.Equal(int64(math.MaxInt64), ft.OptionalInt64)
	require.Equal(int8(math.MaxInt8), *ft.NullableInt8)
	require.Equal(int16(math.MaxInt16), *ft.NullableInt16)
	require.Equal(int32(math.MaxInt32), *ft.NullableInt32)
	require.Equal(int64(math.MaxInt64), *ft.NullableInt64)
}
