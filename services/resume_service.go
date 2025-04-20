package services

import (
	"backend/config"
	"backend/db"
	"backend/models"
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

func GetResume() (*models.Resume, error) {
	var resume models.Resume
	if err := db.DB.First(&resume).Error; err != nil {
		return nil, err
	}
	return &resume, nil
}

func UploadResumeToR2(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	// Get R2 credentials using config.GetEnv
	r2Endpoint := config.GetEnv("R2_ENDPOINT", "")
	r2AccessKey := config.GetEnv("R2_ACCESS_KEY", "")
	r2SecretKey := config.GetEnv("R2_SECRET_KEY", "")
	r2Bucket := config.GetEnv("R2_BUCKET", "")

	if r2Endpoint == "" || r2AccessKey == "" || r2SecretKey == "" || r2Bucket == "" {
		return "", fmt.Errorf("R2 credentials are not properly configured")
	}

	// Create a new AWS session with S3-compatible configuration
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String("auto"),
		Endpoint:         aws.String(r2Endpoint),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      credentials.NewStaticCredentials(r2AccessKey, r2SecretKey, ""),
	})
	if err != nil {
		return "", fmt.Errorf("failed to create AWS session: %w", err)
	}

	s3Client := s3.New(sess)

	// Generate a unique file name
	uniqueFileName := fmt.Sprintf("%s-%s", uuid.New().String(), fileHeader.Filename)

	// Read the file content into a buffer
	buffer := new(bytes.Buffer)
	if _, err := io.Copy(buffer, file); err != nil {
		return "", fmt.Errorf("failed to read file content: %w", err)
	}

	// Upload the file to the R2 bucket
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(r2Bucket),
		Key:    aws.String(uniqueFileName),
		Body:   bytes.NewReader(buffer.Bytes()),
		ACL:    aws.String("private"),
		ContentType: aws.String(fileHeader.Header.Get("Content-Type")),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file to R2: %w", err)
	}

	// Return the file URL
	fileURL := fmt.Sprintf("%s/%s/%s", r2Endpoint, r2Bucket, uniqueFileName)
	return fileURL, nil
}

func DeleteFileFromR2(fileURL string) error {
	// Get R2 credentials using config.GetEnv
	r2Endpoint := config.GetEnv("R2_ENDPOINT", "")
	r2AccessKey := config.GetEnv("R2_ACCESS_KEY", "")
	r2SecretKey := config.GetEnv("R2_SECRET_KEY", "")
	r2Bucket := config.GetEnv("R2_BUCKET", "")

	if r2Endpoint == "" || r2AccessKey == "" || r2SecretKey == "" || r2Bucket == "" {
		return fmt.Errorf("R2 credentials are not properly configured")
	}

	// Create a new AWS session with S3-compatible configuration
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String("auto"),
		Endpoint:         aws.String(r2Endpoint),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      credentials.NewStaticCredentials(r2AccessKey, r2SecretKey, ""),
	})
	if err != nil {
		return fmt.Errorf("failed to create AWS session: %w", err)
	}

	s3Client := s3.New(sess)

	// Extract the file key from the file URL
	fileKey := fileURL[len(fmt.Sprintf("%s/%s/", r2Endpoint, r2Bucket)):] // Remove the bucket URL prefix

	// Delete the file from the R2 bucket
	_, err = s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(r2Bucket),
		Key:    aws.String(fileKey),
	})
	if err != nil {
		return fmt.Errorf("failed to delete file from R2: %w", err)
	}

	return nil
}

func CreateResume(file multipart.File, fileHeader *multipart.FileHeader) (*models.Resume, error) {
	// Upload the file to R2
	fileURL, err := UploadResumeToR2(file, fileHeader)
	if err != nil {
		log.Printf("Error uploading resume to R2: %v", err)
		return nil, err
	}

	// Create a new resume record
	resume := &models.Resume{
		FileName: fileHeader.Filename,
		FileURL:  fileURL,
	}

	if err := db.DB.Create(resume).Error; err != nil {
		log.Printf("Error creating resume in database: %v", err)
		return nil, err
	}

	return resume, nil
}

func UpdateResume(file multipart.File, fileHeader *multipart.FileHeader) (*models.Resume, error) {
	// Retrieve the existing resume
	var resume models.Resume
	if err := db.DB.First(&resume).Error; err != nil {
		return nil, fmt.Errorf("failed to find existing resume: %w", err)
	}

	// Delete the old file from R2
	if err := DeleteFileFromR2(resume.FileURL); err != nil {
		return nil, fmt.Errorf("failed to delete old file from R2: %w", err)
	}

	// Upload the new file to R2
	fileURL, err := UploadResumeToR2(file, fileHeader)
	if err != nil {
		return nil, fmt.Errorf("failed to upload new file to R2: %w", err)
	}

	// Update the resume record in the database
	resume.FileName = fileHeader.Filename
	resume.FileURL = fileURL

	if err := db.DB.Save(&resume).Error; err != nil {
		return nil, fmt.Errorf("failed to update resume in database: %w", err)
	}

	return &resume, nil
}

func DeleteResume(id uint) error {
	// Retrieve the existing resume
	var resume models.Resume
	if err := db.DB.First(&resume, id).Error; err != nil {
		return fmt.Errorf("failed to find resume: %w", err)
	}

	// Delete the file from R2
	if err := DeleteFileFromR2(resume.FileURL); err != nil {
		return fmt.Errorf("failed to delete file from R2: %w", err)
	}

	// Delete the resume record from the database
	if err := db.DB.Delete(&resume).Error; err != nil {
		return fmt.Errorf("failed to delete resume from database: %w", err)
	}

	return nil
}