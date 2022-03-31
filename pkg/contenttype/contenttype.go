package contenttype

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"aws-lambda-in-go-lang/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"go.uber.org/zap"
)

var sessionSvc *s3.S3
var sessionaws *session.Session

//service function interface
type Service interface {
	DetectContentType(input models.RequestParams) (*models.LambdaResponse, error)
	Newsvc(models.SvcInput, bool) (*s3.S3, *session.Session, error)
}

// service object
type ContentType struct {
	Logger       *zap.Logger
	InputRequest models.RequestParams
}

// initiate new ContentType object
func NewContentType(logger *zap.Logger) Service {

	return &ContentType{
		Logger: logger,
	}
}

// function to detect content type
func (cv *ContentType) DetectContentType(input models.RequestParams) (*models.LambdaResponse, error) {
	cv.InputRequest = input

	// if input.Url == "" {
	// 	return nil, fmt.Errorf("source link found as empty")
	// }
	if (input.Source == models.Source{}) {
		return nil, fmt.Errorf("source list found as empty")
	}

	MimeType, err := cv.DetectMimeType(input)
	if err != nil {
		return nil, err
	}

	return &models.LambdaResponse{
		Message: "successfully Detected MIME Type",
		Success: true,
		Data: map[string]interface{}{
			"MimeType": MimeType,
			"Message":  "Success",
		},
	}, nil
}

// Find MIME Type of an s3 Object
func (cv *ContentType) DetectMimeType(req models.RequestParams) (string, error) {

	var svc *s3.S3
	var err error

	// Get aws Credentials from static storage
	if (req.Credentials == models.Credentials{}) {

		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))

		svc = s3.New(sess)

	} else {
		svc, _, err = cv.Newsvc(models.SvcInput{
			Region:    req.Source.BucketRegion,
			AccessKey: req.Credentials.AccessKey,
			SecretKey: req.Credentials.SecretKey,
		}, true)

		if err != nil {
			return "", err
		}
	}

	// read the file from s3
	reqObj, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(req.Source.Bucket),
		Key:    aws.String(req.Source.Path),
	})
	if err != nil {
		return "error in opening s3 url", err
	}
	// Performance issue with large size files
	var reader io.Reader = reqObj.Body
	data, _ := ioutil.ReadAll(reader)
	//find MIME-type of the []byte
	MimeType := http.DetectContentType(data)

	return MimeType, nil
	// return *reqObj.ContentType, nil
}

// Session create
func (cv *ContentType) Newsvc(input models.SvcInput, force bool) (*s3.S3, *session.Session, error) {
	if sessionSvc == nil || force {
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String(input.Region),
			Credentials: credentials.NewStaticCredentials(
				input.AccessKey,
				input.SecretKey,
				"", // a token will be created when the session it's used.
			),
		})
		if err != nil {
			return nil, nil, fmt.Errorf("session create error %v", err)
		}
		sessionaws = sess
		sessionSvc = s3.New(sess)
	}

	return sessionSvc, sessionaws, nil
}
