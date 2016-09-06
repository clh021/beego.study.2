package tusd_test

import (
	"github.com/tus/tusd"
	"github.com/tus/tusd/consullocker"
	"github.com/tus/tusd/filestore"
	"github.com/tus/tusd/limitedstore"
)

func ExampleNewStoreComposer() {
	composer := tusd.NewStoreComposer()

	fs := filestore.New("./data")
	fs.UseIn(composer)

	cl := consullocker.New(nil)
	cl.UseIn(composer)

	ls := limitedstore.New(1024*1024*1024, composer.Core, composer.Terminater)
	ls.UseIn(composer)

	config := tusd.Config{
		StoreComposer: composer,
	}

	_, _ = tusd.NewHandler(config)
}
