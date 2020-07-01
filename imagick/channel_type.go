// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <MagickWand/MagickWand.h>
*/
import "C"

type ChannelType int

const (
	CHANNEL_UNDEFINED  ChannelType = C.UndefinedChannel
	CHANNEL_RED        ChannelType = C.RedChannel
	CHANNEL_GRAY       ChannelType = C.GrayChannel
	CHANNEL_CYAN       ChannelType = C.CyanChannel
	CHANNEL_GREEN      ChannelType = C.GreenChannel
	CHANNEL_MAGENTA    ChannelType = C.MagentaChannel
	CHANNEL_BLUE       ChannelType = C.BlueChannel
	CHANNEL_YELLOW     ChannelType = C.YellowChannel
	CHANNEL_ALPHA      ChannelType = C.AlphaChannel
	CHANNEL_OPACITY    ChannelType = C.OpacityChannel
	CHANNEL_BLACK      ChannelType = C.BlackChannel
	CHANNEL_INDEX      ChannelType = C.IndexChannel
	CHANNEL_TRUE_ALPHA ChannelType = C.TrueAlphaChannel
	CHANNELS_COMPOSITE ChannelType = C.CompositeChannels
	CHANNELS_ALL       ChannelType = C.AllChannels
	CHANNELS_RGB       ChannelType = C.RGBChannels
	CHANNELS_GRAY      ChannelType = C.GrayChannels
	CHANNELS_SYNC      ChannelType = C.SyncChannels
	CHANNELS_DEFAULT   ChannelType = C.DefaultChannels
)

type PixelChannelType int

const (
	PIXEL_CHANNEL_UNDEFINED     PixelChannelType = C.UndefinedPixelChannel
	PIXEL_CHANNEL_RED           PixelChannelType = C.RedPixelChannel
	PIXEL_CHANNEL_CYAN          PixelChannelType = C.CyanPixelChannel
	PIXEL_CHANNEL_GRAY          PixelChannelType = C.GrayPixelChannel
	PIXEL_CHANNEL_L             PixelChannelType = C.LPixelChannel
	PIXEL_CHANNEL_LABEL         PixelChannelType = C.LabelPixelChannel
	PIXEL_CHANNEL_Y             PixelChannelType = C.YPixelChannel
	PIXEL_CHANNEL_A             PixelChannelType = C.aPixelChannel
	PIXEL_CHANNEL_GREEN         PixelChannelType = C.GreenPixelChannel
	PIXEL_CHANNEL_MAGENTA       PixelChannelType = C.MagentaPixelChannel
	PIXEL_CHANNEL_CB            PixelChannelType = C.CbPixelChannel
	PIXEL_CHANNEL_B             PixelChannelType = C.bPixelChannel
	PIXEL_CHANNEL_BLUE          PixelChannelType = C.BluePixelChannel
	PIXEL_CHANNEL_YELLOW        PixelChannelType = C.YellowPixelChannel
	PIXEL_CHANNEL_CR            PixelChannelType = C.CrPixelChannel
	PIXEL_CHANNEL_BLACK         PixelChannelType = C.BlackPixelChannel
	PIXEL_CHANNEL_ALPHA         PixelChannelType = C.AlphaPixelChannel
	PIXEL_CHANNEL_INDEX         PixelChannelType = C.IndexPixelChannel
	PIXEL_CHANNEL_READMASK      PixelChannelType = C.ReadMaskPixelChannel
	PIXEL_CHANNEL_WRITEMASK     PixelChannelType = C.WriteMaskPixelChannel
	PIXEL_CHANNEL_META          PixelChannelType = C.MetaPixelChannel
	PIXEL_CHANNEL_COMPOSITEMASK PixelChannelType = C.CompositeMaskPixelChannel
)
