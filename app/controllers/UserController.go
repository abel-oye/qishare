package controllers

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"github.com/coocood/qbs"
	"github.com/river/qishare/app/models"
	"github.com/river/qishare/app/routes"
	"github.com/robfig/revel"
	"strconv"
	"strings"
)

type UserController struct {
	Application
}

//登录跳转
func (u *UserController) SignInRedirect() revel.Result {
	return u.RenderTemplate("User/signIn.html")
}

//注册跳转
func (u *UserController) SignUpRedirect() revel.Result {
	return u.RenderTemplate("User/signUp.html")
}

//忘记密码
func (u *UserController) ForgotRedirect() revel.Result {
	return u.RenderTemplate("User/Forgot.html")
}

func (u *UserController) SignIn(email, password, remember string) revel.Result {

	u.Validation.Required(email).Message("请输入邮箱")

	if u.Validation.HasErrors() {
		u.Validation.Keep()
		u.FlashParams()
		return u.Redirect(routes.UserController.SignInRedirect())
	}
	//查询用户
	user := new(models.User)
	condition := qbs.NewCondition("email = ?", email).
		And("password = ?", models.EncryptPassword(password))
	//models.EncryptPassword(password)
	u.q.Condition(condition).Find(user)
	if user.Id == 0 {
		u.Validation.Keep()
		u.FlashParams()
		u.Flash.Out["email"] = email
		u.Flash.Error("邮箱或密码错误！")
		return u.Redirect(routes.UserController.SignInRedirect())
	}

	u.Session["userName"] = user.UserName
	u.Session["userId"] = strconv.Itoa(int(user.Id))
	u.Session["isLogin"] = "true"

	//记住密码
	if remember == "1" {
		fmt.Println(remember)
	}
	preUrl, ok := u.Session["preUrl"]
	if ok {
		return u.Redirect(preUrl)
	}
	return u.Redirect(routes.App.Index())
}

//注册
func (u *UserController) SignUp(user models.User) revel.Result {

	user.Validation(u.q, u.Validation)

	if u.Validation.HasErrors() {
		u.Validation.Keep()
		return u.Redirect(routes.UserController.SignUpRedirect())
	}
	notEncrypted := user.Password
	user.Password = models.EncryptPassword(user.Password)
	user.ValidateCode = strings.Replace(uuid.NewUUID().String(), "-", "", -1)
	if !user.Save(u.q) {
		u.Flash.Error("注册用户失败")
		return u.Redirect(routes.UserController.SignUpRedirect())
	}
	subject := "激活账号 —— 奇享-向世界分享我们"
	content := `这封信是由 奇享 发送的。
				您收到这封邮件，是由于在 奇享 获取了新用户注册地址使用 了这个邮箱地址。如果您并没有访问过 奇享，
				或没有进行上述操作，请忽 略这封邮件。
				您不需要退订或进行其他进一步的操作。
				----------------------------------------------------------------------
				新用户注册说明
				----------------------------------------------------------------------
				如果您是 奇享 的新用户，或在修改您的注册 Email 时使用了本地址，我们需 要对您的地址有效性进行验证以避免垃圾邮件或地址被滥用。
				您只需点击下面的链接即可进行用户注册，
				"http://localhost:9000/user/validate/` + user.ValidateCode + `"
				(如果上面不是链接形式，请将该地址手工粘贴到浏览器地址栏再访问)
				 感谢您的访问，祝您使用愉快！`
	//发送验证邮件
	go sendMail(subject, content, []string{user.Email})
	//注册成功登陆
	return u.SignIn(user.Email, notEncrypted, "")

}

//登出
func (u *UserController) LogOut() revel.Result {
	for k := range u.Session {
		delete(u.Session, k)
	}
	return u.Redirect(routes.App.Index())
}

//忘记密码
func (u *UserController) Forgot(mail string) revel.Result {
	//重新生成验证码
	user := new(models.User)
	user.ValidateCode = strings.Replace(uuid.NewUUID().String(), "-", "", -1)
	user.Email = mail
	_, err := u.q.Update(user)
	if err != nil {
		fmt.Println(err)
	}
	subject := "激活账号 —— 奇享-向世界分享我们"
	content := `这封信是由 奇享 发送的。
				您收到这封邮件，是由于在 奇享 获取了新用户注册地址使用 了这个邮箱地址。如果您并没有访问过 奇享，
				或没有进行上述操作，请忽 略这封邮件。
				您不需要退订或进行其他进一步的操作。
				----------------------------------------------------------------------
				找回密码说明
				----------------------------------------------------------------------
				如果您是 奇享 的老用户，或在修改您的注册 Email 时使用了本地址，我们需 要对您的地址有效性进行验证以避免垃圾邮件或地址被滥用。
				您只需点击下面的链接即可进行用户密码找回，
				"http://localhost:9000/user/forgot/` + user.ValidateCode + `"
				(如果上面不是链接形式，请将该地址手工粘贴到浏览器地址栏再访问)
				 感谢您的访问，祝您使用愉快！`
	//发送验证邮件
	go sendMail(subject, content, []string{user.Email})
	return u.Redirect(routes.UserController.ForgotRedirect())
}

//验证用户
func (c *UserController) Validate(code string) revel.Result {
	user := FindUserByCode(c.q, code)
	if user.Id == 0 {
		return c.NotFound("用户不存在或校验码错误")
	}

	user.IsActive = true
	user.Save(c.q)

	c.Flash.Success("您的账号成功激活，请登录！")

	return c.Redirect(routes.UserController.SignInRedirect())
}

//
func FindUserById(q *qbs.Qbs, id int64) *models.User {
	user := new(models.User)
	q.WhereEqual("id", id).Find(user)
	return user
}

//查询验证码是否存在
func FindUserByCode(q *qbs.Qbs, code string) *models.User {
	user := new(models.User)
	q.WhereEqual("validate_code", code).Find(user)
	return user
}
