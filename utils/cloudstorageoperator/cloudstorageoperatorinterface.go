package cloudstorageoperator

import (
	"context"

	"cloud.google.com/go/iam"
	"cloud.google.com/go/storage"
)

type IBucketHandle interface {
	ACL() *storage.ACLHandle
	AddNotification(ctx context.Context, n *storage.Notification) (ret *storage.Notification, err error)
	Attrs(ctx context.Context) (attrs *storage.BucketAttrs, err error)
	Create(ctx context.Context, projectID string, attrs *storage.BucketAttrs) (err error)
	DefaultObjectACL() *storage.ACLHandle
	Delete(ctx context.Context) (err error)
	DeleteNotification(ctx context.Context, id string) (err error)
	GenerateSignedPostPolicyV4(object string, opts *storage.PostPolicyV4Options) (*storage.PostPolicyV4, error)
	IAM() *iam.Handle
	If(conds storage.BucketConditions) *storage.BucketHandle
	LockRetentionPolicy(ctx context.Context) error
	Notifications(ctx context.Context) (n map[string]*storage.Notification, err error)
	Object(name string) *storage.ObjectHandle
	Objects(ctx context.Context, q *storage.Query) *storage.ObjectIterator
	Retryer(opts ...storage.RetryOption) *storage.BucketHandle
	SignedURL(object string, opts *storage.SignedURLOptions) (string, error)
	Update(ctx context.Context, uattrs storage.BucketAttrsToUpdate) (attrs *storage.BucketAttrs, err error)
	UserProject(projectID string) *storage.BucketHandle
}
