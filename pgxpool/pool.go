func (p *Pool) AcquireFunc(ctx context.Context, fn func(*Conn) error) error {
	conn, err := p.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	return fn(conn)
}