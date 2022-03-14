package intern

import "context"

type ProtoServer struct {
}

func (srv *ProtoServer) Update(context.Context, *ExtraUserInfoReq) (*ExtraUserInfoRes, error) {
	return nil, nil
}

func (srv *ProtoServer) mustEmbedUnimplementedUpdateExtraServer() {

}
