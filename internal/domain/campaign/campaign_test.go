package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name    = "Marketing"
	content = "Body"
	email   = []string{"teste@teste"}
)

func Test_NewCampaing_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	result, _ := NewCampaign(name, content, email)

	// t.Run("expected Marketing", func(t *testing.T) {
	// 	if result.Name != "Mkt" {
	// 		t.Errorf("expected Marketing: %s", name)
	// 	}

	// })

	// t.Run("expected content", func(t *testing.T) {
	// 	if content != result.Content {
	// 		t.Errorf("expected Body: %s", name)
	// 	}

	// })
	//assert.Equal(result.ID, "1")
	assert.Equal(result.Name, "Marketing")
}

func Test_NewCampaing_IDIsNotNill(t *testing.T) {
	assert := assert.New(t)

	result, _ := NewCampaign(name, content, email)

	assert.NotNil(result.ID)
}

func Test_NewCampaing_CreatedOnMustBeNowl(t *testing.T) {
	assert := assert.New(t)

	now := time.Now().Add(-time.Minute)

	result, _ := NewCampaign(name, content, email)

	assert.Greater(result.CreatedOn, now)
}

func Test_NewCampaing_MustValidateName(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, email)

	assert.Equal("name is required", err.Error())
}

func Test_NewCampaing_MustValidaContent(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", email)

	assert.Equal("content is required", err.Error())
}

func Test_NewCampaing_MustValidaEmail(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assert.Equal("emails is required", err.Error())
}
