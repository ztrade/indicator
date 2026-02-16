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


# Quick Start: OHLC Indicators

You can quickly create OHLC-based indicators (such as ATR, ADX) using `NewOHLCIndicator`:

```go
atr, err := indicator.NewOHLCIndicator("ATR", 14)
adx, err := indicator.NewOHLCIndicator("ADX", 14)
if err != nil {
    panic(err)
}
for _, candle := range candles {
    atr.UpdateOHLC(candle.Open, candle.High, candle.Low, candle.Close)
    adx.UpdateOHLC(candle.Open, candle.High, candle.Low, candle.Close)
}
fmt.Println("ATR:", atr.Result())
fmt.Println("ADX:", adx.Result())
fmt.Println("+DI:", adx.PlusDI(), "-DI:", adx.MinusDI())
```

Currently supported: ATR, ADX. More OHLC indicators can be added easily.

**Note:**
- For accurate results, always use `UpdateOHLC(open, high, low, close)`.
- `wilderAverage` is the internal smoothing engine for ATR/ADX:
    - Warm-up: uses SMA to initialize the first smoothed value
    - Running: `next = (prev*(period-1) + current) / period`
- During warm-up, ATR/ADX may return 0 until enough data arrives.
- Both indicators also support the simple `Update(price)` interface (close-only fallback).

# Cheers to
Some indicator refer to [Gekko](https://github.com/thrasher-/gocryptotrader)

Some indicator refer to tradingview wiki [StochRSI](https://www.tradingview.com/wiki/Stochastic_RSI_(STOCH_RSI))
