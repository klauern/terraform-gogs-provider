package gogs

import (
	gogs "github.com/gogits/go-gogs-client"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,

		Schema: map[string]*schema.Schema{
			"gogs_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"login": {
				Type:     schema.TypeString,
				Required: true,
			},
			"full_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"avatar_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"is_admin": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceUserCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*gogs.Client)
	create := gogs.CreateUserOption{
		Email:      d.Get("email").(string),
		FullName:   d.Get("full_name").(string),
		LoginName:  d.Get("login").(string),
		Password:   d.Get("password").(string),
		SendNotify: false,
		Username:   d.Get("username").(string),
	}

	user, err := client.AdminCreateUser(create)
	if err != nil {
		return errors.WithMessage(err, "unable to create user")
	}

	if d.Get("is_admin").(bool) {
		return resourceUserUpdate(d, m)
	}

	return setUserResourceData(d, user)
}

func resourceUserRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*gogs.Client)
	user, err := client.GetUserInfo(d.Get("username").(string))
	if err != nil {
		return errors.WithMessage(err, "unable to retrieve user "+d.Get("username").(string))
	}

	return setUserResourceData(d, user)
}

func resourceUserUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*gogs.Client)
	isAdmin := d.Get("is_admin").(bool)
	edit := gogs.EditUserOption{
		Admin:     &isAdmin,
		Email:     d.Get("email").(string),
		FullName:  d.Get("full_name").(string),
		LoginName: d.Get("login").(string),
		Password:  d.Get("password").(string),
	}

	err := client.AdminEditUser(d.Get("username").(string), edit)
	if err != nil {
		return errors.WithMessage(err, "unable to edit user "+d.Get("username").(string))
	}
	return nil
}

func resourceUserDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*gogs.Client)
	return client.AdminDeleteUser(d.Get("username").(string))
}

func setUserResourceData(d *schema.ResourceData, u *gogs.User) error {
	if err := d.Set("avatar_url", u.AvatarUrl); err != nil {
		return errors.WithMessage(err, "cannot set avatar URL")
	}
	if err := d.Set("email", u.Email); err != nil {
		return errors.WithMessage(err, "cannot set email")
	}
	if err := d.Set("full_name", u.FullName); err != nil {
		return errors.WithMessage(err, "cannot set full name")
	}
	if err := d.Set("gogs_id", u.ID); err != nil {
		return errors.WithMessage(err, "cannot set id")
	}
	if err := d.Set("username", u.UserName); err != nil {
		return errors.WithMessage(err, "cannot set username")
	}
	return nil
}
