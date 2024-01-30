// List get list
func List(ctx context.Context, page, pageSize uint64, cond SearchCond) ([]*Model, error) {
	e, err := models.GetDB()
	if err != nil {
		log.Error(ctx, "models.GetDB() failed", zap.Error(err))
		return nil, err
	}

	var list []*Model
	b := e.Bind(tableName).SelectStruct(Model{})
	b = makeWhere(b, cond)
	if page == 0 && pageSize == 0 {
		err = b.All(ctx, &list)
	} else {
		err = b.Limit(pageSize).Offset((page-1)*pageSize).All(ctx, &list)
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return list, nil
		}
		log.Error(ctx, "select db failed", zap.Error(err))
		return nil, err
	}

	return list, nil
}
