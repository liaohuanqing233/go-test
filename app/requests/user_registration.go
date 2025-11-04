package requests

import (
	"goblog/app/models/user"

	"github.com/thedevsaddam/govalidator"
)

// ValidateRegistrationForm 用户注册数据表单校验
func ValidateRegistrationForm(data user.User) map[string][]string {
	rules := govalidator.MapData{
		"name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
		"email":            []string{"required", "email", "min:4", "max:30", "not_exists:users,email"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
	}

	messages := govalidator.MapData{
		"name":             []string{"required:用户名必填", "alpha_num:用户名只允许为英文和数字", "between:用户名长度需在3~20之间"},
		"email":            []string{"required:邮件名必填", "email:邮件格式错误", "min:邮件长度需大于4", "max:邮件长度需小于30"},
		"password":         []string{"required:密码必填", "min:密码长度需大于6"},
		"password_confirm": []string{"required:确认密码必填"},
	}

	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}

	errs := govalidator.New(opts).ValidateStruct()

	if data.Password != data.PasswordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入密码不一致")
	}

	return errs
}
