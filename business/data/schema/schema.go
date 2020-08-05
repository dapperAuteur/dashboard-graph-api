// Package schema provides schema support for the database.
package schema

import (
	"context"
	"encoding/json"
	"regexp"
	"strings"
	"time"

	"github.com/ardanlabs/graphql"
	"github.com/pkg/errors"
)

// document represents the schema for the project.
var document = `
enum AllowedTerminacion {
    AR
    ER
	IR
	OR
  }

  enum AllowedTypeAffix {
    PREFIX
    PREFIXIOID
    INFIX
    CIRCUMFIX
    INTERFIX
    DUPLIFIX
    TRANSFIX
    SIMULFIX
    SUPRAFIX
    DISFIX
    STEM
    SUFFIX
    SUFFIXOID
    NA
  }

  enum AllowedTongue {
    ENGLISH
    SPANISH
  }

  enum AllowedCurrency {
    USD
    ETHEREUM
    SPANK
    BOOTY
    BITCOIN
  }

  enum MediaType {
      AUDIO
      IMAGE
      TEXT
      VIDEO
  }

  # Types

  type Activity {
      id: ID!
      endTime: DateTime @search
      startTime: DateTime @search
      media: [Media] @hasInverse(field: activity)
      name: String @search(by: [hash])
      note: [Note] @hasInverse(field: activity)
      person: [Person] @hasInverse(field: activity)
      tag: [Tag] @hasInverse(field: activity)
  }

  type Affix {
    id: ID!
    example: [String] @search(by: [hash])
    meaning: [String] @search(by: [fulltext, hash])
    media: [Media] @hasInverse(field: affix)
    morpheme: String @search(by: [hash])
    note: [Note] @hasInverse(field: affix)
    tongue: AllowedTongue @search(by: [exact])
    type: [AllowedTypeAffix]
  }

  type BlogPost {
    id: ID!
    author: [Person] @hasInverse(field: blogPost)
    body: String @search(by: [fulltext])
    comment: [Comment] @hasInverse(field: blogPost)
    media: [Media] @hasInverse(field: blogPost)
    note: [Note] @hasInverse(field: blogPost)
    publishDate: DateTime @search
    published: Boolean
    tag: [Tag] @hasInverse(field: blogPost)
    title: String! @search(by: [hash])
  }

  type Budget {
    id: ID!
    budgetName: String! @search(by: [hash])
    budgetValue: Float @search
    media: [Media] @hasInverse(field: budget)
    note: [Note] @hasInverse(field: budget)
    partner: [Person] @hasInverse(field: budget)
    tag: [Tag] @hasInverse(field: budget)
    transaction: [Transaction] @hasInverse(field: budget)
  }

  type Comment {
    id: ID!
    author: [Person] @hasInverse(field: comment)
    blogPost: [BlogPost] @hasInverse(field: comment)
    body: String @search(by: [fulltext])
    comment: [Comment] @hasInverse(field: comment)
    media: [Media] @hasInverse(field: comment)
    note: [Note] @hasInverse(field: comment)
    publishDate: DateTime @search
    published: Boolean
    tag: [Tag] @hasInverse(field: comment)
}

  type FinancialAccount {
    id: ID!
    accountName: String! @search(by: [hash])
    currentValue: Float @search
    media: [Media] @hasInverse(field: financialAccount)
    note: [Note] @hasInverse(field: financialAccount)
    owner: [Person] @hasInverse(field: financialAccount)
    tag: [Tag] @hasInverse(field: financialAccount)
    transaction: [Transaction] @hasInverse(field: financialAccount)
  }

  type Game {
    id: ID!
    attempts: Int @search
    bulls: Int @search
    cows: Int @search
    guess: [String]
    note: [Note] @hasInverse(field: game)
    player: [Person] @hasInverse(field: game)
    score: Int @search
    tag: [Tag] @hasInverse(field: game)
    winningWord: String @search(by: [hash])
    won: Boolean
  }

  type Media {
    id: ID!
    activity: [Activity] @hasInverse(field: media)
    affix: [Affix] @hasInverse(field: media)
    blogPost: [BlogPost] @hasInverse(field: media)
    budget: [Budget] @hasInverse(field: media)
    comment: [Comment] @hasInverse(field: media)
    financialAccount: [FinancialAccount] @hasInverse(field: media)
    link: String
    note: [Note] @hasInverse(field: media)
    occurrance: DateTime @search
    person: [Person] @hasInverse(field: media)
    tag: [Tag] @hasInverse(field: media)
    transaction: [Transaction] @hasInverse(field: media)
    type: MediaType
    vendor: [Vendor] @hasInverse(field: media)
    verbo: [Verbo] @hasInverse(field: media)
    word: [Word] @hasInverse(field: media)
}

type Note {
    id: ID!
    activity: Activity @hasInverse(field: note)
    affix: [Affix] @hasInverse(field: note)
    blogPost: BlogPost @hasInverse(field: note)
    budget: Budget @hasInverse(field: note)
    comment: Comment @hasInverse(field: note)
    financialAccount: FinancialAccount @hasInverse(field: note)
    game: Game @hasInverse(field: note)
    media: Media @hasInverse(field: note)
    note: String @search(by: [term])
    person: Person @hasInverse(field: note)
    tag: Tag @hasInverse(field: note)
    transaction: Transaction @hasInverse(field: note)
    vendor: Vendor @hasInverse(field: note)
    verbo: Verbo @hasInverse(field: note)
    word: Word @hasInverse(field: note)
}

type Person {
    id: ID!
    activity: [Activity] @hasInverse(field: person)
    associate: [Person] @hasInverse(field: associate)
    blogPost: [BlogPost] @hasInverse(field: author)
    budget: [Budget] @hasInverse(field: partner)
    comment: [Comment] @hasInverse(field: author)
    email: String @search(by: [hash])
    financialAccount: [FinancialAccount] @hasInverse(field: owner)
    game: [Game] @hasInverse(field: player)
    guess: [String]
    isUser: Boolean
    media: [Media] @hasInverse(field: person)
    nickname: [String] @search(by: [hash])
    note: [Note] @hasInverse(field: person)
    passwordHash: String
    profileImageUrl: Media
    role: [Int] @search
    tag: [Tag] @hasInverse(field: person)
    transaction: [Transaction] @hasInverse(field: participant)
}

  type Tag {
    id: ID!
    activity: [Activity] @hasInverse(field: tag)
    blogPost: [BlogPost] @hasInverse(field: tag)
    budget: [Budget] @hasInverse(field: tag)
    comment: [Comment] @hasInverse(field: tag)
    financialAccount: [FinancialAccount] @hasInverse(field: tag)
    game: [Game] @hasInverse(field: tag)
    media: [Media] @hasInverse(field: tag)
    note: [Note] @hasInverse(field: tag)
    person: [Person] @hasInverse(field: tag)
    tagName: String @search(by: [hash])
    transaction: [Transaction] @hasInverse(field: tag)
    vendor: [Vendor] @hasInverse(field: tag)
  }

  type Transaction {
    id: ID!
    budget: [Budget] @hasInverse(field: transaction)
    currency: AllowedCurrency @search(by: [hash])
    financialAccount: [FinancialAccount] @hasInverse(field: transaction)
    media: [Media] @hasInverse(field: transaction)
    note: [Note] @hasInverse(field: transaction)
    occurrence: DateTime @search
    participant: [Person] @hasInverse(field: transaction)
    tag: [Tag] @hasInverse(field: transaction)
    transactionEvent: String @search(by: [hash])
    transactionValue: Float @search
    vendor: Vendor @hasInverse(field: transaction)
  }

  type Vendor {
    id: ID!
    media: [Media] @hasInverse(field: vendor)
    note: [Note] @hasInverse(field: vendor)
    tag: [Tag] @hasInverse(field: vendor)
    transaction: Transaction @hasInverse(field: vendor)
    vendorName: String @search(by: [hash])
  }

  type Verbo {
    id: ID!
    cambiar_de_irregular: String @search(by: [hash])
    categoria_de_irregular: String @search(by: [hash])
    english: String @search(by: [hash])
    spanish: String @search(by: [hash])
    grupo: Float @search
    irregular: Boolean
    reflexive: Boolean
    media: [Media] @hasInverse(field: verbo)
    note: [Note] @hasInverse(field: verbo)
    terminacion: AllowedTerminacion @search(by: [exact])
  }

  type Word {
    id: ID!
    definition: String @search(by: [fulltext])
    f_points: Int @search
    s_points: Int @search
    in_game: Boolean
    isFourLetterWord: Boolean
    media: [Media] @hasInverse(field: word)
    note: [Note] @hasInverse(field: word)
    tier: Int @search
    tongue: String @search(by: [hash])
    word: String @search(by: [hash])
  }
`

// Schema error variables.
var (
	ErrNoSchemaExists = errors.New("no schema exists")
	ErrInvalidSchema  = errors.New("schema doesn't match")
)

// Schema provides support for schema operations against the database.
type Schema struct {
	graphql  *graphql.GraphQL
	document string
}

// New constructs a Schema value for use to manage the schema.
func New(graphql *graphql.GraphQL) *Schema {
	schema := Schema{
		graphql:  graphql,
		document: document,
	}

	return &schema
}

// DropAll perform an alter operatation against the configured server
// to remove all the data and schema.
func (s *Schema) DropAll(ctx context.Context) error {
	query := strings.NewReader(`{"drop_all": true}`)
	if err := s.graphql.Do(ctx, "alter", query, nil); err != nil {
		return errors.Wrap(err, "dropping schema and data")
	}

	schema, err := s.retrieve(ctx)
	if err != nil {
		return errors.Wrap(err, "can't validate schema, db not ready")
	}

	if err := s.validate(ctx, schema); err != ErrNoSchemaExists {
		return errors.Wrap(err, "unable to drop schema and data")
	}

	return nil
}

// Create is used create the schema in the database.
func (s *Schema) Create(ctx context.Context) error {
	schema, err := s.retrieve(ctx)
	if err != nil {
		return errors.Wrap(err, "can't create schema, db not ready")
	}

	// If the schema matches against what we know the
	// schema to be, don't try to update it.
	if err := s.validate(ctx, schema); err == nil {
		return nil
	}

	query := `mutation updateGQLSchema($schema: String!) {
		updateGQLSchema(input: {
			set: { schema: $schema }
		}) {
			gqlSchema {
				schema
			}
		}
	}`
	vars := map[string]interface{}{"schema": s.document}

	if err := s.graphql.QueryWithVars(ctx, graphql.CmdAdmin, query, vars, nil); err != nil {
		return errors.Wrap(err, "create schema")
	}

	schema, err = s.retrieve(ctx)
	if err != nil {
		return errors.Wrap(err, "can't create schema, db not ready")
	}

	if err := s.validate(ctx, schema); err != nil {
		return errors.Wrap(err, "invalid schema")
	}

	return nil
}

// retrieve queries the database for the schema and handles situations
// when the database is not ready for schema operations.
func (s *Schema) retrieve(ctx context.Context) (string, error) {
	for {
		schema, err := s.query(ctx)
		if err != nil {
			if strings.Contains(err.Error(), "Server not ready") {

				// If the context deadline exceeded then we are done trying.
				if ctx.Err() != nil {
					return "", errors.Wrap(err, "server not ready")
				}

				// We need to wait for the server to be ready for this :(.
				time.Sleep(2 * time.Second)
				continue
			}

			return "", errors.Wrap(err, "server not ready")
		}

		return schema, nil
	}
}

func (s *Schema) query(ctx context.Context) (string, error) {
	query := `query { getGQLSchema { schema }}`
	result := make(map[string]interface{})
	if err := s.graphql.QueryWithVars(ctx, graphql.CmdAdmin, query, nil, &result); err != nil {
		return "", errors.Wrap(err, "query schema")
	}

	data, err := json.Marshal(result)
	if err != nil {
		return "", errors.Wrap(err, "marshal schema")
	}

	return string(data), nil
}

func (s *Schema) validate(ctx context.Context, schema string) error {
	if schema == `{"getGQLSchema":null}` || schema == `{"getGQLSchema":{"schema":""}}` {
		return ErrNoSchemaExists
	}

	if len(schema) < 27 {
		return ErrInvalidSchema
	}

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return errors.Wrap(err, "regex compile")
	}

	exp := strings.ReplaceAll(s.document, "\\n", "")
	exp = reg.ReplaceAllString(exp, "")
	schema = strings.ReplaceAll(schema[27:], "\\n", "")
	schema = strings.ReplaceAll(schema, "\\t", "")
	schema = reg.ReplaceAllString(schema, "")

	if exp != schema {
		return ErrInvalidSchema
	}

	return nil
}
