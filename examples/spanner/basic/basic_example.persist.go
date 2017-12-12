// This file is generated by protoc-gen-persist
// Source File: examples/spanner/basic/basic_example.proto
// DO NOT EDIT !
package basic

import (
	io "io"

	"cloud.google.com/go/spanner"
	mytime "github.com/tcncloud/protoc-gen-persist/examples/mytime"
	persist_lib "github.com/tcncloud/protoc-gen-persist/examples/spanner/basic/persist_lib"
	hooks "github.com/tcncloud/protoc-gen-persist/examples/spanner/hooks"
	test "github.com/tcncloud/protoc-gen-persist/examples/test"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	gstatus "google.golang.org/grpc/status"
)

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// WARNING ! WARNING ! WARNING ! WARNING !WARNING ! WARNING !
// In order for your code to work you have to create a file
// in this package with the following content:
//
// type MySpannerImpl struct {
//	PERSIST *persist_lib.MySpannerPersistHelper
// }
// WARNING ! WARNING ! WARNING ! WARNING !WARNING ! WARNING !
//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (s *MySpannerImpl) UniaryInsert(ctx context.Context, req *test.ExampleTable) (*test.ExampleTable, error) {
	var err error
	_ = err

	params := &persist_lib.MySpannerUniaryInsertInput{}
	params.Id = req.Id
	params.Name = req.Name
	if params.StartTime, err = (mytime.MyTime{}).ToSpanner(req.StartTime).SpannerValue(); err != nil {
		return nil, gstatus.Errorf(codes.Unknown, "could not convert type: %v", err)
	}

	var res = test.ExampleTable{}
	var iterErr error
	err = s.PERSIST.UniaryInsert(ctx, params, func(row *spanner.Row) {
		if row == nil { // there was no return data
			return
		}
		var Id int64
		if err := row.ColumnByName("id", &Id); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Id = Id
		var StartTime *spanner.GenericColumnValue
		if err := row.ColumnByName("start_time", StartTime); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		{
			local := &mytime.MyTime{}
			if err := local.SpannerScan(StartTime); err != nil {
				iterErr = gstatus.Errorf(codes.Unknown, "could not scan out custom type: %s", err)
				return
			}
			res.StartTime = local.ToProto()
		}
		var Name string
		if err := row.ColumnByName("name", &Name); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Name = Name
	})
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *MySpannerImpl) UniarySelect(ctx context.Context, req *test.ExampleTable) (*test.ExampleTable, error) {
	var err error
	_ = err

	params := &persist_lib.MySpannerUniarySelectInput{}
	params.Id = req.Id
	params.Name = req.Name

	var res = test.ExampleTable{}
	var iterErr error
	err = s.PERSIST.UniarySelect(ctx, params, func(row *spanner.Row) {
		if row == nil { // there was no return data
			return
		}
		var Id int64
		if err := row.ColumnByName("id", &Id); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Id = Id
		var StartTime *spanner.GenericColumnValue
		if err := row.ColumnByName("start_time", StartTime); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		{
			local := &mytime.MyTime{}
			if err := local.SpannerScan(StartTime); err != nil {
				iterErr = gstatus.Errorf(codes.Unknown, "could not scan out custom type: %s", err)
				return
			}
			res.StartTime = local.ToProto()
		}
		var Name string
		if err := row.ColumnByName("name", &Name); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Name = Name
	})
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *MySpannerImpl) UniaryUpdate(ctx context.Context, req *test.ExampleTable) (*test.PartialTable, error) {
	var err error
	_ = err

	params := &persist_lib.MySpannerUniaryUpdateInput{}
	params.Id = req.Id
	params.Name = req.Name
	if params.StartTime, err = (mytime.MyTime{}).ToSpanner(req.StartTime).SpannerValue(); err != nil {
		return nil, gstatus.Errorf(codes.Unknown, "could not convert type: %v", err)
	}

	var res = test.PartialTable{}
	var iterErr error
	err = s.PERSIST.UniaryUpdate(ctx, params, func(row *spanner.Row) {
		if row == nil { // there was no return data
			return
		}
		var Id int64
		if err := row.ColumnByName("id", &Id); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Id = Id
		var StartTime *spanner.GenericColumnValue
		if err := row.ColumnByName("start_time", StartTime); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		{
			local := &mytime.MyTime{}
			if err := local.SpannerScan(StartTime); err != nil {
				iterErr = gstatus.Errorf(codes.Unknown, "could not scan out custom type: %s", err)
				return
			}
			res.StartTime = local.ToProto()
		}
	})
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *MySpannerImpl) UniaryDeleteRange(ctx context.Context, req *test.ExampleTableRange) (*test.ExampleTable, error) {
	var err error
	_ = err

	params := &persist_lib.MySpannerUniaryDeleteRangeInput{}
	params.EndId = req.EndId
	params.StartId = req.StartId

	var res = test.ExampleTable{}
	var iterErr error
	err = s.PERSIST.UniaryDeleteRange(ctx, params, func(row *spanner.Row) {
		if row == nil { // there was no return data
			return
		}
		var Id int64
		if err := row.ColumnByName("id", &Id); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Id = Id
		var StartTime *spanner.GenericColumnValue
		if err := row.ColumnByName("start_time", StartTime); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		{
			local := &mytime.MyTime{}
			if err := local.SpannerScan(StartTime); err != nil {
				iterErr = gstatus.Errorf(codes.Unknown, "could not scan out custom type: %s", err)
				return
			}
			res.StartTime = local.ToProto()
		}
		var Name string
		if err := row.ColumnByName("name", &Name); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Name = Name
	})
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *MySpannerImpl) UniaryDeleteSingle(ctx context.Context, req *test.ExampleTable) (*test.ExampleTable, error) {
	var err error
	_ = err

	params := &persist_lib.MySpannerUniaryDeleteSingleInput{}
	params.Id = req.Id

	var res = test.ExampleTable{}
	var iterErr error
	err = s.PERSIST.UniaryDeleteSingle(ctx, params, func(row *spanner.Row) {
		if row == nil { // there was no return data
			return
		}
		var Id int64
		if err := row.ColumnByName("id", &Id); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Id = Id
		var StartTime *spanner.GenericColumnValue
		if err := row.ColumnByName("start_time", StartTime); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		{
			local := &mytime.MyTime{}
			if err := local.SpannerScan(StartTime); err != nil {
				iterErr = gstatus.Errorf(codes.Unknown, "could not scan out custom type: %s", err)
				return
			}
			res.StartTime = local.ToProto()
		}
		var Name string
		if err := row.ColumnByName("name", &Name); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Name = Name
	})
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *MySpannerImpl) NoArgs(ctx context.Context, req *test.ExampleTable) (*test.ExampleTable, error) {
	var err error
	_ = err

	params := &persist_lib.MySpannerNoArgsInput{}

	var res = test.ExampleTable{}
	var iterErr error
	err = s.PERSIST.NoArgs(ctx, params, func(row *spanner.Row) {
		if row == nil { // there was no return data
			return
		}
		var Id int64
		if err := row.ColumnByName("id", &Id); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Id = Id
		var StartTime *spanner.GenericColumnValue
		if err := row.ColumnByName("start_time", StartTime); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		{
			local := &mytime.MyTime{}
			if err := local.SpannerScan(StartTime); err != nil {
				iterErr = gstatus.Errorf(codes.Unknown, "could not scan out custom type: %s", err)
				return
			}
			res.StartTime = local.ToProto()
		}
		var Name string
		if err := row.ColumnByName("name", &Name); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Name = Name
	})
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// spanner server streaming ServerStream
func (s *MySpannerImpl) ServerStream(req *test.Name, stream MySpanner_ServerStreamServer) error {
	var err error
	_ = err

	params := &persist_lib.MySpannerServerStreamInput{}

	var iterErr error
	err = s.PERSIST.ServerStream(stream.Context(), params, func(row *spanner.Row) {
		if iterErr != nil || row == nil {
			return
		}
		var res test.ExampleTable

		var Id int64
		if err := row.ColumnByName("id", &Id); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Id = Id
		var StartTime *spanner.GenericColumnValue
		if err := row.ColumnByName("start_time", StartTime); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		{
			local := &mytime.MyTime{}
			if err := local.SpannerScan(StartTime); err != nil {
				iterErr = gstatus.Errorf(codes.Unknown, "could not scan out custom type: %s", err)
				return
			}
			res.StartTime = local.ToProto()
		}

		var Name string
		if err := row.ColumnByName("name", &Name); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Name = Name

		if err := stream.Send(&res); err != nil {
			iterErr = err
			return
		}
	})
	if err != nil {
		return err
	} else if iterErr != nil {
		return iterErr
	}
	return nil
}

func (s *MySpannerImpl) ClientStreamInsert(stream MySpanner_ClientStreamInsertServer) error {
	var err error
	_ = err
	feed, stop := s.PERSIST.ClientStreamInsert(stream.Context())
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return gstatus.Errorf(codes.Unknown, err.Error())
		}

		params := &persist_lib.MySpannerClientStreamInsertInput{}
		params.Id = req.Id
		params.Name = req.Name
		if params.StartTime, err = (mytime.MyTime{}).ToSpanner(req.StartTime).SpannerValue(); err != nil {
			return gstatus.Errorf(codes.Unknown, "could not convert type: %v", err)
		}

		feed(params)
	}
	row, err := stop()
	if err != nil {
		return err
	}
	res := test.NumRows{}
	if row != nil {

		var Count int64
		if err := row.ColumnByName("count", &Count); err != nil {
			return gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Count = Count

	}

	if err := stream.SendAndClose(&res); err != nil {
		return err
	}
	return nil
}

func (s *MySpannerImpl) ClientStreamDelete(stream MySpanner_ClientStreamDeleteServer) error {
	var err error
	_ = err
	feed, stop := s.PERSIST.ClientStreamDelete(stream.Context())
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return gstatus.Errorf(codes.Unknown, err.Error())
		}

		params := &persist_lib.MySpannerClientStreamDeleteInput{}
		params.Id = req.Id

		feed(params)
	}
	row, err := stop()
	if err != nil {
		return err
	}
	res := test.NumRows{}
	if row != nil {

		var Count int64
		if err := row.ColumnByName("count", &Count); err != nil {
			return gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Count = Count

	}

	if err := stream.SendAndClose(&res); err != nil {
		return err
	}
	return nil
}

func (s *MySpannerImpl) ClientStreamUpdate(stream MySpanner_ClientStreamUpdateServer) error {
	var err error
	_ = err
	feed, stop := s.PERSIST.ClientStreamUpdate(stream.Context())
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return gstatus.Errorf(codes.Unknown, err.Error())
		}

		params := &persist_lib.MySpannerClientStreamUpdateInput{}
		params.Id = req.Id
		params.Name = req.Name
		if params.StartTime, err = (mytime.MyTime{}).ToSpanner(req.StartTime).SpannerValue(); err != nil {
			return gstatus.Errorf(codes.Unknown, "could not convert type: %v", err)
		}

		feed(params)
	}
	row, err := stop()
	if err != nil {
		return err
	}
	res := test.NumRows{}
	if row != nil {

		var Count int64
		if err := row.ColumnByName("count", &Count); err != nil {
			return gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Count = Count

	}

	if err := stream.SendAndClose(&res); err != nil {
		return err
	}
	return nil
}

func (s *MySpannerImpl) UniaryInsertWithHooks(ctx context.Context, req *test.ExampleTable) (*test.ExampleTable, error) {
	var err error
	_ = err

	beforeRes, err := hooks.UniaryInsertBeforeHook(req)
	if err != nil {
		return nil, grpc.Errorf(codes.Unknown, err.Error())
	}
	if beforeRes != nil {

		return beforeRes, nil
	}

	params := &persist_lib.MySpannerUniaryInsertWithHooksInput{}
	params.Id = req.Id
	params.Name = req.Name
	if params.StartTime, err = (mytime.MyTime{}).ToSpanner(req.StartTime).SpannerValue(); err != nil {
		return nil, gstatus.Errorf(codes.Unknown, "could not convert type: %v", err)
	}

	var res = test.ExampleTable{}
	var iterErr error
	err = s.PERSIST.UniaryInsertWithHooks(ctx, params, func(row *spanner.Row) {
		if row == nil { // there was no return data
			return
		}
		var Id int64
		if err := row.ColumnByName("id", &Id); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Id = Id
		var StartTime *spanner.GenericColumnValue
		if err := row.ColumnByName("start_time", StartTime); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		{
			local := &mytime.MyTime{}
			if err := local.SpannerScan(StartTime); err != nil {
				iterErr = gstatus.Errorf(codes.Unknown, "could not scan out custom type: %s", err)
				return
			}
			res.StartTime = local.ToProto()
		}
		var Name string
		if err := row.ColumnByName("name", &Name); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Name = Name
	})
	if err != nil {
		return nil, err
	}

	err = hooks.UniaryInsertAfterHook(req, &res)
	if err != nil {
		return nil, grpc.Errorf(codes.Unknown, err.Error())
	}

	return &res, nil
}

func (s *MySpannerImpl) UniarySelectWithHooks(ctx context.Context, req *test.ExampleTable) (*test.ExampleTable, error) {
	var err error
	_ = err

	beforeRes, err := hooks.UniaryInsertBeforeHook(req)
	if err != nil {
		return nil, grpc.Errorf(codes.Unknown, err.Error())
	}
	if beforeRes != nil {

		return beforeRes, nil
	}

	params := &persist_lib.MySpannerUniarySelectWithHooksInput{}
	params.Id = req.Id

	var res = test.ExampleTable{}
	var iterErr error
	err = s.PERSIST.UniarySelectWithHooks(ctx, params, func(row *spanner.Row) {
		if row == nil { // there was no return data
			return
		}
		var Id int64
		if err := row.ColumnByName("id", &Id); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Id = Id
		var StartTime *spanner.GenericColumnValue
		if err := row.ColumnByName("start_time", StartTime); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		{
			local := &mytime.MyTime{}
			if err := local.SpannerScan(StartTime); err != nil {
				iterErr = gstatus.Errorf(codes.Unknown, "could not scan out custom type: %s", err)
				return
			}
			res.StartTime = local.ToProto()
		}
		var Name string
		if err := row.ColumnByName("name", &Name); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Name = Name
	})
	if err != nil {
		return nil, err
	}

	err = hooks.UniaryInsertAfterHook(req, &res)
	if err != nil {
		return nil, grpc.Errorf(codes.Unknown, err.Error())
	}

	return &res, nil
}

func (s *MySpannerImpl) UniaryUpdateWithHooks(ctx context.Context, req *test.ExampleTable) (*test.PartialTable, error) {
	var err error
	_ = err

	beforeRes, err := hooks.UniaryUpdateBeforeHook(req)
	if err != nil {
		return nil, grpc.Errorf(codes.Unknown, err.Error())
	}
	if beforeRes != nil {

		return beforeRes, nil
	}

	params := &persist_lib.MySpannerUniaryUpdateWithHooksInput{}
	params.Id = req.Id
	params.Name = req.Name
	if params.StartTime, err = (mytime.MyTime{}).ToSpanner(req.StartTime).SpannerValue(); err != nil {
		return nil, gstatus.Errorf(codes.Unknown, "could not convert type: %v", err)
	}

	var res = test.PartialTable{}
	var iterErr error
	err = s.PERSIST.UniaryUpdateWithHooks(ctx, params, func(row *spanner.Row) {
		if row == nil { // there was no return data
			return
		}
		var Id int64
		if err := row.ColumnByName("id", &Id); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Id = Id
		var StartTime *spanner.GenericColumnValue
		if err := row.ColumnByName("start_time", StartTime); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		{
			local := &mytime.MyTime{}
			if err := local.SpannerScan(StartTime); err != nil {
				iterErr = gstatus.Errorf(codes.Unknown, "could not scan out custom type: %s", err)
				return
			}
			res.StartTime = local.ToProto()
		}
	})
	if err != nil {
		return nil, err
	}

	err = hooks.UniaryUpdateAfterHook(req, &res)
	if err != nil {
		return nil, grpc.Errorf(codes.Unknown, err.Error())
	}

	return &res, nil
}

func (s *MySpannerImpl) UniaryDeleteWithHooks(ctx context.Context, req *test.ExampleTableRange) (*test.ExampleTable, error) {
	var err error
	_ = err

	beforeRes, err := hooks.UniaryDeleteBeforeHook(req)
	if err != nil {
		return nil, grpc.Errorf(codes.Unknown, err.Error())
	}
	if beforeRes != nil {

		return beforeRes, nil
	}

	params := &persist_lib.MySpannerUniaryDeleteWithHooksInput{}
	params.EndId = req.EndId
	params.StartId = req.StartId

	var res = test.ExampleTable{}
	var iterErr error
	err = s.PERSIST.UniaryDeleteWithHooks(ctx, params, func(row *spanner.Row) {
		if row == nil { // there was no return data
			return
		}
		var Id int64
		if err := row.ColumnByName("id", &Id); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Id = Id
		var StartTime *spanner.GenericColumnValue
		if err := row.ColumnByName("start_time", StartTime); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		{
			local := &mytime.MyTime{}
			if err := local.SpannerScan(StartTime); err != nil {
				iterErr = gstatus.Errorf(codes.Unknown, "could not scan out custom type: %s", err)
				return
			}
			res.StartTime = local.ToProto()
		}
		var Name string
		if err := row.ColumnByName("name", &Name); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Name = Name
	})
	if err != nil {
		return nil, err
	}

	err = hooks.UniaryDeleteAfterHook(req, &res)
	if err != nil {
		return nil, grpc.Errorf(codes.Unknown, err.Error())
	}

	return &res, nil
}

// spanner server streaming ServerStreamWithHooks
func (s *MySpannerImpl) ServerStreamWithHooks(req *test.Name, stream MySpanner_ServerStreamWithHooksServer) error {
	var err error
	_ = err

	beforeRes, err := hooks.ServerStreamBeforeHook(req)
	if err != nil {
		return grpc.Errorf(codes.Unknown, err.Error())
	}
	if beforeRes != nil {

		for _, res := range beforeRes {
			err = stream.Send(res)
			if err != nil {
				return err
			}
		}
		return nil
	}

	params := &persist_lib.MySpannerServerStreamWithHooksInput{}

	var iterErr error
	err = s.PERSIST.ServerStreamWithHooks(stream.Context(), params, func(row *spanner.Row) {
		if iterErr != nil || row == nil {
			return
		}
		var res test.ExampleTable

		var Id int64
		if err := row.ColumnByName("id", &Id); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Id = Id
		var StartTime *spanner.GenericColumnValue
		if err := row.ColumnByName("start_time", StartTime); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		{
			local := &mytime.MyTime{}
			if err := local.SpannerScan(StartTime); err != nil {
				iterErr = gstatus.Errorf(codes.Unknown, "could not scan out custom type: %s", err)
				return
			}
			res.StartTime = local.ToProto()
		}

		var Name string
		if err := row.ColumnByName("name", &Name); err != nil {
			iterErr = gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Name = Name

		err = hooks.ServerStreamAfterHook(req, &res)
		if err != nil {
			iterErr = grpc.Errorf(codes.Unknown, err.Error())
		}

		if err := stream.Send(&res); err != nil {
			iterErr = err
			return
		}
	})
	if err != nil {
		return err
	} else if iterErr != nil {
		return iterErr
	}
	return nil
}

func (s *MySpannerImpl) ClientStreamUpdateWithHooks(stream MySpanner_ClientStreamUpdateWithHooksServer) error {
	var err error
	_ = err
	feed, stop := s.PERSIST.ClientStreamUpdateWithHooks(stream.Context())
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return gstatus.Errorf(codes.Unknown, err.Error())
		}

		beforeRes, err := hooks.ClientStreamUpdateBeforeHook(req)
		if err != nil {
			return grpc.Errorf(codes.Unknown, err.Error())
		}
		if beforeRes != nil {
			continue

		}

		params := &persist_lib.MySpannerClientStreamUpdateWithHooksInput{}
		params.Id = req.Id
		params.Name = req.Name

		feed(params)
	}
	row, err := stop()
	if err != nil {
		return err
	}
	res := test.NumRows{}
	if row != nil {

		var Count int64
		if err := row.ColumnByName("count", &Count); err != nil {
			return gstatus.Errorf(codes.Unknown, "could not convert type %v", err)
		}

		res.Count = Count

	}

	err = hooks.ClientStreamUpdateAfterHook(nil, &res)
	if err != nil {
		return grpc.Errorf(codes.Unknown, err.Error())
	}

	if err := stream.SendAndClose(&res); err != nil {
		return err
	}
	return nil
}
