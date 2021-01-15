package models

//"github.com/Qesy/QesyGo"

// func SysGetConf() (map[string]string, error) {
// 	key := lib.Sys_hmset_key()
// 	rs, err := lib.RedisCr.HGetAll(key)
// 	if err != nil || len(rs) == 0 {
// 		return SysCacheConf()
// 	}
// 	return rs, err
// }

// func SysCacheConf() (map[string]string, error) {
// 	key := lib.Sys_hmset_key()
// 	var m QesyDb.Model
// 	m.Table = "sys_conf"
// 	rs, err := m.ExecSelectOne()
// 	if err != nil {
// 		return nil, err
// 	}
// 	lib.RedisCr.HMset(key, rs)
// 	return rs, nil
// }
