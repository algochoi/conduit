// Package filterprocessor docs
package filterprocessor

//go:generate go run ../../../../cmd/conduit-docs/main.go ../../../../conduit-docs/
//go:generate go run ../../../../cmd/readme_config_includer/generator.go

import (
	"github.com/algorand/conduit/conduit/plugins/processors/filterprocessor/expression"
)

//Name: conduit_processors_filter

// SubConfig is the configuration needed for each additional filter
type SubConfig struct {
	/* <code>tag</code> is the tag of the struct to analyze.<br/>
	It can be of the form `txn.*` where the specific ending is determined by the field you wish to filter on.<br/>
	It can also be a field in the ApplyData.
	*/
	FilterTag string `yaml:"tag"`
	/* <code>expression-type</code> is the type of comparison applied between the field, identified by the tag, and the expression.<br/>
	<ul>
		<li>exact</li>
		<li>regex</li>
		<li>less-than</li>
		<li>less-than-equal</li>
		<li>greater-than</li>
		<li>great-than-equal</li>
		<li>equal</li>
		<li>not-equal</li>
	</ul>
	*/
	ExpressionType expression.Type `yaml:"expression-type"`
	// <code>expression</code> is the user-supplied part of the search or comparison.
	Expression string `yaml:"expression"`
}

// Config configuration for the filter processor
type Config struct {
	// <code>search-inner</code> configures the filter processor to recursively search inner transactions for expressions.
	SearchInner bool `yaml:"search-inner"`
	// <code>omit-group-transactions</code> configures the filter processor to return the matched transaction without its grouped transactions.
	OmitGroupTransactions bool `yaml:"omit-group-transactions"`
	/* <code>filters</code> are a list of SubConfig objects with an operation acting as the string key in the map

	filters:
		- [any,all,none]:
			expression: ""
			expression-type: ""
			tag: ""
	*/
	Filters []map[string][]SubConfig `yaml:"filters"`
}
