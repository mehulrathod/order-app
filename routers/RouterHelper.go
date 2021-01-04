package Routers

import (
	"updated_structure/orderapp/models"

	v1con "updated_structure/orderapp/controllers/api/v1"
	v1repo "updated_structure/orderapp/repository/api/v1"
	service "updated_structure/orderapp/services/api/v1"
)

func hUser() *v1con.UserController {
	repo := v1repo.UserRepo{
		DB: models.GetDB(),
	}
	us := service.UserService{
		User:     models.User{},
		UserRepo: repo,
	}
	uc := v1con.UserController{
		UserService: us,
	}
	return &uc
}

func hCategory() *v1con.CategoryController {
	repo := v1repo.CategoryRepo{
		DB: models.GetDB(),
	}
	cs := service.CategoryService{
		Category:     models.Category{},
		CategoryRepo: repo,
	}
	cc := v1con.CategoryController{
		CategoryService: cs,
	}
	return &cc
}

func hMenu() *v1con.MenuController {
	repo := v1repo.MenuRepo{
		DB: models.GetDB(),
	}
	ms := service.MenuService{
		Menu:     models.Menu{},
		MenuRepo: repo,
	}
	mc := v1con.MenuController{
		MenuService: ms,
	}
	return &mc
}

func hOrder() *v1con.OrderController {
	repo := v1repo.OrderRepo{
		DB: models.GetDB(),
	}
	os := service.OrderService{
		Order:     models.Order{},
		OrderRepo: repo,
	}
	oc := v1con.OrderController{
		OrderService: os,
	}
	return &oc
}
