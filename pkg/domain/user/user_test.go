package user_test

import (
	"testing"

	"github.com/hsmtkk/clean_arch_study_0/pkg/domain/user"
	"github.com/hsmtkk/clean_arch_study_0/test"
)

func TestUser(t *testing.T) {
	u := user.New(1, "alpha", "bravo")
	test.AssertEqualInt(t, 1, u.ID())
	test.AssertEqualString(t, "alpha", u.FirstName())
	test.AssertEqualString(t, "bravo", u.LastName())
}
