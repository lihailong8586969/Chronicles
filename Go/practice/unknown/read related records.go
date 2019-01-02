// 代码来自知乎 : https://www.zhihu.com/question/48324702/answer/240258951

// read related records.
func (d *dbBase) ReadBatch(q dbQuerier, qs *querySet, mi *modelInfo, cond *Condition, container interface{}, tz *time.Location, cols []string) (int64, error) {

	val := reflect.ValueOf(container)
	ind := reflect.Indirect(val)

	errTyp := true
	one := true
	isPtr := true

	if val.Kind() == reflect.Ptr {
		fn := ""
		if ind.Kind() == reflect.Slice {
			one = false
			typ := ind.Type().Elem()
			switch typ.Kind() {
			case reflect.Ptr:
				fn = getFullName(typ.Elem())
			case reflect.Struct:
				isPtr = false
				fn = getFullName(typ)
			}
		} else {
			fn = getFullName(ind.Type())
		}
		errTyp = fn != mi.fullName
	}

	if errTyp {
		if one {
			panic(fmt.Errorf("wrong object type `%s` for rows scan, need *%s", val.Type(), mi.fullName))
		} else {
			panic(fmt.Errorf("wrong object type `%s` for rows scan, need *[]*%s or *[]%s", val.Type(), mi.fullName, mi.fullName))
		}
	}

	rlimit := qs.limit
	offset := qs.offset

	Q := d.ins.TableQuote()

	var tCols []string
	if len(cols) > 0 {
		hasRel := len(qs.related) > 0 || qs.relDepth > 0
		tCols = make([]string, 0, len(cols))
		var maps map[string]bool
		if hasRel {
			maps = make(map[string]bool)
		}
		for _, col := range cols {
			if fi, ok := mi.fields.GetByAny(col); ok {
				tCols = append(tCols, fi.column)
				if hasRel {
					maps[fi.column] = true
				}
			} else {
				panic(fmt.Errorf("wrong field/column name `%s`", col))
			}
		}
		if hasRel {
			for _, fi := range mi.fields.fieldsDB {
				if fi.fieldType&IsRelField > 0 {
					if maps[fi.column] == false {
						tCols = append(tCols, fi.column)
					}
				}
			}
		}
	} else {
		tCols = mi.fields.dbcols
	}

	colsNum := len(tCols)
	sep := fmt.Sprintf("%s, T0.%s", Q, Q)
	sels := fmt.Sprintf("T0.%s%s%s", Q, strings.Join(tCols, sep), Q)

	tables := newDbTables(mi, d.ins)
	tables.parseRelated(qs.related, qs.relDepth)

	where, args := tables.getCondSql(cond, false, tz)
	orderBy := tables.getOrderSql(qs.orders)
	limit := tables.getLimitSql(mi, offset, rlimit)
	join := tables.getJoinSql()

	for _, tbl := range tables.tables {
		if tbl.sel {
			colsNum += len(tbl.mi.fields.dbcols)
			sep := fmt.Sprintf("%s, %s.%s", Q, tbl.index, Q)
			sels += fmt.Sprintf(", %s.%s%s%s", tbl.index, Q, strings.Join(tbl.mi.fields.dbcols, sep), Q)
		}
	}

	query := fmt.Sprintf("SELECT %s FROM %s%s%s T0 %s%s%s%s", sels, Q, mi.table, Q, join, where, orderBy, limit)

	d.ins.ReplaceMarks(&query)

	var rs *sql.Rows
	if r, err := q.Query(query, args...); err != nil {
		return 0, err
	} else {
		rs = r
	}

	refs := make([]interface{}, colsNum)
	for i, _ := range refs {
		var ref interface{}
		refs[i] = &ref
	}

	defer rs.Close()

	slice := ind

	var cnt int64
	for rs.Next() {
		if one && cnt == 0 || one == false {
			if err := rs.Scan(refs...); err != nil {
				return 0, err
			}

			elm := reflect.New(mi.addrField.Elem().Type())
			mind := reflect.Indirect(elm)

			cacheV := make(map[string]*reflect.Value)
			cacheM := make(map[string]*modelInfo)
			trefs := refs

			d.setColsValues(mi, &mind, tCols, refs[:len(tCols)], tz)
			trefs = refs[len(tCols):]

			for _, tbl := range tables.tables {
				// loop selected tables
				if tbl.sel {
					last := mind
					names := ""
					mmi := mi
					// loop cascade models
					for _, name := range tbl.names {
						names += name
						if val, ok := cacheV[names]; ok {
							last = *val
							mmi = cacheM[names]
						} else {
							fi := mmi.fields.GetByName(name)
							lastm := mmi
							mmi = fi.relModelInfo
							field := last
							if last.Kind() != reflect.Invalid {
								field = reflect.Indirect(last.Field(fi.fieldIndex))
								if field.IsValid() {
									d.setColsValues(mmi, &field, mmi.fields.dbcols, trefs[:len(mmi.fields.dbcols)], tz)
									for _, fi := range mmi.fields.fieldsReverse {
										if fi.inModel && fi.reverseFieldInfo.mi == lastm {
											if fi.reverseFieldInfo != nil {
												f := field.Field(fi.fieldIndex)
												if f.Kind() == reflect.Ptr {
													f.Set(last.Addr())
												}
											}
										}
									}
									last = field
								}
							}
							cacheV[names] = &field
							cacheM[names] = mmi
						}
					}
					trefs = trefs[len(mmi.fields.dbcols):]
				}
			}

			if one {
				ind.Set(mind)
			} else {
				if cnt == 0 {
					// you can use a empty & caped container list
					// orm will not replace it
					if ind.Len() != 0 {
						// if container is not empty
						// create a new one
						slice = reflect.New(ind.Type()).Elem()
					}
				}

				if isPtr {
					slice = reflect.Append(slice, mind.Addr())
				} else {
					slice = reflect.Append(slice, mind)
				}
			}
		}
		cnt++
	}

	if one == false {
		if cnt > 0 {
			ind.Set(slice)
		} else {
			// when a result is empty and container is nil
			// to set a empty container
			if ind.IsNil() {
				ind.Set(reflect.MakeSlice(ind.Type(), 0, 0))
			}
		}
	}

	return cnt, nil
}