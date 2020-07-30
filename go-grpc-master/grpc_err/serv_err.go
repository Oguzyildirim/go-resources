package main

import (
	"io"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/vladimirvivien/go-grpc/protobuf"
	"github.com/vladimirvivien/go-grpc/util"
)

const port = ":50051"

// CurrencyService implements the pb CurrencyServiceServer interface
type CurrencyService struct {
	mtex sync.Mutex
	data []*pb.Currency
}

func newCurrencyService(data []*pb.Currency) *CurrencyService {
	return &CurrencyService{data: data}
}

// GetCurrencyList searches (by Code or Number) and return CurrencyList
func (c *CurrencyService) GetCurrencyList(
	ctx context.Context,
	req *pb.CurrencyRequest,
) (*pb.CurrencyList, error) {

	if req.GetNumber() == 0 && req.GetCode() == "" {
		// dont return raw errors like the following:
		// return nil, fmt.Errorf("must provide currency number or code")

		// instread return status with informative code:
		return nil, status.Errorf(
			codes.InvalidArgument,
			"must provide currency number or code",
		)

	}

	var items []*pb.Currency
	for _, cur := range c.data {
		if cur.GetNumber() == req.GetNumber() || cur.GetCode() == req.GetCode() {
			items = append(items, cur)
		}
	}

	return &pb.CurrencyList{Items: items}, nil
}

// GetCurrencyStream returns matching Currencies as a server stream
func (c *CurrencyService) GetCurrencyStream(
	req *pb.CurrencyRequest,
	stream pb.CurrencyService_GetCurrencyStreamServer,
) error {

	if req.GetNumber() == 0 && req.GetCode() == "" {
		return status.Errorf(
			codes.InvalidArgument,
			"must provide currency number or code",
		)
	}

	for _, cur := range c.data {
		if cur.GetNumber() == req.GetNumber() || cur.GetCode() == req.GetCode() {
			if err := stream.Send(cur); err != nil {
				return err // err is rcp status
			}
		}
	}

	return nil
}

// SaveCurrencyStream adds Currency values from a stream to currency items
func (c *CurrencyService) SaveCurrencyStream(
	stream pb.CurrencyService_SaveCurrencyStreamServer,
) error {

	curList := new(pb.CurrencyList)
	for {
		cur, err := stream.Recv()

		// ensure we're not done first.
		if err != nil {
			// if done, close sream and return result
			if err == io.EOF {
				c.mtex.Lock()
				copy(c.data, curList.Items) // save to list
				c.mtex.Unlock()
				return stream.SendAndClose(curList)
			}
			return err
		}

		// validation with structured error
		// returned to the client
		if cur.GetName() == "" ||
			cur.Code == "" ||
			cur.Number == 0 || cur.Country == "" {

			stat := status.New(
				codes.InvalidArgument,
				"name, country, code, and number required",
			)
			// attach data structure to status error
			statDetail, err := stat.WithDetails(cur)
			if err != nil {
				return status.Errorf(codes.Internal, err.Error())
			}
			return statDetail.Err()
		}

		curList.Items = append(curList.Items, cur)
	}

}

// FindCurrencyStream sends a stream of CurrencyRequest while
// streaming Currency values from server.
func (c *CurrencyService) FindCurrencyStream(
	stream pb.CurrencyService_FindCurrencyStreamServer,
) error {

	for {
		req, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				return nil // we're done
			}
			return err
		}

		// validate req
		if req.GetNumber() == 0 && req.GetCode() == "" {
			return status.Errorf(
				codes.InvalidArgument,
				"invalid request, must provide number or code",
			)
		}

		// search and stream result
		for _, cur := range c.data {
			if cur.GetCode() == req.GetCode() || cur.GetNumber() == req.GetNumber() {
				if err := stream.Send(cur); err != nil {
					return err
				}
			}
		}
	}
}

func main() {

	// load data into protobuf structures
	data, err := util.LoadPbFromCsv("./../curdata.csv")
	if err != nil {
		log.Fatal(err) // dont start
	}

	lstnr, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to start server:", err)
	}

	// setup and register currency service
	curService := newCurrencyService(data)
	grpcServer := grpc.NewServer()
	pb.RegisterCurrencyServiceServer(grpcServer, curService)

	// start service's server
	log.Println("starting currency rpc service on", port)
	if err := grpcServer.Serve(lstnr); err != nil {
		log.Fatal(err)
	}
}
