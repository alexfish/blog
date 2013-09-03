// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/robfig/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tPost struct {}
var Post tPost


func (_ tPost) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Post.Index", args).Url
}

func (_ tPost) Show(
		id interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Post.Show", args).Url
}

func (_ tPost) Update(
		post interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "post", post)
	return revel.MainRouter.Reverse("Post.Update", args).Url
}

func (_ tPost) Delete(
		id interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Post.Delete", args).Url
}

func (_ tPost) GetUpdate(
		id interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Post.GetUpdate", args).Url
}

func (_ tPost) GetCreate(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Post.GetCreate", args).Url
}

func (_ tPost) PostCreate(
		post interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "post", post)
	return revel.MainRouter.Reverse("Post.PostCreate", args).Url
}


type tUser struct {}
var User tUser


func (_ tUser) GetLogin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.GetLogin", args).Url
}

func (_ tUser) PostLogin(
		Email string,
		Password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "Email", Email)
	revel.Unbind(args, "Password", Password)
	return revel.MainRouter.Reverse("User.PostLogin", args).Url
}

func (_ tUser) GetLogout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.GetLogout", args).Url
}


