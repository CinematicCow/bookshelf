package auth

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning : env file load error", err)
	}
}

func Auth() {
	apiBasePath := "/auth"
	websiteBasePath := "/auth"

	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: os.Getenv("SUPERTOKEN_HOST"),
			APIKey:        os.Getenv("SUPERTOKEN_API_KEY"),
		},
		AppInfo: supertokens.AppInfo{
			AppName:         "bookshelf",
			APIDomain:       "http://localhost:4000",
			WebsiteDomain:   "http://localhost:3000",
			APIBasePath:     &apiBasePath,
			WebsiteBasePath: &websiteBasePath,
		},
		RecipeList: []supertokens.Recipe{
			emailpassword.Init(nil),
			session.Init(nil),
		},
	})

	if err != nil {
		panic("Supertokens Auth Package error: " + err.Error())
	}
}
