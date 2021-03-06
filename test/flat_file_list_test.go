package test

import (
	"fmt"
	diskclient "github.com/yantonov/yandex-disk-restapi-go/src"
	"net/http"
	"reflect"
	"testing"
)

func Test_FlatFileList_Simple(t *testing.T) {
	var client = NewStubResponseClient(`{
  "items": [
      {
        "name": "photo2.png",
        "preview": "https://downloader.disk.yandex.ru/preview/...",
        "created": "2014-04-22T14:57:13+04:00",
        "modified": "2014-04-22T14:57:14+04:00",
        "path": "disk:/foo/photo2.png",
        "md5": "53f4dc6379c8f95ddf11b9508cfea271",
        "type": "file",
        "mime_type": "image/png",
        "size": 54321
      },
      {
        "name": "photo1.png",
        "preview": "https://downloader.disk.yandex.ru/preview/...",
        "created": "2014-04-21T14:57:13+04:00",
        "modified": "2014-04-21T14:57:14+04:00",
        "path": "disk:/foo/photo1.png",
        "md5": "4334dc6379c8f95ddf11b9508cfea271",
        "type": "file",
        "mime_type": "image/png",
        "size": 34567
      }
    ],
    "limit": 20,
    "offset": 0
  }`, http.StatusOK)
	response, err := client.NewFlatFileListRequest().Exec()
	if err != nil {
		t.Error(fmt.Sprintf("unexpected error %s", err.Error()))
	}
	var expected = &diskclient.FilesResourceListResponse{}
	var resource1 = diskclient.ResourceInfoResponse{
		Name:          "photo2.png",
		Preview:       "https://downloader.disk.yandex.ru/preview/...",
		Created:       "2014-04-22T14:57:13+04:00",
		Modified:      "2014-04-22T14:57:14+04:00",
		Path:          "disk:/foo/photo2.png",
		Md5:           "53f4dc6379c8f95ddf11b9508cfea271",
		Resource_type: "file",
		Mime_type:     "image/png",
		Size:          54321,
	}
	var resource2 = diskclient.ResourceInfoResponse{
		Name:          "photo1.png",
		Preview:       "https://downloader.disk.yandex.ru/preview/...",
		Created:       "2014-04-21T14:57:13+04:00",
		Modified:      "2014-04-21T14:57:14+04:00",
		Path:          "disk:/foo/photo1.png",
		Md5:           "4334dc6379c8f95ddf11b9508cfea271",
		Resource_type: "file",
		Mime_type:     "image/png",
		Size:          34567,
	}
	expected.Items = []diskclient.ResourceInfoResponse{resource1, resource2}
	var limit uint64 = 20
	var offset uint64 = 0
	expected.Offset = &offset
	expected.Limit = &limit

	if !reflect.DeepEqual(response, expected) {
		t.Errorf("should match\nactual   = %v\nexpected = %v", response, expected)
	}
}

func Test_FlatFileList_NoItemsInResponse(t *testing.T) {
	var client = NewStubResponseClient(`{
    "limit": 20,
    "offset": 0
  }`, http.StatusOK)
	response, err := client.NewFlatFileListRequest().Exec()
	if err != nil {
		t.Error(fmt.Sprintf("unexpected error %s", err.Error()))
	}
	var expected = &diskclient.FilesResourceListResponse{}
	expected.Items = []diskclient.ResourceInfoResponse{}
	var limit uint64 = 20
	var offset uint64 = 0
	expected.Offset = &offset
	expected.Limit = &limit

	if !reflect.DeepEqual(response, expected) {
		t.Errorf("should match\nactual   = %v\nexpected = %v", response, expected)
	}
}

func Test_FlatFileListRequest_Limit(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var limit uint32 = 123456
	var options = diskclient.FlatFileListRequestOptions{
		Limit: &limit,
	}
	request := client.NewFlatFileListRequest(options).Request()

	if request.Parameters["limit"] == nil {
		t.Errorf("limit is undefined")
	}
	request_limit := (request.Parameters["limit"]).(*uint32)
	if *request_limit != 123456 {
		t.Errorf("invalid limit, actual : %d", *request_limit)
	}
}

func Test_FlatFileListRequest_NoLimit(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var options = diskclient.FlatFileListRequestOptions{}
	request := client.NewFlatFileListRequest(options).Request()

	if request.Parameters["limit"] != nil {
		t.Errorf("limit must be undefined")
	}
}

func Test_FlatFileListRequest_Offset(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var offset uint32 = 123456
	var options = diskclient.FlatFileListRequestOptions{
		Offset: &offset,
	}
	request := client.NewFlatFileListRequest(options).Request()

	if request.Parameters["offset"] == nil {
		t.Errorf("offset is undefined")
	}
	request_offset := (request.Parameters["offset"]).(*uint32)
	if *request_offset != 123456 {
		t.Errorf("invalid offset, actual : %d", *request_offset)
	}
}

func Test_FlatFileListRequest_NoOffset(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var options = diskclient.FlatFileListRequestOptions{}
	request := client.NewFlatFileListRequest(options).Request()

	if request.Parameters["offset"] != nil {
		t.Errorf("offset must be undefined")
	}
}

func Test_FlatFileListRequest_PreviewSize(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var options = diskclient.FlatFileListRequestOptions{
		Preview_size: (&diskclient.PreviewSize{}).PredefinedSizeM(),
	}
	request := client.NewFlatFileListRequest(options).Request()

	size := request.Parameters["preview_size"]
	if size != "M" {
		t.Errorf("invalid preview_size, actual : %d", size)
	}
}

func Test_FlatFileListRequest_NoPreviewSize(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var options = diskclient.FlatFileListRequestOptions{}
	request := client.NewFlatFileListRequest(options).Request()

	size := request.Parameters["preview_size"]
	if size != nil {
		t.Errorf("preview size must be undefined")
	}
}

func Test_FlatFileListRequest_PreviewCrop(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	crop := true
	var options = diskclient.FlatFileListRequestOptions{
		Preview_crop: &crop,
	}
	request := client.NewFlatFileListRequest(options).Request()

	extracted_crop := request.Parameters["preview_crop"].(*bool)
	if *extracted_crop != true {
		t.Errorf("invalid preview_crop, actual : %v", *extracted_crop)
	}
}

func Test_FlatFileListRequest_NoPreviewCrop(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var options = diskclient.FlatFileListRequestOptions{}
	request := client.NewFlatFileListRequest(options).Request()

	if request.Parameters["preview_crop"] != nil {
		t.Errorf("preview_crop must be undefined")
	}
}

func Test_FlatFileListRequest_FieldsList(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var fields = []string{"a", "b"}
	var options = diskclient.FlatFileListRequestOptions{
		Fields: fields,
	}
	request := client.NewFlatFileListRequest(options).Request()

	extracted_param := request.Parameters["fields"]
	if extracted_param != "a,b" {
		t.Errorf("invalid fields param, actual : %v", extracted_param)
	}
}

func Test_FlatFileListRequest_EmptyFieldsList(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var options = diskclient.FlatFileListRequestOptions{}
	request := client.NewFlatFileListRequest(options).Request()

	extracted_param := request.Parameters["fields"]
	if extracted_param != nil {
		t.Errorf("fields param must be undefined")
	}
}

func Test_FlatFileListRequest_MediaTypes(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var options = diskclient.FlatFileListRequestOptions{}
	options.Media_type = []diskclient.MediaType{
		*(&diskclient.MediaType{}).Audio(),
		*(&diskclient.MediaType{}).Backup(),
	}
	request := client.NewFlatFileListRequest(options).Request()

	extracted_param := request.Parameters["media_type"]
	if extracted_param != "audio,backup" {
		t.Errorf(fmt.Sprintf("invalid media type, actual = %v", request.Parameters["media_type"]))
	}
}

func Test_FlatFileListRequest_NoMediaTypes(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var options = diskclient.FlatFileListRequestOptions{}
	request := client.NewFlatFileListRequest(options).Request()

	extracted_param := request.Parameters["media_type"]
	if extracted_param != nil {
		t.Errorf("media type is not expected")
	}
}
