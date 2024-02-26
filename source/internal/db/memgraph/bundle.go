package memgraph

import (
	"context"
	"fmt"
	"strings"

	"github.com/LearningMotors/platform/service/platenv"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"

	"github.com/vvijayan-suki/memgraph-notes/source/internal/db/constant/label"
	"github.com/vvijayan-suki/memgraph-notes/source/internal/db/constant/property"
)

var driver neo4j.DriverWithContext

type Bundle struct{}

func New() *Bundle {
	return &Bundle{}
}

func (b *Bundle) Initialize() {
	ctx := context.Background()

	initialize(ctx)

	createIndex(ctx)
}

func initialize(ctx context.Context) {
	// ENV variables
	const (
		envMemgraphScheme  = "MEMGRAPH_SCHEME"
		envMemgraphAddress = "MEMGRAPH_ADDR"
	)

	// Default values
	const (
		defaultScheme  = "bolt"
		defaultAddress = "localhost:7687"
	)

	var err error

	scheme := platenv.GetEnvWithDefaultAsString(envMemgraphScheme, defaultScheme)
	address := platenv.GetEnvWithDefaultAsString(envMemgraphAddress, defaultAddress)

	uri := fmt.Sprintf("%s://%s", scheme, address)

	driver, err = neo4j.NewDriverWithContext(uri, neo4j.NoAuth())
	if err != nil {
		panic(err)
	}
	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	}
}

func createIndex(ctx context.Context) {
	// cannot use execute query for creating indexes https://memgraph.com/docs/client-libraries/go#automatic-transaction-management
	// using auto-commit transactions instead as explained by the link
	session := driver.NewSession(ctx, neo4j.SessionConfig{
		BookmarkManager: driver.ExecuteQueryBookmarkManager(),
	})

	defer func() {
		err := session.Close(ctx)
		if err != nil {
			panic(err)
		}
	}()

	const (
		CreateIndexOn = "CREATE INDEX ON"
	)

	type index struct {
		labels     []string
		properties []string
	}

	indices := []index{
		{
			[]string{label.Note},
			[]string{property.SecondaryID},
		},
		{
			[]string{label.Note},
			[]string{property.ID},
		},
		{
			[]string{label.Note},
			[]string{},
		},
		{
			[]string{label.NMSEvent},
			[]string{property.SecondaryID},
		},
		{
			[]string{label.Section},
			[]string{property.SecondaryID},
		},
		{
			[]string{label.Section},
			[]string{property.ID},
		},
		{
			[]string{label.Metadata},
			[]string{property.SecondaryID},
		},
		{
			[]string{label.MetadataEntry},
			[]string{property.SecondaryID},
		},
		{
			[]string{label.SectionContent},
			[]string{property.SecondaryID},
		},
		{
			[]string{label.MacroContent},
			[]string{property.SecondaryID},
		},
		{
			[]string{label.VersionedComposition},
			[]string{property.SecondaryID},
		},
		{
			[]string{label.SectionS2Entry},
			[]string{property.SecondaryID},
		},
	}

	for _, index := range indices {
		labels := fmt.Sprintf(":%s", strings.Join(index.labels, ":"))
		properties := ""
		if len(index.properties) > 0 {
			properties = fmt.Sprintf("(%s)", strings.Join(index.properties, ","))
		}

		query := fmt.Sprintf("%s %s%s", CreateIndexOn, labels, properties)

		_, err := session.Run(ctx, query, nil)
		if err != nil {
			panic(err)
		}
	}
}

func (b *Bundle) Run() {}

func (b *Bundle) Shutdown() {
	ctx := context.Background()

	err := driver.Close(ctx)
	if err != nil {
		panic(err)
	}
}
