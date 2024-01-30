// Insert add new row
func Insert(ctx context.Context, row db.Row) (int64,error) {
	e, err := models.GetDB()
	if err != nil {
		log.Error(ctx, "models.GetDB() failed", zap.Error(err))
		return 0, err
	}

	res, err := e.Bind(tableName).Insert(rows...).Exec(ctx)
	if err != nil {
		log.Error(ctx, "select db failed", zap.Error(err))
		return 0, err
	}

	return res.RowsAffected()
}
