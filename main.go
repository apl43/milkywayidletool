package main

import (
	"github.com/apl43/milkywayidletool/tool"
	"github.com/apl43/milkywayidletool/model/alchemymodel"
)

func main() {
	cfg := alchemymodel.AlchemyTransformationConfig{}
	tool.CalculateTheTransformationOutput(&cfg)
}