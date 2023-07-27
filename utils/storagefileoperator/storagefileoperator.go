package storagefileoperator

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func StorageFileUpload(file multipart.File, bucket string, objectName string, cloudStorage *storage.Client) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	w := os.Stdout
	wc := cloudStorage.Bucket(bucket).Object(objectName).NewWriter(ctx)

	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}

	if err := wc.Close(); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}

	fmt.Fprintf(w, "%v uploaded to %v. \n", objectName, bucket)

	return nil
}

func StorageFileUrlGet(bucket string, objectName string, cloudStorage *storage.Client) (string, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	attrs, _ := cloudStorage.Bucket(bucket).Object(objectName).Attrs(ctx)
	if attrs != nil {
		storageLink := "storage_link"
		objectPublicUrl := storageLink + objectName

		fmt.Println("objectPublicUrl: ", objectPublicUrl)
		return objectPublicUrl, nil
	} else {
		fmt.Println("Object do not exist.")
		return "", fmt.Errorf("Object do not exist.")
	}
}

func StorageGetFileList(bucket string, prefix string, cloudStorage *storage.Client) ([]interface{}, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// ex. Prefix: "ECGFiles/"
	objects := cloudStorage.Bucket(bucket).Objects(ctx, &storage.Query{Prefix: prefix, Delimiter: "/"})
	tableData := make([]interface{}, 0)
	for {
		obj, err := objects.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		tableData = append(tableData, obj.Prefix)
		fmt.Println("obj.Name: ", obj.Prefix)
	}
	fmt.Println("tableData: ", tableData)
	return tableData, nil
}
