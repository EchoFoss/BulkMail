package campaign

import (
	"fmt"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	name     = "Campaign X"
	content  = "Email body"
	contacts = []string{"JohnDoe@email.test", "FIzzbuzz@emai.text", "Barbaz@email.test"}
)

func TestNewCampaign(t *testing.T) {
	assert := assert2.New(t)

	campaign, _ := NewCampain(name, content, contacts)

	fmt.Println(campaign.ID)
	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Equal(len(contacts), len(campaign.Contacts))
}

func TestNewCampaignIdIsNotNil(t *testing.T) {
	assert := assert2.New(t)
	name = "Campaign X"
	content = "Email body"
	contacts = []string{"JohnDoe@email.test", "FIzzbuzz@emai.text", "Barbaz@email.test"}

	campaign, _ := NewCampain(name, content, contacts)

	assert.NotNil(campaign.ID)
}
func TestNewCampainCreatedOnIsValid(t *testing.T) {

	assert := assert2.New(t)
	name = "Campaign X"
	content = "Email body"
	contacts = []string{"JohnDoe@email.test", "FIzzbuzz@emai.text", "Barbaz@email.test"}

	before := time.Now().Add(-time.Minute)
	campaign, _ := NewCampain(name, content, contacts)
	now := time.Now().Add(time.Minute)

	assert.Greater(now, campaign.CreatedAt)
	assert.Less(before, campaign.CreatedAt)

}

func TestNewCampainMustValidateName(t *testing.T) {

	assert := assert2.New(t)

	_, err := NewCampain("", content, contacts)

	assert.Equal(err, nameIsRequiredErr)

}

func TestNewCampainMustValidateContent(t *testing.T) {

	assert := assert2.New(t)

	_, err := NewCampain(name, "", contacts)

	assert.Equal(err, contentIsRequiredErr)

}

func TestNewCampainMustValidateContacts(t *testing.T) {

	assert := assert2.New(t)

	_, err := NewCampain(name, content, []string{})

	assert.Equal(err, contactIsRequiredErr)

}
