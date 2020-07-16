// Package affix provides CRUD access to the database
package affix

import (
	"context"
	"fmt"

	"github.com/ardanlabs/graphql"
	"github.com/pkg/errors"
)

// Set of error variables for CRUD operations.
var (
	ErrNotExits = errors.New("affix does not exist")
	ErrExits    = errors.New("affix exist")
	ErrNotFound = errors.New("affix not found")
)

// Add adds a new affix to the database. If the affix already exits
// this function will fail but the found affix is returned. If the affix is
// being added, the affix with the id from the databse is returned.
func Add(ctx context.Context, gql *graphql.GraphQL, na NewAffix) (Affix, error) {
	a := Affix{
		Example:   na.Example,
		Meaning:   na.Meaning,
		Media:     na.Media,
		Morpheme:  na.Morpheme,
		Note:      na.Note,
		Tongue:    na.Tongue,
		AffixType: na.AffixType,
	}

	a, err := add(ctx, gql, a)
	if err != nil {
		return Affix{}, errors.Wrap(err, "adding affix to database")
	}
	return a, nil
}

// One returns the specified affix from the database by the media id.
func One(ctx context.Context, gql *graphql.GraphQL, affixId string) (Affix, error) {
	query := fmt.Sprintf(`
	query {
		getAffix(id: %q) {
			id
			example
			meaning
			media
			morpheme
			note
			tongue
			affix_type
		}
	}
	`, affixId)

	var result struct {
		GetAffix Affix `json:"getAffix"`
	}
	if err := gql.Query(ctx, query, &result); err != nil {
		return Affix{}, errors.Wrap(err, "query failed")
	}

	if result.GetAffix.ID == "" {
		return Affix{}, ErrNotFound
	}

	return result.GetAffix, nil
}

// OneByMorpheme returns the specified affix from the database by morpheme.
func OneByMorpheme(ctx context.Context, gql *graphql.GraphQL, morpheme string) (Affix, error) {
	query := fmt.Sprintf(`
query {
	queryAffix(filter: {morpheme: {eq: %q}}) {
		id
		example
		meaning
		media
		morpheme
		note
		tongue
		affix_type
	}
}`, morpheme)

	var result struct {
		QueryAffix []Affix `json:"queryAffix"`
	}
	if err := gql.Query(ctx, query, &result); err != nil {
		return Affix{}, errors.Wrap(err, "query failed")
	}

	if len(result.QueryAffix) != 1 {
		return Affix{}, ErrNotFound
	}

	return result.QueryAffix[0], nil
}

// ==================================================

func add(ctx context.Context, gql *graphql.GraphQL, affix Affix) (Affix, error) {
	mutation, result := prepareAdd(affix)
	if err := gql.Query(ctx, mutation, &result); err != nil {
		return Affix{}, errors.Wrap(err, "failed to add affix")
	}

	if len(result.AddAffix.Affix) != 1 {
		return Affix{}, errors.New("affix id not returned")
	}

	affix.ID = result.AddAffix.Affix[0].ID
	return affix, nil
}

func prepareAdd(affix Affix) (string, addResult) {
	var result addResult
	mutation := fmt.Sprintf(`
mutation {
	addAffix(input: [{
		id: %q
		example: %q
		meaning: %q
		media: %q
		morpheme: %q
		note: %q
		tongue: %q
		affix_type: %q
	}])
	%s
}`, affix.example, affix.meaning, affix.media, affix.morpheme, affix.note, affix.tongue, affix.affix_type, result.document())

	return mutation, result
}
