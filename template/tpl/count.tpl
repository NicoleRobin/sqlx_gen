// Count get count
func Count(ctx context.Context, cond SearchCond) (int64, error) {
	e, err := models.GetDB()
	if err != nil {
		log.Error(ctx, "models.GetDB() failed", zap.Error(err))
		return 0, err
	}

	b := e.Bind(tableName).SelectStruct(Model{})
	b = makeWhere(b, cond)
	count, err := b.Count(ctx)
	if err != nil {
		log.Error(ctx, "query count from db failed", zap.Error(err))
		return 0, err
	}
	return count, nil
}

