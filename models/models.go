package models

type LambdaResponse struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

// Credentials
type Credentials struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

//source param
type Source struct {
	Path         string `json:"path"`
	Bucket       string `json:"bucket"`
	BucketRegion string `json:"region"`
}

// input request params
type RequestParams struct {
	Source      Source      `json:"source"`
	Credentials Credentials `json:"credentials"`
}

// input payload
type InputPayLoadParams struct {
	PayLoad string `json:"payload"`
}

type SvcInput struct {
	Bucket, Region, AccessKey, SecretKey string
}
