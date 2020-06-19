// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickCore/MagickCore.h>
*/
import "C"
import (
	"runtime"
	"unsafe"
)

type ChannelStatistics struct {
	Depth             int
	Area              float64
	Minima            float64
	Maxima            float64
	Sum               float64
	SumSquared        float64
	SumCubed          float64
	SumFourthPower    float64
	Mean              float64
	Variance          float64
	StandardDeviation float64
	Kurtosis          float64
	Skewness          float64
	Entropy           float64
}

func newChannelStatistics(ccs *C.ChannelStatistics) *ChannelStatistics {
	cs := &ChannelStatistics{
		Depth:             int(ccs.depth),
		Area:              float64(ccs.area),
		Minima:            float64(ccs.minima),
		Maxima:            float64(ccs.maxima),
		Sum:               float64(ccs.sum),
		SumSquared:        float64(ccs.sum_squared),
		SumCubed:          float64(ccs.sum_cubed),
		SumFourthPower:    float64(ccs.sum_fourth_power),
		Mean:              float64(ccs.mean),
		Variance:          float64(ccs.variance),
		StandardDeviation: float64(ccs.standard_deviation),
		Kurtosis:          float64(ccs.kurtosis),
		Skewness:          float64(ccs.skewness),
		Entropy:           float64(ccs.entropy),
	}
	runtime.KeepAlive(ccs)

	return cs
}

func cChannelStatisticsArrayToSlice(p *C.ChannelStatistics) []ChannelStatistics {
	var css []ChannelStatistics
	q := uintptr(unsafe.Pointer(p))
	for i := 0; i < C.MaxPixelChannels; i++ {
		p = (*C.ChannelStatistics)(unsafe.Pointer(q))
		css = append(css, *newChannelStatistics(p))
		q += unsafe.Sizeof(q)
	}
	return css
}
