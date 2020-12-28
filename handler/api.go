package handler

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)


//Todo
type Todo struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
}

 var todos = []Todo{
	{Id: 1, Name:"Animation", Completed: false},
	{Id: 2, Name:"Aws ", Completed: false },
	{Id: 3, Name:"Digital Ocean ", Completed: false},
	{Id: 4, Name:"Linode ", Completed: false},

 }


 func AllTodo(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":  true,
		"todo": todos,
	})
}


 

func CreateTodo(c *fiber.Ctx) error {
	type Request struct {
		Name string `json:"name"`
	}
	var body Request
	err := c.BodyParser(&body)
	if err == nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}
	//creating todo

	todo := Todo{
		Id : len(todos) + 1,
		Name: body.Name,
		Completed: false,
	}
	todos = append(todos, todo)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success":  true,
		"todo": todos,
	})
}



func GetOne(c *fiber.Ctx) error {
	todoId := c.Params("id")
	id, err := strconv.Atoi(todoId)

	if err != nil {
	return 	c.Status( fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id not valid",
		})
	}

	for _, todo := range todos {
		if todo.Id == id {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success":  true,
				"todo": todo,
			})
		}
		return 	c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":"todo Not found",
			})


	}

	return nil
}

func DeleteTodo(c *fiber.Ctx) error{
	paramsId := c.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse id",
		})


	}

	for i, todo := range todos {
		if todo.Id == id {
			todos = append(todos[0:1], todos[i+1:]...)
			c.Status(fiber.StatusNoContent)

		}
		c.Status(fiber.StatusNotFound)
	}
	return nil
}

//func UpdateTodo(c *fiber.Ctx) error  {
//	type Request struct {
//		Name string `json:"name"`
//		Complete bool `json:"complete"`
//	}
//
//	paramsId := c.Params("id")
//	id, err := strconv.Atoi(paramsId)
//
//	if err != nil {
//		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"error": "Cannot parse id",
//		})
//	}
//
//	//
//	var body Request
//	err = c.BodyParser(&body)
//	if err != nil {
//		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"error": "Cannot parse id",
//		})
//	}
//	var todos Todo
//	for _,  t:= range todos {
//		if t.Id == id {
//			todo = t
//			break
//		}
//	}
//
//	if  todo.Id == 0{
//		c.Status(fiber.StatusNotFound)
//	}
//
//	if body.Name != nil {
//		todo.Name = *body.Name
//	}
//
//}