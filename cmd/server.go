package cmd

import (
	"os"

	"github.com/issfriends/isspay/internal/app"
	"github.com/issfriends/isspay/internal/delivery"
	"github.com/issfriends/isspay/internal/delivery/bot"
	"github.com/issfriends/isspay/internal/delivery/restful"
	"github.com/issfriends/isspay/internal/repository/database"
	"github.com/issfriends/isspay/pkg/config"
	"github.com/spf13/cobra"
	"github.com/vx416/gox/log"
	"go.uber.org/fx"
)

var (
	// Server run iss pay server command
	Server = &cobra.Command{
		Use:   "server",
		Short: "run isspay http api server",
		Run:   runServer,
	}
	migration bool
	testData  bool
)

func init() {
	Server.Flags().BoolVarP(&migration, "migration", "m", false, "run migration")
	Server.Flags().BoolVarP(&testData, "testdata", "t", false, "create test data")
}

func runServer(cmd *cobra.Command, args []string) {
	defer recoverPanic()

	cfg, err := config.Init()
	if err != nil {
		log.Get().Errorf("server: init config failed, err:%+v", err)
		os.Exit(1)
	}

	if _, err = cfg.Log.Build(); err != nil {
		log.Get().Errorf("server: init logger failed, err:%+v", err)
		os.Exit(1)
	}

	opts := fx.Options(
		cfg.ProvideInfra(),
		fx.Provide(
			database.New,
			app.New,
			restful.New,
			bot.New,
		),
		fx.Invoke(delivery.RunServer),
	)

	if migration {
		opts = fx.Options(
			opts,
			fx.Invoke(runMigrations),
		)
	}

	if testData {
		opts = fx.Options(
			opts,
			fx.Invoke(initTestData),
		)
	}

	app := fx.New(opts)

	runApp(app)
}
