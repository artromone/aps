package main

import (
    "math"
    "math/rand"
)

type PoissonGenerator struct {
    lambda float64  // интенсивность (среднее число событий в единицу времени)
}

func NewPoissonGenerator(lambda float64) *PoissonGenerator {
    return &PoissonGenerator{
        lambda: lambda,
    }
}

// Генерирует количество событий за заданный интервал времени
func (pg *PoissonGenerator) GetEventsCountForInterval(intervalSeconds float64) int {
    L := pg.lambda * intervalSeconds
    p := math.Exp(-L)
    X := 0
    s := p
    u := rand.Float64()
    
    for s < u {
        X++
        p = p * L / float64(X)
        s += p
    }
    
    return X
}
