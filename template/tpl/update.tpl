// Update modify a row
func Update(ctx context.Context, id int64, row db.Row) (int64, error) {
	e, err := models.GetDB()
	if err != nil {
		log.Error(ctx, "models.GetDB() failed", zap.Error(err))
		return 0, err
	}

	res, err := e.Bind(tableName).Update(row).Where(db.Eq("id", id)).Exec(ctx)
	if err != nil {
		log.Error(ctx, "update db failed", zap.Error(err))
		return 0, err
	}

	return res.RowsAffected()
}
