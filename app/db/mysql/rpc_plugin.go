package mysql

import "context"

type RpcPlugin struct{}

func (rp *RpcPlugin) PostCall(ctx context.Context, serviceName, methodName string, args, reply interface{}, err error) (interface{}, error) {
	var (
		instance = &DB{
			ctx: ctx,
		}
		tx = instance._TxCtx(ctx)
	)
	if tx == nil {
		return reply, nil
	}

	if err != nil {
		tx.Rollback()
		return reply, nil
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return reply, nil
}
