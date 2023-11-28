package v1alpha1

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/golang/protobuf/jsonpb"
	"github.com/stretchr/testify/assert"

	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
)

func TestMarshalUnMarshal_BranchTask(t *testing.T) {
	r, err := ioutil.ReadFile("testdata/branch.json")
	assert.NoError(t, err)
	o := NodeSpec{}
	err = json.Unmarshal(r, &o)
	assert.NoError(t, err)
	assert.NotNil(t, o.BranchNode.If)
	assert.Equal(t, core.ComparisonExpression_GT, o.BranchNode.If.Condition.BooleanExpression.GetComparison().Operator)
	assert.Equal(t, 1, len(o.InputBindings))
	raw, err := json.Marshal(o)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, raw)
	}
}

// TestBranchNodeSpecMethods tests the methods of the BranchNodeSpec struct.
func TestErrorMarshalAndUnmarshalJSON(t *testing.T) {
	coreError := &core.Error{
		FailedNodeId: "TestNode",
		Message:      "Test error message",
	}

	err := Error{Error: coreError}
	data, jErr := err.MarshalJSON()
	assert.Nil(t, jErr)

	// Unmarshalling the JSON back to a new core.Error struct
	var newCoreError core.Error
	uErr := jsonpb.Unmarshal(bytes.NewReader(data), &newCoreError)
	assert.Nil(t, uErr)
	assert.Equal(t, coreError.Message, newCoreError.Message)
	assert.Equal(t, coreError.FailedNodeId, newCoreError.FailedNodeId)
}

func TestBranchNodeSpecMethods(t *testing.T) {
	// Creating a core.BooleanExpression instance for testing
	boolExpr := &core.BooleanExpression{}

	// Creating an Error instance for testing
	errorMessage := &core.Error{
		Message: "Test error",
	}

	ifNode := NodeID("ifNode")
	elifNode := NodeID("elifNode")
	elseNode := NodeID("elseNode")

	// Creating a BranchNodeSpec instance for testing
	branchNodeSpec := BranchNodeSpec{
		If: IfBlock{
			Condition: BooleanExpression{
				BooleanExpression: boolExpr,
			},
			ThenNode: &ifNode,
		},
		ElseIf: []*IfBlock{
			{
				Condition: BooleanExpression{
					BooleanExpression: boolExpr,
				},
				ThenNode: &elifNode,
			},
		},
		Else:     &elseNode,
		ElseFail: &Error{Error: errorMessage},
	}

	assert.Equal(t, boolExpr, branchNodeSpec.If.GetCondition())

	assert.Equal(t, &ifNode, branchNodeSpec.If.GetThenNode())

	assert.Equal(t, &branchNodeSpec.If, branchNodeSpec.GetIf())

	assert.Equal(t, &elseNode, branchNodeSpec.GetElse())

	elifs := branchNodeSpec.GetElseIf()
	assert.Equal(t, 1, len(elifs))
	assert.Equal(t, boolExpr, elifs[0].GetCondition())
	assert.Equal(t, &elifNode, elifs[0].GetThenNode())

	assert.Equal(t, errorMessage, branchNodeSpec.GetElseFail())

	branchNodeSpec.ElseFail = nil
	assert.Nil(t, branchNodeSpec.GetElseFail())
}
