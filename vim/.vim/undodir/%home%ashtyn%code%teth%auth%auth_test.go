Vim�UnDo� Fa*�Kd�D����)��?j������E1��͎   J                                  `W�t    _�                             ����                                                                                                                                                                                                                                                                                                                                                  V        `W�,     �                 Generated TestDBCreds_Init   Generated TestNetwork_Init   Generated TestAuth5�_�                            ����                                                                                                                                                                                                                                                                                                                                                  V        `W�-    �               J   package auth       import (   	"database/sql"   
	"reflect"   
	"testing"       ,	"github.com/ethereum/go-ethereum/ethclient"   	_ "github.com/lib/pq"   )       %func TestDBCreds_Init(t *testing.T) {   	type fields struct {   		Host     string   		Password string   		Port     string   		Database string   		User     string   	}   	tests := []struct {   		name   string   		fields fields   		want   *sql.DB   	}{   		// TODO: Add test cases.   	}   	for _, tt := range tests {   %		t.Run(tt.name, func(t *testing.T) {   			db := DBCreds{   				Host:     tt.fields.Host,   !				Password: tt.fields.Password,   				Port:     tt.fields.Port,   !				Database: tt.fields.Database,   				User:     tt.fields.User,   			}   :			if got := db.Init(); !reflect.DeepEqual(got, tt.want) {   :				t.Errorf("DBCreds.Init() = %v, want %v", got, tt.want)   			}   		})   	}   }       %func TestNetwork_Init(t *testing.T) {   	tests := []struct {   		name string   		n    Network   		want *ethclient.Client   	}{   		// TODO: Add test cases.   	}   	for _, tt := range tests {   %		t.Run(tt.name, func(t *testing.T) {   <			if got := tt.n.Init(); !reflect.DeepEqual(got, tt.want) {   :				t.Errorf("Network.Init() = %v, want %v", got, tt.want)   			}   		})   	}   }       func TestAuth(t *testing.T) {   	tests := []struct {   		name string   		want Creds   	}{   		// TODO: Add test cases.   	}   	for _, tt := range tests {   %		t.Run(tt.name, func(t *testing.T) {   7			if got := Auth(); !reflect.DeepEqual(got, tt.want) {   2				t.Errorf("Auth() = %v, want %v", got, tt.want)   			}   		})   	}   }5�_�                    1       ����                                                                                                                                                                                                                                                                                                                                                             `W�f     �   1   3   J    5�_�                    2        ����                                                                                                                                                                                                                                                                                                                                                             `W�h     �   1   2           5�_�                    1       ����                                                                                                                                                                                                                                                                                                                                                             `W�m     �   0   3   J      		// TODO: Add test cases.5�_�                    2       ����                                                                                                                                                                                                                                                                                                                                                             `W�r     �   1   2          			if test5�_�                     2       ����                                                                                                                                                                                                                                                                                                                                                             `W�s    �               J   package auth       import (   	"database/sql"   
	"reflect"   
	"testing"       ,	"github.com/ethereum/go-ethereum/ethclient"   	_ "github.com/lib/pq"   )       %func TestDBCreds_Init(t *testing.T) {   	type fields struct {   		Host     string   		Password string   		Port     string   		Database string   		User     string   	}   	tests := []struct {   		name   string   		fields fields   		want   *sql.DB   	}{   		// TODO: Add test cases.   	}   	for _, tt := range tests {   %		t.Run(tt.name, func(t *testing.T) {   			db := DBCreds{   				Host:     tt.fields.Host,   !				Password: tt.fields.Password,   				Port:     tt.fields.Port,   !				Database: tt.fields.Database,   				User:     tt.fields.User,   			}   :			if got := db.Init(); !reflect.DeepEqual(got, tt.want) {   :				t.Errorf("DBCreds.Init() = %v, want %v", got, tt.want)   			}   		})   	}   }       %func TestNetwork_Init(t *testing.T) {   	tests := []struct {   		name string   		n    Network   		want *ethclient.Client   	}{   		// TODO: Add test cases.   	}   	for _, tt := range tests {   %		t.Run(tt.name, func(t *testing.T) {   <			if got := tt.n.Init(); !reflect.DeepEqual(got, tt.want) {   :				t.Errorf("Network.Init() = %v, want %v", got, tt.want)   			}   		})   	}   }       func TestAuth(t *testing.T) {   	tests := []struct {   		name string   		want Creds   	}{   		// TODO: Add test cases.   	}   	for _, tt := range tests {   %		t.Run(tt.name, func(t *testing.T) {   7			if got := Auth(); !reflect.DeepEqual(got, tt.want) {   2				t.Errorf("Auth() = %v, want %v", got, tt.want)   			}   		})   	}   }5��