package routers

import (
	"sync"

	"github.com/arnavmahajan630/cfcli/api"
	"github.com/arnavmahajan630/cfcli/api/controller"
	"github.com/arnavmahajan630/cfcli/api/models"
)


func HandleProfileCommand(username string) error{
	var profile models.UserProfile
    var wg sync.WaitGroup
	errchan := make(chan error, 3)
	wg.Add(3)
	go func () {
		defer wg.Done()
		err := controller.FetchUserInfo(username, &profile)
		errchan <- err
	}()

	go func() {
		defer wg.Done()
		err := controller.FetchUserStatus(username, &profile)
		errchan <- err
	}()

	go func(){
		defer wg.Done()
		err := controller.FetchUserRatings(username, &profile)
		errchan <- err
	}()
	wg.Wait()
	close(errchan)

	for err := range errchan {
		if err != nil {
			return err;
		}
	}
	// Export final profile for UI
	api.UserProfile = profile
	return nil
}
