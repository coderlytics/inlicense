package inlicense

import (
	"testing"
	"time"

	"coderlytics.at/inlicense/internal"
	"github.com/stretchr/testify/assert"
)

func TestLicenseIsExpiredBeforeStartDate(t *testing.T) {
	lic := internal.License{
		ValidFrom: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		ValidTo:   time.Date(2021, 7, 1, 0, 0, 0, 0, time.UTC),
	}

	license := License{
		license: lic,
	}

	now := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	assert.True(t, license.isExpired(now))
}

func TestLicenseIsExpiredAfterEndDate(t *testing.T) {
	lic := internal.License{
		ValidFrom: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		ValidTo:   time.Date(2021, 7, 1, 0, 0, 0, 0, time.UTC),
	}

	license := License{
		license: lic,
	}

	now := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	assert.True(t, license.isExpired(now))
}

func TestLicenseIsNotExpired(t *testing.T) {
	lic := internal.License{
		ValidFrom: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		ValidTo:   time.Date(2021, 7, 1, 0, 0, 0, 0, time.UTC),
	}

	license := License{
		license: lic,
	}

	now := time.Date(2021, 5, 1, 0, 0, 0, 0, time.UTC)
	assert.False(t, license.isExpired(now))
}
