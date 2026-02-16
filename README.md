# indicator
The trade indicator.

# Indicator Support
| indicator | support |
|----------|------|
| EMA   | Yes  |
| SMA   | Yes  |
| SMMA  | Yes  |
| Stoch  | No Test |
| StochRSI | Yes |
| ATR | Yes |
| ADX | Yes |

# ATR / ADX Usage
`ATR` and `ADX` are OHLC-based indicators. For accurate results, use `UpdateOHLC(high, low, close)`.

```go
atr := indicator.NewATR(14)
adx := indicator.NewADX(14)

for _, candle := range candles {
	atr.UpdateOHLC(candle.High, candle.Low, candle.Close)
	adx.UpdateOHLC(candle.High, candle.Low, candle.Close)
}

fmt.Println("ATR:", atr.Result())
fmt.Println("ADX:", adx.Result())
fmt.Println("+DI:", adx.PlusDI(), "-DI:", adx.MinusDI())
```

`wilderAverage` is the internal smoothing engine used by ATR/ADX:
- warm-up phase: use SMA to initialize the first smoothed value
- running phase: `next = (prev*(period-1) + current) / period`

Because of warm-up, ATR/ADX may return 0 in early samples before enough data arrives.
Both indicators also keep compatibility with the common `Update(price)` interface (close-only fallback).

# Cheers to
Some indicator refer to [Gekko](https://github.com/thrasher-/gocryptotrader)

Some indicator refer to tradingview wiki [StochRSI](https://www.tradingview.com/wiki/Stochastic_RSI_(STOCH_RSI))
