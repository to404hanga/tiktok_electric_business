package ioc

import (
	"fmt"
	"tiktok_electric_business/user/repository/dao"

	"github.com/spf13/viper"
	prometheus2 "github.com/to404hanga/pkg404/gormx/callbacks/prometheus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
	"gorm.io/plugin/prometheus"
)

func InitDB() *gorm.DB {
	type Config struct {
		DSN string `yaml:"dsn"`
	}
	c := &Config{
		DSN: "root:123456@tcp(localhost:3306)/teb_user",
	}
	err := viper.UnmarshalKey("db", &c)
	if err != nil {
		panic(fmt.Errorf("the initial configuration failed: %v", err))
	}
	db, err := gorm.Open(mysql.Open(c.DSN), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %v", err))
	}
	// 接入 prometheus
	err = db.Use(prometheus.New(prometheus.Config{
		DBName:          "tiktok_electric_business",
		RefreshInterval: 15,
		MetricsCollector: []prometheus.MetricsCollector{
			&prometheus.MySQL{
				VariableNames: []string{"Threads_running"},
			},
		},
	}))
	if err != nil {
		panic(err)
	}
	err = db.Use(tracing.NewPlugin(tracing.WithoutMetrics()))
	if err != nil {
		panic(err)
	}
	prom := prometheus2.Callbacks{
		Namespace:  "to404hanga_lsh",
		Subsystem:  "tiktok_electric_business",
		Name:       "gorm",
		InstanceId: "my-instance-1",
		Help:       "gorm DB 查询",
	}
	err = prom.Initialize(db)
	if err != nil {
		panic(err)
	}
	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}
