package testutils

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func AssertError(t *testing.T, err error, want error) {
	t.Helper()
	if err != nil && want != nil {
		if diff := cmp.Diff(err.Error(), want.Error()); diff != "" {
			t.Errorf("Returned error is not want. \nDiff: \n%v", diff)
		}
	}
	if (err != nil && want == nil) || (err == nil && want != nil) {
		if diff := cmp.Diff(err, want); diff != "" {
			t.Errorf("Returned error is not want. \nDiff: \n%v", diff)
		}
	}
}

// xo で自動生成したmodelの非公開フィールドを無視して比較しないといけない
// 使用例：testutils.AssertResponse(t, got, tt.want, models.Group{})
func AssertResponse(t *testing.T, got, want any, ignoreUnexportedTypes ...any) {
	t.Helper()
	var diff string
	if len(ignoreUnexportedTypes) > 0 {
		diff = cmp.Diff(got, want, cmpopts.IgnoreUnexported(ignoreUnexportedTypes[0]))
	} else {
		diff = cmp.Diff(got, want)
	}
	if diff != "" {
		t.Errorf("Returned value is not want.\nDiff: \n%s", diff)
	}
}

func AssertResponseWithOption(t *testing.T, got, want any, opt cmp.Option) {
	t.Helper()
	if diff := cmp.Diff(got, want, opt); diff != "" {
		t.Errorf("Returned value is not want.\nDiff: \n%s", diff)
	}
}

func AssertResponseWithOptions(t *testing.T, got, want any, opts ...cmp.Option) {
	t.Helper()
	if diff := cmp.Diff(got, want, opts...); diff != "" {
		t.Errorf("Returned value is not want.\nDiff: \n%s", diff)
	}
}

func DefaultIgnoreFieldsOptions[T any](target T) cmp.Options {
	return cmp.Options{
		cmpopts.IgnoreFields(
			target,
			"ID",
			"CreatedAt",
			"UpdatedAt",
		),
	}
}

func GenerateIgnoreFieldsOptions[T any](target T, fields ...string) cmp.Options {
	return cmp.Options{
		cmpopts.IgnoreFields(
			target,
			fields...,
		),
	}
}

func GenerateIgnoreUnexportedTypesOptions(types ...any) cmp.Options {
	return cmp.Options{
		cmpopts.IgnoreUnexported(types...),
	}
}

func ToPtr[T any](value T) *T {
	return &value
}
