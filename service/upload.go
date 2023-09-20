package upload

import (
	"context"
	"io"
	"os"
	"time"

	"google.golang.org/grpc"

	uploadpb "github.com/ahsoromdoni/stream-upload/grpc"
)

type Client struct {
	client uploadpb.UploadServiceClient
}

func NewClient(conn grpc.ClientConnInterface) Client {
	return Client{
		client: uploadpb.NewUploadServiceClient(conn),
	}
}

func (c Client) Upload(ctx context.Context, file string) (string, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(10*time.Second))
	defer cancel()

	stream, err := c.client.Upload(ctx)
	if err != nil {
		return "", err
	}

	fil, err := os.Open(file)
	if err != nil {
		return "", err
	}

	// Maximum 1KB size per stream.
	buf := make([]byte, 1024)

	for {
		num, err := fil.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		if err := stream.Send(&uploadpb.UploadRequest{Chunk: buf[:num]}); err != nil {
			return "", err
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return "", err
	}

	return res.GetName(), nil
}
