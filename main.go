package main

import (
	"fmt"
	"main/compiler"
	"main/visitor"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Post("/compile", func(c *fiber.Ctx) error {

		code := c.FormValue("code")

		fmt.Println(code)

		lexer := compiler.NewTSwiftLexer(antlr.NewInputStream(code))

		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		parser := compiler.NewTSwiftLanguage(stream)
		parser.BuildParseTrees = true

		tree := parser.Program()

		replVisitor := visitor.NewVisitor()
		replVisitor.Visit(tree)

		replVisitor.ScopeTrace.Print()

		return c.SendString("compiled")

	})

	app.Listen(":3000")

}
