package db

type MarketGroup struct {
	Description     string `json:"description,omitempty" yaml:"description"`
	HasTypes        bool   `json:"hasTypes,omitempty" yaml:"hasTypes"`
	IconID          int    `json:"iconID,omitempty" yaml:"iconID"`
	MarketGroupID   int    `json:"marketGroupID,omitempty" yaml:"marketGroupID"`
	MarketGroupName string `json:"marketGroupName,omitempty" yaml:"marketGroupName"`
	ParentGroupID   int    `json:"parentGroupID,omitempty" yaml:"parentGroupID"`
}

func (tt marketGroups) Breadcrumb(id int) []*MarketGroup {
	var (
		mg         = make([]*MarketGroup, 0)
		rootMarket = &MarketGroup{MarketGroupName: "Market"}
	)
	// if this is root create virtual Top Marketgroup
	// and be done
	if id == 0 {
		mg = append(mg, rootMarket)
		return mg
	}
	// does this marketgroup exist at all?
	// if no return nil so that the server
	// can react with 404 not found
	current, ok := tt[id]
	if !ok{
		return nil
	}
	mg=append(mg, current)
	for {
		// get the parent of the current Marketgroup
		// and make it the current
		current, ok = tt[current.ParentGroupID]
		// if there is no parent we are done and
		// append the virtual Top Marketgroup for the breadcrumb
		if !ok{
			return append(mg, rootMarket)
		}
		mg = append(mg, current)
	}
}

/*type marketGroups map[int]*MarketGroup

func (tt marketGroups) GetByKey(key int) interface{} {
	return tt[key]
}
func (tt marketGroups) SaveToDB() error {
	for k, v := range tt {
		if err := insert(tt.FileName(), k, v); err != nil {
			return err
		}
	}
	return nil
}
func (tt marketGroups) FileName() string {
	return "invMarketGroups"
}
func (tt marketGroups) TableName() string {
	return strings.ToLower(tt.FileName())
}
func (tt marketGroups) New(id int, data []byte) error {
	var (
		err     error
		newItem = new(MarketGroup)
	)
	if err = json.Unmarshal(data, newItem); err != nil {
		return fmt.Errorf("Can't load into Table:%s ID:%v %s", tt.TableName(), id, data)
	}
	tt[id] = newItem
	return nil
}
func loadFromDB(tt mapper)error{
	var (
		rows *sql.Rows
		err  error
		id   int
		data []byte
	)
	rows, err = pg.Query(fmt.Sprintf("SELECT * FROM %s", tt.FileName()))
	for rows.Next() {
		if err = rows.Scan(&id, &data); err != nil {
			return fmt.Errorf("Can't read value from %s: %s", tt.FileName(), err)
		}

		if err = tt.New(id, data); err != nil {
			return err
		}
	}
	return nil
}
func (tt marketGroups) LoadFromDB() error {
	return loadFromDB(tt)
	/*var (
		rows *sql.Rows
		err  error
		id   int
		data []byte
	)
	rows, err = pg.Query(fmt.Sprintf("SELECT * FROM %s", tt.FileName()))
	for rows.Next() {
		if err = rows.Scan(&id, &data); err != nil {
			return fmt.Errorf("Can't read value from %s: %s", tt.FileName(), err)
		}

		if err = tt.New(id, data); err != nil {
			return err
		}
	}
	return nil*/
//}
/*func loadObjectFromYAML(m mapper){
	var (
		err error
		arr = make([]*MarketGroup, 0)
		f *os.File
		path = fmt.Sprintf("%s/%s/%s.%s",basePath,yamlObjectPath,m.FileName(),yamlExt)
	)
}*/
/*func (tt marketGroups) LoadFromYAML() error {
	path := fmt.Sprintf("%s/%s/%s.%s",basePath,yamlObjectPath,tt.FileName(),yamlExt)
	var (
		err error
		arr = make([]*MarketGroup, 0)
		f *os.File
	)
	if f, err = os.OpenFile(path, os.O_RDONLY, 0644); err != nil {
		return err
	}
	defer f.Close()
	dec:= yaml.NewDecoder(f)
	dec.SetStrict(true)
	if err = dec.Decode(&arr); err != nil {
		return err
	}
	for _, v := range arr {
		tt[v.MarketGroupID] = v
	}
	return nil
}
func (tt marketGroups) Breadcrumb(id int) []*MarketGroup {
	var (
		mg         = make([]*MarketGroup, 0)
		rootMarket = &MarketGroup{MarketGroupName: "Market"}
	)
	// if this is root create virtual Top Marketgroup
	// and be done
	if id == 0 {
		mg = append(mg, rootMarket)
		return mg
	}
	// does this marketgroup exist at all?
	// if no return nil so that the server
	// can react with 404 not found
	current, ok := tt[id]
	if !ok{
		return nil
	}
	mg=append(mg, current)
	for {
		// get the parent of the current Marketgroup
		// and make it the current
		current, ok = tt[current.ParentGroupID]
		// if there is no parent we are done and
		// append the virtual Top Marketgroup for the breadcrumb
		if !ok{
			return append(mg, rootMarket)
		}
		mg = append(mg, current)
	}
}
/*func (tt marketGroups) ItemBreadcrumb(id int) []*MarketGroup {
	if
}*/
