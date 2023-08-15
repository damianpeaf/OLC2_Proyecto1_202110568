package main

import (
	"fmt"
	"main/compiler"
	"main/repl"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Post("/compile", func(c *fiber.Ctx) error {

		code := c.FormValue("code")

		fmt.Println(code)

		lexer := compiler.NewTSwiftLexer(antlr.NewInputStream(code))

		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		parser := compiler.NewTSwiftLanguage(stream)
		parser.BuildParseTrees = true

		tree := parser.Program()

		dclVisitor := repl.NewDclVisitor()
		dclVisitor.Visit(tree)

		replVisitor := repl.NewVisitor(dclVisitor)

		replVisitor.Visit(tree)

		replVisitor.ScopeTrace.Print()
		replVisitor.Console.Show()

		//
		return c.JSON(struct {
			Errors []repl.Error `json:"errors"`
			Output string       `json:"output"`
		}{
			Errors: replVisitor.ErrorTable.Errors,
			Output: replVisitor.Console.GetOutput(),
			// TODO: CST report
			// TODO: scope trace report
		})

	})

	app.Listen(":3000")

}
