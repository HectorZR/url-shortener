package shortener

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"

	recaptcha "cloud.google.com/go/recaptchaenterprise/v2/apiv1"
	recaptchapb "cloud.google.com/go/recaptchaenterprise/v2/apiv1/recaptchaenterprisepb"
	"github.com/HectorZR/url-shortener/shared"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

/*
 * ShortenedURL model
 */
type ShortenedURL struct {
	gorm.Model
	OriginalURL    string `gorm:"not null;unique"`
	ExpirationDate time.Time
}

func ShortenURL(url string, db *gorm.DB) ShortenedURL {
	var existingUrl ShortenedURL
	db.First(&existingUrl, "original_url = ?", url)

	if existingUrl.ID != 0 {
		return existingUrl
	}

	newUrl := ShortenedURL{OriginalURL: url, ExpirationDate: time.Now().Add(time.Hour * 24)}
	db.Create(&newUrl)
	return newUrl
}

func GetOriginalURL(shortCode string, db *gorm.DB) (ShortenedURL, error) {
	id := shared.DecodeBase62(shortCode)
	var shortened ShortenedURL
	db.Where("expiration_date > NOW()").First(&shortened, id)

	if shortened.ID == 0 {
		return ShortenedURL{}, errors.New("Short URL not found")
	}

	return shortened, nil
}

func ValidateURL(u string) error {
	if u == "" {
		return errors.New("URL cannot be empty")
	}

	if strings.Contains(u, " ") {
		return errors.New("URL cannot contain spaces")
	}

	_, err := url.ParseRequestURI(u)
	if err != nil {
		return errors.New("URL is not valid")
	}

	re := regexp.MustCompile(`^(http|https):\/\/[^\s/$.?#].[^\s]*$`)
	if !re.MatchString(u) {
		return errors.New("URL format is invalid")
	}

	return nil
}

/**
 * Create an assessment to analyze the risk of a UI action.
 *
 * @param projectID: Your Google Cloud Project ID.
 * @param recaptchaKey: The reCAPTCHA key associated with the site/app
 * @param token: The generated token obtained from the client.
 * @param recaptchaAction: Action name corresponding to the token.
 * @return bool: true if validation passed, false otherwise
 * @return error: any error that occurred during validation
 */
func createAssessment(projectID string, recaptchaKey string, token string, recaptchaAction string) (bool, error) {
	env := shared.GetEnvVars()
	ctx := context.Background()

	jsonCreds := env.GoogleCredentialsJson
	option := option.WithCredentialsJSON([]byte(jsonCreds))

	// Create the reCAPTCHA client.
	client, err := recaptcha.NewClient(ctx, option)

	if err != nil {
		fmt.Printf("Error creating reCAPTCHA client\n")
		return false, fmt.Errorf("Error creating reCAPTCHA client: %v\n", err)
	}

	defer client.Close()

	// Set the properties of the event to be tracked.
	event := &recaptchapb.Event{
		Token:   token,
		SiteKey: recaptchaKey,
	}

	assessment := &recaptchapb.Assessment{
		Event: event,
	}

	// Build the assessment request.
	request := &recaptchapb.CreateAssessmentRequest{
		Assessment: assessment,
		Parent:     fmt.Sprintf("projects/%s", projectID),
	}

	response, err := client.CreateAssessment(ctx, request)

	if err != nil {
		return false, fmt.Errorf("error calling CreateAssessment: %v", err)
	}

	// Check if the token is valid.
	if !response.TokenProperties.Valid {
		fmt.Printf("reCAPTCHA token invalid for reasons: %v\n", response.TokenProperties.InvalidReason)
		return false, nil
	}

	// Check if the expected action was executed.
	if response.TokenProperties.Action != recaptchaAction {
		fmt.Printf("reCAPTCHA action mismatch. Expected: %s, Got: %s\n", recaptchaAction, response.TokenProperties.Action)
		return false, nil
	}

	// Get the risk score and log it
	score := response.RiskAnalysis.Score
	fmt.Printf("reCAPTCHA score: %v\n", score)

	// Log any risk analysis reasons
	for _, reason := range response.RiskAnalysis.Reasons {
		fmt.Printf("Risk reason: %s\n", reason.String())
	}

	// Consider scores above 0.5 as valid (you can adjust this threshold)
	if score >= 0.5 {
		return true, nil
	}

	fmt.Printf("reCAPTCHA score too low: %v (threshold: 0.5)\n", score)
	return false, nil
}

/**
 * Validates reCAPTCHA token and returns validation result
 *
 * @param projectID: Your Google Cloud Project ID.
 * @param recaptchaKey: The reCAPTCHA key associated with the site/app
 * @param token: The generated token obtained from the client.
 * @param recaptchaAction: Action name corresponding to the token.
 */
func validateRecaptcha(projectID string, recaptchaKey string, token string, recaptchaAction string) error {
	// Validate reCAPTCHA token first
	if token == "" {
		return errors.New("reCAPTCHA verification required")
	}

	isValid, err := createAssessment(projectID, recaptchaKey, token, recaptchaAction)

	if err != nil {
		return err
	}

	if !isValid {
		return errors.New("reCAPTCHA verification failed. Please try again.")
	}

	return nil
}
