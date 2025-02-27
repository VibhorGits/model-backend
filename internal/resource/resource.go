package resource

import (
	"context"
	"fmt"
	"strings"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type NamespaceType string

const (
	User         NamespaceType = "users"
	Organization NamespaceType = "organizations"
)

type Namespace struct {
	NsType NamespaceType
	NsUid  uuid.UUID
}

func (ns Namespace) String() string {
	return fmt.Sprintf("%s/%s", ns.NsType, ns.NsUid.String())
}

// ExtractFromMetadata extracts context metadata given a key
func ExtractFromMetadata(ctx context.Context, key string) ([]string, bool) {
	data, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return []string{}, false
	}
	return data[strings.ToLower(key)], true
}

// GetRequestSingleHeader get a request header, the header has to be single-value HTTP header
func GetRequestSingleHeader(ctx context.Context, header string) string {
	metaHeader := metadata.ValueFromIncomingContext(ctx, strings.ToLower(header))
	if len(metaHeader) != 1 {
		return ""
	}
	return metaHeader[0]
}

func GetDefinitionID(name string) (string, error) {
	id := strings.TrimPrefix(name, "model-definitions/")
	if !strings.HasPrefix(name, "model-definitions/") || id == "" {
		return "", status.Error(codes.InvalidArgument, "Error when extract model-definitions resource id")
	}
	return id, nil
}

// GetRscNameID returns the resource ID given a resource name
func GetRscNameID(path string) (string, error) {
	id := path[strings.LastIndex(path, "/")+1:]
	if id == "" {
		return "", fmt.Errorf("error when extract resource id from resource name '%s'", path)
	}
	return id, nil
}

// GetRscPermalinkUID returns the resource UID given a resource permalink
func GetRscPermalinkUID(path string) (uuid.UUID, error) {
	splits := strings.Split(path, "/")
	if len(splits) < 2 {
		return uuid.Nil, fmt.Errorf("error when extract resource id from resource permalink '%s'", path)
	}

	return uuid.FromStringOrNil(splits[1]), nil
}

func UserUidToUserPermalink(userUid uuid.UUID) string {
	return fmt.Sprintf("users/%s", userUid.String())
}

func GetOperationID(name string) (string, error) {
	id := strings.TrimPrefix(name, "operations/")
	if !strings.HasPrefix(name, "operations/") || id == "" {
		return "", status.Error(codes.InvalidArgument, "Error when extract operations resource id")
	}
	return id, nil
}
