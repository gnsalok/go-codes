package myapi

import "context"

func Greeter(ctx context.Context, api API, id string) (string, error) {
	u, err := api.GetUser(ctx, id)
	if err != nil {
		return "", err
	}
	return "Hello, " + u.Name, nil
}
