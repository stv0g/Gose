package server

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Setup initializes the S3 bucket (life-cycle rules & CORS)
func (s *Server) Setup() error {

	corsRule := &s3.CORSRule{
		AllowedHeaders: aws.StringSlice([]string{"Authorization"}),
		AllowedOrigins: aws.StringSlice([]string{"*"}),
		MaxAgeSeconds:  aws.Int64(3000),
		AllowedMethods: aws.StringSlice([]string{"PUT", "GET"}),
		ExposeHeaders:  aws.StringSlice([]string{"ETag"}),
	}

	if _, err := s.PutBucketCors(&s3.PutBucketCorsInput{
		Bucket: aws.String(s.Config.Bucket),
		CORSConfiguration: &s3.CORSConfiguration{
			CORSRules: []*s3.CORSRule{
				corsRule,
			},
		},
	}); err != nil {
		return fmt.Errorf("failed to set bucket %s's CORS rules: %w", s.Config.Bucket, err)
	}

	lcRules := []*s3.LifecycleRule{
		{
			ID:     aws.String("Abort Multipart Uploads"),
			Status: aws.String("Enabled"),
			AbortIncompleteMultipartUpload: &s3.AbortIncompleteMultipartUpload{
				DaysAfterInitiation: aws.Int64(31),
			},
			Filter: &s3.LifecycleRuleFilter{
				Prefix: aws.String("/"),
			},
		},
	}

	for _, cls := range s.Config.Expiration {
		lcRules = append(lcRules, &s3.LifecycleRule{
			ID:     aws.String(fmt.Sprintf("Expiration after %s", cls.Title)),
			Status: aws.String("Enabled"),
			Filter: &s3.LifecycleRuleFilter{
				Tag: &s3.Tag{
					Key:   aws.String("expiration"),
					Value: aws.String(cls.ID),
				},
			},
			Expiration: &s3.LifecycleExpiration{
				Days: aws.Int64(cls.Days),
			},
		})
	}

	if _, err := s.PutBucketLifecycleConfiguration(&s3.PutBucketLifecycleConfigurationInput{
		Bucket: aws.String(s.Config.Bucket),
		LifecycleConfiguration: &s3.BucketLifecycleConfiguration{
			Rules: lcRules,
		},
	}); err != nil {
		return fmt.Errorf("failed to set bucket %s's lifecycle rules: %w", s.Config.Bucket, err)
	}

	// lc, err := svc.GetBucketLifecycleConfiguration(&s3.GetBucketLifecycleConfigurationInput{
	// 	Bucket: aws.String(.Bucket),
	// })
	// if err != nil {
	// 	return fmt.Errorf("failed get life-cycle rules: %w", err)
	// }
	// log.Printf("Life-cycle rules: %+#v\n", lc)

	return nil
}
