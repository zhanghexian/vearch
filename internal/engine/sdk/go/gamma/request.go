/**
 * Copyright 2019 The Vearch Authors.
 *
 * This source code is licensed under the Apache License, Version 2.0 license
 * found in the LICENSE file in the root directory of this source tree.
 */

package gamma

import (
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/vearch/vearch/v3/internal/engine/idl/fbs-gen/go/gamma_api"
	"github.com/vearch/vearch/v3/internal/proto/vearchpb"
)

type DistanceMetricType uint8

type TermFilter struct {
	Field   string
	Value   []byte
	IsUnion int32
}

type RangeFilter struct {
	Field        string
	LowerValue   []byte
	UpperValue   []byte
	IncludeLower bool
	IncludeUpper bool
}

type VectorQuery struct {
	Name      string
	Value     []byte
	MinScore  float64
	MaxScore  float64
	IndexType string
}

type Request struct {
	ReqNum           int32
	TopN             int32
	BruteForceSearch int32
	VecFields        []VectorQuery
	Fields           []string
	RangeFilters     []RangeFilter
	TermFilters      []TermFilter
	IndexParams      string
	MultiVectorRank  int32
	L2Sqrt           bool
	Ranker           string
}

func SearchRequestSerialize(request *vearchpb.SearchRequest) []byte {
	builder := flatbuffers.NewBuilder(0)
	indexParams := builder.CreateString(request.IndexParams)

	var fields, vectorQuerys, rangeFilters, termFilters []flatbuffers.UOffsetT
	fields = make([]flatbuffers.UOffsetT, len(request.Fields))
	vectorQuerys = make([]flatbuffers.UOffsetT, len(request.VecFields))
	rangeFilters = make([]flatbuffers.UOffsetT, len(request.RangeFilters))
	termFilters = make([]flatbuffers.UOffsetT, len(request.TermFilters))

	for i := 0; i < len(request.Fields); i++ {
		fields[i] = builder.CreateString(request.Fields[i])
	}

	for i := 0; i < len(request.VecFields); i++ {
		name := builder.CreateString(request.VecFields[i].Name)
		index_type := builder.CreateString(request.VecFields[i].IndexType)
		gamma_api.VectorQueryStartValueVector(builder, len(request.VecFields[i].Value))
		for j := len(request.VecFields[i].Value) - 1; j >= 0; j-- {
			builder.PrependByte(request.VecFields[i].Value[j])
		}
		value := builder.EndVector(len(request.VecFields[i].Value))
		gamma_api.VectorQueryStart(builder)
		gamma_api.VectorQueryAddName(builder, name)
		gamma_api.VectorQueryAddValue(builder, value)
		gamma_api.VectorQueryAddMinScore(builder, request.VecFields[i].MinScore)
		gamma_api.VectorQueryAddMaxScore(builder, request.VecFields[i].MaxScore)
		gamma_api.VectorQueryAddIndexType(builder, index_type)
		vectorQuerys[i] = gamma_api.VectorQueryEnd(builder)
	}

	for i := 0; i < len(request.RangeFilters); i++ {
		field := builder.CreateString(request.RangeFilters[i].Field)
		gamma_api.RangeFilterStartLowerValueVector(builder, len(request.RangeFilters[i].LowerValue))
		for j := len(request.RangeFilters[i].LowerValue) - 1; j >= 0; j-- {
			builder.PrependByte(request.RangeFilters[i].LowerValue[j])
		}
		lowerValue := builder.EndVector(len(request.RangeFilters[i].LowerValue))

		gamma_api.RangeFilterStartUpperValueVector(builder, len(request.RangeFilters[i].UpperValue))
		for j := len(request.RangeFilters[i].UpperValue) - 1; j >= 0; j-- {
			builder.PrependByte(request.RangeFilters[i].UpperValue[j])
		}
		upperValue := builder.EndVector(len(request.RangeFilters[i].UpperValue))

		gamma_api.RangeFilterStart(builder)
		gamma_api.RangeFilterAddField(builder, field)
		gamma_api.RangeFilterAddLowerValue(builder, lowerValue)
		gamma_api.RangeFilterAddUpperValue(builder, upperValue)
		gamma_api.RangeFilterAddIncludeLower(builder, request.RangeFilters[i].IncludeLower)
		gamma_api.RangeFilterAddIncludeUpper(builder, request.RangeFilters[i].IncludeUpper)
		rangeFilters[i] = gamma_api.RangeFilterEnd(builder)
	}

	for i := 0; i < len(request.TermFilters); i++ {
		field := builder.CreateString(request.TermFilters[i].Field)
		gamma_api.TermFilterStartValueVector(builder, len(request.TermFilters[i].Value))
		for j := len(request.TermFilters[i].Value) - 1; j >= 0; j-- {
			builder.PrependByte(request.TermFilters[i].Value[j])
		}
		value := builder.EndVector(len(request.TermFilters[i].Value))
		gamma_api.TermFilterStart(builder)
		gamma_api.TermFilterAddField(builder, field)
		gamma_api.TermFilterAddValue(builder, value)
		gamma_api.TermFilterAddIsUnion(builder, request.TermFilters[i].IsUnion)
		termFilters[i] = gamma_api.TermFilterEnd(builder)
	}

	gamma_api.RequestStartFieldsVector(builder, len(request.Fields))
	for i := 0; i < len(request.Fields); i++ {
		builder.PrependUOffsetT(fields[i])
	}
	f := builder.EndVector(len(request.Fields))

	gamma_api.RequestStartVecFieldsVector(builder, len(request.VecFields))
	for i := 0; i < len(request.VecFields); i++ {
		builder.PrependUOffsetT(vectorQuerys[i])
	}
	v := builder.EndVector(len(request.VecFields))

	gamma_api.RequestStartRangeFiltersVector(builder, len(request.RangeFilters))
	for i := 0; i < len(request.RangeFilters); i++ {
		builder.PrependUOffsetT(rangeFilters[i])
	}
	r := builder.EndVector(len(request.RangeFilters))

	gamma_api.RequestStartTermFiltersVector(builder, len(request.TermFilters))
	for i := 0; i < len(request.TermFilters); i++ {
		builder.PrependUOffsetT(termFilters[i])
	}
	t := builder.EndVector(len(request.TermFilters))

	ranker := builder.CreateString(request.Ranker)

	gamma_api.RequestStart(builder)
	gamma_api.RequestAddReqNum(builder, request.ReqNum)
	gamma_api.RequestAddTopn(builder, request.TopN)
	gamma_api.RequestAddBruteForceSearch(builder, request.IsBruteSearch)
	gamma_api.RequestAddFields(builder, f)
	gamma_api.RequestAddVecFields(builder, v)
	gamma_api.RequestAddRangeFilters(builder, r)
	gamma_api.RequestAddTermFilters(builder, t)
	gamma_api.RequestAddIndexParams(builder, indexParams)
	gamma_api.RequestAddMultiVectorRank(builder, request.MultiVectorRank)
	gamma_api.RequestAddL2Sqrt(builder, request.L2Sqrt)
	gamma_api.RequestAddRanker(builder, ranker)
	gamma_api.RequestAddTrace(builder, request.Trace)

	builder.Finish(builder.EndObject())
	return builder.FinishedBytes()
}

func QueryRequestSerialize(request *vearchpb.QueryRequest) []byte {
	builder := flatbuffers.NewBuilder(0)
	var fields, vectorQuerys, rangeFilters, termFilters []flatbuffers.UOffsetT
	vectorQuerys = make([]flatbuffers.UOffsetT, 0)
	fields = make([]flatbuffers.UOffsetT, len(request.Fields))
	rangeFilters = make([]flatbuffers.UOffsetT, len(request.RangeFilters))
	termFilters = make([]flatbuffers.UOffsetT, len(request.TermFilters))

	for i := 0; i < len(request.Fields); i++ {
		fields[i] = builder.CreateString(request.Fields[i])
	}

	for i := 0; i < len(request.RangeFilters); i++ {
		field := builder.CreateString(request.RangeFilters[i].Field)
		gamma_api.RangeFilterStartLowerValueVector(builder, len(request.RangeFilters[i].LowerValue))
		for j := len(request.RangeFilters[i].LowerValue) - 1; j >= 0; j-- {
			builder.PrependByte(request.RangeFilters[i].LowerValue[j])
		}
		lowerValue := builder.EndVector(len(request.RangeFilters[i].LowerValue))

		gamma_api.RangeFilterStartUpperValueVector(builder, len(request.RangeFilters[i].UpperValue))
		for j := len(request.RangeFilters[i].UpperValue) - 1; j >= 0; j-- {
			builder.PrependByte(request.RangeFilters[i].UpperValue[j])
		}
		upperValue := builder.EndVector(len(request.RangeFilters[i].UpperValue))

		gamma_api.RangeFilterStart(builder)
		gamma_api.RangeFilterAddField(builder, field)
		gamma_api.RangeFilterAddLowerValue(builder, lowerValue)
		gamma_api.RangeFilterAddUpperValue(builder, upperValue)
		gamma_api.RangeFilterAddIncludeLower(builder, request.RangeFilters[i].IncludeLower)
		gamma_api.RangeFilterAddIncludeUpper(builder, request.RangeFilters[i].IncludeUpper)
		rangeFilters[i] = gamma_api.RangeFilterEnd(builder)
	}

	for i := 0; i < len(request.TermFilters); i++ {
		field := builder.CreateString(request.TermFilters[i].Field)
		gamma_api.TermFilterStartValueVector(builder, len(request.TermFilters[i].Value))
		for j := len(request.TermFilters[i].Value) - 1; j >= 0; j-- {
			builder.PrependByte(request.TermFilters[i].Value[j])
		}
		value := builder.EndVector(len(request.TermFilters[i].Value))
		gamma_api.TermFilterStart(builder)
		gamma_api.TermFilterAddField(builder, field)
		gamma_api.TermFilterAddValue(builder, value)
		gamma_api.TermFilterAddIsUnion(builder, request.TermFilters[i].IsUnion)
		termFilters[i] = gamma_api.TermFilterEnd(builder)
	}
	gamma_api.RequestStartFieldsVector(builder, len(request.Fields))
	for i := 0; i < len(request.Fields); i++ {
		builder.PrependUOffsetT(fields[i])
	}
	f := builder.EndVector(len(request.Fields))

	gamma_api.RequestStartVecFieldsVector(builder, len(vectorQuerys))
	for i := 0; i < len(vectorQuerys); i++ {
		builder.PrependUOffsetT(vectorQuerys[i])
	}
	v := builder.EndVector(len(vectorQuerys))

	gamma_api.RequestStartRangeFiltersVector(builder, len(request.RangeFilters))
	for i := 0; i < len(request.RangeFilters); i++ {
		builder.PrependUOffsetT(rangeFilters[i])
	}
	r := builder.EndVector(len(request.RangeFilters))

	gamma_api.RequestStartTermFiltersVector(builder, len(request.TermFilters))
	for i := 0; i < len(request.TermFilters); i++ {
		builder.PrependUOffsetT(termFilters[i])
	}
	t := builder.EndVector(len(request.TermFilters))

	indexParams := builder.CreateString("")
	ranker := builder.CreateString("")

	gamma_api.RequestStart(builder)
	gamma_api.RequestAddReqNum(builder, 1)
	gamma_api.RequestAddTopn(builder, request.Limit)
	gamma_api.RequestAddBruteForceSearch(builder, 0)
	gamma_api.RequestAddFields(builder, f)
	gamma_api.RequestAddVecFields(builder, v)
	gamma_api.RequestAddRangeFilters(builder, r)
	gamma_api.RequestAddTermFilters(builder, t)
	gamma_api.RequestAddIndexParams(builder, indexParams)
	gamma_api.RequestAddMultiVectorRank(builder, 0)
	gamma_api.RequestAddL2Sqrt(builder, false)
	gamma_api.RequestAddRanker(builder, ranker)
	gamma_api.RequestAddTrace(builder, request.Trace)

	builder.Finish(builder.EndObject())
	return builder.FinishedBytes()
}
